package camComponents

import (
	"github.com/go-cam/cam/core/camBase"
	"github.com/go-cam/cam/core/camConfigs"
	"github.com/go-cam/cam/core/camModels"
	"github.com/go-cam/cam/core/camUtils"
	"github.com/gorilla/sessions"
	"net/http"
	"reflect"
	"strconv"
)

// http server component
type HttpServer struct {
	Base

	config *camConfigs.HttpServer

	controllerDict       map[string]reflect.Type    // controller reflect.Type dict
	controllerActionDict map[string]map[string]bool // map[controllerName]map[actionName]

	store *sessions.FilesystemStore
}

// init
func (component *HttpServer) Init(configInterface camBase.ConfigComponentInterface) {
	component.Base.Init(configInterface)

	configValue := reflect.ValueOf(configInterface)
	var config *camConfigs.HttpServer
	if configValue.Kind() == reflect.Ptr {
		config = configValue.Interface().(*camConfigs.HttpServer)
	} else if configValue.Kind() == reflect.Struct {
		configStruct := configValue.Interface().(camConfigs.HttpServer)
		config = &configStruct
	} else {
		panic("illegal config")
	}

	component.name = component.getComponentName(configInterface.GetComponent())
	component.config = config
	component.controllerDict, component.controllerActionDict = common.getControllerDict(component.config.ControllerList)
	component.store = component.getFilesystemStore()
}

// start
func (component *HttpServer) Start() {
	component.Base.Start()

	if !component.config.IsSslOnly {
		go component.listenAndServe()
	}
	if component.config.IsSslOn {
		go component.listenAndServeTLS()
	}
}

// stop
func (component *HttpServer) Stop() {
	component.Base.Stop()
}

// Receive http request, Call controller action, Send http response
func (component *HttpServer) handlerFunc(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if rec := recover(); rec != nil {
			w.WriteHeader(500)
			panic(rec)
			//_, _ = w.Write([]byte(rec.(string)))
		}
	}()
	response := []byte("")

	url := r.URL.String()
	session, err := component.store.Get(r, component.config.SessionName)
	sessionModel := camModels.NewHttpSession(session)
	contextModel := component.config.NewContext()
	contextModel.SetSession(sessionModel)
	if err != nil {
		panic("get session fail:" + err.Error())
	}

	dirs := camUtils.Url.SplitUrl(url)
	dirLength := len(dirs)
	var controllerName string
	var actionName string
	if dirLength == 0 {
		// TODO default route
		panic("404")
	} else if dirLength == 1 {
		controllerName = camUtils.Url.UrlToHump(dirs[0])
		actionName = ""
	} else {
		controllerName = camUtils.Url.UrlToHump(dirs[0])
		actionName = camUtils.Url.UrlToHump(dirs[1])
	}

	hasAction := false // 动作是否存在
	if actionName == "" {
		hasAction = true
	} else if actionDict, has := component.controllerActionDict[controllerName]; has {
		_, hasAction = actionDict[actionName]
	}

	if hasAction {
		response = component.callControllerAction(controllerName, actionName, contextModel, w, r)
	}

	err = session.Save(r, w)
	if err != nil {
		panic("session save failed!" + err.Error())
	}

	w.WriteHeader(200)
	_, _ = w.Write(response)
}

// call controller action
func (component *HttpServer) callControllerAction(controllerName string, actionName string, context camBase.ContextInterface, w http.ResponseWriter, r *http.Request) []byte {
	response := []byte("")

	controllerType := component.controllerDict[controllerName]
	controllerValue := reflect.New(controllerType.Elem())
	controllerInterface := controllerValue.Interface().(camBase.ControllerInterface)
	if controllerInterface == nil {
		panic("controller must be implement base.ControllerInterface")
	}

	// set controller params
	controllerInterface.Init()
	controllerInterface.SetApp(component.app)
	controllerInterface.SetContext(context)
	controllerInterface.SetHttpValues(w, r)
	if actionName == "" {
		actionName = controllerInterface.GetDefaultAction()
		if actionName == "" {
			panic("404")
		}
	}

	// BeforeAction
	if !controllerInterface.BeforeAction(actionName) {
		panic("illegal request")
	}

	// DoAction
	action := controllerValue.MethodByName(actionName)
	_ = action.Call([]reflect.Value{})
	response = controllerInterface.Read()
	// AfterAction
	response = controllerInterface.AfterAction(actionName, response)

	return response
}

// get session store
func (component *HttpServer) getFilesystemStore() *sessions.FilesystemStore {
	runtimeDir := camUtils.File.GetRunPath() + "/runtime/session"
	if !camUtils.File.Exists(runtimeDir) {
		err := camUtils.File.Mkdir(runtimeDir)
		if err != nil {
			panic("create runtime dir failed! " + err.Error())
		}
	}
	return sessions.NewFilesystemStore(runtimeDir, []byte("none"))
}

// enable server
func (component *HttpServer) listenAndServe() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", component.handlerFunc)
	server := &http.Server{
		Addr:    ":" + strconv.FormatUint(uint64(component.config.Port), 10),
		Handler: mux,
	}
	err := server.ListenAndServe()
	camUtils.Error.Panic(err)
}

// enable server with SSl
func (component *HttpServer) listenAndServeTLS() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", component.handlerFunc)
	server := &http.Server{
		Addr:    ":" + strconv.FormatUint(uint64(component.config.SslPort), 10),
		Handler: mux,
	}
	err := server.ListenAndServeTLS(component.config.SslCertFile, component.config.SslKeyFile)
	camUtils.Error.Panic(err)
}
