package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	// 读取proto文件
	content, err := ioutil.ReadFile("D:/code/demo/gin/proto/hello.proto")
	if err != nil {
		panic(err)
	}

	// 匹配配置了 HTTP 方法的服务和接口
	regex := regexp.MustCompile(`(?s)\s*(service|rpc)\s+(\S+)\s*\{\s*\n((?:\s*//.*\n)*)\s*(.*)\n\s*\}`)
	matches := regex.FindAllSubmatch(content, -1)
	for _, match := range matches {
		fmt.Println(string(match[4]))

		// 获取服务名或接口名
		name := string(match[2])
		// 获取代码块中的内容
		block := string(match[4])

		// 匹配 HTTP 方法
		httpRegex := regexp.MustCompile(`(?s)\n\s*//\s*google\.api\.http:\s*\{(.+?)\}\s*\n`)
		httpMatches := httpRegex.FindStringSubmatch(block)

		// 如果匹配成功，说明是配置了 HTTP 方法的服务或接口
		if len(httpMatches) > 0 {
			options := httpMatches[1]
			// 匹配路由注释中的 HTTP 方法和路径
			httpMethodRegex := regexp.MustCompile(`^\s*([A-Z]+)\s+(.+)$`)
			httpMethodMatches := httpMethodRegex.FindStringSubmatch(options)

			// 如果匹配成功，并且是 POST、GET、DELETE 或 PUT 方法，则输出服务名或接口名和路径
			if len(httpMethodMatches) > 0 && (httpMethodMatches[1] == "POST" || httpMethodMatches[1] == "GET" || httpMethodMatches[1] == "DELETE" || httpMethodMatches[1] == "PUT") {
				fmt.Println("Name:", name)
				fmt.Println("Path:", httpMethodMatches[2])
			}
		}
	}
}
