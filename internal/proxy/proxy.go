package proxy

import (
	"dip/internal/logger"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type DipProxy struct {
	Namespace   string            `json:"namespace" yaml:"namespace"`
	Name        string            `json:"name" yaml:"name"`
	ServiceName string            `json:"serviceName" yaml:"serviceName"`
	Path        string            `json:"path" yaml:"path"`
	PathType    string            `json:"pathType" yaml:"pathType"`
	Tag         map[string]string `json:"tag" yaml:"tag"`
	Upstream    Upstream
}

type Upstream struct {
	ServiceName string `json:"serviceName" yaml:"serviceName"`
	Endpoint    string `json:"endpoint" yaml:"endpoint"`
}

var (
	PrefixPathType  = "Prefix"
	ExtractPathType = "Extract"
)

type DipProxyContext struct {
}

type DipProxyTransport struct {
	proxyLogger *logger.ProxyLogger
	sourceUrl   string
	targetUrl   string
}

func (transport DipProxyTransport) RoundTrip(r *http.Request) (*http.Response, error) {
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
		Path:         transport.sourceUrl,
		ProxyTarget:  transport.targetUrl,
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
	reverseProxy.Transport = DipProxyTransport{
		proxyLogger: proxyLogger,
	}
	return reverseProxy, nil
}
