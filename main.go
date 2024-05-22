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
	 _   _      _ _            
	| | | | ___| | | ___ 
	| |_| |/ _ \ | |/ _ \
	|  _  |  __/ | | (_) |
	|_| |_|\___|_|_|\___/ 
	_______________________

	WELCOME TO ASCII IMAGE GENERATOR.
	ENTER YOUR PROMPT TO GENERATE ASCII IMAGE.
	`))
	prompt := StringPrompt("Enter the prompt:")

	image, err := internal.GenerateImage(prompt)

	if err != nil {
		fmt.Println(chalk.Red.Color("Error:"), err)
		return
	}

	fmt.Println(chalk.Green.Color("Image generated successfully."))
	fmt.Println(chalk.Green.Color("Image URL:"), image)
}
