package service

import (
	"chat/conf"
	"chat/ret"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

func (manager *ClientManager) Start() {
	for {
		log.Println("<---监听管道通信--->")
		select {
		case client := <-Manager.Online: // 建立连接
			log.Printf("建立新连接: %v", client.ID)
			Manager.Clients[client.ID] = client
			replyMsg := &ReplyMsg{
				Code:    ret.WebsocketSuccess,
				Content: "已连接至服务器",
			}
			msg, _ := json.Marshal(replyMsg)
			_ = client.WsConn.WriteMessage(websocket.TextMessage, msg)
		case client := <-Manager.Offline: // 断开连接
			log.Printf("连接失败:%v", client.ID)
			if _, ok := Manager.Clients[client.ID]; ok {
				replyMsg := &ReplyMsg{
					Code:    ret.WebsocketEnd,
					Content: "连接已断开",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = client.WsConn.WriteMessage(websocket.TextMessage, msg)
				close(client.Send)
				delete(Manager.Clients, client.ID)
			}
		case broadcast := <-Manager.Broadcast: // 广播信息
			message := broadcast.Message
			sendId := broadcast.Client.SendID
			flag := false // 默认对方不在线
			for id, client := range Manager.Clients {
				if id != sendId {
					continue
				}
				select {
				case client.Send <- message:
					flag = true
				default:
					close(client.Send)
					delete(Manager.Clients, client.ID)
				}
			}
			id := broadcast.Client.ID
			if flag {
				log.Println("对方在线应答")
				replyMsg := &ReplyMsg{
					Code:    ret.WebsocketOnlineReply,
					Content: "对方在线应答",
				}
				msg, err := json.Marshal(replyMsg)
				_ = broadcast.Client.WsConn.WriteMessage(websocket.TextMessage, msg)
				err = InsertMsg(conf.MongoDBName, id, string(message), 1, int64(3*month))
				if err != nil {
					fmt.Println("InsertOneMsg Err", err)
				}
			} else {
				log.Println("对方不在线")
				replyMsg := ReplyMsg{
					Code:    ret.WebsocketOfflineReply,
					Content: "对方不在线应答",
				}
				msg, err := json.Marshal(replyMsg)
				_ = broadcast.Client.WsConn.WriteMessage(websocket.TextMessage, msg)
				err = InsertMsg(conf.MongoDBName, id, string(message), 0, int64(3*month))
				if err != nil {
					fmt.Println("InsertOneMsg Err", err)
				}
			}
		}
	}
}
