//go:build ignore

package main

import "fmt"

const GitHubActionsYaml = `
name: Go CI

on:
  push:
    branches: [ "main" ]
  pull_request:
    branches: [ "main" ]

jobs:
  build-and-test:
    runs-on: ubuntu-latest
    
    steps:
      - name: Checkout code
        uses: actions/checkout@v3

      - name: Setup Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.20'
      
      - name: Download dependencies
        run: go mod download

      - name: Run tests
        run: go test ./...
`

func main() {
	fmt.Println("=== GitHub Actions (.github/workflows/test.yml) ===")
	fmt.Print(GitHubActionsYaml)
	fmt.Println("===================================================")
}
