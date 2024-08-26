package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"github.com/gorilla/websocket"
)

var html = []byte(
	`<html>
		<body>
			<h1>Hello Go!</h1>
			<code></code>
			<script>
				var ws = new WebSocket("ws://127.0.0.1:7777/ws")
				ws.onmessage = function(e) {
					document.querySelector("code").innerHTML += e.data + "<br>"
				}
			</script>
		</body>
	</html>
	`)

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write(html)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("", err)))
		return
	}
	defer ws.Close()
	ws.WriteMessage(1, []byte("About to run the command...\n"))

	// execute and get a pipe
	cmd := exec.Command("tasklist")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println(err)
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Println(err)
		return
	}

	if err := cmd.Start(); err != nil {
		log.Println(err)
		return
	}

	s := bufio.NewScanner(io.MultiReader(stdout, stderr))
	for s.Scan() {
		ws.WriteMessage(1, s.Bytes())
	}

	if err := cmd.Wait(); err != nil {
		log.Println(err)
		return
	}

	ws.WriteMessage(1, []byte("Done\n"))
}

func main() {
	fmt.Println("Hello Go WebSocket")
	fmt.Println("Web server listening to 7777")

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", serveWs)
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		fmt.Printf("Error occurred - %s\n", err)
	}
}
