package logs

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// 清理day天之前的日志
func Clean(day int) {
	// 致少要保持3天的日志
	if day < 3 {
		return
	}

	// 读取所有日志文件
	files, err := ioutil.ReadDir(log.dir)
	if nil != err {
		Error(err)
		return
	}

	// 枚举所在日志文件
	for _, file := range files {
		// 判断如果为目录就跳过
		if file.IsDir() {
			continue
		}

		// 判断如果后缀不是.log就跳过
		if false == strings.HasSuffix(file.Name(), ".log") {
			continue
		}

		// 获取名字的时间差
		diff := compareNameTime(file.Name())
		// 如果时间差小于指定天数，就退出
		if diff < day {
			continue
		}

		// 删除文件
		Info("auto remove:", file.Name())
		// 删除文件
		os.Remove(filepath.Join(log.dir, file.Name()))
	}
}

// 获取文件的时间差
func compareNameTime(name string) int {
	// 去掉.log
	name = name[:len(name)-4]
	// 解析
	sp := strings.Split(name, "-")
	// 没有分隔，就认为不是我们的日志文件
	if len(sp) < 4 {
		return -1
	}

	if len(sp) > 4 {
		sp = sp[len(sp)-4:]
	}

	// 解析年份
	year, _ := strconv.Atoi(sp[1])
	month, _ := strconv.Atoi(sp[2])
	day, _ := strconv.Atoi(sp[3])

	old := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local).Unix()
	now := time.Now().Unix()

	return int((now - old) / (3600 * 24))
}
