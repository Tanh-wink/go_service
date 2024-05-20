package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"time"
	"go/sevice/config"
)

var(
	conf = config.GetConfig()
)

type LogLevel uint16

//日志常量
const (
	UNKNOW LogLevel = iota
	TRACE
	DEBUG
	INFO
	WARN
	ERROR
	DPANIC
	PANIC
	FATAL

)

func (l LogLevel) String() string {
	switch l {
	case TRACE:
		return "TRACE"
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case DPANIC:
		return "DPANIC"
	case PANIC:
		return "PANIC"
	case FATAL:
		return "FATAL"
	default:
		panic(fmt.Sprint("Incorrect log level:", int(l)))
	}
}

//解析日志级别
func paraLogLevel(level string) (LogLevel, error) {
	level = strings.ToLower(level)
	switch level {
		case "tarce": return TRACE, nil
		case "debug": return DEBUG,nil
		case "info": return INFO, nil
		case "warn": return WARN, nil
		case "error": return ERROR, nil
		case "dpanic": return DPANIC, nil
		case "panic": return PANIC, nil
		case "fatal": return FATAL, nil
		default: 
			err := fmt.Errorf(fmt.Sprintf("Invalid log level: %s", level))
			return UNKNOW, err
	}
}


type ILogger interface{
	// Name 当前 Logger 的名称
	Name() string
	setLevel() LogLevel

	// Log 打印指定 level 的日志，其中：
	//
	//  - ctx 是上下文。如不存在上下文时可以传 nil，或通过 NoContext()
	//    获取 RawLogger 使用。
	//
	//  - message 是日志消息。
	//
	//    在 json 输出模式下通常是 "message" 字段；
	//
	//    在文本输出模式下是主体消息的开始部分
	//
	//  - fields 是日志的补充字段。
	//
	//    在 json 输出模式下是 "message" 字段的同级字段，字段 key 通过
	//    LogKey 或 LogEntry 接口获得；
	//
	//    在文本输出模式下，是在主体消息中跟随在 message 参数之后使用 "key=value"
	//    形式输出每一个值，其中 key 通过 LogKey 或 LogEntry 接口获得，value
	//    通过 LogValue 或 LogEntry 接口获得。
	Log(level LogLevel, msg string, fields ...any)

	// Logf 打印指定 level 的日志，其中：
	//
	//  - ctx 是上下文。如不存在上下文时可以传 nil，或通过 NoContext()
	//    获取 RawLogger 使用。
	//  - format 是 fmt.Sprintf 中的 format。
	//  - a 是是 fmt.Sprintf 中的 a，此处 a 仍然支持通过实现 LogValue 或 LogEntry
	//    接口定制输出值。
	Logf(level LogLevel, format string, a ...any)

	Trace(message string, fields ...any)
	Tracef(format string, a ...any)
	Debug(message string, fields ...any)
	Debugf(format string, a ...any)
	Info(message string, fields ...any)
	Infof(format string, a ...any)
	Warn(message string, fields ...any)
	Warnf(format string, a ...any)
	Error(message string, fields ...any)
	Errorf(format string, a ...any)

	IsTraceEnabled() bool
	IsDebugEnabled() bool
	IsInfoEnabled() bool
	IsWarnEnabled() bool
	IsErrorEnabled() bool
}


// Logger implements ILogger which Logger Log to Write and Logf to Writef.
type Logger struct {
	Name_  string
	Level LogLevel
	Write  func(level LogLevel, name string, message string, fields ...any)
	Writef func(level LogLevel, name string, format string, a ...any)
}

func (l *Logger) Name() string {
	return l.Name_
}

func (l *Logger) setLevel() LogLevel {
	return GetLogLevel()
}

func (l *Logger) Trace(message string, fields ...any) {
	l.Log(TRACE, message, fields...)
}
func (l *Logger) Tracef(format string, a ...any) {
	l.Logf(TRACE, format, a...)
}

func (l *Logger) Debug(message string, fields ...any) {
	l.Log(DEBUG, message, fields...)
}
func (l *Logger) Debugf(format string, a ...any) {
	l.Logf(DEBUG, format, a...)
}

func (l *Logger) Info(message string, fields ...any) {
	l.Log(INFO, message, fields...)
}
func (l *Logger) Infof(format string, a ...any) {
	l.Logf(INFO, format, a...)
}

func (l *Logger) Warn(message string, fields ...any) {
	l.Log(WARN, message, fields...)
}
func (l *Logger) Warnf(format string, a ...any) {
	l.Logf(WARN, format, a...)
}

func (l *Logger) Error(message string, fields ...any) {
	l.Log(ERROR, message, fields...)
}
func (l *Logger) Errorf(format string, a ...any) {
	l.Logf(ERROR, format, a...)
}

func (l *Logger) Log(level LogLevel, message string, fields ...any) {
	l.Write(level, l.Name_, message, fields...)
}

func (l *Logger) Logf(level LogLevel, format string, a ...any) {
	l.Writef(level, l.Name_, format, a...)
}

func (l *Logger) IsTraceEnabled() bool {
	return GetLogLevel() <= TRACE
}
func (l *Logger) IsDebugEnabled() bool {
	return GetLogLevel() <= DEBUG
}
func (l *Logger) IsInfoEnabled() bool {
	return GetLogLevel() <= INFO
}
func (l *Logger) IsWarnEnabled() bool {
	return GetLogLevel() <= WARN
}
func (l *Logger) IsErrorEnabled() bool {
	return GetLogLevel() <= ERROR
}

// log config

type LogConfig struct {
	// 日志输出器
	Appender string `json:"appender"`
	// 日志输出文件路径
	FilePath string `json:"file-path"`
	// 每个日志文件保存的大小 单位:M
	MaxSize int `json:"max-size"`
	// 文件最多保存多少天
	MaxAge int `json:"max-age"`
	// 日志文件最多保存多少个备份
	MaxBackups int `json:"max-backups"`
	// 是否压缩
	Compress bool `json:"compress"`
	// 日志级别
	LogLevel string `json:"log-level"`
	// error日志是否打印堆栈
	StacktraceLevel string `json:"stacktrace-level"`
	// 是否禁用日志缓冲
	DisableBuffer bool `json:"disable-buffer"`
	// 缓冲区大小 单位：KB
	BufferSize int `json:"buffer-size"`
	// 缓冲刷新间隔 单位：秒
	FlushInterval int64 `json:"flush-interval"`
	// logname维度的个性化配置
	LogNameConfigs map[string]*LogConfig `json:"name" configx:"name"`
}

var (
	logConf      *LogConfig
)

func GetLogLevel() LogLevel {
	if logConf == nil {
		return INFO
	}
	level, err := paraLogLevel(logConf.LogLevel)
	if err == nil {
		return level
	}
	return INFO
}


// get logger
func GetLogger(name string) ILogger  {
	var logFile string = ""
	if conf != nil {
		logFile = conf.LogFile 
	}
	logger := NewLogger(name, logFile)
	return logger
}

// build a logger
func NewLogger(name string, logFile string) ILogger {
	if len(logFile) != 0 {
		dayChangeLock := sync.RWMutex{}
		dayChangeLock.Lock()
		defer dayChangeLock.Unlock()
		now := time.Now()
		postFix := now.Format("2006-01-02")
		logFile := logFile + "." + postFix
		logOut, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err != nil {
			panic(err)
		} else {
			multiWriter := io.MultiWriter(os.Stdout, logOut)
			l := log.New(multiWriter, name+" ", log.Ldate|log.Lmicroseconds|log.Lmsgprefix)
			return NewLoggerWithWriter(name, l.Print, l.Printf)
		}
	} else {
		l := log.New(log.Writer(), name+" ", log.Ldate|log.Lmicroseconds|log.Lmsgprefix)
		return NewLoggerWithWriter(name, l.Print, l.Printf)
	}
	
}

func NewLoggerWithWriter(name string, write func(v ...any), writef func(format string, a ...any)) *Logger {
	return &Logger{
		Name_: name,
		Write: func(level LogLevel, name string, message string, fields ...any) {
			sb := strings.Builder{}
			sb.WriteString(level.String())
			// TODO: add trace id
			// TODO: add key biz items
			sb.WriteString(" -- ")
			sb.WriteString(message)
			for i, field := range fields {
				switch field := field.(type) {
				case LogEntry:
					k, v := field.LogValue()
					sb.WriteString(fmt.Sprintf(", %s=%v", k, v))
				case interface {
					LogValue
					LogKey
				}:
					sb.WriteString(fmt.Sprintf(", %s=%v", field.LogKey(), field.LogValue()))
				case LogValue:
					sb.WriteString(fmt.Sprintf(", #%d=%v", i, field.LogValue()))
				case LogKey:
					sb.WriteString(fmt.Sprintf(", %s=%v", field.LogKey(), field))
				default:
					sb.WriteString(fmt.Sprintf(", #%d=%v", i, field))
				}
			}
			write(sb.String())
		},
		Writef: func(level LogLevel, name string, format string, a ...any) {
			// TODO: add trace id
			// TODO: add key biz items
			writef(level.String()+" -- "+format, a...)
		},
	}
}

// LogKey 实现此接口，可以自定义当前 type 在日志中的 key。
// 如果需要同时定义 key 和 value，可实现 LogEntry 的 `LogValue() (string, any)` 方法。
type LogKey = interface {
	// LogKey 提供当前 type 在日志中的 key。
	// 如果不提供，默认是当前 type 的名字；
	// 如果当前 type 没有名字，则使用 "#i"（i 为日志参数下标） 作为 key。
	LogKey() string
}

// LogValue 实现此接口，可以自定义当前 type 在日志中的 value。
// 如果需要同时定义 key 和 value，可实现 LogEntry 的 `LogValue() (string, any)` 方法。
type LogValue = interface {
	// LogValue 提供当前 type 在日志中的 value。
	// 如果不提供，默认使用底层日志框架行为。
	//
	// 例如，可通过此功能实现脱敏能力：
	//
	//	type Foo struct {
	//		phone string
	//	}
	//	func (this *Foo) LogValue() any {
	//		return map[string]string {
	//			"phone": mask(this.phone)
	//		}
	//	}
	//
	// ℹ️ 如果需要同时定义日志的 key，可以实现 `LogValue() (string, any)`，见：LogEntry.LogValue
	LogValue() any
}

// LogEntry 是 LogKey 和 LogValue 的便捷声明，可通过只实现一个 `LogValue() (string, any)` 方法来减少代码量。
type LogEntry = interface {
	// LogValue 实现此方法同时返回 key 和 value 来自定义某个特定 type 在日志中的 key 和 value。
	//
	// 例如，可通过此功能实现脱敏能力：
	//
	//	type Foo struct {
	//		phone string
	//	}
	//	func (this *Foo) LogValue() (string, any) {
	//		return "foo", map[string]string {
	//			"phone": mask(this.phone)
	//		}
	//	}
	LogValue() (key string, value any)
}
