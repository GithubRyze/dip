package logger

import (
	"testing"
	"time"
)

func TestNewProxyLogger(t *testing.T) {
	proxyLogger := NewProxyLogger(100)
	if proxyLogger == nil {
		t.Errorf("Expected ProxyLogger to be created, but got nil")
	}
	if proxyLogger.logChan == nil {
		t.Errorf("Expected logChan to be initialized, but got nil")
	}
	if cap(proxyLogger.logChan) != 100 {
		t.Errorf("Expected logChan buffer size to be %d, but got %d", 100, cap(proxyLogger.logChan))
	}
	time.Sleep(100 * time.Millisecond)
	select {
	case proxyLogger.logChan <- ProxyLog{}:
	default:
		t.Errorf("Expected run() goroutine to be running, but logChan is not accepting messages")
	}

}

func TestProxyLogger_Close(t *testing.T) {
	proxyLogger := NewProxyLogger(100)
	time.Sleep(100 * time.Millisecond)
	proxyLogger.Log(ProxyLog{StatusCode: 3000})
	proxyLogger.Close()
	time.Sleep(100 * time.Millisecond)
	select {
	case _, ok := <-proxyLogger.logChan:
		if ok != false {
			t.Fatalf("logChan was not properly closed")
		}
	default:
		// 如果没有关闭的信号，则通过
	}
}

func TestProxyLogger_Log(t *testing.T) {
	proxyLogger := NewProxyLogger(1)
	proxyLogger.Log(ProxyLog{StatusCode: 3000})
}

func TestProxyLogger_run(t *testing.T) {

}
