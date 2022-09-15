package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

func main() {
	renameFiles("C:\\Users\\Admin\\Desktop\\Nova pasta", randomName)

	// fmt.Println(gimmePrimes(1000))
}

func randomName() string {
	strong := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	alfabet := []rune(strong)
	name := make([]rune, 25)
	for i := 0; i < 25; i++ {
		name[i] = alfabet[rand.Intn(len(strong))]
	}

	return string(name)
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

func ffmpegVideoCompress(path, string, file string) {
	oldPath := path + "\\" + file
	newPath := path + "\\compressed\\" + file

	if !existsDir(newPath) {
		os.Mkdir(newPath, os.ModePerm)
	}

	runCommand("ffmpeg -i "+oldPath+" -vcodec libx264 -crf 24 "+newPath, path)
}

func renameFiles(path string, callback func() string) {
	for _, file := range readFiles(path) {
		if file.IsDir() {
			continue
		}
		newName := callback()
		splot := strings.Split(file.Name(), ".")
		fmt.Println(file.Name() + " -> " + newName)

		err := os.Rename(path+"\\"+file.Name(), path+"\\"+newName+"."+splot[1])
		if err != nil {
			fmt.Println(err)
		}
	}
}

func readFiles(path string) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	return files
}

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

func gimmePrimes(threshold int) []int {
	var number int = 3
	var i int = 1
	var guard bool = true
	primes := []int{}
	primes = append(primes, 2)
	for i < threshold {
		guard = true
		root := int(math.Sqrt(float64(number)))
		for _, prime := range primes {
			if prime > root {
				break
			}

			if number%prime == 0 {
				guard = false
			}
		}

		if guard {
			primes = append(primes, number)
			i += 1
		}
		number++
	}
	return primes
}

func gitCommands(path string) {
	runCommand("git add .", path)
	runCommand("git commit -m '<data>'", path)
	runCommand("git push", path)
}
