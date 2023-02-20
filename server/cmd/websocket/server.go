package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/ws", handleWebsocket)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func handleWebsocket(w http.ResponseWriter, r *http.Request) {
	upgrade := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("error upgrading websocket error: %v\n", err)
		return
	}
	defer conn.Close()

	done := make(chan struct{})
	go func() {
		for {
			m := make(map[string]interface{})
			err = conn.ReadJSON(&m)
			if err != nil {
				if !websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
					fmt.Printf("unexpected websocket error: %v\n", err)
					return
				}
				fmt.Printf("can not json %v\n", err)
				done <- struct{}{}
				break
			}
			fmt.Printf("received websocket msg: %v\n", m)
		}
	}()

	for i := 0; i < 10; i++ {
		err = conn.WriteJSON(map[string]interface{}{
			"hello": "websocket",
			"id":    i,
		})
		if err != nil {
			fmt.Printf("error writing websocket %v\n", err)
		}
		select {
		case <-time.After(200 * time.Millisecond):
		case <-done:
			return
		}
	}
}
