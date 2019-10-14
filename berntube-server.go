package main

import (
	"github.com/bweben/berntube-server/config"
	"github.com/bweben/berntube-server/web"
	"github.com/bweben/berntube-server/web/socket"
	"github.com/plimble/ace"
	"github.com/plimble/ace-contrib/cors"
)

const (
	ApiEndpoint = "/api/v1"

	RoomEndpoint  = ApiEndpoint + "/room/:id"
	RoomsEndpoint = ApiEndpoint + "/rooms"

	Address = ":5000"
)

func main() {
	a := ace.Default()

	a.Use(cors.Cors(config.CorsOptions))

	a.GET(RoomEndpoint, web.RoomHandler)
	a.GET(RoomsEndpoint, web.RoomsHandler)

	// have to use http as ace is to strict in the HandleFunc type
	socket.ServeServer()

	a.Run(Address)
}
