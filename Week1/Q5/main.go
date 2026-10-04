package main

import "fmt"

// 1. スタックに収まる関数
// 単なる値を返すため、関数終了時に変数は関数スタックと一緒に破棄される
func createOnStack() ___ {
	x := 10
	return ___
}

// 2. ヒープにエスケープする関数
// ポインタを返すため、関数終了後も変数の中身が残る必要がある
func createOnHeap() ___ {
	y := 20
	return ___
}

// （参考）インターフェースに入れる場合もエスケープすることが多いです
func escapeByInterface() {
	z := 30
	fmt.Println(z) // fmt.Println は引数として `any` (interface{}) を受け取るためエスケープが発生しやすい
}

func main() {
	val := createOnStack()
	ptr := createOnHeap()

	fmt.Printf("Stack value: %d\n", val)
	fmt.Printf("Heap value: %d\n", *ptr)
	escapeByInterface()
}
// `go build -gcflags="-m" main.go` を実行してエスケープ解析の結果を確認してみよう
