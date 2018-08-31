package exceptions

import (
	"github.com/go-clog/clog"
	"github.com/wuleying/silver-xwlb/globals"
	"runtime"
)

// CheckError 错误检查
func CheckError(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		clog.Error(globals.ClogSkip, "%s:%d %s", file, line, err.Error())
	}
}