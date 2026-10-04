//go:build ignore

package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func heavyComputation() {
	var sum int
	for i := 0; i < 10000000; i++ {
		sum += i
	}
}

func main() {
	go func() {
		fmt.Println("Starting pprof server on :6060")
		if err := http.ListenAndServe(":6060", nil); err != nil {
			fmt.Println("pprof server error:", err)
		}
	}()

	fmt.Println("Main application running... (Wait 5 seconds to simulate work)")
	
	for i := 0; i < 5; i++ {
		heavyComputation()
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Done")
}
