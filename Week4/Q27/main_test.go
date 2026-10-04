package main

import (
	"strings"
	"testing"
)

func TestGitHubActionsYAML(t *testing.T) {
	compact := strings.ReplaceAll(GitHubActionsYaml, " ", "")
	compact = strings.ReplaceAll(compact, "\n", "")

	if !strings.Contains(compact, "branches:[\"main\"]") {
		t.Errorf("Missing target branch 'main'")
	}

	if !strings.Contains(compact, "runs-on:ubuntu-latest") {
		t.Errorf("Runs-on should be ubuntu-latest")
	}

	if !strings.Contains(compact, "run:gomoddownload") {
		t.Errorf("Missing 'go mod download' command")
	}

	if !strings.Contains(compact, "run:gotest./...") {
		t.Errorf("Missing 'go test ./...' command")
	}
}
