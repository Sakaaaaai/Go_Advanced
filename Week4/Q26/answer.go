//go:build ignore

package main

import "fmt"

const DockerfileContent = `
# ========== Stage 1: Build ==========
FROM golang:1.20 AS builder

WORKDIR /app

COPY main.go .

RUN go build -o server main.go


# ========== Stage 2: Runtime ==========
FROM alpine:latest

WORKDIR /

COPY --from=builder /app/server /server

CMD ["/server"]
`

func main() {
	fmt.Println("=== Multi-stage Dockerfile ===")
	fmt.Print(DockerfileContent)
	fmt.Println("==============================")
}
