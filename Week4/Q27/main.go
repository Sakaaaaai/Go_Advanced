package main

import "fmt"

const GitHubActionsYaml = `
name: Go CI

# 1. トリガーの設定
on:
  push:
    branches: [ "___" ]
  pull_request:
    branches: [ "___" ]

jobs:
  build-and-test:
    # 2. 実行環境
    runs-on: ___
    
    steps:
      - name: Checkout code
        uses: actions/checkout@v3

      - name: Setup Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.20'
      
      # 3. 依存関係のダウンロード
      - name: Download dependencies
        run: ___ mod ___

      # 4. テストの実行
      - name: Run tests
        run: ___ test ___
`

func main() {
	fmt.Println("=== GitHub Actions (.github/workflows/test.yml) ===")
	fmt.Println(GitHubActionsYaml)
	fmt.Println("===================================================")
}
