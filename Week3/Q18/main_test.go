package main

import (
	"bytes"
	"context"
	"io"
	"os"
	"strings"
	"testing"
)

func TestInterceptor(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	dummyHandler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return "Success", nil
	}
	
	info := &UnaryServerInfo{FullMethod: "/test.Service/Method"}
	loggingInterceptor(context.Background(), nil, info, dummyHandler)

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "[Global Interceptor] Method called: /test.Service/Method") {
		t.Errorf("Pre-handler interceptor log is missing or incorrect. Got:\n%s", output)
	}

	if !strings.Contains(output, "[Global Interceptor] Method execution finished.") {
		t.Errorf("Post-handler interceptor log is missing or incorrect. Got:\n%s", output)
	}
}
