package main

import (
	"fmt"
	"strings"
)

type server struct {
	rooms    map[string]*room
	commands chan command
}

func NewServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) Run() {
	for _, r := range s.rooms {
		go r.WaitMessage()
	}
	go s.WaitInput()
}

func (s *server) AddRoom(r *room) {
	s.rooms[r.name] = r
}

func (s *server) GetRoom(name string) *room {
	return s.rooms[name]
}

func (s *server) WaitInput() {
	for cmd := range s.commands {
		switch cmd.command[1:] {
		case "rename":
			if len(cmd.args) < 1 {
				cmd.client.sendInfo("No name entered.\n")
				continue
			}
			s.renameUser(cmd)
		case "rooms":
			s.listRooms(cmd)
		case "members":
			s.listMembers(cmd)
		case "conn":
			cmd.client.SendText(cmd.client.conn.RemoteAddr().String())
		case "chroom":
			if len(cmd.args) < 1 {
				cmd.client.sendInfo("No room entered.\n")
				continue
			}
			if s.GetRoom(cmd.args[0]) == nil {
				cmd.client.sendInfo("Not a valid room.\n")
				continue
			}
			s.changeClientRoom(cmd)
		case "help":
			cmd.client.SendText("Commands:")
			cmd.client.SendText("\n- .rename  = update your nickname")
			cmd.client.SendText("\n- .rooms   = list all availible rooms")
			cmd.client.SendText("\n- .members = list all members from your current room")
			cmd.client.SendText("\n- .conn    = show your connection remote address")
			cmd.client.SendText("\n- .chroom  = change your current room")
			cmd.client.SendText("\n- .quit    = close your connection\n")
		case "quit":
			cmd.client.conn.Close()
			s.GetRoom(cmd.client.room_name).RemoveClient(cmd.client)
		default:
			cmd.client.sendInfo("Not a valid command.\n")
		}
		fmt.Printf("> %s : %s.\n", cmd.command, strings.Join(cmd.args, ", "))
	}
}

func (s *server) renameUser(cmd command) {
	cmd.client.nick = cmd.args[0]
	cmd.client.sendInfo("Your nick is now: " + cmd.args[0] + "\n")
}

func (s *server) listRooms(cmd command) {
	cmd.client.SendText("availiable rooms:\n")
	for r_name, r := range s.rooms {
		cmd.client.SendText(fmt.Sprintf("%s - %d online\n", r_name, len(r.members)))
	}
	cmd.client.BreakLine()
}

func (s *server) listMembers(cmd command) {
	cmd.client.SendText(cmd.client.room_name + "'s members:\n")
	for _, c := range s.GetRoom(cmd.client.room_name).members {
		cmd.client.SendText("- " + c.nick + "\n")
	}
	cmd.client.BreakLine()
}

func (s *server) changeClientRoom(cmd command) {
	oldRoom := cmd.client.room_name
	newRoom := cmd.args[0]

	s.GetRoom(newRoom).AddClient(cmd.client)
	s.GetRoom(oldRoom).RemoveClient(cmd.client)
	cmd.client.sendInfo(fmt.Sprintf("Your are now in %s.\n", cmd.client.room_name))
}
