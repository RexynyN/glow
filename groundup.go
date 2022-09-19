package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

// func main() {
// 	renameFiles("C:\\Users\\Admin\\Desktop\\Nova pasta", randomName)

// 	// fmt.Println(gimmePrimes(1000))
// }

func runCommand(command string, path string) (string, error) {
	cmd := exec.Command(command)

	if path != "" {
		cmd.Dir = path
	}

	cmd.Wait()
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(stdout), err
}

func ffmpegVideoCompress(path, string, file string) {
	oldPath := path + "\\" + file
	newPath := path + "\\compressed\\" + file

	if !existsDir(newPath) {
		os.Mkdir(newPath, os.ModePerm)
	}

	runCommand("ffmpeg -i "+oldPath+" -vcodec libx264 -crf 24 "+newPath, path)
}

func treatFiles(path string, callback func(string, string)) {
	for _, file := range readFiles(path) {
		if file.IsDir() {
			continue
		}
		callback(path, file.Name())
	}
}

func existsDir(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func readFiles(path string) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	return files
}
