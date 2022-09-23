/*
Copyright © 2022 Breno Nogueira breno.s.nogueira@hotmail.com
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
	toName string
)

var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename a file or a directory of files using various utilities.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if random, _ := cmd.Flags().GetBool("random"); random {
			cwd, _ := os.Getwd()
			renameFiles(cwd, randomName)
			return
		}

	},
}

func init() {
	FileCmd.AddCommand(renameCmd)

	// Selectors
	renameCmd.Flags().String("contains", "", "Selects all files which contains the given literal.")
	renameCmd.Flags().String("startsWith", "", "Selects all files which starts with the given literal.")
	renameCmd.Flags().String("endsWith", "", "Selects all files which ends with the given literal.")
	renameCmd.Flags().String("extensions", "", "Selects files by the given pool of file extensions. (separated by comma)")

	// Operations
	renameCmd.Flags().String("iterate", "", "Type of value to append to '--to' flag (number, letter, mixed), '--to' must have {} to be replaced by the value.")
	renameCmd.Flags().BoolP("random", "r", false, "Renames all selected files to a random string of characters and numbers.")
	renameCmd.Flags().String("replace", "", "Replace all instances of the given expression, if found. (--to flag is required)")
	renameCmd.Flags().String("replaceOnce", "", "Replace first instance of the given expression, if found. (--to flag is required)")
	renameCmd.Flags().String("to", "", "The value to replace, or the name to be set.")

	// String Cases
	renameCmd.Flags().Bool("toUpper", false, "Flips all selected files to Upper Case (after all replace and rename operations)")
	renameCmd.Flags().Bool("toLower", false, "Flips all selected files to Lower Case (after all replace and rename operations)")
	renameCmd.Flags().Bool("toTitle", false, "Flips all selected files to Title Case (after all replace and rename operations)")

	// Tools
	renameCmd.Flags().Bool("revert", false, "Revert the last rename operation in the current folder, if any.")
}

// glow file rename --extensions "mp4,png" --startsWith "abc" --endsWith "123" --replace "0a--" --to ""
// glow file rename --iterate number --to "Bogus Volume {}"

func randomName() string {
	strong := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	alfabet := []rune(strong)
	name := make([]rune, 35)
	for i := 0; i < 35; i++ {
		name[i] = alfabet[rand.Intn(len(strong))]
	}

	return string(name)
}

func replaceName(filename string, expression string, replacer string) string {
	return strings.ReplaceAll(filename, expression, replacer)
}

func renameFiles(path string, nameGetter func() string) {
	for _, file := range readFiles(path) {
		if file.IsDir() {
			continue
		}
		newName := nameGetter()
		splot := strings.Split(file.Name(), ".")
		fmt.Println(file.Name() + " -> " + newName)

		err := os.Rename(path+"\\"+file.Name(), path+"\\"+newName+"."+splot[1])
		if err != nil {
			panic(err)
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
