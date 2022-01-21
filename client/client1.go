package main

import (
	"fmt"
	"golang.org/x/net/websocket"
)

func main() {
	resp, _ := websocket.Dial("ws://localhost:3000/ws?uid=1&toUid=2", "", "http://localhost:3000")
	fmt.Println("response: ", resp)
}
