package middlewares

import (
	"demo/docker/config"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var ipLock sync.Map

const (
	_TIME  = "_time"
	_COUNT = "_count"
)

func IPRate(ctx *gin.Context) {
	ip := ctx.ClientIP()
	currentTime := time.Now()

	// 检查是否超过请求限制
	if !checkIpLimit(ip, currentTime) {
		resp := &config.GWResponse{
			Code:    -10005,
			Message: "系统繁忙",
		}
		ctx.AbortWithStatusJSON(http.StatusTooManyRequests, resp)
		return
	}
	ctx.Next()
}

func checkIpLimit(ip string, currentTime time.Time) bool {
	// 获取该IP的锁
	lock, _ := ipLock.LoadOrStore(ip, &sync.Mutex{})
	mutex := lock.(*sync.Mutex)

	// 对该IP的请求进行加锁，保证同一时间只能有一个请求处理
	mutex.Lock()
	defer mutex.Unlock()

	// 检查上次请求的时间
	lastRequestTime, exists := ipLock.Load(fmt.Sprintf("%v%v", ip, _TIME))
	if exists {
		if lastRequestTime.(time.Time).Add(time.Minute).After(currentTime) {
			// 如果距离上次请求不足1分钟，则判断是否超过请求限制
			count, _ := ipLock.Load(fmt.Sprintf("%v%v", ip, _COUNT))
			if count.(int) >= config.IP_RATE_MINUTE {
				return false
			}

			// 更新请求次数
			ipLock.Store(fmt.Sprintf("%v%v", ip, _COUNT), count.(int)+1)
			return true
		}
	}

	// 如果不存在上次请求的时间或者距离上次请求超过1分钟，则重置计数器
	ipLock.Store(fmt.Sprintf("%v%v", ip, _TIME), currentTime)
	ipLock.Store(fmt.Sprintf("%v%v", ip, _COUNT), 1)

	return true
}
