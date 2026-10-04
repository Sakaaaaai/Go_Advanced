package main

import (
	"fmt"
	"time"
)

// ========= モック環境 =========
type LogRequest struct{}
type LogResponse struct{ Message string }

// 自動生成されるストリームインターフェースのモック
type LogService_DownloadLogsServer interface {
	Send(*LogResponse) error
}

type mockStream struct{}
func (s *mockStream) Send(res *LogResponse) error {
	fmt.Printf("[Stream Sending...] %s\n", res.Message)
	return nil
}
// =============================

type server struct{}

// サーバーストリーミングの実装
func (s *server) DownloadLogs(req *LogRequest, stream LogService_DownloadLogsServer) error {
	logs := []string{
		"Log 1: Server started",
		"Log 2: Processing data",
		"Log 3: Error encountered",
		"Log 4: Shutting down",
	}

	fmt.Println("Starting log stream...")

	// 1. ログを一つずつ Send でクライアントに送信するループ
	for _, l := range logs {
		// 少し時間がかかる処理をシミュレート
		time.Sleep(100 * time.Millisecond)

		// 2. stream.Send を使ってメッセージを一つずつクライアントへ送る
		err := stream.___(&___{Message: ___})
		if err != nil {
			return err
		}
	}

	fmt.Println("Stream finished successfully.")

	// 3. 全て送り終えたら nil を返してストリームを閉じる
	return ___
}

func main() {
	s := &server{}
	stream := &mockStream{}
	
	err := s.DownloadLogs(&LogRequest{}, stream)
	if err != nil {
		fmt.Println("Stream Error:", err)
	}
}
