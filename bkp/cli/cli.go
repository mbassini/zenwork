package cli

import (
	"flag"
	"fmt"
	"os"
)

func Run(args []string) {
	switch true {
	case args[0] == "list" || args[0] == "l":
		fmt.Println("Lists command")
		handleListCommand(args[1:])
	case args[0] == "task" || args[0] == "t":
		fmt.Println("Tasks command")
		//handleTaskCommand(args[1:])
	default:
		fmt.Println("Unknown command. Use one of: 'list', 'l', 'task' or 't'.")
		os.Exit(1)
	}
}

func handleListCommand(args []string) {
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	var addFlag *string
	listCmd.StringVar(addFlag, "add", "", "Add a new list")

	fmt.Println(addFlag)
	fmt.Println(args)
	fmt.Println(len(args))
}
