package main

import (
	"fmt"
)

type message struct {
	data     []byte
	room     string
	username string
}

type subscription struct {
	conn *connection
	room string
	user string
}

// hub maintains the set of active connections and broadcasts messages to the
// connections.
type hub struct {
	// Registered connections.
	rooms map[string]map[userconnection]bool

	// Inbound messages from the connections.
	broadcast chan message

	// Register requests from the connections.
	register chan subscription

	// Unregister requests from connections.
	unregister chan subscription
}

type userconnection struct {
	connection *connection
	username   string
}

var h = hub{
	broadcast:  make(chan message),
	register:   make(chan subscription),
	unregister: make(chan subscription),
	rooms:      make(map[string]map[userconnection]bool),
}

func listRooms() []string {
	arr := make([]string, 0)
	fmt.Println(h.rooms)
	for k, _ := range h.rooms {
		arr = append(arr, k)
	}
	return arr
}

func removeUser(room, username string) bool {
	if val, ok := h.rooms[room]; ok {
		for user, _ := range val {
			if user.username == username {
				delete(val, user)
				return true
			}
		}
	}
	return false
}

func listUser(room string) []string {
	if val, ok := h.rooms[room]; ok {
		arr := make([]string, 0)
		for k, v := range val {
			if v {
				arr = append(arr, k.username)
			}
		}
		return arr
	}
	return []string{}
}

func deleteRooms(room string) bool {
	if _, ok := h.rooms[room]; ok {
		delete(h.rooms, room)
		return true
	}
	return false
}

func (h *hub) run() {
	for {
		select {
		case s := <-h.register:
			fmt.Println("Reg: ", s)
			connections := h.rooms[s.room]
			if connections == nil {

				connections := make(map[userconnection]bool)
				h.rooms[s.room] = connections
			}
			usrconn := userconnection{
				connection: s.conn,
				username:   s.user,
			}
			h.rooms[s.room][usrconn] = true

		case s := <-h.unregister:
			connections := h.rooms[s.room]
			if connections != nil {
				usrconn := userconnection{
					connection: s.conn,
					username:   s.user,
				}
				if _, ok := connections[usrconn]; ok {
					delete(connections, usrconn)
					close(s.conn.send)
					if len(connections) == 0 {
						delete(h.rooms, s.room)
					}
				}
			}
		case m := <-h.broadcast:
			fmt.Println("msg : ", m)
			connections := h.rooms[m.room]
			for c := range connections {
				select {
				case c.connection.send <- m.data:
				default:
					close(c.connection.send)
					delete(connections, c)
					if len(connections) == 0 {
						delete(h.rooms, m.room)
					}
				}
			}
		}
	}
}
