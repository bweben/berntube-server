package socket

import (
	socketio "github.com/googollee/go-socket.io"
	"log"
)

func CreateServer() *socketio.Server {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	go server.Serve()
	defer server.Close()

	return server
}
