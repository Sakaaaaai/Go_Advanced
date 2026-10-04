//go:build ignore

package main

import (
	"fmt"
	"sync"
)

var (
	once sync.Once
	sm   sync.Map
)

func initialize() {
	once.Do(func() {
		fmt.Println("System initialized! (This should only appear once)")
	})
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			initialize()

			sm.Store(id, fmt.Sprintf("Data for ID %d", id))
		}(i)
	}

	wg.Wait()

	if val, ok := sm.Load(5); ok {
		fmt.Println("Result:", val.(string))
	} else {
		fmt.Println("Not found")
	}
}
