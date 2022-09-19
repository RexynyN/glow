package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func blop() {
	videos := getVideos()

	for i, video := range videos {
		videos[i].Description = video.Description + "\n Remember to like and subscribe!"
	}

	fmt.Println(videos)

	saveVideos(videos)
}

type video struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Imageurl    string `json:"imageurl"`
	Url         string `json:"url"`
}

func getVideos() (videos []video) {
	filebytes, err := ioutil.ReadFile("./videos.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(filebytes, &videos)
	if err != nil {
		panic(err)
	}

	return videos
}

func saveVideos(videos []video) {
	videoBytes, err := json.Marshal(videos)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./videos-updated.json", videoBytes, 0644)
	if err != nil {
		panic(err)
	}
}
