package main

import (
	"strings"
	"testing"
)

func TestGraphQLSchema(t *testing.T) {
	compact := strings.ReplaceAll(GraphQLSchema, " ", "")
	compact = strings.ReplaceAll(compact, "\n", "")
	compact = strings.ReplaceAll(compact, "\r", "")

	// 必須フィールドが含まれているかチェック
	if !strings.Contains(compact, "typeUser{") {
		t.Errorf("Missing 'type User'")
	}
	if !strings.Contains(compact, "id:ID!") {
		t.Errorf("Missing 'id: ID!' in User")
	}
	if !strings.Contains(compact, "name:String!") {
		t.Errorf("Missing 'name: String!' in User")
	}
	if !strings.Contains(compact, "age:Int") && !strings.Contains(compact, "age:Int}") {
		// int! ではなく int であること
		if strings.Contains(compact, "age:Int!") {
			t.Errorf("age should not be required (!)")
		} else {
			t.Errorf("Missing 'age: Int' in User")
		}
	}

	if !strings.Contains(compact, "typeQuery{") {
		t.Errorf("Missing 'type Query'")
	}
	if !strings.Contains(compact, "user(id:ID!):User") {
		t.Errorf("Missing or incorrect 'user(id: ID!): User' in Query")
	}
}
