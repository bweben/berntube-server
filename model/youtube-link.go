package model

import (
	"math"
	"regexp"
	"strconv"
	"time"
)

const YoutubeTimeRegex = "[?&]t=(\\d*)s?"

type YoutubeLink string

func Increase(link YoutubeLink, oldTime time.Time) YoutubeLink {
	convertedLink := string(link)
	compiledTimeRegex := regexp.MustCompile(YoutubeTimeRegex)

	timePart := compiledTimeRegex.FindStringSubmatch(convertedLink)
	roundedTimeSince := math.Round(time.Now().Sub(oldTime).Seconds())

	return YoutubeLink(compiledTimeRegex.ReplaceAllStringFunc(convertedLink, func(s string) string {
		timeInSecs, err := strconv.Atoi(timePart[1])

		if err != nil {
			return s
		}

		return "?t=" + strconv.Itoa(timeInSecs+int(roundedTimeSince))
	}))
}
