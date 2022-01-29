package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

func main() {
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:3000/ws?uid=1&toUid=2", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			fmt.Println("正在读取消息...")
			_, msg, err := c.ReadMessage()
			if err != nil {
				log.Fatal("read:", err)
			}
			log.Println("recv: ", string(msg))
			var sendString string
			fmt.Print("输入要发送的内容:")
			_, _ = fmt.Scan(&sendString)
			_ = c.WriteMessage(websocket.TextMessage, []byte(sendString))
		}
	}()
	wg.Wait()
}
