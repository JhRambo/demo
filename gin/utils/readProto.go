package utils

import (
	"io/ioutil"
	"regexp"
	"strings"
)

func ReadProto(filePath string) []map[string]string {
	// 读取proto文件的内容
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	// 正则匹配go_package包的路径
	packRegex := regexp.MustCompile(`option\s+go_package\s*=\s*"(.+?)"\s*;`)
	// 定义正则表达式来匹配服务名、接口名、路由和请求类型
	serviceRegex := regexp.MustCompile(`service\s+(\w+)\s*{\s*rpc\s+(\w+)\s*\(\s*([^\)]*)\s*\)\s*returns\s*\(([^\)]*)\)\s*{[^}]*option`)
	routeRegex := regexp.MustCompile(`rpc\s+(\w+)\s*\(\s*([^\)]*)\s*\)\s*returns\s*\(([^\)]*)\)\s*{[^}]*option\s*\(\s*google\.api\.http\s*\)\s*=\s*\{\s*(get|post|delete|put):\s*["'](/?[^\s"']*)["'];?\s*`)
	// 正则匹配stream binary
	streamRegex := regexp.MustCompile(`stream\s+(\w+)`)

	packName := ""
	// 1.获取包名
	packMatches := packRegex.FindAllSubmatch(data, -1)
	for _, v := range packMatches {
		packNames := strings.Split(string(v[1]), "/")
		packName = packNames[len(packNames)-1]
		break
	}

	ms := []map[string]string{}
	serviceName := ""
	// 2.获取service服务名
	serviceMatches := serviceRegex.FindAllSubmatch(data, -1)
	for _, v := range serviceMatches {
		serviceName = string(v[1])
		break
	}
	// 3.获取路由（uri）及请求类型（post|get|delete|put）请求、响应参数
	routerMatches := routeRegex.FindAllSubmatch(data, -1)
	for _, v := range routerMatches {
		m := map[string]string{}
		m["packName"] = packName
		m["serviceName"] = serviceName
		m["rpcName"] = string(v[1])
		m["request"] = strings.TrimSpace(string(v[2]))
		// stream binary二进制流
		// 获取二进制流参数名
		if strings.Contains(m["request"], "stream ") {
			m["streamBinary"] = "stream"
			requestName := ""
			requestMatches := streamRegex.FindStringSubmatch(m["request"])
			if len(requestMatches) > 1 {
				requestName = requestMatches[1]
			}
			m["request"] = requestName
			bytesName := ""
			bytesNameRegex := regexp.MustCompile(`message\s+` + requestName + `\s*{\s*bytes\s+(\w+)\s*=`)
			bytesNameMatches := bytesNameRegex.FindAllSubmatch(data, -1)
			for _, v := range bytesNameMatches {
				bytesName = strings.TrimSpace(string(v[1]))
				break
			}
			m["bytesName"] = strings.Title(bytesName)
		}
		m["response"] = string(v[3])
		m["method"] = strings.ToUpper(string(v[4]))
		m["uri"] = string(v[5])
		ms = append(ms, m)
	}
	return ms
}
