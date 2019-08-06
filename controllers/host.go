package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/dinhnguyen138/poker-backend/models"
	"github.com/dinhnguyen138/poker-backend/utilities"
	"github.com/kataras/golog"
)

var hosts []string

func RegisterHost(w http.ResponseWriter, r *http.Request) {
	request := new(models.RegisterHostMsg)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&request)
	golog.Info(request.IpAddress)
	hosts = append(hosts, request.IpAddress)
	w.WriteHeader(http.StatusOK)
}

func PickHost() string {
	for i, s := 0, len(hosts); i < s; i++ {
		if utilities.CheckPing(hosts[i]) == true {
			return hosts[i]
		} else {
			hosts = append(hosts[:i], hosts[i+1:]...)
			i--
			s--
		}
	}
	return ""
}
