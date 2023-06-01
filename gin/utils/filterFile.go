package utils

import (
	"bufio"
	"os"
	"regexp"
)

// 过滤文件的注释代码
func FilterFile(filePath string) string {
	outFile := "D:/code/demo/gin/doc/p.proto"

	//使用正则表达式匹配多行注释
	multilineCommentRegexp := regexp.MustCompile(`(?s)/\*.*?\*/`)
	//使用正则表达式匹配单行注释
	inlineCommentRegexp := regexp.MustCompile(`//.*`)

	// 打开文件并读取内容
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 创建输出文件
	out, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}
	defer out.Close()
	writer := bufio.NewWriter(out)

	// 逐行读取，并过滤掉注释和其他特定的代码
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// 处理多行注释
		line := multilineCommentRegexp.ReplaceAllString(scanner.Text(), "")
		// 处理单行注释
		line = inlineCommentRegexp.ReplaceAllString(line, "")
		// 处理其他特定的代码注释在这里添加你需要过滤的正则表达式
		// 然后将过滤后的内容写入输出文件中
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}
	// 刷新缓存
	err = writer.Flush()
	if err != nil {
		panic(err)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return outFile
}
