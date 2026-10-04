package main

import (
	"fmt"
	"net/http"
	// 1. pprofパッケージをブランクインポート（副作用としての初期化のみ）する
	_ "net/http/___"
	"time"
)

// 無駄にCPUを消費する重い処理
func heavyComputation() {
	var sum int
	for i := 0; i < 10000000; i++ {
		sum += i
	}
}

func main() {
	// 2. pprof 用のHTTPサーバーをゴルーチンで起動する
	go func() {
		fmt.Println("Starting pprof server on :6060")
		// ポート6060で起動、ハンドラはnilを指定
		if err := http.___(":___", ___); err != nil {
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
