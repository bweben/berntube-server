package helper

import (
	"fmt"
	"github.com/martini-contrib/render"
)

func HandleRoomNotExistingError(render render.Render, err error, param string) {
	if err != nil {
		fmt.Printf("room '%s' doesn't exist\n", param)
		render.Redirect("/api/v1/rooms")
	}
}
