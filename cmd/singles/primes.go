package singles

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"

	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
var PrimesCmd = &cobra.Command{
	Use:        "primes [threshold] [flags]",
	Example:    "glow primes 100 --json",
	Args:       cobra.MinimumNArgs(1),
	ArgAliases: []string{"threshold"},
	Short:      "Shows you a lot of prime numbers, as much as you need.",
	Long:       ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("cmd.Args: %v\n", cmd.Args)
		threshold, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("The threshold of primes is in the incorrect format, use an integer number.")
			return
		}

		for _, prime := range gimmePrimes(threshold) {
			fmt.Println(prime)
		}
	},
}

func init() {
	PrimesCmd.Flags().BoolP("json", "j", false, "Sets the output of primes to a json file")
	PrimesCmd.Flags().BoolP("txt", "t", false, "Sets the output of primes to a txt file")

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

func gimmePrimes(threshold int) []int {
	var number int = 3
	var i int = 1
	var guard bool = true
	primes := []int{2}
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

func saveToJson(primes []int) {
	videoBytes, err := json.Marshal(primes)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./videos-updated.json", videoBytes, 0644)
	if err != nil {
		panic(err)
	}
}

func saveToTxt() {

}
