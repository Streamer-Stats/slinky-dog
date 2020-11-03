package socket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// ERR is const error of socket rcv
const ERR = -1

// Socket connect socket to brain and get live data from it
type Socket struct {
	Clients   map[string]*websocket.Conn
	Broadcast chan string
	Upgrader  *websocket.Upgrader
	Request   *http.Request
}

func (s *Socket) stringToByte(msg string) []byte {
	return []byte(msg)
}
func (s *Socket) byteToString(msg []byte) string {
	return string(msg)
}

func (s *Socket) getClient() string {
	return s.Request.Header.Get("Ws-Client-Name")
}

func (s *Socket) recv(client string) (int, string) {
	msgType, msg, err := s.Clients[client].ReadMessage()
	if err != nil {
		log.Println(err)
		return ERR, ""
	}

	log.Println("Recv from "+client+": ", s.byteToString(msg))
	return msgType, s.byteToString(msg)
}

func (s *Socket) write(client string, msg string, msgType int) {
	err := s.Clients[client].WriteMessage(msgType, s.stringToByte(msg))

	if err != nil {
		log.Println(err)
		return
	}
}

// Handle map my connections with my socket
func (s *Socket) Handle() {
	client := s.getClient()
	var msgType int
	var msg string
	for {
		msgType, msg = s.recv(client)
		if msgType != ERR && client != "brain" {
			s.write("brain", msg, msgType)
		} else if client == "brain" {
			s.write(client, msg, msgType)
		}
	}
}

// CreateConn map my connections with my socket
func (s *Socket) CreateConn(w http.ResponseWriter, r *http.Request) *Socket {

	// Upgrade initial GET request to a websocket
	ws, err := s.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	s.Request = r
	s.Clients[s.getClient()] = ws

	return s

}

// NewSocket constructor
func NewSocket() *Socket {
	return &Socket{
		Clients:   make(map[string]*websocket.Conn),
		Broadcast: make(chan string),
		Upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		},
	}
}
