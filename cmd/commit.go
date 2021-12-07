package cmd

import (
	"fmt"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Generates a commit message & commits it.",
	Run: func(cmd *cobra.Command, args []string) {
		content, err := exec.Command("git", "diff", "--numstat").Output()
		if err != nil {
			pterm.Error.Println(err)
		}

		out := strings.Fields(string(content))
		message, file, verb, _, _ := rules(out)
		input := fmt.Sprintf("%s(%s): %s", message, file, verb)

		commitOut, commitErr := exec.Command("git", "commit", "-m", fmt.Sprintf("'%s'", input)).Output()

		// if there is an error with our execution handle it here
		if err != nil {
			pterm.Error.Println(commitErr)
		}

		output := string(commitOut[:])

		fmt.Println(output)
		pterm.Success.Println("Command Successfully Executed")
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
