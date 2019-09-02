package model

type Room struct {
	Id      int
	Name    string
	Running YoutubeLink
	Queue   []YoutubeLink
}

var Rooms = []Room{
	{
		Id:      123,
		Name:    "Sample",
		Running: "https://youtu.be/Yw6u6YkTgQ4",
		Queue:   []YoutubeLink{},
	},
}
