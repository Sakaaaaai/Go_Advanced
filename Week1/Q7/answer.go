//go:build ignore

package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var leakySlice [][]byte

func memoryLeak() {
	for {
		data := make([]byte, 1024*1024)
		leakySlice = append(leakySlice, data)
		
		fmt.Printf("Current leakySlice size: %d MB\n", len(leakySlice))
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go func() {
		fmt.Println("Starting pprof server on :6060")
		if err := http.ListenAndServe(":6060", nil); err != nil {
			fmt.Println("pprof server error:", err)
		}
	}()

	fmt.Println("Starting memory leak simulation...")
	memoryLeak()
}
