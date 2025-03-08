package cmd

import (
	"fmt"
	"slices"
	"strings"
)

func Run(args []string) {
	fmt.Printf("[ARGS]: %v\n", args)
	// cleanArgs(&args)
	inputCommand, positionArg, options, err := parseCommand(args)
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Printf("[Command]: %v\n", inputCommand)
	fmt.Printf("[PosArgs]: %v\n", positionArg)
	fmt.Printf("[Length PosArgs]: %v\n", len(positionArg))
	fmt.Printf("[Options]: %v\n", options)

	command, exists := getCommands()[inputCommand]
	if !exists {
		fmt.Printf("Unknown command %v\n", inputCommand)
	}

	if len(options) > 0 {
		err = validateOptions(options, command.allowedOptions)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}
	fmt.Printf("[Options Length]: %v", command)
	fmt.Printf("[Options Length]: %v", len(options))

	err = command.callback(positionArg, options)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

//	func cleanArgs(args *[]string) {
//		for i, arg := range *args {
//			(*args)[i] = strings.ToLower(strings.TrimSpace(arg))
//		}
//	}

func parseCommand(args []string) (command string, positionArg string, options map[string]string, err error) {
	if len(args) == 0 {
		usageStr := fmt.Sprintln("Usage: zenwork <command> [<args>]")
		helpStr := fmt.Sprintln("See zenwork help")
		return "", "", nil, fmt.Errorf("%s%s", usageStr, helpStr)
	}

	command = args[0]
	options = make(map[string]string)

	for i := 1; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "--") {
			key := strings.TrimPrefix(arg, "--")
			if i+1 < len(args) && !strings.HasPrefix(args[i+1], "--") {
				options[key] = args[i+1]
				i++
			} else {
				options[key] = "true"
			}
		} else {
			if positionArg == "" {
				positionArg = arg
			} else {
				return "", "", nil, fmt.Errorf("only one position argument is allowed\n")
			}
		}
	}

	return command, positionArg, options, nil
}

type cliCommand struct {
	name           string
	description    string
	allowedOptions []string
	callback       func(positionArg string, options map[string]string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"add": {
			name:           "add",
			description:    "Add a new task",
			allowedOptions: []string{"description", "status", "priority"},
			callback:       commandAdd,
		},
	}
}

func validateOptions(options map[string]string, allowedOptions []string) error {
	for option, _ := range options {
		if !slices.Contains(allowedOptions, option) {
			return fmt.Errorf("%s is not an allowed option", option)
		}
	}
	return nil
}
