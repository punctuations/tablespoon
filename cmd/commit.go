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
		unstaged, _ := cmd.Flags().GetBool("unstaged")
		ncomment, _ := cmd.Flags().GetBool("no-comment")
		selectFlag, _ := cmd.Flags().GetString("select")
		coauth, _ := cmd.Flags().GetString("co-author")

		//tbsp: add optional `--unstaged` flag
		var inf []byte
		var differentErr error
		if unstaged {
			content, err := exec.Command("git", "diff", "--numstat").Output()
			inf = content
			differentErr = err
		} else {
			content, err := exec.Command("git", "diff", "--staged", "--numstat").Output()
			inf = content
			differentErr = err
		}

		if differentErr != nil {
			pterm.Error.Println("500: An error occurred when running git diff; ", differentErr.Error())
			return
		}

		out := strings.Fields(string(inf))
		message, file, short, files, diffs, rulesErr := rules(out, unstaged, ncomment, selectFlag)
		if rulesErr != nil {
			pterm.Error.Println(rulesErr)
			return
		}
		input := fmt.Sprintf("%s(%s): %s", message, file, short)
		desc := "\n\n" //tbsp: init desc var

		if full {
			for f := range files {
				desc = desc + fmt.Sprintf("- %s - %d changes\n", files[f], diffs[f])
			}
			username, usernameErr := exec.Command("git", "config", "user.name").Output()
			email, emailErr := exec.Command("git", "config", "user.email").Output()

			if usernameErr != nil {
				pterm.Error.Println("500: An error occurred while accessing global git config username; ", usernameErr.Error())
				return
			}
			if emailErr != nil {
				pterm.Error.Println("500: An error occurred while accessing global git config email; ", emailErr.Error())
				return
			}

			desc = desc + fmt.Sprintf("\nAuthored-by: %s <%s>\n", strings.Fields(string(username))[0], strings.Fields(string(email))[0])
			if coauth != "" {
				var addr string
				if len(strings.Split(coauth, ":")) > 1 {
					addr = strings.Split(coauth, ":")[1]
				} else {
					addr = coauth + "@users.noreply.github.com"
				}
				desc = desc + fmt.Sprintf("\nCo-Authored-by: %s <%s>\n", strings.Split(coauth, ":")[0], addr)
			}
		}

		if unstaged {
			_, resErr := exec.Command("git", "add", "*").Output()

			if resErr != nil {
				pterm.Error.Println("500: An error occurred when running git add; ", resErr.Error())
				return
			}
		}

		//#!: don't commit if a field is missing
		if input != "" || desc != "" {
			commitOut, commitErr := exec.Command("git", "commit", "-m", fmt.Sprintf("%s%s", input, desc)).Output()

			// if there is an error with our execution handle it here
			if commitErr != nil {
				pterm.Error.Println("500: An error occurred when running git commit; ", commitErr.Error())
				return
			}

			output := string(commitOut[:])

			fmt.Println(output)
			pterm.Success.Println("✨ Command Successfully Executed")
		} else {
			pterm.Error.Println("☕️ Empty fields in commit message")
		}
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)

	commitCmd.Flags().BoolP("full", "f", false, "full length commit")
	commitCmd.Flags().BoolP("unstaged", "u", false, "generate message for all changed files")
	commitCmd.Flags().BoolP("no-comment", "n", false, "prompt user for short description")
	commitCmd.Flags().StringP("select", "s", "", "choose file to showcase in short commit message")
	commitCmd.Flags().StringP("co-author", "c", "", "add co-author to commit")
}
