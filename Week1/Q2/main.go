package main

import (
	"fmt"
	"time"
)

// 1. worker関数
// jobs から値を取り出し、2倍にして results に送信します
func worker(id int, jobs <-chan int, results chan<- int) {
	// 2. jobs チャネルが閉じられるまでタスクを取り出し続ける
	for j := ___ ___ {
		fmt.Printf("Worker %d started  job %d\n", id, j)
		time.Sleep(time.Second) // タスクの処理にかかる時間をシミュレート
		fmt.Printf("Worker %d finished job %d\n", id, j)
		
		results <- j * ___
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 3. 3つのワーカー（ゴルーチン）を起動する
	for w := 1; w <= 3; w++ {
		go ___(___, ___, ___)
	}

	// 4. 全てのジョブを送信し、チャネルを閉じる
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// 5. 結果を受信して出力する
	for a := 1; a <= numJobs; a++ {
		<-results
	}
	
	fmt.Println("All jobs finished!")
}
