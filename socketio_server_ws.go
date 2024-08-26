package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	fmt.Println("Hello Golang Socket.io")

	server := socketio.NewServer(nil)

	server.OnConnect("/", func(sock socketio.Conn) error {
		sock.SetContext("")
		fmt.Println("SocketIO Connected:", sock.ID())
		return nil
	})

	server.OnError("/", func(sock socketio.Conn, err error) {
		fmt.Println("SocketIO Error:", err)
	})

	server.OnDisconnect("/", func(sock socketio.Conn, reason string) {
		fmt.Println("SocketIO Closed", reason)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	log.Println("Serving at localhost:7070...")
	log.Fatal(http.ListenAndServe(":7070", nil))
}
