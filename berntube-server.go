package main

import (
	"fmt"
	"github.com/bweben/berntube-server/config"
	"github.com/bweben/berntube-server/model"
	"github.com/bweben/berntube-server/web"
	"github.com/bweben/berntube-server/web/socket"
	engineio "github.com/googollee/go-engine.io"
	"github.com/googollee/go-engine.io/transport"
	"github.com/googollee/go-engine.io/transport/polling"
	"github.com/googollee/go-engine.io/transport/websocket"
	"github.com/googollee/go-socket.io"
	"github.com/plimble/ace"
	"github.com/plimble/ace-contrib/cors"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	ApiEndpoint = "/api/v1"

	RoomEndpoint  = ApiEndpoint + "/room/:id"
	RoomsEndpoint = ApiEndpoint + "/rooms"

	ConnRoomEndpoint = ApiEndpoint + "/conn/room/:id"
	SocketIoConn     = "/socket.io/"

	Address         = ":5000"
	SocketIoAddress = ":5001"
)

func main() {
	a := ace.Default()

	a.Use(cors.Cors(config.CorsOptions))

	a.GET(RoomEndpoint, web.RoomHandler)
	a.GET(RoomsEndpoint, web.RoomsHandler)
	a.GET(ConnRoomEndpoint, socket.ConnectHandler)

	go createSocketIoServer()

	a.Run(Address)
}

func createSocketIoServer() {
	pt := polling.Default
	wt := websocket.Default

	wt.CheckOrigin = func(req *http.Request) bool {
		return true
	}

	server, err := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			pt,
			wt,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "join", func(s socketio.Conn, msg string) {
		fmt.Println("room:")
		fmt.Printf("%v", s.Rooms())
		fmt.Println(msg)
		s.Join(msg)
		fmt.Printf("%v", s.Rooms())
		fmt.Printf("%v", s.Context())
		fmt.Println("---")

		socketRoomNumber, err := strconv.Atoi(msg)

		if err == nil {
			indexInRoom := findRoomIndex(socketRoomNumber)

			if indexInRoom != -1 {
				room := model.Rooms[indexInRoom]
				room.Running = model.Increase(room.Running, room.StartTime)
				s.Emit("link-update", room)
			}
		}
	})

	server.OnEvent("/", "link", func(s socketio.Conn, msg string) {
		fmt.Println("link: " + msg)
		s.SetContext(msg)
		for _, socketRoom := range s.Rooms() {
			socketRoomNumber, err := strconv.Atoi(socketRoom)

			if err == nil {
				indexInRooms := findRoomIndex(socketRoomNumber)

				if indexInRooms == -1 {
					model.Rooms = append(model.Rooms, model.Room{
						Id:        socketRoomNumber,
						Name:      "",
						Running:   model.YoutubeLink(msg),
						Queue:     []model.YoutubeLink{},
						StartTime: time.Now(),
					})

					indexInRooms = len(model.Rooms) - 1
				} else {
					model.Rooms[indexInRooms].Running = model.YoutubeLink(msg)
				}

				server.BroadcastToRoom(socketRoom, "link-update", model.Rooms[indexInRooms])
			}
		}
	})

	server.OnError("/", func(e error) {
		fmt.Println("meet error", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, msg string) {
		fmt.Println("closed", msg)
	})

	go server.Serve()
	defer server.Close()

	http.Handle(SocketIoConn, corsMiddleware(server))
	log.Fatal(http.ListenAndServe(SocketIoAddress, nil))
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

func findRoomIndex(id int) (roomToFind int) {
	roomToFind = -1
	for index, room := range model.Rooms {
		if room.Id == id {
			roomToFind = index
		}
	}
	return
}
