package logs

import (
	"fmt"
	"runtime"
)

type GWLog struct{}

var gwlog *GWLog

func GetErrorLocation(err ...interface{}) string {
	return gwlog.GetErrorLocation(fmt.Sprint(err...))
}

// 记录错误发生的位置
func (l *GWLog) GetErrorLocation(str string) string {
	pc, _, line, _ := runtime.Caller(2)
	p := runtime.FuncForPC(pc)
	return fmt.Sprintf("%s(%d): %s", p.Name(), line, str)
}
