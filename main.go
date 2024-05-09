package main

import (
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	s := NewServer()
	s.AddRoom(NewRoom("#root"))
	s.AddRoom(NewRoom("#mu"))
	s.AddRoom(NewRoom("#pol"))

	s.Run()
	port := "8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	listener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		log.Fatalln("unable to start server:", err)
	}

	defer listener.Close()
	log.Println("server started on port :" + port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("connection error: ", err.Error())
			continue
		}

		code := strings.Split(conn.RemoteAddr().String(), ":")[1]
		c := client{
			conn:     conn,
			nick:     "anon" + code,
			commands: s.commands,
		}
		c.SendText("Welcome to Terminal Golang Channel!")
		c.SendText("\nType .help to see all availible commands\n")
		s.GetRoom("#root").AddClient(&c)
		go c.ReadInput()
	}
}
