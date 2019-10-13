package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	regex := "[?&]t=(\\d*)s?"
	link := "https://youtu.be/7i-QNuLwWbU?t=776s"
	index := regexp.MustCompile(regex).FindStringSubmatch(link)
	fmt.Printf("%v", index)
	newString := regexp.MustCompile(regex).ReplaceAllStringFunc(link, func(s string) string {
		fmt.Println(s)
		timeInSecs, err := strconv.Atoi(index[1])

		if err != nil {
			return s
		}

		return "?t=" + strconv.Itoa(timeInSecs+123)
	})
	fmt.Println(newString)
}
