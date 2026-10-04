package main

import (
	"fmt"
	"net/http"
)

// ========= モック環境 =========
type wsConn struct {
	closed bool
}
// ReadMessageはダミーのメッセージを1回だけ返し、次はエラーを返すようにモック化
var readCount int
func (c *wsConn) ReadMessage() (int, []byte, error) {
	if readCount > 0 {
		return 0, nil, fmt.Errorf("client disconnected")
	}
	readCount++
	return 1, []byte("Hello WebSocket!"), nil
}
func (c *wsConn) WriteMessage(messageType int, data []byte) error {
	fmt.Printf("Echoing message: %s\n", string(data))
	return nil
}
func (c *wsConn) Close() error {
	c.closed = true
	fmt.Println("Connection closed.")
	return nil
}

// アップグレーダーのモック
type Upgrader struct{}
func (u *Upgrader) Upgrade(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*wsConn, error) {
	fmt.Println("Upgraded to WebSocket!")
	return &wsConn{}, nil
}
var upgrader = Upgrader{}
// =============================

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// 1. HTTPリクエストをWebSocketにアップグレード
	conn, err := upgrader.___(___, ___, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	// ハンドラ終了時に忘れずに閉じる
	___ conn.___()

	// 2. 無限ループで送受信を行うエコーサーバー
	for {
		// メッセージを読み込む
		msgType, msg, err := conn.___()
		if err != nil {
			fmt.Println("Read error (or disconnected):", err)
			___ // ループを抜ける
		}

		// 受け取ったメッセージをそのまま送り返す（Echo）
		err = conn.___(msgType, msg)
		if err != nil {
			fmt.Println("Write error:", err)
			___ // ループを抜ける
		}
	}
}

func main() {
	// ダミーのリクエストでハンドラを直接呼び出す
	fmt.Println("Starting WebSocket Echo handler...")
	wsHandler(nil, nil)
}
