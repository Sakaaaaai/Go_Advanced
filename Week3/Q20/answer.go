//go:build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	messageQueue := make(chan string, 10)

	go func() {
		for msg := range messageQueue {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Sending email to:", msg)
		}
		fmt.Println("Subscriber closed.")
	}()

	fmt.Println("API: User 1 registered.")
	messageQueue <- "user1@example.com"
	
	fmt.Println("API: User 2 registered.")
	messageQueue <- "user2@example.com"

	fmt.Println("API: Finished immediately. (Emails are sent in background)")

	close(messageQueue)

	time.Sleep(2 * time.Second)
}
