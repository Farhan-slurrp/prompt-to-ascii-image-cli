package main

import (
	"bufio"
	"farhan-slurrp/ascii-image-cli/internal"
	"fmt"
	"os"
	"strings"

	"github.com/ttacon/chalk"
)

func StringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, chalk.Bold.TextStyle(label)+"\n")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func main() {
	fmt.Println(chalk.Green.Color(`
=============================================
          WELCOME TO ASCII ART GENERATOR
=============================================

Transform your images into stunning ASCII art with ease!

Whether you're looking to create unique text-based art, share fun images with friends, 
or just explore the world of ASCII, this tool makes it simple for you.

-------------------------------------------------------------
Enter your prompt to get started!!!
-------------------------------------------------------------
	`))
	prompt := StringPrompt("Enter the prompt:")
	fmt.Println(chalk.Blue.Color("\nGenerating image..."))
	image, err := internal.GenerateImage(prompt)

	if err != nil {
		fmt.Println(chalk.Red.Color("Error:"), err)
		panic(err)
	}

	internal.PrintAsciiImage(image)

}
