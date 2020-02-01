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

// http服务
type HttpServer struct {
	Base

	config *camConfigs.HttpServer

	controllerDict       map[string]reflect.Type    // 控制器反射map
	controllerActionDict map[string]map[string]bool // 控制器 => 方法 => 是否存在（注册时记录）

	store *sessions.FilesystemStore
}

// 使用配置初始化数据
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

// 启动
func (component *HttpServer) Start() {
	component.Base.Start()

	mux := http.NewServeMux()
	mux.HandleFunc("/", component.handlerFunc)
	server := &http.Server{
		Addr:    ":" + strconv.FormatUint(uint64(component.config.Port), 10),
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func (component *HttpServer) Stop() {
	component.Base.Stop()
}

// http 处理方法
func (component *HttpServer) handlerFunc(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if rec := recover(); rec != nil {
			w.WriteHeader(500)
			panic(rec)
			//_, _ = w.Write([]byte(rec.(string)))
		}
	}()
	// 返回数据
	response := []byte("")

	url := r.URL.String()
	session, err := component.store.Get(r, component.config.SessionName)
	sessionModel := camModels.NewHttpSession(session)
	contextModel := camModels.NewContext(sessionModel)
	if err != nil {
		panic("get session fail:" + err.Error())
	}

	dirs := camUtils.Url.SplitUrl(url)
	dirLength := len(dirs)
	if dirLength == 0 {
		// TODO 默认路由
		panic("404")
	} else if dirLength == 1 {
		// TODO 控制器默认路由
		panic("404")
	}

	controllerName := camUtils.Url.UrlToHump(dirs[0])
	actionName := camUtils.Url.UrlToHump(dirs[1])
	hasAction := false // 动作是否存在
	if actionDict, has := component.controllerActionDict[controllerName]; has {
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

// 调用控制器处理
func (component *HttpServer) callControllerAction(controllerName string, actionName string, context *camModels.Context, w http.ResponseWriter, r *http.Request) []byte {
	response := []byte("")

	controllerType := component.controllerDict[controllerName]
	controllerValue := reflect.New(controllerType.Elem())
	controllerInterface := controllerValue.Interface().(camBase.ControllerInterface)
	if controllerInterface == nil {
		panic("controller must be implement base.ControllerInterface")
	}

	// 设置控制器数据
	controllerInterface.Init()
	controllerInterface.SetApp(component.app)
	controllerInterface.SetContext(context)
	controllerInterface.SetHttpValues(w, r)

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

// 获取文件 session store
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
