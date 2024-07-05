package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var AllRooms RoomMap

func CreateRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Origin", "*")
	roomId := AllRooms.CreateRoom()
	type resp struct {
		RoomId string `json:"room_id"`
	}
	json.NewEncoder(w).Encode(resp{RoomId: roomId})
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type brodcastMsg struct {
	Message map[string]interface{}
	RoomId  string
	Client  *websocket.Conn
}

var brodast = make(chan brodcastMsg)

func brodaster() {
	for {
		msg := <-brodast
		for _, client := range AllRooms.Map[msg.RoomId] {
			if client.Conn != msg.Client {
				err := client.Conn.WriteJSON(msg.Message)
				if err != nil {
					log.Fatal(err)
					client.Conn.Close()
				}
			}
		}

	}
}

func JoinRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	roomId, ok := r.URL.Query()["roomId"]
	if !ok {
		log.Println("RoomID missing in URL Parameters")
		return
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	go brodaster()
	for {
		var msg brodcastMsg
		err := ws.ReadJSON(&msg.Message)
		if err != nil {
			log.Fatal("read err:", err)
		}

		msg.Client = ws
		msg.RoomId = roomId[0]
		log.Println(msg.Message)
		brodast <- msg
	}
}
