package cmd

import (
	"fmt"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

// alias represents the config command
var alias = &cobra.Command{
	Use:   "alias",
	Short: "creates an alias for tablespoon.",
	Run: func(cmd *cobra.Command, args []string) {
		shorthand, _ := cmd.Flags().GetString("shorthand")
		if shorthand == "" {
			shorthand = "tbsp"
		}

		shellCmd, shellErr := exec.Command("echo $SHELL").Output()
		if shellErr != nil {
			pterm.Error.Println("500: Unable to identify shell; ", shellErr)
			return
		}

		shell := strings.Split(string(shellCmd), "/")[len(strings.Split(string(shellCmd), "/"))-1]

		if shell == "bash" || shell == "sh" {
			fmt.Printf("function %s () {{\n  $TBSP_CMD=$(\n    echo tablespoon $@\n  ) && eval $TBSP_CMD\n}}", shorthand)
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
	rootCmd.AddCommand(alias)

	alias.Flags().StringP("shorthand", "s", "", "add a shorthand for the tablespoon command")
}
