package proxy

import (
	"bytes"
	"dip/internal/logger"
	"net/http"
	"testing"
	"time"
)

func TestDipProxy_DoProxy(t *testing.T) {
	proxy := DipProxy{
		Path:     "/test/001",
		PathType: ExactPathType,
		Upstream: Upstream{
			Endpoint: "https://example.com/dcms/test/0101",
		},
	}
	_, err := proxy.DoProxy("/test/path", nil)
	if err != nil {
		t.Errorf("expect no err, but get %s", err.Error())
	}
	proxy = DipProxy{
		Path:     "/test/001",
		PathType: ExactPathType,
		Upstream: Upstream{
			Endpoint: "http://example.com/dcms/test/0101%xx",
		},
	}
	_, err = proxy.DoProxy("/test/path", nil)
	if err == nil {
		t.Errorf("expect err, but get nil")
	}

}

func Test_dipProxyTransport_RoundTrip(t *testing.T) {
	dip := dipProxyTransport{
		proxyLogger:    logger.NewProxyLogger(1),
		SourceUrl:      "/test/path",
		TargetUrl:      "http://localhost:30688/test/path",
		SourceServicer: "source",
		TargetServicer: "target",
	}
	handler := http.NewServeMux()
	handler.HandleFunc("/test/path", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	server := &http.Server{
		Addr:    ":30688", // 指定端口 8080
		Handler: handler,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			t.Errorf("http server start err: %s", err)
		}
	}()
	time.Sleep(2 * time.Second)
	r, _ := http.NewRequest("GET", "http://localhost:30688/test/path", bytes.NewBuffer([]byte(`{"key":"value"}`)))
	res, err := dip.RoundTrip(r)
	if err != nil {
		t.Errorf("expected no err, but get %s", err.Error())
	}
	if res.StatusCode != 200 {
		t.Errorf("expected get 200, but get %d", res.StatusCode)
	}
}
func Test_dipProxyTransport_getTarget(t *testing.T) {
	proxy := DipProxy{
		Path:     "/test/001",
		PathType: ExactPathType,
		Upstream: Upstream{
			Endpoint: "https://example.com/dcms/test/0101",
		},
	}
	url := proxy.getTargetUrl("/test/path")
	if url != "https://example.com/dcms/test/0101" {
		t.Errorf("expect https://example.com/dcms/test/0101, but get %s", url)
	}
	proxy = DipProxy{
		Path:     "/test/001",
		PathType: PrefixPathType,
		Upstream: Upstream{
			Endpoint: "https://example.com",
		},
	}
	url = proxy.getTargetUrl("/test/path")
	if url != "https://example.com/test/path" {
		t.Errorf("expect https://example.com/test/path, but get %s", url)
	}
	proxy = DipProxy{
		Path:     "/test/001",
		PathType: "PrefixPathType",
		Upstream: Upstream{
			Endpoint: "https://example.com",
		},
	}
	url = proxy.getTargetUrl("/test/path")
	if url != "" {
		t.Errorf("expect \"\", but get %s", url)
	}
}
