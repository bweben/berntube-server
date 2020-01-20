package web

import (
	"fmt"
	"github.com/bweben/berntube-server/model"
	"github.com/bweben/berntube-server/web/helper"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"strconv"
)

func RoomHandler(params martini.Params, render render.Render) {
	idParam := params["id"]
	id, err := strconv.ParseInt(idParam, 10, 64)

	helper.HandleRoomNotExistingError(render, err, idParam)

	if id >= int64(len(model.Rooms)) {
		fmt.Printf("array index out of range for id: %d\n", id)
		render.Redirect("/api/v1/rooms")
	} else {
		render.JSON(200, model.Rooms[id])
	}
}
