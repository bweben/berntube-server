package main

import (
	"github.com/bweben/berntube-server/web"
	"github.com/plimble/ace"
)

func main() {
	a := ace.New()
	a.GET("/api/v1/room/:id", web.RoomHandler)
	a.GET("/api/v1/rooms", web.RoomsHandler)
	a.Run(":8180")
}
