package sync

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// Struc to represent a Git Repo
type Repo struct {
	Path   string `json:"path"`
	Branch string `json:"branch"`
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

func readRepos() (repos []Repo) {
	cwd, err := os.Executable()
	if err != nil {
		panic(err)
	}

	fileBytes, err := os.ReadFile(filepath.Join(cwd, "sync_repos.json"))
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &repos)
	if err != nil {
		panic(err)
	}

	return repos
}

func saveRepos(repos []Repo) {
	repoBytes, err := json.Marshal(repos)
	if err != nil {
		panic(err)
	}

	cwd, err := os.Executable()
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(filepath.Join(cwd, "sync_repos.json"), repoBytes, 0644)
	if err != nil {
		panic(err)
	}
}
