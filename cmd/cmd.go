package cmd

import (
	"os"
	"os/exec"
	"strings"
)

// Command Type
type Command struct {
	Bin   string
	Args  []string
	Dir   string
	Envir []string
}

// Exec command
func Exec(c Command) error {
	cm := exec.Command(c.Bin, c.Args...)
	cm.Dir = c.Dir
	cm.Env = append(os.Environ(), c.Envir...)
	cm.Stderr = os.Stderr
	cm.Stdin = os.Stdin
	cm.Stdout = os.Stdout
	return cm.Run()
}

// Cout string output of command
func Cout(cc []string) ([]string, error) {
	c, err := exec.Command(cc[0], cc[1:]...).Output()
	if err != nil {
		return nil, err
	}

	return strings.Split(string(c[:]), "\n"), nil
}
