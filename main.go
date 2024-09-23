package main

import (
	"server-streaming-rpc/client"
	"server-streaming-rpc/server"
)

func main() {
	go server.Run()
	client.Run()
}
