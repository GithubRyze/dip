package logger

import (
	"log"
	"time"
)

type ProxyLog struct {
	Timestamp    time.Time
	ClientIP     string
	Method       string
	Path         string
	ProxyTarget  string
	RequestBody  string
	ResponseBody string
	StatusCode   int
	Latency      time.Duration
}

// Logger 异步日志处理器
type ProxyLogger struct {
	logChan  chan ProxyLog
	stopChan chan struct{}
}

// NewLogger 创建一个 Logger 实例
func NewProxyLogger(bufferSize int) *ProxyLogger {
	l := &ProxyLogger{
		logChan:  make(chan ProxyLog, bufferSize),
		stopChan: make(chan struct{}),
	}
	go l.run()
	return l
}

// Log 将日志信息发送到日志处理器
func (l *ProxyLogger) Log(entry ProxyLog) {
	l.logChan <- entry
}

// Close 关闭日志处理器
func (l *ProxyLogger) Close() {
	close(l.logChan)
	<-l.stopChan
}

// run 异步处理日志
func (l *ProxyLogger) run() {
	for {
		select {
		case entry := <-l.logChan:
			// 处理日志
			log.Printf(
				"[%s] Client: %s | Method: %s | Path: %s | ProxyTarget: %s | Status: %d | Latency: %s | RequestBody: %s | ResponseBody: %s\n",
				entry.Timestamp.Format(time.RFC3339),
				entry.ClientIP, entry.Method, entry.Path, entry.ProxyTarget, entry.StatusCode,
				entry.Latency, entry.RequestBody, entry.ResponseBody,
			)
		case <-l.stopChan:
			// 收到停止信号，退出日志处理
			close(l.stopChan)
			return
		}
	}
}
