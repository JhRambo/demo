package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"golang.org/x/sync/errgroup"
)

func main() {
	NewRedis()
	fu1()
	time.Sleep(100 * time.Second)
}

var rdb *redis.Client
var ctx = context.Background()
var mutex sync.Mutex

func NewRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:36379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
func Lock(key string) error {
	mutex.Lock()
	defer mutex.Unlock()
	_, err := rdb.SetNX(key, 1, 1*time.Second).Result()
	if err != nil {
		log.Println(err.Error())
	}
	return err
}
func UnLock(key string) error {
	_, err := rdb.Del(key).Result()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return err
}
func Expire(key string) error {
	_, err := rdb.Expire(key, 1*time.Second).Result()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return err
}
func fu1() error {
	ch := make(chan bool)
	// 加锁
	err := Lock("lock_key")
	if err != nil {
		return err
	}
	//解锁
	defer func() {
		err = UnLock("lock_key")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("release redis lock success")
	}()

	g, _ := errgroup.WithContext(context.Background())
	//...主业务代码
	g.Go(func() error {
		time.Sleep(15 * time.Second)
		ch <- true
		return nil
	})

	//锁续期
	g.Go(func() error {
		ticker := time.NewTicker(time.Second * 1)
		for {
			select {
			// 任务还没执行完 每秒续期
			case <-ticker.C:
				Expire("lock_key")
				fmt.Println(time.Now())
				//收到执行完的 channel 就关闭time定时任务
			case <-ch:
				ticker.Stop()
				return nil
			}
		}
	})
	//等待信号量
	if err = g.Wait(); err != nil {
		return err
	}
	close(ch)
	return nil
}
