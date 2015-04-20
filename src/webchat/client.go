package main

import (
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type client struct {
	socket   *websocket.Conn
	send     chan *message
	room     *room
	userData map[string]interface{}
}

func (c *client) read() {
	for {
		var msg *message
		if err := c.socket.ReadJSON(&msg); err == nil {
			c.room.forward <- msg
			msg.Name = c.userData["name"].(string)
			msg.When = time.Now()

			if url, ok := c.userData["avatar_url"].(string); ok {
				msg.AvatarURL = url
			}

			c.room.tracer.Trace(" client " + c.socket.RemoteAddr().String() + " send message.\n")
		} else {
			log.Println(err)
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteJSON(msg); err != nil {
			break
		}
		log.Println("Send msg to " + c.socket.RemoteAddr().String())
	}
	c.socket.Close()
}
