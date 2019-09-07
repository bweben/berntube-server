package web

import (
	"fmt"
	"github.com/bweben/berntube-server/model"
	"github.com/bweben/berntube-server/web/helper"
	"github.com/plimble/ace"
	"strconv"
)

func RoomHandler(c *ace.C) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	helper.HandleRoomNotExistingError(c, err, idParam)

	if id >= int64(len(model.Rooms)) {
		fmt.Printf("array index out of range for id: %d\n", id)
		c.Redirect("/api/v1/rooms")
	} else {
		c.JSON(200, model.Rooms[id])
	}
}
