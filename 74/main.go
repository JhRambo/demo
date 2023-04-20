package main
 
import (
	"context"
	"fmt"
	"net"
	"time"
 
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)
 
funcUnaryTest() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler)(resp interface{}, err error) {
		start := time.Now()
		// 请求日期
		requestDate := start.Format(time.RFC3339)
		var res interface{}
		deferfunc() {
			// metadata
			md, _ := metadata.FromIncomingContext(ctx)
			// User Agent 和 host
			ua, host := extractFromMD(md)
			// 请求IP
			clientIp := getPeerAddr(ctx)
			// 请求耗时
			delay := time.Since(start).Milliseconds()
			// 请求调用的rpc方法 i.e., /package.service/method.
			fullMethod := info.FullMethod
			// 请求内容
			requestBody := req
			// 响应的状态码
			responseStatus := int(status.Code(err))
			// 响应数据
			responseBody := res
		}()
		res, err = handler(ctx, req)
		return res, err
	}
}
 
funcextractFromMD(md metadata.MD)(ua string, host string) {
	if v, ok := md["x-forwarded-user-agent"]; ok {
		ua = fmt.Sprintf("%v", v)
	} else {
		ua = fmt.Sprintf("%v", md["user-agent"])
	}
	if v, ok := md[":authority"]; ok && len(v) > 0 {
		host = fmt.Sprintf("%v", v[0])
	}
	return ua, host
}
 
funcgetPeerAddr(ctx context.Context)string {
	var addr string
	if pr, ok := peer.FromContext(ctx); ok {
		if tcpAddr, ok := pr.Addr.(*net.TCPAddr); ok {
			addr = tcpAddr.IP.String()
		} else {
			addr = pr.Addr.String()
		}
	}
	return addr
}