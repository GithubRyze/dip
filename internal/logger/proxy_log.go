package logger

import (
	"log"
	"time"
)

type ProxyLog struct {
	Timestamp      time.Time
	ClientIP       string
	Method         string
	SourceServicer string
	SourcePath     string
	TargetServicer string
	TargetPath     string
	RequestBody    string
	ResponseBody   string
	StatusCode     int
	Latency        time.Duration
}

type ProxyLogger struct {
	logChan chan ProxyLog
}

func NewProxyLogger(bufferSize int) *ProxyLogger {
	l := &ProxyLogger{
		logChan: make(chan ProxyLog, bufferSize),
	}
	go l.run()
	return l
}

func (l *ProxyLogger) Log(entry ProxyLog) {
	l.logChan <- entry
}

func (l *ProxyLogger) Close() {
	close(l.logChan)
}

func (l *ProxyLogger) run() {

	for entry := range l.logChan {
		log.Printf(
			"[%s] Client: %s | Method: %s | SourcePath: %s | TargetPath: %s | Status: %d | Latency: %s | RequestBody: %s | ResponseBody: %s\n",
			entry.Timestamp.Format(time.RFC3339),
			entry.ClientIP, entry.Method, entry.SourcePath, entry.TargetPath, entry.StatusCode,
			entry.Latency, entry.RequestBody, entry.ResponseBody,
		)
	}
	log.Print("logChan closed, stopping logger\n")

}
