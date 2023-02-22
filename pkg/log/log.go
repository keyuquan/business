package log

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var (
	logTmFmtWithMS = "2006-01-02 15:04:05.000"
	Logger         *zap.Logger
)

func InitLog() {
	// 自定义时间输出格式
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + t.Format(logTmFmtWithMS) + "]")
	}
	// 自定义日志级别显示
	customLevelEncoder := func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + level.CapitalString() + "]")
	}

	// 自定义文件：行号输出项
	customCallerEncoder := func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + caller.TrimmedPath() + "]")
	}

	encoderConf := zapcore.EncoderConfig{
		CallerKey:      "caller_line", // 打印文件名和行数
		LevelKey:       "level_name",
		MessageKey:     "msg",
		TimeKey:        "ts",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     customTimeEncoder,   // 自定义时间格式
		EncodeLevel:    customLevelEncoder,  // 小写编码器
		EncodeCaller:   customCallerEncoder, // 全路径编码器
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 自定义格式
	// core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConf), zap.NewAtomicLevelAt(l.logMinLevel))
	//

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConf), os.Stdout, zap.NewAtomicLevelAt(zapcore.DebugLevel))
	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

type Trace interface {
	GetTraceID() string
}

func Info(msg string) {
	Logger.Info(msg)
}

func Infof(msg string, f ...interface{}) {
	Logger.Info(fmt.Sprintf(msg, f...))
}

func Error(msg string) {
	Logger.Error(msg)
}

func Errorf(msg string, f ...interface{}) {
	Logger.Error(fmt.Sprintf(msg, f...))
}

func Fatalf(msg string, f ...interface{}) {
	Logger.Fatal(fmt.Sprintf(msg, f...))
}

type Log struct {
	context.Context
	Trace string
}

func NewLog(trace string) *Log {
	return &Log{
		Context: context.Background(),
		Trace:   trace,
	}
}

func (l Log) Info(msg string) {
	Logger.With(zap.String("trace", l.Trace)).Info(msg)
}
func (l Log) Infof(msg string, f ...interface{}) {
	Logger.With(zap.String("trace", l.Trace)).Info(fmt.Sprintf(msg, f...))
}
func (l Log) Error(msg string) {
	Logger.With(zap.String("trace", l.Trace)).Error(msg)
}
func (l Log) Errorf(msg string, f ...interface{}) {
	Logger.With(zap.String("trace", l.Trace)).Error(fmt.Sprintf(msg, f...))
}
func (l Log) Fatalf(msg string, f ...interface{}) {
	Logger.With(zap.String("trace", l.Trace)).Fatal(fmt.Sprintf(msg, f...))
}
