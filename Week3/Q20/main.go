package main

import (
	"fmt"
	"time"
)

func main() {
	// RabbitMQの代わりとなるメッセージキュー
	messageQueue := make(chan string, 10)

	// 1. Subscriber（購読者）を別ゴルーチンで立ち上げる
	___ func() {
		// 2. キューからメッセージを受け取り続ける
		for msg := ___ ___ {
			// 重い処理（メール送信など）のシミュレート
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Sending email to:", msg)
		}
		fmt.Println("Subscriber closed.")
	}()

	fmt.Println("API: User 1 registered.")
	// 3. Publisher（送信者）としてキューにメッセージを入れる
	___ <- "user1@example.com"
	
	fmt.Println("API: User 2 registered.")
	___ <- "user2@example.com"

	fmt.Println("API: Finished immediately. (Emails are sent in background)")

	// 4. メッセージの送信完了を伝えるためにチャネルを閉じる
	___(___)

	// Subscriberが処理を終えるのを少し待つ（本来はWaitGroupを使いますが今回はSleepで簡略化）
	time.Sleep(2 * time.Second)
}
