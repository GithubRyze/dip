package internal

import (
	"dip/internal/errors"
	"dip/internal/logger"
	"dip/internal/proxy"
	"net/http"
)

var (
	proxyLogger = logger.NewProxyLogger(10000)
)

type DipHttpServer struct {
}

func (DipHttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestPath := r.URL.Path
	dipProxy, err := proxy.ProxyManager.MatchPath(requestPath)
	// 找到路由
	if err != errors.NotMatchRouterError {
		reverseProxy, err := dipProxy.DoProxy(requestPath, proxyLogger)
		if err != nil {
			//todo 加入日志
			w.Header().Set("server", "dip")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("dip internal error"))
			return
		}
		reverseProxy.ServeHTTP(w, r)
		return
	}
	// 继续匹配 workflow 的触发器

}

type DipProxyTransport struct {
}

func (transport DipProxyTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	//start := time.Now()
	//requestDump, err := httputil.DumpRequestOut(r, true)
	//if err != nil {
	//	return nil, err
	//}
	return http.DefaultTransport.RoundTrip(r)
	//if err != nil {
	//	return nil, err
	//}
	//responseDump, err := httputil.DumpResponse(response, true)
	//if err != nil {
	//	// copying the response body did not work
	//	return nil, err
	//}
	//transport.proxyLogger.Log(logger.ProxyLog{
	//	Timestamp:    time.Now(),
	//	ClientIP:     r.RemoteAddr,
	//	Method:       r.Method,
	//	Path:         transport.sourceUrl,
	//	ProxyTarget:  transport.targetUrl,
	//	RequestBody:  string(requestDump),
	//	ResponseBody: string(responseDump),
	//	StatusCode:   response.StatusCode,
	//	Latency:      time.Since(start),
	//})
	//return response, err
}
