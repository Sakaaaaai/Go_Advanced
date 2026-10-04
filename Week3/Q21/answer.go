//go:build ignore

package main

import (
	"fmt"
	"net/http"
)

type wsConn struct {
	closed bool
}
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

type Upgrader struct{}
func (u *Upgrader) Upgrade(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*wsConn, error) {
	fmt.Println("Upgraded to WebSocket!")
	return &wsConn{}, nil
}
var upgrader = Upgrader{}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read error (or disconnected):", err)
			break
		}

		err = conn.WriteMessage(msgType, msg)
		if err != nil {
			fmt.Println("Write error:", err)
			break
		}
	}
}

func main() {
	fmt.Println("Starting WebSocket Echo handler...")
	wsHandler(nil, nil)
}
