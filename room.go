package main

import "fmt"

type room struct {
	name     string
	members  map[string]*client
	messages chan message
}

func NewRoom(name string) *room {
	return &room{
		name:     name,
		members:  make(map[string]*client),
		messages: make(chan message),
	}
}

func (r *room) AddClient(c *client) {
	c.room_name = r.name
	c.messages = r.messages
	r.members[c.conn.RemoteAddr().String()] = c
	r.broadcastInfo(c.GetAddress(), fmt.Sprintf("User %s entered the room.\n", c.nick))
}

func (r *room) RemoveClient(c *client) {
	delete(r.members, c.GetAddress())
	r.broadcastInfo(c.GetAddress(), fmt.Sprintf("User %s leave the room.\n", c.nick))
}

func (r *room) WaitMessage() {
	for msg := range r.messages {
		r.Broadcast(msg)
	}
}

func (r *room) Broadcast(msg message) {
	for _, c := range r.members {
		if c.conn == msg.client.conn {
			continue
		}
		c.SendMsg(msg)
	}
}

func (r *room) broadcastInfo(sender_addr, txt string) {
	for _, c := range r.members {
		if c.GetAddress() == sender_addr {
			continue
		}
		c.sendInfo(txt)
	}
}
