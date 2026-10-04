package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestCleanArchitecture(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "Test User") {
		t.Errorf("Architecture test failed, got: %s", output)
	}

	// 構造がインターフェースを介しているかチェック（コンパイルと実行が通れば大体OKですが念押し）
	repo := &mockUserRepository{}
	var _ UserRepository = repo // must implement the interface
	
	usecase := NewUserUsecase(repo)
	var _ UserUsecase = usecase // must implement the interface

	// GetUser の異常系
	_, err := usecase.GetUser("")
	if err == nil || !strings.Contains(err.Error(), "id is required") {
		t.Errorf("Expected 'id is required' error for empty ID")
	}
}
