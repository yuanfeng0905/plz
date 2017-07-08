package app

import (
	"github.com/v2pro/plz/logging"
	"os"
)

func Run(f func() int, kv ...interface{}) {
	logging.LoggerOf("metric", "counter", "begin", "app").
		Info("app begin", kv...)
	defer func() {
		recovered := recover()
		if recovered != nil {
			code := -1
			for _, handle := range AfterPanic {
				code = handle(recovered, kv)
			}
			for _, handle := range BeforeFinish {
				handle(kv)
			}
			os.Exit(code)
			return
		}
	}()
	code := f()
	for _, handle := range BeforeFinish {
		handle(kv)
	}
	os.Exit(code)
}

var AfterPanic = []func(recovered interface{}, kv []interface{}) int{
	func(recovered interface{}, kv []interface{}) int {
		logging.LoggerOf("metric", "counter", "panic", "app").
			Error("app panic", append(kv, "recovered", recovered)...)
		return 1
	},
}

var BeforeFinish = []func(kv []interface{}){
	func(kv []interface{}) {
		logging.LoggerOf("metric", "counter", "finish", "app").
			Info("app finish", kv...)
	},
}