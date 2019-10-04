package components

import (
	"cin/base"
	"cin/configs"
	"net/http"
	"reflect"
	"strconv"
)

// http服务
type HttpServer struct {
	Base

	port              uint16                     // http 端口
	handlerDict       map[string]reflect.Type    // 控制器反射map
	handlerActionDict map[string]map[string]bool // 控制器 => 方法 => 是否存在（注册时记录）
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
	component.port = config.Port
	component.handlerDict, component.handlerActionDict = common.getControllerDict(config.HandlerList)
}

// 启动
func (component *HttpServer) Start() {
	component.Base.Start()

	mux := http.NewServeMux()
	mux.HandleFunc("/", component.handlerFunc)
	server := &http.Server{
		Addr:    "0.0.0.0:" + strconv.FormatUint(uint64(component.port), 10),
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
	//url := r.URL.String()
}
