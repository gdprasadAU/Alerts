package main

import (
	"Alerts/controller"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/alerts/service_id={service_id}&start_ts={start_ts}&end_ts={end_ts}", controller.ReadFromDB).Methods("GET")
	router.HandleFunc("/alerts", controller.WriteToDB).Methods("POST")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}
