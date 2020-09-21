package comet

import (
	"fmt"
	"im/config"
	"log"
	"net/http"
)

func StartWS() {
	imC := config.Config.Im
	hub := newHub()
	go hub.run()
	http.HandleFunc(imC.Index, func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	fmt.Println("ws服务器开始启动...")
	err := http.ListenAndServe(fmt.Sprintf(":%d", imC.Port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
