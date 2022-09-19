package sync

import (
	"fmt"

	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new repository to sync.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Breno")
	},
}

func init() {
	// VideoCmd.Flags().StringVarP(&urlPath, "url", "u", "", "The url to ping")
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

func addRepo()
