package main

import (
	"fmt"
	"sync"
	"time"
)

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func worker(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			// 重い処理をシミュレート
			time.Sleep(100 * time.Millisecond)
			out <- n * n
		}
		close(out)
	}()
	return out
}

// 2. 複数のチャネルを1つにまとめる merge 関数
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// 一つのチャネルの中身を out に移す関数
	output := func(c <-chan ___) {
		for n := ___ {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range ___ {
		go output(___)
	}

	// WaitGroupで待機し、全て終わったら out を閉じる
	go func() {
		wg.___()
		close(___)
	}()

	return out
}

func main() {
	in := generator(1, 2, 3, 4, 5)

	// 1. Fan-out: 1つの入力チャネルから2つのワーカーにタスクを分散する
	out1 := worker(___)
	out2 := worker(___)

	// 2. Fan-in: 2つの出力チャネルを1つにまとめる
	for result := range merge(___, ___) {
		fmt.Println(result)
	}
}
