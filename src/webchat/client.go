package main

import (
	"github.com/gorilla/websocket"
	"log"
)

type client struct {
	socket *websocket.Conn
	send   chan []byte
	room   *room
}

func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
			c.room.tracer.Trace(" client " + c.socket.RemoteAddr().String() + " send message.\n")
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
		log.Println("Send msg to " + c.socket.RemoteAddr().String())
	}
	c.socket.Close()
}
