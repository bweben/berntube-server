package web

import (
	"fmt"
	"github.com/bweben/berntube-server/model"
	"github.com/plimble/ace"
	"strconv"
)

func RoomHandler(c *ace.C) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		fmt.Printf("room '%s' doesn't exist\n", idParam)
		c.Redirect("/api/v1/rooms")
	} else if id >= int64(len(model.Rooms)) {
		fmt.Printf("array index out of range for id: %d\n", id)
		c.Redirect("/api/v1/rooms")
	} else {
		c.JSON(200, model.Rooms[id])
	}
}
