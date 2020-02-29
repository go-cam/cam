package camHttp

import (
	"github.com/go-cam/cam/camBase"
	"github.com/go-cam/cam/camModels"
	"github.com/go-cam/cam/camPluginContext"
	"github.com/go-cam/cam/camPluginRouter"
	"github.com/go-cam/cam/camUtils"
	"github.com/gorilla/sessions"
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// http server component
type HttpComponent struct {
	camBase.Component
	camPluginRouter.RouterPlugin
	camPluginContext.ContextPlugin

	config *HttpComponentConfig
	store  *sessions.FilesystemStore
}

// init
func (component *HttpComponent) Init(configInterface camBase.ComponentConfigInterface) {
	component.Component.Init(configInterface)

	configValue := reflect.ValueOf(configInterface)
	var config *HttpComponentConfig
	if configValue.Kind() == reflect.Ptr {
		config = configValue.Interface().(*HttpComponentConfig)
	} else if configValue.Kind() == reflect.Struct {
		configStruct := configValue.Interface().(HttpComponentConfig)
		config = &configStruct
	} else {
		panic("illegal config")
	}

	component.config = config
	component.RouterPlugin.Init(&config.RouterPluginConfig)
	component.ContextPlugin.Init(&config.ContextPluginConfig)

	//component.controllerDict, component.controllerActionDict = camComponents.Common.GetControllerDict(component.config.ControllerList)
	component.store = component.getFilesystemStore()
}

// start
func (component *HttpComponent) Start() {
	component.Component.Start()

	if !component.config.SslOnly {
		go component.listenAndServe()
	}
	if component.config.IsSslOn {
		go component.listenAndServeTLS()
	}
}

// stop
func (component *HttpComponent) Stop() {
	component.Component.Stop()
}

// Receive http request, Call controller action, Send http response
func (component *HttpComponent) handlerFunc(responseWriter http.ResponseWriter, request *http.Request) {
	defer func() {
		if rec := recover(); rec != nil {
			responseWriter.WriteHeader(500)
			panic(rec)
			//_, _ = responseWriter.Write([]byte(rec.(string)))
		}
	}()

	route := ""
	url := request.URL.String()
	dirs := camUtils.Url.SplitUrl(url)
	dirLen := len(dirs)
	if dirLen == 1 {
		route = dirs[0]
	} else {
		route = dirs[0] + "/" + dirs[1]
	}
	controller, action := component.RouterPlugin.GetControllerAction(route)
	if controller == nil || action == nil {
		panic("404")
	}

	session, err := component.store.Get(request, component.config.SessionName)
	if err != nil {
		panic("get session fail:" + err.Error())
	}
	sessionModel := camModels.NewHttpSession(session)
	contextModel := component.NewContext()
	contextModel.SetSession(sessionModel)
	values := component.getRequestValues(request)

	controller.Init()
	controller.SetApp(component.App)
	controller.SetContext(contextModel)
	controller.SetValues(values)

	if !controller.BeforeAction(action) {
		panic("invalid request")
	}
	action.Call()
	response := controller.AfterAction(action, controller.GetResponse())

	err = session.Save(request, responseWriter)
	if err != nil {
		panic(err)
	}

	responseWriter.WriteHeader(200)
	_, err = responseWriter.Write(response)
	if err != nil {
		panic(err)
	}
}

// get session store
func (component *HttpComponent) getFilesystemStore() *sessions.FilesystemStore {
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
func (component *HttpComponent) listenAndServe() {
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
func (component *HttpComponent) listenAndServeTLS() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", component.handlerFunc)
	server := &http.Server{
		Addr:    ":" + strconv.FormatUint(uint64(component.config.SslPort), 10),
		Handler: mux,
	}
	err := server.ListenAndServeTLS(component.config.SslCertFile, component.config.SslKeyFile)
	camUtils.Error.Panic(err)
}

// get request params
func (component *HttpComponent) getRequestValues(request *http.Request) map[string]interface{} {
	values := map[string]interface{}{}

	// parse params from request url
	_ = request.ParseForm()
	for key, value := range request.Form {
		values[key] = value
	}

	// parse params from form data
	contentType := request.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "multipart/form-data") {
		// multipart/form-data; boundary=----WebKitFormBoundaryDumfytNg1NzoZq2r
		boundaryRegexp, _ := regexp.Compile("boundary=([-|0-9a-zA-Z]+)")
		boundaries := boundaryRegexp.FindStringSubmatch(contentType)
		if len(boundaries) < 2 {
			panic("fail to parse form values")
		}
		boundary := "--" + boundaries[1]

		bytes, _ := ioutil.ReadAll(request.Body)
		bodyStr := string(bytes)
		paramsStrList := strings.Split(bodyStr, boundary)

		for _, row := range paramsStrList {
			if row == "" || !strings.Contains(row, "\"") {
				// exclude row
				continue
			}

			repl := "Content-Disposition: form-data; name=\"([0-9a-zA-Z|_]+)\""
			keyRegexp, _ := regexp.Compile(repl)
			keyList := keyRegexp.FindStringSubmatch(row)
			key := keyList[1]

			valueRow := keyRegexp.ReplaceAllString(row, "")
			value := strings.Trim(valueRow, "\n")
			value = strings.Trim(value, "\r")
			value = strings.Trim(value, "\r\n")
			value = strings.Trim(value, " ")

			values[key] = values
		}
	}

	return values
}
