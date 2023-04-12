package utils

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func HandleError(err error, messages ...string) {
	PrintErrorAndExit(1, err, messages...)
}

func PrintErrorAndExit(exitCode int, err error, messages ...string) {
	printError(err)

	if len(messages) > 0 {
		for _, message := range messages {
			fmt.Println(message)
		}
	}

	supportMsg := fmt.Sprintf("\n\nIf this issue continues, get support at https://symeo.io")
	fmt.Fprintln(os.Stderr, supportMsg)

	os.Exit(exitCode)
}

func printError(e error) {
	color.New(color.FgRed).Fprintf(os.Stderr, "Hmm, we ran into an error: %v", e)
}
