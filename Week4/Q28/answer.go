//go:build ignore

package main

import (
	"context"
	"fmt"
)

type mockSpan struct { name string }
func (s *mockSpan) End() { fmt.Printf("[Trace] Span '%s' ended.\n", s.name) }
func (s *mockSpan) RecordError(err error) { fmt.Printf("[Trace] Error recorded on span '%s': %v\n", s.name, err) }

type mockTracer struct{}
func (t *mockTracer) Start(ctx context.Context, name string) (context.Context, *mockSpan) {
	fmt.Printf("[Trace] Span '%s' started.\n", name)
	return ctx, &mockSpan{name: name}
}

func FetchData(ctx context.Context, tracer *mockTracer) error {
	_, span := tracer.Start(ctx, "FetchDatabase")
	
	defer span.End()

	fmt.Println("Fetching data from DB...")
	
	err := fmt.Errorf("database timeout")
	
	if err != nil {
		span.RecordError(err)
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
