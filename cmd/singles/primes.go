package singles

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
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

		primes := gimmePrimes(threshold)

		json, _ := cmd.Flags().GetBool("json")
		txt, _ := cmd.Flags().GetBool("txt")
		if json {
			saveToJson(primes)
		}

		if txt {
			saveToTxt(primes)
		}

		if !json && !txt {
			for _, prime := range primes {
				fmt.Println(prime)
			}
		}
	},
}

func init() {
	PrimesCmd.Flags().BoolP("json", "j", false, "Sets the output of primes to a json file")
	PrimesCmd.Flags().BoolP("txt", "t", false, "Sets the output of primes to a txt file")
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
	primeBytes, err := json.Marshal(primes)
	if err != nil {
		panic(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filepath.Join(cwd, "primes.json"), primeBytes, 0644)
	if err != nil {
		panic(err)
	}
}

func saveToTxt(primes []int) {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	f, err := os.Create(filepath.Join(cwd, "primes.txt"))
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	val := ""
	for _, prime := range primes {
		val += fmt.Sprint(prime) + "\n"
	}
	data := []byte(val)

	_, err2 := f.Write(data)
	if err2 != nil {
		log.Fatal(err2)
	}
}
