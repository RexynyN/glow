/*
Utilities to work with files, like rename, sort, etc.
Copyright © 2022 Breno Nogueira breno.s.nogueira@hotmail.com
*/
package file

import (
	"encoding/json"
	"fmt"
	"glow/common"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

type operationValues struct {
	OriginalName string `json:"originalName"`
	ChangedName  string `json:"changedName"`
}
type renameOperation struct {
	Path       string            `json:"path"`
	Operations []operationValues `json:"values"`
}

var (
	iterateEnum []string = []string{"number", "letter", "mixed"}
	blacklist   []rune   = []rune{'/', '\\', ':', '?', '*', '<', '>', '|', '"'}
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
	cwd := common.GetCwd()
	if revert, _ := cmd.Flags().GetBool("revert"); revert {
		revertLastChanges()
		return
	}

	files := common.ReadFiles(cwd)
	// Selectors
	if contains, _ := cmd.Flags().GetString("contains"); contains != "" {
		files = filterFiles(files, func(name string) bool {
			return strings.Contains(common.GetPureFilename(name), contains)
		})
	}

	if startsWith, _ := cmd.Flags().GetString("startsWith"); startsWith != "" {
		files = filterFiles(files, func(name string) bool {
			return strings.HasPrefix(common.GetPureFilename(name), startsWith)
		})
	}

	if endsWith, _ := cmd.Flags().GetString("endsWith"); endsWith != "" {
		files = filterFiles(files, func(name string) bool {
			return strings.HasSuffix(common.GetPureFilename(name), endsWith)
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
	changedFiles := getFileNames(files)
	common.NaturalSort(changedFiles)
	originalFiles := make([]string, len(changedFiles))
	copy(originalFiles, changedFiles)

	toValue, _ := cmd.Flags().GetString("to")
	if !checkForbiddenRunes(toValue) {
		fmt.Println("New file name cannot contain the following characters: / \\ : ? * < > | \"")
		return
	}

	replace, _ := cmd.Flags().GetString("replace")
	if random, _ := cmd.Flags().GetBool("random"); random {
		changedFiles = randomizeFiles(changedFiles)
	} else if iterate, _ := cmd.Flags().GetString("iterate"); iterate != "" && toValue != "" {
		if !checkEnum(iterate, iterateEnum) {
			fmt.Println("'--iterate' flag does not contain a valid option.")
			return
		}

		if !strings.Contains(toValue, "{}") {
			fmt.Println("'--to' flag does not contain the '{}' token to be replaced.")
			return
		}

		changedFiles = iterateFiles(changedFiles, toValue, iterate)
	} else if replaceOnce, _ := cmd.Flags().GetString("replaceOnce"); replaceOnce != "" || replace != "" {
		if replaceOnce != "" {
			changedFiles = replaceFiles(changedFiles, replaceOnce, 1, toValue)
		} else {
			changedFiles = replaceFiles(changedFiles, replace, -1, toValue)
		}
	}

	// String Cases
	if toUpper, _ := cmd.Flags().GetBool("toUpper"); toUpper {
		changedFiles = filesToCase(changedFiles, strings.ToUpper)
	} else if toLower, _ := cmd.Flags().GetBool("toLower"); toLower {
		changedFiles = filesToCase(changedFiles, strings.ToLower)
	} else if toTitle, _ := cmd.Flags().GetBool("toTitle"); toTitle {
		changedFiles = filesToCase(changedFiles, strings.ToTitle)
	}

	if changesWereMade(originalFiles, changedFiles) {
		makeRecord(originalFiles, cwd, changedFiles)
		renameFiles(originalFiles, cwd, changedFiles)
	} else {
		fmt.Println("No changes were made. If that's not intentional, check your filters and try again.")
	}
}

func init() {
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

func changesWereMade(originalFiles []string, changedFiles []string) (sentinel bool) {
	sentinel = false
	for index := range originalFiles {
		if originalFiles[index] != changedFiles[index] {
			sentinel = true
			break
		}
	}

	return sentinel
}

func saveRecords(records renameOperation) {
	recordsBytes, err := json.Marshal(records)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filepath.Join(common.GetExePath(), "rename_operations.json"), recordsBytes, 0644)
	if err != nil {
		panic(err)
	}
}

func makeRecord(originalFiles []string, cwd string, changedFiles []string) {
	records := make([]operationValues, len(originalFiles))

	for index := range records {
		records[index] = operationValues{OriginalName: originalFiles[index], ChangedName: changedFiles[index]}
	}

	saveRecords(renameOperation{
		Path:       cwd,
		Operations: records,
	})
}

func replaceFiles(files []string, literal string, numChanges int, replace string) []string {
	for index := range files {
		files[index] = strings.Trim(strings.Replace(common.GetPureFilename(files[index]), literal, replace, numChanges)+common.GetFileExtension(files[index]), " ")
	}

	return files
}

func iterateFiles(files []string, toValue string, iterate string) []string {
	for index := range files {
		files[index] = strings.Trim(strings.ReplaceAll(toValue, "{}", fmt.Sprint(index+1))+common.GetFileExtension(files[index]), " ")
	}

	return files
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

func revertLastChanges() {
	reverter := renameOperation{}

	filebytes, err := os.ReadFile(filepath.Join(common.GetExePath(), "rename_operations.json"))
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(filebytes, &reverter)
	if err != nil {
		fmt.Println("An error occured while reading the renaming operations history.")
	}

	source := make([]string, len(reverter.Operations))
	modified := make([]string, len(reverter.Operations))
	for index, value := range reverter.Operations {
		modified[index] = value.OriginalName
		source[index] = value.ChangedName
	}

	renameFiles(source, reverter.Path, modified)
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

func randomizeFiles(files []string) (changedFiles []string) {
	for i := 0; i < len(files); i++ {
		changedFiles = append(changedFiles, randomName()+common.GetFileExtension(files[i]))
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

func renameFiles(files []string, path string, newFiles []string) {
	if len(files) != len(newFiles) {
		log.Fatal("The number of files in the list are not equal, cannot map old filenames to new filenames")
	}

	for index, file := range files {
		err := os.Rename(filepath.Join(path, file), filepath.Join(path, newFiles[index]))
		if err != nil {
			fmt.Println("An error occured while trying to rename the files: ", err)
		}
	}
}

func checkForbiddenRunes(name string) (sentinel bool) {
	sentinel = true
	for _, runer := range blacklist {
		if strings.ContainsRune(name, runer) {
			sentinel = false
			break
		}
	}
	return sentinel
}
