/*
Utilities to work with files, like rename, sort, etc.
Copyright © 2022 Breno Nogueira breno.s.nogueira@hotmail.com
*/
package file

import (
	"fmt"
	"glow/common"
	"io/fs"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	toName      string
	iterateEnum []string = []string{"number", "letter", "mixed"}
)

var renameCmd = &cobra.Command{
	Use:     "rename",
	Short:   "Rename a file or a directory of files using various utilities.",
	Example: "glow file rename --extensions \"mp4,png\" --startsWith \"abc\" --endsWith \"123\" --replace \"abc\" --to \"\"\nglow file rename --iterate number --to \"BOGUS VOLUME {}\" --toTitle",
	Long:    ``,
	Run:     run,
}

func run(cmd *cobra.Command, args []string) {
	// Tools
	cwd, _ := os.Getwd()
	if revert, _ := cmd.Flags().GetBool("revert"); revert {
		revertLastChanges(cwd)
	}

	files := readFiles(cwd)
	// Selectors
	if contains, _ := cmd.Flags().GetString("contains"); contains != "" {
		files = filterFiles(files, func(name string) bool { return strings.Contains(name, contains) })
	}

	if startsWith, _ := cmd.Flags().GetString("startsWith"); startsWith != "" {
		files = filterFiles(files, func(name string) bool { return strings.HasPrefix(name, startsWith) })
	}

	if endsWith, _ := cmd.Flags().GetString("endsWith"); endsWith != "" {
		files = filterFiles(files, func(name string) bool {
			name = strings.Split(name, ".")[0]
			return strings.HasSuffix(name, endsWith)
		})
	}

	if extensions, _ := cmd.Flags().GetString("extensions"); extensions != "" {
		extensions = strings.ReplaceAll(extensions, ".", "")
		extensionsSlice := strings.FieldsFunc(extensions, func(c rune) bool { return c == ',' })

		files = filterFiles(files, func(name string) (sentinel bool) {
			sentinel = false
			for _, extension := range extensionsSlice {
				if strings.HasSuffix(name, extension) {
					sentinel = true
					break
				}
			}
			return
		})
	}

	// Operations
	changedFiles := make([]string, len(files))
	toValue, _ := cmd.Flags().GetString("to")
	if random, _ := cmd.Flags().GetBool("random"); random {
		changedFiles = randomizeFiles(len(files))
	} else if iterate, _ := cmd.Flags().GetString("iterate"); iterate != "" && toValue != "" {
		if !checkEnum(iterate, iterateEnum) {
			fmt.Println("'--iterate' flag does not contain a valid option.")
			return
		}

		if !strings.Contains(toValue, "{}") {
			fmt.Println("'--to' flag does not contain the '{}' token to be replaced.")
			return
		}

		// TODO: Parei aqui!
		changedFiles = getFileNames(files)
		common.NaturalSort(changedFiles)

		originalFiles := make([]string, len(changedFiles))
		copy(originalFiles, changedFiles)

		iterateFiles(changedFiles, toValue, iterate)

	}

	// String Cases
	if toUpper, _ := cmd.Flags().GetBool("toUpper"); toUpper {
		changedFiles = filesToCase(changedFiles, strings.ToUpper)
	} else if toLower, _ := cmd.Flags().GetBool("toLower"); toLower {
		changedFiles = filesToCase(changedFiles, strings.ToLower)
	} else if toTitle, _ := cmd.Flags().GetBool("toTitle"); toTitle {
		changedFiles = filesToCase(changedFiles, strings.ToTitle)
	}

	// renameFiles(files, cwd, changedFiles)
}

func iterateFiles(changedFiles []string, toValue string, iterate string) []string {
	for index := range changedFiles {
		changedFiles[index] = strings.ReplaceAll(toValue, "{}", fmt.Sprint(index+1))
	}

	return changedFiles
}

func getFileNames(files []fs.FileInfo) (filenames []string) {
	for _, file := range files {
		filenames = append(filenames, file.Name())
	}
	return
}

func checkEnum(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func revertLastChanges(cwd string) {
	panic("unimplemented")
}

func filesToCase(files []string, toCase func(string) string) (filenames []string) {
	for _, file := range files {
		filenames = append(filenames, toCase(file))
	}
	return
}

func filterFiles(files []fs.FileInfo, filter func(string) bool) (filtered []fs.FileInfo) {
	for _, file := range files {
		if filter(file.Name()) {
			filtered = append(filtered, file)
		}
	}

	return filtered
}

func init() {
	FileCmd.AddCommand(renameCmd)

	// Selectors
	renameCmd.Flags().String("contains", "", "Selects all files which contains the given literal.")
	renameCmd.Flags().String("startsWith", "", "Selects all files which starts with the given literal.")
	renameCmd.Flags().String("endsWith", "", "Selects all files which ends with the given literal (excluding the file extension).")
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

func randomizeFiles(size int) (changedFiles []string) {
	for i := 0; i < size; i++ {
		changedFiles = append(changedFiles, randomName())
	}

	return
}

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

func renameFiles(files []string, path string, nameGetter func() string) {
	for _, file := range files {

		newName := nameGetter()
		splot := strings.Split(file, ".")
		fmt.Println(file + " -> " + newName)

		err := os.Rename(path+"\\"+file, path+"\\"+newName+"."+splot[1])
		if err != nil {
			panic(err)
		}
	}
}

func readFiles(path string) (files []os.FileInfo) {
	dirFiles, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range dirFiles {
		if !file.IsDir() {
			files = append(files, file)
		}
	}

	return files
}
