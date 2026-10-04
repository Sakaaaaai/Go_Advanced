//go:build ignore

package main

import (
	"fmt"
	"time"
)

type LogRequest struct{}
type LogResponse struct{ Message string }

type LogService_DownloadLogsServer interface {
	Send(*LogResponse) error
}

type mockStream struct{}
func (s *mockStream) Send(res *LogResponse) error {
	fmt.Printf("[Stream Sending...] %s\n", res.Message)
	return nil
}

type server struct{}

func (s *server) DownloadLogs(req *LogRequest, stream LogService_DownloadLogsServer) error {
	logs := []string{
		"Log 1: Server started",
		"Log 2: Processing data",
		"Log 3: Error encountered",
		"Log 4: Shutting down",
	}

	fmt.Println("Starting log stream...")

	for _, l := range logs {
		time.Sleep(100 * time.Millisecond)

		err := stream.Send(&LogResponse{Message: l})
		if err != nil {
			return err
		}
	}

	fmt.Println("Stream finished successfully.")

	return nil
}

func main() {
	s := &server{}
	stream := &mockStream{}
	
	err := s.DownloadLogs(&LogRequest{}, stream)
	if err != nil {
		fmt.Println("Stream Error:", err)
	}
}
