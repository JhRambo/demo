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

	// 定义正则表达式来匹配服务名、接口名、路由和请求类型
	serviceRegex := regexp.MustCompile(`service\s+(\w+)\s*{\s*rpc\s+(\w+)\s*\(\s*([^\)]*)\s*\)\s*returns\s*\(([^\)]*)\)\s*{[^}]*option`)
	routeRegex := regexp.MustCompile(`rpc\s+(\w+)\s*\(\s*([^\)]*)\s*\)\s*returns\s*\(([^\)]*)\)\s*{[^}]*option\s*\(\s*google\.api\.http\s*\)\s*=\s*\{\s*(get|post|delete|put):\s*["'](/?[^\s"']*)["'];?\s*`)

	ms := []map[string]string{}
	serviceName := ""
	// 1.获取service服务名
	serviceMatches := serviceRegex.FindAllSubmatch(data, -1)
	for _, v := range serviceMatches {
		serviceName = string(v[1])
		break
	}
	// 2.获取路由（uri）及请求类型（post|get|delete|put）请求、响应参数
	routerMatches := routeRegex.FindAllSubmatch(data, -1)
	for _, v := range routerMatches {
		m := map[string]string{}
		m["serviceName"] = serviceName
		m["rpcName"] = string(v[1])
		m["request"] = string(v[2])
		m["response"] = string(v[3])
		m["method"] = strings.ToUpper(string(v[4]))
		m["uri"] = string(v[5])
		ms = append(ms, m)
	}
	return ms
}
