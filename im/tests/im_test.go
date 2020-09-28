package tests

import (
	"github.com/gorilla/websocket"
	"im/config"
	"log"
	"net/url"
	"testing"
	"time"
	"utils"
)

func init() {
	if err := utils.ScanConfig(&config.Config); err != nil {
		panic(err)
	}
	config.Init()
}

func TestComet(t *testing.T) {
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8888", Path: "/im"}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("err:", err)
				return
			}
			log.Printf("收到消息: %s", message)
		}
	}()

	for {
		str := "hello"
		err := conn.WriteMessage(websocket.TextMessage, []byte(str))
		if err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second)
	}
}
