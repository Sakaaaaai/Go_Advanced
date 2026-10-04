package main

import (
	"strings"
	"testing"
)

func TestDockerfileContent(t *testing.T) {
	compact := strings.ToLower(strings.ReplaceAll(DockerfileContent, " ", ""))
	compact = strings.ReplaceAll(compact, "\n", "")

	if !strings.Contains(compact, "fromgolang:1.20asbuilder") {
		t.Errorf("Missing builder setup")
	}

	if !strings.Contains(compact, "gobuild-oservermain.go") {
		t.Errorf("Missing go build command")
	}

	if !strings.Contains(compact, "fromalpine:latest") && !strings.Contains(compact, "fromalpine") {
		t.Errorf("Missing alpine runtime image")
	}

	if !strings.Contains(compact, "copy--from=builder/app/server") {
		t.Errorf("Missing copy from builder")
	}

	if !strings.Contains(compact, "cmd[\"/server\"]") && !strings.Contains(compact, "entrypoint[\"/server\"]") {
		t.Errorf("Missing cmd/entrypoint instruction")
	}
}
