package model

import (
	"math"
	"strconv"
	"time"
)

type YoutubeLink string

func Increase(link YoutubeLink, oldTime time.Time) YoutubeLink {
	return YoutubeLink(string(link) + "&t=" + strconv.FormatFloat(math.Round(time.Now().Sub(oldTime).Seconds()), 'f', 0, 64) + "s")
}
