package http

import (
	"io"
	"net/http"
	"os"
	"runtime"
	"strconv"
	mercari "webhook/src/controller"
	"webhook/src/global/enum"
	"webhook/src/logger"
)

func Start() {
	if enum.CMD.Port > 0 {
		enum.CONFIG.Server.Port = enum.CMD.Port
	}
	server := http.Server{Addr: enum.CONFIG.Server.Address + ":" + strconv.FormatUint(enum.CONFIG.Server.Port, 10)}
	// 欢迎页
	http.HandleFunc("/", handler(welcome))
	// favicon.ico
	http.HandleFunc("/favicon.ico", favicon)
	//注册路由  func是回调函数，用于路由的响应，
	http.HandleFunc(enum.CONFIG.Server.Context, handler(mercari.GetItems))
	//请求指定路由，返回指定结果
	http.HandleFunc("/name", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, "这是/name返回的数据")
	})
	// 禁用客户端连接缓存到连接池
	//http.DefaultTransport.(*http.Transport).DisableKeepAlives = true
	logger.Infof("Http server listen %s:%+v", enum.CONFIG.Server.Address, enum.CONFIG.Server.Port)
	if err := server.ListenAndServe(); err != nil {
		logger.Error("Http server 服务失败\n", err)
		return
	}
}

func handler(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		logger.Infof("%+v %+v %+v", request.Method, request.RemoteAddr, request.Header.Get("User-Agent"))
		//logger.Infof("%+v %+v\n%+v", request.Method, request.RemoteAddr, util.ToJsonString(request.Header))
		handlerFunc(response, request)
	}
}

func welcome(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Add("content-type", "text/html; charset=UTF-8")
	writer.Header().Add("server", runtime.Version())
	serverName := "WebHook"
	_, _ = io.WriteString(writer, "<meta charset=\"UTF-8\"><style>*{ padding: 0; margin: 0; } .think_default_text{ padding: 4px 48px;} a{color:#2E5CD5;cursor: pointer;text-decoration: none} a:hover{text-decoration:underline; } body{ background: #fff; font-family: \"Century Gothic\",\"Microsoft yahei\",serif; color: #333;font-size:18px} h1{ font-size: 100px; font-weight: normal; margin-bottom: 12px; } p{ line-height: 1.6em; font-size: 42px }</style><div style=\"padding: 24px 48px;\"> <h1>:)</h1><p>"+serverName+" <span style=\"font-size:30px\">v"+enum.VERSION+"</span><br/><span style=\"font-size:30px\">"+runtime.GOOS+" "+runtime.GOARCH+" "+runtime.Version()+"</span></p><span style=\"font-size:22px;\">[ 由涛哥倾情奉献 - 异步 协程 高性能 网络通信引擎 ]</span></div><think parse=\"1\" style=\"display: block; overflow: hidden;\"><div class=\"think_default_text\"></div></think><script>document.title='"+serverName+"';</script>")
}

func favicon(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Add("content-type", "image/x-icon")
	writer.Header().Add("server", runtime.Version())
	data, _ := os.ReadFile(enum.RootPath + string(os.PathSeparator) + "favicon.ico")
	_, _ = writer.Write(data)
}
