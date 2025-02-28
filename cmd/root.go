package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zenwork",
	Short: "A CLI task manager for developers",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
