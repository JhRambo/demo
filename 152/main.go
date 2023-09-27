package main

import "sync"

var (
	mutex *sync.Mutex
	wg    *sync.WaitGroup
)

func main() {
	fg()
	wg.Wait()
}

func fg() {
	wg.Add(4)
	go func() {
		for i := 0; i < 4; i++ {
			mutex.Lock()
			defer mutex.Unlock()
			defer wg.Done()
		}
	}()
}
