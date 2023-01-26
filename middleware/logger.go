package middleware

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	_ "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var lg *zap.Logger

func InitLogger() (err error) {
	// 配置zapcore
	encoder := getEncoder()
	writeSyncer := getLogWriter()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte("debug"))
	
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	if true {
		// 进入开发模式，日志输出到终端
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(		// 多个输出
			zapcore.NewCore(encoder, writeSyncer, l),		// 往日志文件里面写
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),	// 终端输出
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, l)
	}

	
	// 实例化logger
	lg = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg)
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./log/test.log",
		MaxSize:    10,		// 日志文件最大大小（mb）
		MaxBackups: 5,      // 旧文件保存天数
		MaxAge:     30,		// 旧文件保存时间
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// zap日志中间件
func GinLogger(c *gin.Context) {
	start := time.Now()
	path := c.Request.URL.Path
	query := c.Request.URL.RawQuery
	// 挂起，待路由处理函数执行完后再往下执行
	c.Next()

	// 路由处理函数耗时
	cost := time.Since(start)
	// 设置日志输出信息
	lg.Info(path,
		zap.Int("status", c.Writer.Status()),
		zap.String("method", c.Request.Method),
		zap.String("path", path),
		zap.String("query", query),
		zap.Duration("cost", cost),
	)
}

// recover项目出现的panic，并用zap记录
// stack意味是否输出栈日志，栈日志很长，但可以轻易找到错误发生处
func GinRecover(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 判断客户端是否断开连接
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				// 获取请求
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				// 若链接断开就不会有状态码返回
				if brokenPipe {
					lg.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					c.Error(err.(error))
					// 不再继续处理
					c.Abort()
					return
				}

				if stack {
					lg.Error("[Recovery from panic]",
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					lg.Error("[Recovery from panic]",
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}

}
