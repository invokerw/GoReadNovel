package logger

import (
	"github.com/op/go-logging"
	"os"
	"sync"
)

var alog *logging.Logger
var mu sync.Mutex

// Example format string. Everything except the message has a custom color
// which is dependent on the log level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli second.
var format = logging.MustStringFormatter(
	`%{color}%{time:2006-01-02T15:04:05} %{shortfile} %{shortfunc}  ▶   %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

// Password is just an example type implementing the Redactor interface. Any
// time this is logged, the Redacted() function will be called.
type Password string

func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}

func ALogger() *logging.Logger {
	if alog == nil {
		alog = InitAlog()
		return alog
	} else {
		return alog
	}
}
func InitAlog() *logging.Logger {
	mu.Lock()

	alog = logging.MustGetLogger("read")

	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	// For messages written to backend2 we want to add some additional
	// information to the output, including the used log level and the name of
	// the function.
	backend2Formatter := logging.NewBackendFormatter(backend2, format)

	// Only errors and more severe messages should be sent to backend1
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.ERROR, "")

	// Set the backends to be used.
	logging.SetBackend(backend1Leveled, backend2Formatter)
	mu.Unlock()

	return alog
}

/*  测试例子
func main() {
	// For demo purposes, create two backend for os.Stderr.
	ALogger().Error("errrrrrr")
	alog.Debugf("debug %s", Password("secret"))
	alog.Info("info")
	alog.Notice("notice")
	alog.Warning("warning")
	alog.Error("err")
	alog.Critical("crit")
}
*/
