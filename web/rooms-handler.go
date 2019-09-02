package web

import (
	"github.com/bweben/berntube-server/model"
	"github.com/plimble/ace"
)

func RoomsHandler(c *ace.C) {
	c.JSON(200, model.Rooms)
}
