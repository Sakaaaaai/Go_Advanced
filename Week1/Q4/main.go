package main

import (
	"fmt"
	"sync"
)

var (
	// 1. sync.Once の変数を定義
	once sync.___
	
	// 2. sync.Map の変数を定義
	sm sync.___
)

func initialize() {
	// 1回だけ実行されるように once.Do で包む
	once.___(___() {
		fmt.Println("System initialized! (This should only appear once)")
	})
}

func main() {
	var wg sync.WaitGroup

	// 10個のゴルーチンから同時にアクセス
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			// 初期化関数を呼ぶ（1回しか実行されないはず）
			initialize()

			// 3. sync.Map にデータを保存 (Store)
			sm.___(id, fmt.Sprintf("Data for ID %d", id))
		}(i)
	}

	wg.Wait()

	// 4. sync.Map からデータを取り出す (Load)
	if val, ok := sm.___(___); ok {
		// valの型は any なので、string型として使う場合は型アサーションが必要
		fmt.Println("Result:", val.(string))
	} else {
		fmt.Println("Not found")
	}
}
