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
		full, _ := cmd.Flags().GetBool("full")
		ncomment, _ := cmd.Flags().GetBool("no-comment")
		content, err := exec.Command("git", "diff", "--staged", "--numstat").Output()
		if err != nil {
			pterm.Error.Println(err)
		}

		out := strings.Fields(string(content))
		message, file, verb, files, diffs := rules(out, ncomment)
		input := fmt.Sprintf("%s(%s): %s", message, file, verb)
		desc := "" // init

		if full {
			for f := range files {
				desc = desc + fmt.Sprintf("- %s - %d changes\n", files[f], diffs[f])
			}
			username, usernameErr := exec.Command("git", "config", "user.name").Output()
			email, emailErr := exec.Command("git", "config", "user.email").Output()

			if usernameErr != nil {
				pterm.Error.Println(usernameErr)
			}
			if emailErr != nil {
				pterm.Error.Println(emailErr)
			}

			desc = desc + fmt.Sprintf("\nAuthored-by: %s <%s>\n", strings.Fields(string(username))[0], strings.Fields(string(email))[0])
		}

		commitOut, commitErr := exec.Command("git", "commit", "-m", fmt.Sprintf("'%s\n\n%s'", input, desc)).Output()

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

	commitCmd.Flags().BoolP("full", "f", false, "full length commit")
	commitCmd.Flags().BoolP("no-comment", "c", false, "prompt user for short description")
}
