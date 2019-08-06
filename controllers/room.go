package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dinhnguyen138/poker-backend/db"
	"github.com/dinhnguyen138/poker-backend/models"
)

var rooms []models.Room

func GetRooms(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	data, _ := json.Marshal(db.GetRooms())
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func CreateRoom(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	request := new(models.CreateRoomMsg)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&request)
	fmt.Println(request)
	host := PickHost()
	var room *models.Room
	if host != "" {
		room = db.CreateRoom(request.Amount, request.NumPlayer, host)
	} else {
		room = nil
	}
	if room == nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		data, _ := json.Marshal(room)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}

}

func QuickFind(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	user := r.Context().Value("user")
	claim := user.(*jwt.Token).Claims.(jwt.MapClaims)
	userId, _ := claim["sub"].(string)
	foundUser := db.GetUser(userId)
	room := db.FindRoom(foundUser.Amount)
	data, _ := json.Marshal(room)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
