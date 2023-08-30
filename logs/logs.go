package logs

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"time"
)

const (
	LOG_ERROR = iota
	LOG_WARING
	LOG_INFO
	LOG_DEBUG
)

var log *mylog

/*
 * 初始化
 */
func init() {
	log = newMylog()
	// 执行日志清理函数
	go clean30DayLog()
}

// 清除30天的日志
func clean30DayLog() {
	// 启动时暂停10秒
	time.Sleep(time.Second * 10)
	// 暂停改为tikcer机制
	tick := time.NewTicker(time.Hour * 12)
	defer tick.Stop()
	for {
		Clean(30)
		// 暂停一小时
		<-tick.C
	}
}

func Init(dir string, file string, level int, savefile bool) {
	log.setDir(dir)
	log.setFile(file)
	log.setLevel(level)
	log.setSavefile(savefile)
}

// 打印错误日志
func Error(err ...interface{}) {
	log.write(LOG_ERROR, fmt.Sprint(err...))
}

func Errorf(format string, a ...interface{}) {
	log.write(LOG_ERROR, fmt.Sprintf(format, a...))
}

// 打印警告日志
func Waring(war ...interface{}) {
	log.write(LOG_WARING, fmt.Sprint(war...))
}

func Waringf(format string, a ...interface{}) {
	log.write(LOG_WARING, fmt.Sprintf(format, a...))
}

// 打印信息日志
func Info(info ...interface{}) {
	log.write(LOG_INFO, fmt.Sprint(info...))
}

func Infof(format string, a ...interface{}) {
	log.write(LOG_INFO, fmt.Sprintf(format, a...))
}

// 打印调试日志
func Debug(deb ...interface{}) {
	log.write(LOG_DEBUG, fmt.Sprint(deb...))
}

func Debugf(format string, a ...interface{}) {
	log.write(LOG_DEBUG, fmt.Sprintf(format, a...))
}

func SetLogFile(logFile bool) {
	log.logFile = logFile
}

func SetLogDate(logDate bool) {
	log.logDate = logDate
}

func SetLogTerminal(logTerminal bool) {
	log.logTerminal = logTerminal
}

// 跟踪错误日志
func Trace() {
	err := recover()
	if nil == err {
		return
	}

	TraceError(err)
}

func TraceError(err interface{}) {
	// 记录开始时间
	var strTime string
	t := time.Now()
	if log.logDate {
		strTime = fmt.Sprintf("[%04d-%02d-%02d %02d:%02d:%02d]",
			t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	} else {
		strTime = fmt.Sprintf("[%02d:%02d:%02d]",
			t.Hour(), t.Minute(), t.Second())
	}

	buf := &bytes.Buffer{}
	// 第一行
	buf.WriteString(strTime)
	buf.WriteString(" [E] ")
	buf.WriteString(fmt.Sprint(err))
	buf.WriteString("\n")

	// 输出路径
	for i := 3; i < 20; i++ {
		pc, _, line, ok := runtime.Caller(i)
		if false == ok {
			break
		}

		p := runtime.FuncForPC(pc)
		// 每个栈输出一行
		buf.WriteString("    ")
		buf.WriteString(p.Name())
		buf.WriteString(": ")
		buf.WriteString(fmt.Sprint(line))
		buf.WriteString("\n")
	}

	// 如果保存到文件，同时也要输出到控制台
	str := string(buf.Bytes())
	// 输出到控制台
	fmt.Print(str)
	// 输出到日志文件
	if log.savefile {
		log.log <- str
	}
}

/*
 * 日志执行函数
 */
type mylog struct {
	log      chan string // 日志chan
	dir      string      // 日志存放目录
	file     string      // 日志文件名
	savefile bool        // 是否保存到文件
	level    int         // 日志级别

	logFile     bool
	logDate     bool
	logTerminal bool
}

func newMylog() *mylog {
	log := &mylog{}

	log.log = make(chan string, 4096)
	log.dir = "/opt/app/log"
	log.file = "out"
	log.savefile = false
	log.logFile = true
	log.logDate = true

	go log.run()
	return log
}

func (l *mylog) setDir(dir string) {
	l.dir = dir
}

func (l *mylog) setFile(file string) {
	l.file = file
}

func (l *mylog) setSavefile(b bool) {
	l.savefile = b
}

func (l *mylog) setLevel(level int) {
	l.level = level
}

func (l *mylog) getLevelString(level int) string {
	switch level {
	case LOG_ERROR:
		return "E"
	case LOG_WARING:
		return "W"
	case LOG_INFO:
		return "I"
	case LOG_DEBUG:
		return "D"
	}

	return "U"
}

func (l *mylog) write(level int, str string) {
	// 判断级别
	if level > l.level {
		return
	}

	var strTime string
	t := time.Now()
	if l.logDate {
		strTime = fmt.Sprintf("[%04d-%02d-%02d %02d:%02d:%02d]",
			t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	} else {
		strTime = fmt.Sprintf("[%02d:%02d:%02d]",
			t.Hour(), t.Minute(), t.Second())
	}

	// 输出日志
	pc, _, line, _ := runtime.Caller(2)
	p := runtime.FuncForPC(pc)
	if l.logFile {
		// 输出带文件信息
		str = fmt.Sprintf("%s [%s] %s(%d): %s\n", strTime,
			l.getLevelString(level), p.Name(), line, str)
	} else {
		// 输出不带文件信息
		str = fmt.Sprintf("%s [%s] %s(%d): %s\n", strTime,
			l.getLevelString(level), path.Base(p.Name()), line, str)
	}

	// 输出到控制台
	if false == l.savefile {
		fmt.Print(str)
		return
	}

	// 如果保存到文件，同时也要输出到控制台
	if l.logTerminal {
		fmt.Print(str)
	}

	// 输出到文件, 防止写入失败导致程序阻塞
	select {
	case l.log <- str:
	default:
		return
	}
}

func (l *mylog) run() {
	for {
		str := <-l.log

		// 判断文件夹是否存在
		_, err := os.Stat(l.dir)
		if nil != err {
			os.MkdirAll(l.dir, os.ModePerm)
		}

		// 获取时间
		t := time.Now()
		path := fmt.Sprintf("%s/%s-%04d-%02d-%02d.log", l.dir, l.file,
			t.Year(), t.Month(), t.Day())
		// 判断文件大小
		state, err := os.Stat(path)
		// 如果文件大于8M,就截断文件，重新生成新的文件
		if nil == err && state.Size() > 2*1024 {
			// 重新写入新的文件
			str = fmt.Sprintf(">>>>>>>>trunc<<<<<<<<: %02d:%02d:%02d.%d\n",
				t.Hour(), t.Minute(), t.Second(), t.Nanosecond()/1000000) + str
			ioutil.WriteFile(path, []byte(str), 0644)
		} else {
			// 追加到文件
			fp, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
			if nil == err {
				fp.WriteString(str)
				fp.Close()
			}
		}
	}
}
