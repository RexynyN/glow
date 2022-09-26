package singles

import (
	"fmt"
	"glow/common"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
var ImageUriCmd = &cobra.Command{
	Use:   "imageuri",
	Short: "Makes operations with images and datauris, for transforming images in text.",
	Long:  ``,
	Run:   run,
}

func init() {
	ImageUriCmd.Flags().BoolP("convert", "c", false, "Convert images to a data uri file (.txt or .json).")
	ImageUriCmd.Flags().BoolP("revert", "r", false, "Revert data uri files to images.")

}

// glow imageuri convert

func run(cmd *cobra.Command, args []string) {
	cwd := common.GetCwd()
	files := common.ReadFilesByExtension(common.GetCwd(), []string{"jpg", "jpeg", "jfif", "png", "webp"})

	for _, file := range files {
		datauri := bogusBinted(filepath.Join(cwd, file.Name()))
		fmt.Println(datauri)
	}
}

func bogusBinted(filepath string) string {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Couldn't load the '" + filepath + "', skipping. ")
		return ""
	}

	mimeType := http.DetectContentType(bytes)
	mimeType = strings.Split(mimeType, "/")[1]

	datauri := "data:image/" + mimeType + ";base64,"
	datauri += common.ToBase64(bytes)

	return datauri
}

func bogusBintedWeb(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Couldn't reach the '" + url + "' url, skipping. ")
		return ""
	}

	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Couldn't load '" + url + "', skipping. ")
		return ""
	}

	mimeType := http.DetectContentType(bytes)
	mimeType = strings.Split(mimeType, "/")[1]

	datauri := "data:image/" + mimeType + ";base64,"
	datauri += common.ToBase64(bytes)

	return datauri
}
