package main

import (
	"github.com/bweben/berntube-server/config"
	"github.com/bweben/berntube-server/web"
	"github.com/bweben/berntube-server/web/socket"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"
	"github.com/martini-contrib/render"
)

const (
	ApiEndpoint = "/api/v1"

	RoomEndpoint  = ApiEndpoint + "/room/:id"
	RoomsEndpoint = ApiEndpoint + "/rooms"

	SocketIoConn = "/socket.io/"

	Address         = ":5000"
	SocketIoAddress = ":5001"
)

func main() {
	server := martini.Classic()

	server.Use(cors.Allow(config.CorsOptions))
	server.Use(render.Renderer())

	server.Get(RoomEndpoint, web.RoomHandler)
	server.Get(RoomsEndpoint, web.RoomsHandler)

	go socket.CreateSocketIoServer(SocketIoConn, SocketIoAddress)

	server.RunOnAddr(Address)
}
