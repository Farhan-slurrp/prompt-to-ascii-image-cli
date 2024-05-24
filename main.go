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
@@@@@@@@@@@@@=:::-*@@- .:*@@@@@@@@@@@@***%@@%****@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@+  @@@@= :@@#  =##****+%% -@@@@# -@@@@@#..*@@@@%. =@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@+  @@@@= .@@@  #@@@@@@@@# +@@@@* -@@@@%.:%@@@@@@#  #@@@@@@@@@@@@@@
@@@@@@@@@@@@@@+  ::::  .@@%  --==+**@@* +@@@@+ =@@@@* .@@@@@@@@  +@@@@@@@@@@@@@@
%@@@@@@@@@@@@@=  @@@@: .@@%  ###****@@+ *@@@@+ =@@@@+ .@@@@@@@%  +@@@@@@@@@@@@@@
%@@@@@@@@@@@@@=  @@@@- .@@%  @@@@@@@@@+ *@@@@* +@@@@#  %@@@@@@+  #@@@@@@@@@@@@@@
%@@@@@@@@@@@@@= :@@@@+ .@@@  @@@@@@@@%# +@@@%* +@@@@%. :%@@@@+ .*@@@@@@@@@@@@@@@
%@@@@@@@@@@@@@= -@@@@*  @@@- .::::--.=.  .:: .  .:::..#-..::.:=%@@@@@@@@@@@@@@%%


WELCOME TO ASCII IMAGE GENERATOR!!
ENTER YOUR PROMPT TO GENERATE YOUR IMAGE.
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
