package main

import "fmt"

// 1. generator関数: 可変長引数で整数を受け取り、チャネルに一つずつ送信する
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- ___
		}
		// データを全て送信し終えたら必ずチャネルを閉じる
		___(out)
	}()
	return out
}

// 2. squarer関数: inチャネルから受け取った値を2乗して、outチャネルに送信する
// 引数の in は受信専用チャネル `<-chan int` にします
func squarer(in <-chan ___) <-chan int {
	out := make(chan int)
	go func() {
		// inチャネルが閉じられるまで受信し続ける
		for n := ___ {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	// 3. パイプラインを構築する
	// generator で 2, 3, 4 のストリームを作る
	in := ___(2, 3, 4)

	// squarer に渡して2乗のストリームにする
	out := ___(in)

	// 出力チャネルから値を取り出して表示する
	for result := range out {
		fmt.Println(result)
	}
}
