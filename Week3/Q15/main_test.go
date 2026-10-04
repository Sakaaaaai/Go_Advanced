package main

import (
	"strings"
	"testing"
)

func TestProtoFileContent(t *testing.T) {
	content := strings.ToLower(ProtoFileContent)

	if !strings.Contains(content, "syntax = \"proto3\"") {
		t.Errorf("Missing syntax = \"proto3\"")
	}

	if !strings.Contains(content, "message user") {
		t.Errorf("Missing message User definition")
	}

	if !strings.Contains(content, "int32 id = 1") {
		t.Errorf("Missing int32 id = 1 field")
	}

	if !strings.Contains(content, "string name = 2") {
		t.Errorf("Missing string name = 2 field")
	}

	if !strings.Contains(content, "service userservice") {
		t.Errorf("Missing service UserService definition")
	}

	// Remove whitespaces to check rpc definition flexibly
	compact := strings.ReplaceAll(content, " ", "")
	if !strings.Contains(compact, "rpcgetuser(userrequest)returns(user)") {
		t.Errorf("Missing or incorrect rpc GetUser definition")
	}
}
