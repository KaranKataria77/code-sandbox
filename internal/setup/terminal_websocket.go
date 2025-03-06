package setup

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

var runningDir string

type Client struct {
	ID     string
	roomID string
	Send   chan []byte
	Conn   *websocket.Conn
}

type Room struct {
	Clients map[string]*Client
	Mutex   sync.Mutex
}

type RoomManager struct {
	Rooms map[string]*Room
	Mutex sync.Mutex
}

var roomManager = &RoomManager{
	Rooms: make(map[string]*Room),
}

// add client to room
func AddClientToRoom(client *Client, RoomID string) {
	// check room id exists of not
	roomManager.Mutex.Lock()
	defer roomManager.Mutex.Unlock()
	if _, exists := roomManager.Rooms[RoomID]; !exists {
		log.Println("RoomId " + RoomID + " does not exists, creating new one ")
		roomManager.Rooms[RoomID] = &Room{
			Clients: make(map[string]*Client),
		}
	}

	room := roomManager.Rooms[RoomID]
	room.Mutex.Lock()
	defer room.Mutex.Unlock()

	room.Clients[client.ID] = client
	log.Println("ClientId " + client.ID + " added to Room " + RoomID)
}

// write message for every client
func WriteMessage(client *Client) {
	for msg := range client.Send {
		err := client.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("Error while writing message")
			break
		}
	}
}

// remove client from room
func RemoveClientFromRoom(client *Client, RoomID string) {
	roomManager.Mutex.Lock()
	defer roomManager.Mutex.Unlock()

	// check room exists or not
	room, exists := roomManager.Rooms[RoomID]
	if !exists {
		log.Println("Room with ID " + RoomID + " does not exists")
		return
	}

	room.Mutex.Lock()
	defer room.Mutex.Unlock()

	// check client exists
	if _, exists := room.Clients[client.ID]; !exists {
		log.Println("ClientId " + client.ID + " was not in room ")
		return
	}
	delete(room.Clients, client.ID)
}

// read message for every client
func ReadMessage(client *Client) {
	// if fail to read message remove client from room
	defer func() {
		log.Println("Removing client" + client.ID + " from roomID " + client.roomID)
		RemoveClientFromRoom(client, client.roomID)
	}()
	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message from client " + client.ID + " = " + err.Error())
			break
		}
		BroadCastToRoom(message, client.roomID, client.ID)
	}
}

// broadcast message to all members in room
func BroadCastToRoom(message []byte, RoomID string, clientID string) {
	// check room exists
	roomManager.Mutex.Lock()
	defer roomManager.Mutex.Unlock()

	room, exists := roomManager.Rooms[RoomID]
	if !exists {
		log.Println("RoomId " + RoomID + " does not exists")
		return
	}

	room.Mutex.Lock()
	defer room.Mutex.Unlock()

	for _, client := range room.Clients {
		if client.ID == clientID {
			continue
		}
		client.Send <- message
	}
}
