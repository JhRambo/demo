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
	serviceRegex := regexp.MustCompile(`service\s+(\w+)\s+{`)
	rpcRegex := regexp.MustCompile(`rpc\s+(\w+)\s*\(\s*([^\)]*)\s*\)\s+returns\s+\(([^\)]*)\)`)
	// routeRegex := regexp.MustCompile(`rpc\s+(\w+)\s*\(\s*(\w+)Request\s*\)\s*returns\s*\(\s*(\w+)Response\s*\)\s*{[^}]*option\s*\(\s*google\.api\.http\s*\)\s*=\s*\{\s*(get|post|delete|put):\s*["'](/?[^\s"']*)["'];?\s*`)
	routeRegex := regexp.MustCompile(`rpc\s+(\w+)\s*\(\s*([^\)]*)\s*\)\s+returns\s+\(([^\)]*)\)\s*{[^}]*option\s*\(\s*google\.api\.http\s*\)\s*=\s*\{\s*(get|post|delete|put):\s*["'](/?[^\s"']*)["'];?\s*`)

	ms := []map[string]string{}
	serviceName := ""
	// 1.获取service服务名
	serviceMatches := serviceRegex.FindAllSubmatch(data, -1)
	services := make([]string, len(serviceMatches))
	for i, m := range serviceMatches {
		services[i] = string(m[1])
	}
	for i := 0; i < len(services); i++ {
		regex := regexp.MustCompile(`Http`)
		if regex.MatchString(services[i]) {
			serviceName = services[i]
			break
		}
	}
	// 2.获取RPC接口名
	rpcMatches := rpcRegex.FindAllSubmatch(data, -1)
	// fmt.Println(len(rpcMatches)) //长度为rpc接口的数量，包含http和grpc
	for _, v := range rpcMatches {
		regex := regexp.MustCompile(`DBRequest`)
		if regex.MatchString(string(v[0])) {
			continue
		}
		m := map[string]string{}
		m["serviceName"] = serviceName
		m["rpcName"] = string(v[1])
		m["request"] = string(v[2])
		m["response"] = string(v[3])
		ms = append(ms, m)
	}
	// 3.获取路由（uri）及请求类型（post|get|delete|put）
	routerMatches := routeRegex.FindAllSubmatch(data, -1)
	for i := 0; i < len(ms); i++ {
		ms[i]["method"] = strings.ToUpper(string(routerMatches[i][4]))
		ms[i]["uri"] = string(routerMatches[i][5])
	}
	return ms
}
