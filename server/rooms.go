package server

import (
	"math/rand"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Participant struct {
	Host bool
	Conn *websocket.Conn
}

type RoomMap struct {
	Mutex sync.RWMutex
	Map   map[string][]Participant
}

func (r *RoomMap) Init() {
	r.Map = make(map[string][]Participant)
}

func (r *RoomMap) CreateRoom() string {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	rand.Seed(time.Now().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	RoomId := string(b)
	r.Map[RoomId] = []Participant{}
	return RoomId
}

func (r *RoomMap) InsertRoom(roomId string, host bool, conn *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	s := Participant{}
	s.Conn = conn
	s.Host = host
	r.Map[roomId] = append(r.Map[roomId], s)
}

func (r *RoomMap) DeleteRoom(roomId string, host bool, conn *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	delete(r.Map, roomId)
}
