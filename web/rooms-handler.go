package web

import (
	"github.com/bweben/berntube-server/model"
	"github.com/martini-contrib/render"
)

func RoomsHandler(render render.Render) {
	render.JSON(200, model.Rooms)
}
