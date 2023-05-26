package main

import (
	"log"
	"sync"
)

type Singleton struct {
	// ...
}

var instance *Singleton
var once sync.Once

// GetInstance returns the singleton instance.
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func main() {
	log.Println(GetInstance())
}
