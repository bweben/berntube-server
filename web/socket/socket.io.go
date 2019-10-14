package socket

import (
	"fmt"
	engineio "github.com/googollee/go-engine.io"
	"github.com/googollee/go-engine.io/transport"
	"github.com/googollee/go-engine.io/transport/websocket"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

const (
	ConnRoomEndpoint = "/socket.io/"
	AddressWS        = ":5050"
)

func ServeServer() {
	wt := websocket.Default
	wt.CheckOrigin = func(req *http.Request) bool {
		return true
	}

	server, err := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			wt,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	server.OnConnect("/", func(conn socketio.Conn) error {
		conn.SetContext("")
		fmt.Print(conn.ID())
		return nil
	})

	server.OnEvent("/", "test", func(s socketio.Conn, msg string) {
		fmt.Print(msg)
	})

	go server.Serve()
	defer server.Close()

	http.Handle(ConnRoomEndpoint, corsMiddleware(server))
	http.ListenAndServe(AddressWS, nil)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, PATCH, GET, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", allowHeaders)

		next.ServeHTTP(w, r)
	})
}
