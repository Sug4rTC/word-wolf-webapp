package handlers

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var (
	roomClients   = make(map[string]map[*websocket.Conn]bool)
	roomClientsMu sync.RWMutex
)

func broadcast(roomID string, v interface{}) {
	roomClientsMu.RLock()
	defer roomClientsMu.RUnlock()
	for conn := range roomClients[roomID] {
		if err := conn.WriteJSON(v); err != nil {
			log.Printf("broad cast error (room %s): %v", roomID, err)
		}
	}
}

func WebSocketHandler(c echo.Context) error {
	roomID := c.QueryParam("roomID")
	if roomID == "" {
		return c.JSON(http.StatusBadRequest, "roomID is required")
	}

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Printf("WebSocketアップグレードエラー: %v", err)
		return err
	}
	defer ws.Close()

	RoomsMu.RLock()
	_, ok := Rooms[roomID]
	RoomsMu.RUnlock()
	if !ok {
		return c.JSON(http.StatusNotFound, "room not found")
	}

	roomClientsMu.Lock()
	if roomClients[roomID] == nil {
		roomClients[roomID] = make(map[*websocket.Conn]bool)
	}
	roomClients[roomID][ws] = true
	roomClientsMu.Unlock()

	broadcast(roomID, map[string]interface{}{
		"type": "join",
	})

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("メッセージ受信エラー: %v", err)
			break
		}

		broadcast(roomID, map[string]interface{}{
			"type":     "message",
			"playload": string(msg),
		})
	}

	roomClientsMu.Lock()
	delete(roomClients[roomID], ws)
	roomClientsMu.Unlock()
	broadcast(roomID, map[string]interface{}{
		"type": "leave",
	})

	return nil
}
