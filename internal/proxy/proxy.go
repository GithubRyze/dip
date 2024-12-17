package proxy

import (
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
	PrefixPathType  = "Prefix"
	ExtractPathType = "Extract"
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
		return nil, err
	}
	response, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		return nil, err
	}
	responseDump, err := httputil.DumpResponse(response, true)
	if err != nil {
		// copying the response body did not work
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
	var targetUrl *url.URL
	var parseUrlError error
	if dipProxy.PathType == ExtractPathType {
		targetUrl, parseUrlError = url.Parse(dipProxy.Upstream.Endpoint)
	}
	if dipProxy.PathType == PrefixPathType {
		targetUrl, parseUrlError = url.Parse(dipProxy.Upstream.Endpoint + requestPath)
	}
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
