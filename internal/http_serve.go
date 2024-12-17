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
	// 找到代理路由
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
