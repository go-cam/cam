package camHttp

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camUtils"
	"github.com/go-cam/cam/plugin/camPluginContext"
	"github.com/go-cam/cam/plugin/camPluginRouter"
	"github.com/gorilla/sessions"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"syscall"
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
func (comp *HttpComponent) Init(configI camBase.ComponentConfigInterface) {
	comp.Component.Init(configI)

	var ok bool
	comp.config, ok = configI.(*HttpComponentConfig)
	if !ok {
		camBase.App.Error("HttpComponent", "invalid config")
	}
	comp.RouterPlugin.Init(&comp.config.RouterPluginConfig)
	comp.ContextPlugin.Init(&comp.config.ContextPluginConfig)
	comp.store = comp.getFilesystemStore()
}

// start
func (comp *HttpComponent) Start() {
	comp.Component.Start()

	if !comp.config.SslOnly {
		camBase.App.Info("HttpComponent", "listen http://:"+strconv.FormatUint(uint64(comp.config.Port), 10))
		go comp.listenAndServe()
	}
	if comp.config.IsSslOn {
		camBase.App.Info("HttpComponent", "listen https://:"+strconv.FormatUint(uint64(comp.config.SslPort), 10))
		go comp.listenAndServeTLS()
	}
}

// stop
func (comp *HttpComponent) Stop() {
	comp.Component.Stop()
}

// Receive http request, Call controller action, Send http response
func (comp *HttpComponent) handlerFunc(responseWriter http.ResponseWriter, request *http.Request) {
	defer func() {
		if rec := recover(); rec != nil {
			panic(rec)
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
	handler := comp.getCustomRoute(route)
	if handler != nil {
		handler(responseWriter, request)
		return
	}

	controller, action := comp.GetControllerAction(route)
	if controller == nil || action == nil {
		panic("404")
	}

	storeSession := comp.getStoreSession(request)
	context := comp.NewContext()
	session := NewHttpSession(storeSession)
	values := comp.getRequestValues(request)

	controller.Init()
	controller.SetContext(context)
	controller.SetSession(session)
	controller.SetValues(values)

	if !controller.BeforeAction(action) {
		panic("invalid request")
	}
	action.Call()
	response := controller.AfterAction(action, controller.GetResponse())

	err := storeSession.Save(request, responseWriter)
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
func (comp *HttpComponent) getFilesystemStore() *sessions.FilesystemStore {
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
func (comp *HttpComponent) listenAndServe() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", comp.handlerFunc)
	server := &http.Server{
		Addr:    ":" + strconv.FormatUint(uint64(comp.config.Port), 10),
		Handler: mux,
	}
	err := server.ListenAndServe()
	camUtils.Error.Panic(err)
}

// enable server with SSl
func (comp *HttpComponent) listenAndServeTLS() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", comp.handlerFunc)
	server := &http.Server{
		Addr:    ":" + strconv.FormatUint(uint64(comp.config.SslPort), 10),
		Handler: mux,
	}
	err := server.ListenAndServeTLS(comp.config.SslCertFile, comp.config.SslKeyFile)
	camUtils.Error.Panic(err)
}

// get request params
func (comp *HttpComponent) getRequestValues(request *http.Request) map[string]interface{} {
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

// get http session
func (comp *HttpComponent) getStoreSession(request *http.Request) *sessions.Session {
	session, err := comp.store.Get(request, comp.config.SessionName)
	if err != nil {
		osPathErr, ok := err.(*os.PathError)
		if !ok {
			panic(err.Error())
		}
		syscallErr, ok := osPathErr.Err.(syscall.Errno)
		if !ok {
			panic(osPathErr.Err.Error())
		}

		// allow error: syscall.ERROR_FILE_NOT_FOUND
		if syscallErr != syscall.ERROR_FILE_NOT_FOUND {
			panic(syscallErr.Error())
		}
	}

	return session
}

// get custom route handler
func (comp *HttpComponent) getCustomRoute(route string) camBase.HttpRouteHandler {
	handler, has := comp.config.routeHandlerDict[route]
	if !has {
		return nil
	}
	return handler
}
