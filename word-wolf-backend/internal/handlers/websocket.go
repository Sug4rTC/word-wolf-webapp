package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocketHandler(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Printf("WebSocketアップグレードエラー: %v", err)
		return err
	}
	defer ws.Close()

	for {
		messageType, msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("メッセージ受信エラー: %v", err)
			break
		}
		log.Printf("受信メッセージ: %s", msg)

		if err := ws.WriteMessage(messageType, msg); err != nil {
			log.Printf("メッセージ送信エラー: %v", err)
			break
		}
	}
	return nil
}
