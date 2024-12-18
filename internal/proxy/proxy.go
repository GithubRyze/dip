package proxy

import (
	"dip/bootstrap/dip_logger"
	"dip/internal/logger"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type DipProxy struct {
	Namespace   string
	Name        string
	ServiceName string
	Path        string
	PathType    string
	Tag         map[string]string
	Upstream    Upstream
}

type Upstream struct {
	ServiceName string
	Endpoint    string
}

var (
	PrefixPathType = "Prefix"
	ExactPathType  = "Exact"
)

type dipProxyTransport struct {
	proxyLogger    *logger.ProxyLogger
	SourceUrl      string
	TargetUrl      string
	SourceServicer string
	TargetServicer string
}

func (transport dipProxyTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	start := time.Now()
	requestDump, err := httputil.DumpRequestOut(r, true)
	if err != nil {
		dip_logger.Errorf("DumpRequestOut error: %s", err.Error())
		return nil, err
	}
	response, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		dip_logger.Errorf("RoundTrip error: %s", err.Error())
		return nil, err
	}
	responseDump, err := httputil.DumpResponse(response, true)
	if err != nil {
		dip_logger.Errorf("DumpResponse error: %s", err.Error())
		return nil, err
	}
	transport.proxyLogger.Log(logger.ProxyLog{
		Timestamp:    time.Now(),
		ClientIP:     r.RemoteAddr,
		Method:       r.Method,
		SourcePath:   transport.SourceUrl,
		TargetPath:   transport.TargetUrl,
		RequestBody:  string(requestDump),
		ResponseBody: string(responseDump),
		StatusCode:   response.StatusCode,
		Latency:      time.Since(start),
	})
	return response, err
}

func (dipProxy *DipProxy) DoProxy(requestPath string, proxyLogger *logger.ProxyLogger) (*httputil.ReverseProxy, error) {
	targetUrl, parseUrlError := url.Parse(dipProxy.getTargetUrl(requestPath))
	if parseUrlError != nil {
		return nil, parseUrlError
	}
	reverseProxy := httputil.NewSingleHostReverseProxy(targetUrl)
	reverseProxy.Transport = dipProxyTransport{
		proxyLogger:    proxyLogger,
		SourceUrl:      requestPath,
		TargetUrl:      targetUrl.Scheme + "://" + targetUrl.Host + targetUrl.Path,
		SourceServicer: dipProxy.ServiceName,
		TargetServicer: dipProxy.Upstream.ServiceName,
	}
	return reverseProxy, nil
}

func (dipProxy *DipProxy) getTargetUrl(requestPath string) string {
	if dipProxy.PathType == ExactPathType {
		return dipProxy.Upstream.Endpoint
	}
	if dipProxy.PathType == PrefixPathType {
		return dipProxy.Upstream.Endpoint + requestPath
	}
	return ""
}
