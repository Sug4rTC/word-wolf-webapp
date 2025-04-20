package handlers

import (
	"crypto/rand"
	"errors"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
)

type Participant struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var (
	RoomParticipants   = make(map[string][]Participant)
	RoomParticipantsMu sync.RWMutex
)

func GenerateParticipantID(digit uint32) (string, error) {
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

func CreateParticipant(c echo.Context) error {
	roomID := c.Param("roomID")
	if roomID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "roomID is required"})
	}

	RoomsMu.RLock()
	_, exists := Rooms[roomID]
	RoomsMu.RUnlock()
	if !exists {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "room not found"})
	}

	var req struct {
		Name string `json:"name"`
	}
	if err := c.Bind(&req); err != nil || req.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "name is required"})
	}

	participantID, err := GenerateParticipantID(6)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	newParticipant := Participant{
		ID:   participantID,
		Name: req.Name,
	}

	RoomParticipantsMu.Lock()
	defer RoomParticipantsMu.Unlock()
	RoomParticipants[roomID] = append(RoomParticipants[roomID], newParticipant)

	return c.JSON(http.StatusCreated, newParticipant)

}
