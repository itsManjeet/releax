package msg

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

var (
	whiteBold = color.New(color.FgWhite).Add(color.Bold).SprintFunc()
	redBold   = color.New(color.FgRed).Add(color.Bold).SprintFunc()
	blueBold  = color.New(color.FgBlue).Add(color.Bold).SprintFunc()
	greenBold = color.New(color.FgGreen).Add(color.Bold).SprintFunc()
	white     = color.New(color.FgWhite).SprintFunc()
	red       = color.New(color.FgRed).SprintFunc()
)

func Error(a ...interface{}) (int, error) {
	return fmt.Fprintf(os.Stdout, "%s %s %s\n", whiteBold("=>"), redBold("err:"), red(a...))
}

func Notice(a ...interface{}) (int, error) {
	return fmt.Fprintf(os.Stdout, "%s %s %s\n", whiteBold("=>"), blueBold("notice:"), white(a...))
}

func Write(a ...interface{}) (int, error) {
	return fmt.Fprintf(os.Stdout, "%s\n", white(a...))
}

func Process(a ...interface{}) (int, error) {
	return fmt.Fprintf(os.Stdout, "%s %s %s\n", whiteBold("=>"), greenBold("process:"), white(a...))
}
