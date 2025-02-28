package main

import (
	"fmt"
	"os"

	"github.com/mbassini/zenwork/cli"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: zenwork <command> [<args>]")
		fmt.Println("Commands: list, l, task, t")
		os.Exit(1)
	}
	cli.Run(os.Args[1:])
}
