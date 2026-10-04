package main

import "fmt"

const DockerfileContent = `
# ========== Stage 1: Build ==========
# 1. builder という名前でビルド環境を作る
FROM golang:1.20 AS ___

WORKDIR /app

# コマンドの仮の想定 (go.mod等のコピー処理は簡略化しています)
COPY main.go .

# 2. Goのプログラムをビルドして server というバイナリを作成
RUN go build -o ___ main.go


# ========== Stage 2: Runtime ==========
# 3. 非常に軽量な alpine をランタイムとして使う
FROM ___

WORKDIR /

# 4. builderステージから、ビルド済みのバイナリ(/app/server)だけをコピーしてくる
COPY --from=___ /app/server /server

# 5. コンテナ起動時の実行コマンド
CMD ["/___"]
`

func main() {
	fmt.Println("=== Multi-stage Dockerfile ===")
	fmt.Println(DockerfileContent)
	fmt.Println("==============================")
}
