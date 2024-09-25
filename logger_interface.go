package icelogrus

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
)

// GormLogrusLogger 使用Logrus实现Gorm的Logger接口
//
// Use Logrus to implement GORM2's Logger Interface
type GormLogrusLogger struct {
	log *logrus.Logger
}

func (g *GormLogrusLogger) LogMode(level logger.LogLevel) logger.Interface {
	// 根据 Gorm 日志级别设置 Logrus 的输出级别
	switch level {
	case logger.Silent:
		g.log.SetLevel(logrus.PanicLevel)
	case logger.Error:
		g.log.SetLevel(logrus.ErrorLevel)
	case logger.Warn:
		g.log.SetLevel(logrus.WarnLevel)
	case logger.Info:
		g.log.SetLevel(logrus.InfoLevel)
	}
	return g
}

func (g *GormLogrusLogger) Info(ctx context.Context, msg string, args ...interface{}) {
	g.log.WithContext(ctx).Infof(msg, args...)
}

func (g *GormLogrusLogger) Warn(ctx context.Context, msg string, args ...interface{}) {
	g.log.WithContext(ctx).Warnf(msg, args...)
}

func (g *GormLogrusLogger) Error(ctx context.Context, msg string, args ...interface{}) {
	g.log.WithContext(ctx).Errorf(msg, args...)
}

func (g *GormLogrusLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil {
		g.log.WithContext(ctx).WithFields(logrus.Fields{
			"elapsed": elapsed,
			"rows":    rows,
		}).Errorf("SQL Error: %s", sql)
	} else {
		g.log.WithContext(ctx).WithFields(logrus.Fields{
			"elapsed": elapsed,
			"rows":    rows,
		}).Infof("Executed SQL: %s", sql) // 确保这里是 Info 而不是 Debug
	}
}

// NewGormLogrusLogger 创建一个新的 GormLogrusLogger 实例，用于将Logrus转换成GORM2可用的logger
//
// Create a new GormLogrusLogger instance to convert Logrus into a logger that GORM2 can use
func NewGormLogrusLogger(log *logrus.Logger) *GormLogrusLogger {
	return &GormLogrusLogger{log: log}
}
