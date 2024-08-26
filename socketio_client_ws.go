package main

import (
	"log"
	"runtime"
	"time"

	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

type Channel struct {
	Channel string `json:"channel"`
}

type Message struct {
	Id      int    `json:"id"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func sendJoin(c *gosocketio.Client) {
	log.Println("Acking to the server")
	result, err := c.Ack("golang", Channel{"golang"}, time.Second*5)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Ack result to /join: ", result)
	}
}

func main() {

	log.Println("Hello Go SocketIo Client")
	runtime.GOMAXPROCS(runtime.NumCPU())

	c, err := gosocketio.Dial(
		gosocketio.GetUrl("localhost", 3004, false),
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		log.Fatal(err)
	}

	err = c.On(gosocketio.OnConnection, func(h *gosocketio.Channel) {
		log.Println("SocketIO Client Connected")
	})
	if err != nil {
		log.Fatal(err)
	}

	err = c.On(gosocketio.OnDisconnection, func(h *gosocketio.Channel) {
		log.Fatal("SocketIO Client Disconnected")
	})
	if err != nil {
		log.Fatal(err)
	}

	// err = c.On("/message", func(h *gosocketio.Channel, args Message) {
	// 	log.Println("--- Got chat message: ", args)
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err = c.Emit("golang", Channel{"main"})
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(30 * time.Second)

	go sendJoin(c)
	go sendJoin(c)
	go sendJoin(c)
	go sendJoin(c)
	go sendJoin(c)

	time.Sleep(60 * time.Second)
	c.Close()

	log.Println(" [x] Complete")
}