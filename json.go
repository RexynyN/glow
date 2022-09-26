package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

func blop() {
	videos := getVideos()

	for i, video := range videos {
		videos[i].Description = video.Description + "\n Remember to like and subscribe!"
	}

	fmt.Println(videos)

	saveVideos(videos)
}

func notAMoronWayToDoIt() {
	coronaVirusJSON := `{
        "name" : "covid-11",
        "country" : "China",
        "city" : "Wuhan",
        "reason" : "Non vedge Food"
    }`

	// Declared an empty map interface
	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(coronaVirusJSON), &result)

	// Print the data type of result variable
	fmt.Println(reflect.TypeOf(result))

	// Reading each value by its key
	fmt.Println("Name :", result["name"],
		"\nCountry :", result["country"],
		"\nCity :", result["city"],
		"\nReason :", result["reason"])
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
