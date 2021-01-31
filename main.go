package main

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	// fmt.Println(servers.Ws)
	serverPort := ":9527"

	server := websocket.Server{
		Handler: wsHandler,
	}

	http.Handle("/ws", server)

	// Server聽一個Port就好，一個Port可以同時使用不同的通訊協定
	fmt.Println("Server啟動!")
	if err := http.ListenAndServe(serverPort, nil); err != nil {
		fmt.Println("Server啟動失敗，錯誤如下：")
		fmt.Println(err.Error())
	}
}

// 對應websocket服務的Handler
func wsServer() {

	serverPort := ":9527"

	http.Handle("/ws", websocket.Handler(wsHandler))

	// WebSocket Server，並監聽
	fmt.Println("WebSocket Server啟動" + serverPort)
	err := http.ListenAndServe(serverPort, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func wsHandler(ws *websocket.Conn) {

	// 預先定義error
	var err error
	var msg string

	//想要保持連線的話，就必需要包在一個無限迴圈裡面，不然執行完就會向Client發送一個關閉訊號
	for {
		msg = ""
		// 接收Client
		if err = websocket.Message.Receive(ws, &msg); err != nil {
			if err != io.EOF {
				fmt.Println("接收訊息錯誤" + err.Error())
			}
			break
		}
		fmt.Println("收到Client訊息：" + msg)

		// 發送訊息
		// msg = "heyhey"
		// if err = websocket.Message.Send(ws, msg); err != nil {
		// 	if err != io.EOF {
		// 		fmt.Println("接收訊息錯誤" + err.Error())
		// 	}
		// 	break
		// }
		// fmt.Println("發送訊息：" + msg)
	}
}
