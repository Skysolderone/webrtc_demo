package main

import (
	"log"
	"net/http"

	"v1/server"
)

func main() {
	server.AllRooms.Init()
	http.HandleFunc("/create", server.CreateRoomRequestHandler)
	http.HandleFunc("/join", server.JoinRoomRequestHandler)
	log.Println("Listern 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
