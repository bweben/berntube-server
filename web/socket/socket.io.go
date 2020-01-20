package socket

import (
	"fmt"
	"github.com/bweben/berntube-server/model"
	"github.com/bweben/berntube-server/web/helper"
	engineio "github.com/googollee/go-engine.io"
	"github.com/googollee/go-engine.io/transport"
	"github.com/googollee/go-engine.io/transport/polling"
	"github.com/googollee/go-engine.io/transport/websocket"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func CreateSocketIoServer(socketIoConn, socketIoAddress string) {
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
		return nil
	})

	server.OnEvent("/", "join", func(s socketio.Conn, msg string) {
		s.Join(msg)

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

	server.OnEvent("/", "play", func(s socketio.Conn, msg bool) {
		for _, socketRoom := range s.Rooms() {
			server.BroadcastToRoom(socketRoom, "playing", true)
		}
	})

	server.OnEvent("/", "pause", func(s socketio.Conn, msg bool) {
		for _, socketRoom := range s.Rooms() {
			server.BroadcastToRoom(socketRoom, "playing", false)
		}
	})

	server.OnEvent("/", "ended", func(s socketio.Conn, msg bool) {
		fmt.Println("ended")
		// todo: play next in queue
	})

	server.OnError("/", func(e error) {
		fmt.Println("meet error", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, msg string) {
		s.LeaveAll()
		s.Close()
	})

	go server.Serve()
	defer server.Close()

	http.Handle(socketIoConn, helper.CorsMiddleware(server))
	log.Fatal(http.ListenAndServe(socketIoAddress, nil))
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
