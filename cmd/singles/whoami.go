package singles

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
var WhoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Tells you who am I.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(iAmHim())
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

func iAmHim() string {
	return `
	I am the Glow CLI tool, made for everything you need (and everything you don't need too!) 

	I was created by the beautiful (and single) Breno Nogueira, a brazilian student who 
	should really focus in one programming language and finish his projects.

	If you really use this tool in your day-to-day work, I was programmed to thank you, as my 
	creator just made me out of pure boredom and his obssession with automating things that
	are better not automated. 
	
	Also, I'm named upon the best party in the best campus of the best university ever. 
	I am a computer, I just work with facts. 
	
	If you want to visit my home to find help, report a misconduct in my actions (you humans call it bugs),
	or want to help me grow by feeding me more code, here's the address:

	https://github.com/RexynyN/glow

	If you want to find my creator and shame him for his rank in competitive games and his inability to
	finish any of his writing projects, or just to find a new programming friend, here's how to find him:

	Twitter: @breno_nogs
	Instagram: @breno.nogs
	GitHub: RexynyN
	Discord: RexynyN#
	
	He really appreciate anyone who enjoys his quirky words that does stuff on the computer, 
	so throw a thank you text at him, it would make his day! 

	Beep boop, pleased to meet you!
	Glow CLI Tool
	`
}
