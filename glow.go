/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"glow/cmd"

	"github.com/fatih/color"
)

func main() {
	color.Red("Welcome to the jungle!")
	text := color.BlueString("Hallucinate when you call my name!")
	fmt.Println(text)
	cmd.Execute()
}
