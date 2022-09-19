/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package file

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	urlPath string
	dir     bool
)

var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename a file or a directory of files using various utilities.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print: " + strings.Join(args, " "))
		con, _ := cmd.Flags().GetString("url")
		fmt.Print(con)
	},
}

func init() {
	FileCmd.AddCommand(renameCmd)

	renameCmd.Flags().StringVarP(&urlPath, "url", "u", ".", "The url to ping")
	renameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// if err := VideoCmd.MarkFlagRequired("url"); err != nil {
	// 	fmt.Println(err)
	// }

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
