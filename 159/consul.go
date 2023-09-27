package consul

import (
	"fmt"
	"sync"

	consulapi "github.com/hashicorp/consul/api"
)

var (
	consulInstance  *ConsulCenter
	consulMutex     sync.Mutex
	clientConnCache sync.Map
)

type ConsulCenter struct {
	KV     *consulapi.KV
	Config string
}

func init() {
	consulInstance = NewConsul("192.168.10.103:38500", "alarm-robot", "123456")
}

func NewConsul(addr string, conf string, token string) *ConsulCenter {
	config := consulapi.DefaultConfig()
	config.Address = addr
	config.Token = token

	client, err := consulapi.NewClient(config)
	if err != nil {
		return nil
	}

	consulInstance := &ConsulCenter{KV: client.KV()}
	consulInstance.Config = conf
	return consulInstance
}

// 获取Consul实例
func GetConsulInstance() *ConsulCenter {
	consulMutex.Lock()
	defer consulMutex.Unlock()

	return consulInstance
}

// 监控consul key/value 变化
func (s *ConsulCenter) WatchConsulKeyChanges(configKey string) bool {
	// 查询键值对
	key, _, err := s.KV.Get(configKey, nil)
	if err != nil {
		fmt.Println(err)
		return false
	}
	// 处理键值对的变化
	if key != nil {
		v, ok := clientConnCache.Load(configKey)
		if ok {
			if v != string(key.Value) {
				clientConnCache.Store(configKey, string(key.Value))
				return true
			}
			return false
		}
		clientConnCache.Store(configKey, string(key.Value))
	}
	return false
}
