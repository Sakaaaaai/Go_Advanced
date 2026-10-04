package main

import (
	"context"
	"fmt"
)

// ========= モック環境 =========
type mockSpan struct { name string }
func (s *mockSpan) End() { fmt.Printf("[Trace] Span '%s' ended.\n", s.name) }
func (s *mockSpan) RecordError(err error) { fmt.Printf("[Trace] Error recorded on span '%s': %v\n", s.name, err) }

type mockTracer struct{}
func (t *mockTracer) Start(ctx context.Context, name string) (context.Context, *mockSpan) {
	fmt.Printf("[Trace] Span '%s' started.\n", name)
	return ctx, &mockSpan{name: name}
}
// =============================

func FetchData(ctx context.Context, tracer *mockTracer) error {
	// 1. 新しいSpan（トレース区間）を開始する
	_, span := tracer.___(___, "___")
	
	// 2. 処理が終わる時に必ずSpanを終了させる
	___ span.___()

	fmt.Println("Fetching data from DB...")
	
	// DBエラーが発生したと仮定
	err := fmt.Errorf("database timeout")
	
	// 3. エラーが起きたことをトレース（Span）に記録する
	if err != nil {
		span.___(err)
		return err
	}

	return nil
}

func main() {
	fmt.Println("=== OpenTelemetry Distributed Tracing (Simulation) ===")
	tracer := &mockTracer{}
	ctx := context.Background()

	err := FetchData(ctx, tracer)
	if err != nil {
		fmt.Println("Command returned error:", err)
	}
}
