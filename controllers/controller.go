package controllers

import (
	"encoding/json"
	"net/http"

	m "eksplorasi/models"
)

// GetAllRooms..
func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT id, room_name FROM rooms"

	rows, err := db.Query(query)
	if err != nil {
		sendModifiedResponse(w, 400, "Query Error")
		return
	}

	var room m.Room
	var rooms []m.Room
	var result m.Rooms

	for rows.Next() {
		if err := rows.Scan(&room.ID, &room.RoomName); err != nil {
			sendModifiedResponse(w, 400, "Scan Error")
			return
		} else {
			rooms = append(rooms, room)
		}
	}

	w.Header().Set("Content-Type", "application/json")

	var response m.RoomsResponse
	response.Status = 200
	response.Message = "Success"
	result.Rooms = rooms
	response.Data = result
	json.NewEncoder(w).Encode(response)
}

// Modified Response Function
func sendModifiedResponse(w http.ResponseWriter, stat int, msg string) {
	var response m.Response
	response.Status = stat
	response.Message = msg
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
