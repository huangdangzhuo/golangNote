package main

import (
	"log"
	"sync"
	"time"
)

func test() {
	var mut sync.Mutex
	maxSize := 10
	counter := 0

	// 排水口
	go func() {
		for {
			mut.Lock()
			if counter == maxSize {
				for i := 0; i < maxSize; i++ {
					counter--
					log.Printf(" OUTPUT counter = %d", counter)
				}
			}
			mut.Unlock()
			time.Sleep(1 * time.Second)
		}
	}()

	// 注水口
	for {
		mut.Lock()
		if counter == 0 {
			for i := 0; i < maxSize; i++ {
				counter++
				log.Printf(" INPUT counter = %d", counter)
			}
		}
		mut.Unlock()
		time.Sleep(1 * time.Second)
	}
}