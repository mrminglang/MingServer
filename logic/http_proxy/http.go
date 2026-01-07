package http_proxy

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"server/repositories/models"
	"strconv"

	"github.com/mrminglang/tools/strings"
	"gitlab.upchinaproduct.com/taf/tafgo/taf"
	"gitlab.upchinaproduct.com/taf/tafgo/taf/util/sync"
	"gitlab.upchinaproduct.com/upgo/utils/server_utils/log"
)

// 输出日志
var logHttpProxy = taf.GetDayLogger("http_proxy", 1)
var logAuthByHttp = taf.GetDayLogger("auth_by_http", 1)
var logSendByHttp = taf.GetDayLogger("send_by_http", 1)

// httpProxy http服务单例
var httpProxy *HttpProxy
var once sync.Once

type HttpProxy struct {
	Servant string // httpServant名字
}

// GetHttpProxy http
func GetHttpProxy(httpServant string) *HttpProxy {
	_ = once.Do(func() error {
		logHttpProxy.Infof("GetHttpProxy start......")
		httpProxy = &HttpProxy{
			Servant: httpServant,
		}
		return nil
	})

	return httpProxy
}

// Run the application
func (hp *HttpProxy) Run() {
	logHttpProxy.Debugf("Run httpServer start......")

	cfg := taf.GetServerConfig()

	httpSCfg, has := cfg.Adapters[hp.Servant+"Adapter"]
	if !has {
		log.Conf.Errorf("Run httpServant conf not found:%s\n", hp.Servant)
		os.Exit(0)
	}
	logHttpProxy.Debugf("Run [serverConfig]|%+v", httpSCfg)

	// 初始化路由
	mux := hp.initRouter()

	logHttpProxy.Debugf("Run router init done......")

	// 避免阻塞主进程
	go func() {
		_ = http.ListenAndServe(fmt.Sprintf("%s:%s", httpSCfg.Endpoint.Host, strconv.Itoa(int(httpSCfg.Endpoint.Port))), mux)
	}() // 设置监听的端口

	logHttpProxy.Debugf("Run httpServer is running.......")
}

// InitRouter 初始化路由
func (hp *HttpProxy) initRouter() *http.ServeMux {
	mux := http.NewServeMux()
	// debug
	mux.HandleFunc("/debug/pprof", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "pprof is disabled", http.StatusNotFound)
	})
	// 认证接口
	mux.HandleFunc("/apush/gateway/api/auth", AuthByHttp)

	return mux
}

// AuthByHttp 认证接口
func AuthByHttp(w http.ResponseWriter, r *http.Request) {
	// 回包json数据
	rspHandler := models.NewHTTPResponseHandler(w, "AuthByHttp")
	logAuthByHttp.Infof("{%s r.URL::%+v start......}", rspHandler.GetFuncName(), r.URL)
	// 返回数据
	rsp := models.AuthRsp{}
	rsp.Code = -1

	// 请求参数
	req := models.AuthReq{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		rsp.Note = "解析请求参数错误!!"
		logAuthByHttp.Errorf("{%s Decode req error::%s}", rspHandler.GetFuncName(), err.Error())
		_ = rspHandler.WriteResponse(rsp)
		return
	}
	logAuthByHttp.Infof("{%s req::%s}", rspHandler.GetFuncName(), strings.Display(req))

	// 校验请求参数
	if req.LoginId == "" || req.Sign == "" {
		rsp.Note = "参数loginId、sign不能为空。"
		logAuthByHttp.Errorf("{%s req param error::%s}", rspHandler.GetFuncName(), rsp.Note)
		_ = rspHandler.WriteResponse(rsp)
		return
	}

	// 认证接口
	//todo 具体业务内容

	if err := rspHandler.WriteResponse(rsp); err != nil {
		rsp.Note = "编码返回结果错误!"
		logAuthByHttp.Errorf("{%s Encode rsp::%s error::%s}", rspHandler.GetFuncName(), strings.Display(rsp), rsp.Note)
		return
	}
	logAuthByHttp.Infof("{%s rsp::%s end......}", rspHandler.GetFuncName(), strings.Display(rsp))
	return
}
