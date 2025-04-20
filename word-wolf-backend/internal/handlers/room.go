package handlers

import (
	"crypto/rand"
	"errors"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
)

type Room struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var (
	Rooms   = map[string]Room{}
	RoomsMu sync.RWMutex
)

func generateRoomID(digit uint32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	//乱数を生成
	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error..." + err.Error())
	}
	result := make([]byte, digit)
	for i, v := range b {
		result[i] = letters[int(v)%len(letters)]
	}
	return string(result), nil
}

func CreateRoom(c echo.Context) error {
	roomID, err := generateRoomID(10)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	room := Room{
		ID:   roomID,
		Name: "New Room",
	}

	RoomsMu.Lock()
	defer RoomsMu.Unlock()
	Rooms[roomID] = room

	return c.JSON(http.StatusCreated, room)
}
