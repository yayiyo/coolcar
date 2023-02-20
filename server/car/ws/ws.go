package ws

import (
	"context"
	"fmt"
	"net/http"

	"coolcar/car/mq"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

func Handler(u *websocket.Upgrader, sub mq.Subscriber, logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := u.Upgrade(w, r, nil)
		if err != nil {
			logger.Error("error upgrading websocket", zap.Error(err))
			return
		}
		defer conn.Close()

		msgs, cleanup, err := sub.Subscribe(context.Background())
		defer cleanup()
		if err != nil {
			logger.Error("error subscribing", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		done := make(chan struct{})
		go func() {
			for {
				_, _, err = conn.ReadMessage()
				if err != nil {
					if !websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
						fmt.Printf("unexpected websocket error: %v\n", err)
						return
					}
					fmt.Printf("can not json %v\n", err)
					done <- struct{}{}
					break
				}
			}
		}()

		for {
			select {
			case msg := <-msgs:
				logger.Info("received message, sending to websocket")
				err = conn.WriteJSON(msg)
				if err != nil {
					logger.Error("error writing websocket", zap.Error(err))
				}
			case <-done:
				return
			}
		}
	}
}
