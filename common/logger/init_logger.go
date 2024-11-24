package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var globalLogger *zap.Logger // 全局 Logger

// InitLogger 初始化全局日志器，支持文件切分和终端输出
func InitLogger(logFile string, maxSize, maxBackups, maxAge int, compress bool, level zapcore.Level) {
	// 配置 lumberjack 实现日志切分
	logWriter := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    maxSize,    // 每个日志文件最大大小（MB）
		MaxBackups: maxBackups, // 保留旧文件的最大数量
		MaxAge:     maxAge,     // 保留旧文件的最大天数
		Compress:   compress,   // 是否压缩旧日志文件
	}

	// 文件日志输出 core
	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), // JSON 格式
		zapcore.AddSync(logWriter),                               // 文件输出
		level,                                                    // 日志级别
	)

	// 终端日志输出 core
	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()), // 控制台友好格式
		zapcore.AddSync(os.Stdout),                                   // 终端输出
		level,                                                        // 日志级别
	)

	// 合并多个 core
	core := zapcore.NewTee(fileCore, consoleCore)

	// 创建 logger，添加调用信息
	globalLogger = zap.New(core, zap.AddCaller())

	// 设置全局 Logger
	zap.ReplaceGlobals(globalLogger)
}

// L 获取全局日志器
func L() *zap.Logger {
	return globalLogger
}
