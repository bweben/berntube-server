package helper

import (
	"fmt"
	"github.com/plimble/ace"
)

func HandleRoomNotExistingError(c *ace.C, err error, param string) {
	if err != nil {
		fmt.Printf("room '%s' doesn't exist\n", param)
		c.Redirect("/api/v1/rooms")
	}
}
