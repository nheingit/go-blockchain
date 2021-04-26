package main

import (
	"os"

	"github.com/nheingit/learnGo/cli"
)

func main() {
	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()

}
