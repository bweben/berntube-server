package socket

import (
	"github.com/bweben/berntube-server/model"
	"github.com/bweben/berntube-server/web/helper"
	"github.com/plimble/ace"
	"strconv"
)

var hubs model.Hubs = make(model.Hubs)

func ConnectHandler(c *ace.C) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	helper.HandleRoomNotExistingError(c, err, idParam)

	var hub *Hub
	if hubs.Has(id) {
		hub, _ = hubs.Get(id)
	} else {
		hub = NewHub()
		go hub.Run()
		_ = hubs.Set(id, hub)
	}

	ServeWS(hub, c.Writer, c.Request, id)
}
