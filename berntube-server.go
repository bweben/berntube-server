package main

import (
	"github.com/bweben/berntube-server/config"
	"github.com/bweben/berntube-server/web"
	"github.com/bweben/berntube-server/web/socket"
	"github.com/plimble/ace"
	"github.com/plimble/ace-contrib/cors"
)

func main() {
	a := ace.Default()

	a.Use(cors.Cors(config.CorsOptions))

	a.GET("/api/v1/room/:id", web.RoomHandler)
	a.GET("/api/v1/rooms", web.RoomsHandler)
	a.GET("/api/v1/conn/room/:id", socket.ConnectHandler)

	a.Run(":5000")
}
