package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// aliasCmd represents the alias command
var aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "creates an alias for tablespoon.",
	Run: func(cmd *cobra.Command, args []string) {
		shorthand, _ := cmd.Flags().GetString("shorthand")
		if shorthand == "" {
			shorthand = "tbsp"
		}

		shellPath := os.Getenv("SHELL")

		shell := strings.Split(shellPath, "/")[len(strings.Split(shellPath, "/"))-1]

		if shell == "bash" || shell == "sh" {
			fmt.Printf("function %s () {\n  $TBSP_CMD=$(\n    echo tablespoon $@\n  ) && eval $TBSP_CMD\n}", shorthand)
		} else if shell == "zsh" {
			fmt.Printf("%s () {{\n  $TBSP_CMD=$(\n    echo tablespoon $@\n  ) && eval $TBSP_CMD\n}}", shorthand)
		} else if shell == "fish" {
			fmt.Printf("function %s\n  tablespoon $argv\nend", shorthand)
		} else if shell == "tsch" || shell == "csh" {
			fmt.Printf("alias %s 'tablespoon \\!*'", shorthand)
		} else {
			// use windows protocol
			fmt.Printf("function %s {\n  tablespoon $args\n}", shorthand)
		}

	},
}

func init() {
	rootCmd.AddCommand(aliasCmd)

	aliasCmd.Flags().StringP("shorthand", "s", "", "add a shorthand for the tablespoon command")
}
