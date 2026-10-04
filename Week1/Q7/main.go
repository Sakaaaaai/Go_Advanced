package main

import (
	"fmt"
	"net/http"
	// 1. pprof をインポートする
	_ "net/http/___"
	"time"
)

// メモリリークを引き起こすグローバル変数
var leakySlice [][]byte

// 2. メモリリーク関数を完成させる
func memoryLeak() {
	for {
		// 1MBのバイトスライスを作成して追加する
		data := make([]byte, 1024*1024)
		// スライスに要素を追加するには append です
		leakySlice = ___(leakySlice, data)
		
		fmt.Printf("Current leakySlice size: %d MB\n", len(leakySlice))
		
		// 1秒に1回だけ実行する（一瞬でメモリを食い尽くさないように）
		time.___(___ * time.Second)
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
	// ここでブロックしてメモリリークを実行し続ける
	memoryLeak()
}
