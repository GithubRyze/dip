package internal

import (
	"dip/bootstrap/dip_logger"
	"dip/internal/errors"
	"dip/internal/logger"
	"dip/internal/proxy"
	"dip/pkg"
	"net/http"
	"runtime/debug"
	"time"
)

var (
	proxyLogger = logger.NewProxyLogger(10000)
)

type DipHttpServer struct {
}

func (DipHttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	clientIP := pkg.GetClientIP(r)
	requestPath := r.URL.Path
	dip_logger.InfofAccess("[Dip] %v | %s | %s | %s",
		time.Now().Format("2006/01/02 - 15:04:05"),
		clientIP,
		method,
		requestPath,
	)
	defer func() {
		if err := recover(); err != nil {
			dip_logger.Errorf("Panic recovered: %v\nStack trace:\n%s", err, string(debug.Stack()))
			http.Error(w, errors.InternalServerError.Error(), http.StatusInternalServerError)
		}
	}()
	dipProxy, err := proxy.ProxyManager.MatchPath(requestPath)
	// 找到代理路由
	if err != errors.NotMatchRouterError {
		reverseProxy, err := dipProxy.DoProxy(requestPath, proxyLogger)
		if err != nil {
			dip_logger.Errorf("reverseProxy error: %s", err.Error())
			http.Error(w, errors.InternalServerError.Error(), http.StatusInternalServerError)
			return
		}
		reverseProxy.ServeHTTP(w, r)
		return
	}
	// 继续匹配 workflow 的触发器
	http.Error(w, errors.NotMatchRouterError.Error(), http.StatusNotFound)
}
