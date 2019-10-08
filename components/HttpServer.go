package components

import (
	"cin/base"
	"cin/configs"
	"cin/models"
	"cin/utils"
	"github.com/gorilla/sessions"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

// http服务
type HttpServer struct {
	Base

	config *configs.HttpServer

	handlerDict       map[string]reflect.Type    // 控制器反射map
	handlerActionDict map[string]map[string]bool // 控制器 => 方法 => 是否存在（注册时记录）

	store *sessions.FilesystemStore
}

// 使用配置初始化数据
func (component *HttpServer) Init(configInterface base.ConfigComponentInterface) {
	component.Base.Init(configInterface)

	configValue := reflect.ValueOf(configInterface)
	var config *configs.HttpServer
	if configValue.Kind() == reflect.Ptr {
		config = configValue.Interface().(*configs.HttpServer)
	} else if configValue.Kind() == reflect.Struct {
		configStruct := configValue.Interface().(configs.HttpServer)
		config = &configStruct
	} else {
		panic("illegal config")
	}

	component.name = component.getComponentName(configInterface.GetComponent())
	component.config = config
	component.handlerDict, component.handlerActionDict = common.getControllerDict(component.config.HandlerList)
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
			_, _ = w.Write([]byte(rec.(string)))
		}
	}()
	// 返回数据
	response := []byte("")

	url := r.URL.String()
	session, err := component.store.Get(r, component.config.SessionName)
	sessionModel := models.NewHttpSession(session)
	contextModel := models.NewHttpContext(sessionModel)
	if err != nil {
		panic("get session fail:" + err.Error())
	}

	tmpStr := strings.Split(url, "/")
	if len(tmpStr) < 3 {
		// TODO 这里添加自定义路由
		panic("404")
	}

	handlerName := tmpStr[1]
	actionName := utils.String.UrlToHump(tmpStr[2])
	hasAction := false // 动作是否存在
	if actionDict, has := component.handlerActionDict[handlerName]; has {
		_, hasAction = actionDict[actionName]
	}

	if hasAction {
		response = component.callHandler(handlerName, actionName, contextModel)
	}


	err = component.store.Save(r, w, session)
	if err != nil {
		panic("session save failed!")
	}

	w.WriteHeader(200)
	_, _ = w.Write(response)
}

// 调用控制器处理
func (component *HttpServer) callHandler(handlerName string, actionName string, context *models.HttpContext) []byte {
	response := []byte("")

	handlerType := component.handlerDict[handlerName]
	handlerValue := reflect.New(handlerType.Elem())
	httpHandlerInterface := handlerValue.Interface().(base.HttpHandlerInterface)
	if httpHandlerInterface == nil {
		panic("controller must be implement base.WebsocketHandlerInterface")
	}
	handlerInterface := handlerValue.Interface().(base.HandlerInterface)
	if handlerInterface == nil {
		panic("controller must be implement base.HandlerInterface")
	}

	// 设置控制器数据
	handlerInterface.SetContext(context)

	// BeforeAction
	if !handlerInterface.BeforeAction(actionName) {
		panic("illegal request")
	}

	// DoAction
	action := handlerValue.MethodByName(actionName)
	retValues := action.Call([]reflect.Value{})
	if len(retValues) != 1 || retValues[0].Kind() != reflect.String {
		panic("only one argument of type []byte can be returned")
	}
	response = retValues[0].Interface().([]byte)

	// AfterAction
	response = handlerInterface.AfterAction(actionName, response)

	return response
}

// 获取文件 session store
func (component *HttpServer) getFilesystemStore() *sessions.FilesystemStore {
	runtimeDir := utils.File.GetRunPath() + "/runtime"
	if !utils.File.Exists(runtimeDir) {
		err := utils.File.Mkdir(runtimeDir)
		if err != nil {
			panic("create runtime dir failed!")
		}
	}
	return sessions.NewFilesystemStore(runtimeDir)
}