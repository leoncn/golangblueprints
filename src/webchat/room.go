package main

import (
	"github.com/gorilla/websocket"
	"github.com/stretchr/objx"
	"log"
	"net/http"
	"os"
	"trace"
)

type room struct {
	forward chan *message

	join, leave chan *client

	clients map[*client]bool
	tracer  trace.Tracer
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
			r.tracer.Trace("client " + client.socket.RemoteAddr().String() + " joined.\n")
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
			r.tracer.Trace("client " + client.socket.RemoteAddr().String() + " left.\n")
		case msg := <-r.forward:
			r.tracer.Trace("Recived message ", msg.Message, "\n")
			for client := range r.clients {
				select {
				case client.send <- msg:
				default:
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)

	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}

	authCookie, err := req.Cookie("auth")
	if err != nil {
		log.Fatal("Fail to get auth cookie:", err)
		return
	}

	client := &client{
		socket:   socket,
		send:     make(chan *message, messageBufferSize),
		room:     r,
		userData: objx.MustFromBase64(authCookie.Value),
	}

	r.join <- client

	defer func() {
		r.leave <- client
	}()

	go client.write()
	client.read()
}

func newRoom() *room {
	return &room{
		forward: make(chan *message),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		tracer:  trace.New(os.Stdout),
	}
}
