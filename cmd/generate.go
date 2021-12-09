package cmd

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a commit message.",
	Run: func(cmd *cobra.Command, args []string) {
		full, _ := cmd.Flags().GetBool("full")
		ncomment, _ := cmd.Flags().GetBool("no-comment")
		content, err := exec.Command("git", "diff", "--staged", "--numstat").Output()
		if err != nil {
			pterm.Error.Println(err)
		}

		out := strings.Fields(string(content))
		message, file, short, files, diffs := rules(out, ncomment)

		fmt.Printf("%s(%s): %s\n\n\n", message, file, short)

		if full {
			for f := range files {
				fmt.Println("-", files[f], "-", diffs[f], "changes")
			}
			username, usernameErr := exec.Command("git", "config", "user.name").Output()
			email, emailErr := exec.Command("git", "config", "user.email").Output()

			if usernameErr != nil {
				pterm.Error.Println(usernameErr)
			}
			if emailErr != nil {
				pterm.Error.Println(emailErr)
			}

			fmt.Printf("\nAuthored-by: %s <%s>\n", strings.Fields(string(username))[0], strings.Fields(string(email))[0])
		}
		pterm.Success.Println("Command Successfully Executed")
	},
}

func rules(input []string, ncomment bool) (message string, file string, short string, files []string, diffs []int) {
	var adds int
	var dels int

	for i := range input {
		deletions, _ := strconv.Atoi(input[i])
		pattern := regexp.MustCompile("0|[1-9][0-9]*")
		if (i+1)%3 != 0 {
			if i == 0 { // initialize all the additions + deletions and store
				del, _ := strconv.Atoi(input[i+1])
				adds = deletions // not actually deletions in this case
				dels = del
			}
			if i >= 1 {
				if pattern.MatchString(input[i-1]) {
					additions, _ := strconv.Atoi(input[i-1])
					changes := deletions + additions
					diffs = append(diffs, []int{changes}...)
					files = append(files, []string{input[i+1]}...)
					if additions+deletions > adds+dels {
						adds = additions
						dels = deletions
					}
				}
			}
		}
	}

	selected := []string{""} // initializing value

	for n := range diffs {
		s, _ := strconv.Atoi(selected[0])
		if n == 0 {
			selected = []string{strconv.Itoa(diffs[n]), files[n]}
		} else if diffs[n] > s {
			selected = []string{strconv.Itoa(diffs[n]), files[n]}
		}
	}

	prompt := promptui.Select{
		Label: fmt.Sprintf("Select type for %s", selected[1]),
		Items: []string{"feat", "fix", "docs", "style", "refactor",
			"test", "chore"},
	}

	_, message, err := prompt.Run()

	if err != nil {
		pterm.Error.Println(err)
	}

	t := []string{""}

	if len(strings.Split(selected[1], "/")) <= 2 {
		file = selected[1]
	} else {
		t = []string{strings.Split(selected[1], "/")[len(strings.Split(selected[1], "/"))-1], strings.Split(selected[1], "/")[len(strings.Split(selected[1], "/"))-2]}
		file = fmt.Sprintf("%s/%s", t[0], t[1])
	}

	wordDiffs, diffErr := exec.Command("git", "diff", "--word-diff=porcelain", file).Output()
	if diffErr != nil {
		pterm.Error.Println(diffErr)
	}

	var in string

	//tbsp: add new `--no-comment flag`
	if ncomment {
		userShort := promptui.Prompt{
			Label:   fmt.Sprintf("What was changed in %s?", file),
			Default: in,
		}

		shortened, shortErr := userShort.Run()
		if shortErr != nil {
			pterm.Error.Println(shortErr)
		}

		short = shortened
	} else {
		if len(strings.Split(string(wordDiffs), "tbsp: ")) < 2 {
			userShort := promptui.Prompt{
				Label:   fmt.Sprintf("What was changed in %s?", file),
				Default: in,
			}

			shortened, shortErr := userShort.Run()
			if shortErr != nil {
				pterm.Error.Println(shortErr)
			}

			short = shortened
		} else {
			short = strings.Split(strings.Split(string(wordDiffs), "tbsp: ")[1], "\n")[0]
		}
	}
	return
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().BoolP("full", "f", false, "full length commit")
	commitCmd.Flags().BoolP("no-comment", "c", false, "prompt user for short description")
}
