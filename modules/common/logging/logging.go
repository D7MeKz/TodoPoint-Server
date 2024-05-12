package logging

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var (
	setupOnce sync.Once
)

type AppLogger struct {
	aLogger *logrus.Logger
}

func NewAppLogger() *AppLogger {
	var appLogger = &AppLogger{}
	setupOnce.Do(func() {
		appLogger.aLogger = logrus.New()
		appLogger.aLogger.SetLevel(logrus.DebugLevel)
		appLogger.aLogger.SetFormatter(&logrus.TextFormatter{
			DisableColors: true,
			FullTimestamp: true,
		})
		appLogger.aLogger.SetReportCaller(true)
		appLogger.aLogger.SetOutput(os.Stdout)
	})
	return appLogger
}

// NewLoggingField creates a new logrus.Fields object with the given context and code
func NewLoggingField(url string, code int) *logrus.Fields {
	return &logrus.Fields{
		// url is extracted from context
		"url": url,
		// code is the Error code.
		"code": code,
	}
}

func (l *AppLogger) Info(ctx *gin.Context, code int, msg string) {
	l.aLogger.WithFields(*NewLoggingField(ctx.Request.URL.Path, code)).Info(msg)
}

func (l *AppLogger) Debug(ctx *gin.Context, code int, msg string) {
	l.aLogger.WithFields(*NewLoggingField(ctx.Request.URL.Path, code)).Debug(msg)
}

func (l *AppLogger) Error(ctx *gin.Context, code int, msg string) {
	l.aLogger.WithFields(*NewLoggingField(ctx.Request.URL.Path, code)).Error(msg)
}

func (l *AppLogger) Warn(ctx *gin.Context, code int, msg string) {
	l.aLogger.WithFields(*NewLoggingField(ctx.Request.URL.Path, code)).Warn(msg)
}
