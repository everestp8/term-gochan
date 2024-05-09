package main

type message struct {
	client *client
	data   string
}

type command struct {
	client  *client
	command string
	args    []string
}

func NewMessage(c *client, data string) message {
	return message{
		client: c,
		data:   data,
	}
}

func NewCommand(c *client, data string, args []string) command {
	return command{
		client:  c,
		command: data,
		args:    args,
	}
}
