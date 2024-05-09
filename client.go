package main

import (
	"bufio"
	"net"
	"strings"
)

type client struct {
	conn      net.Conn
	nick      string
	room_name string
	messages  chan<- message
	commands  chan<- command
}

func (c *client) ReadInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}
		msg = strings.Trim(msg, "\r\n")
		msg = strings.TrimSpace(msg)
		if len(msg) == 0 {
			continue
		}

		if msg[0] == '.' {
			cmdStr := strings.Split(msg, " ")
			c.commands <- NewCommand(c, cmdStr[0], cmdStr[1:])
		} else {
			c.messages <- NewMessage(c, msg)
		}
	}
}

func (c *client) GetAddress() string {
	return c.conn.RemoteAddr().String()
}

func (c *client) SendText(txt string) {
	c.conn.Write([]byte(txt))
}

func (c *client) BreakLine() {
	c.SendText("")
}

func (c *client) sendInfo(txt string) {
	c.SendText("\033[93m: " + txt + "\033[0m")
}

func (c *client) SendMsg(msg message) {
	c.SendText(msg.client.nick + ": " + msg.data + "\n")
}
