package cmd

import (
	"errors"
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
		unstaged, _ := cmd.Flags().GetBool("unstaged")
		ncomment, _ := cmd.Flags().GetBool("no-comment")
		selectFlag, _ := cmd.Flags().GetString("select")
		coauth, _ := cmd.Flags().GetString("co-author")

		var stage string
		if unstaged {
			stage = ""
		} else {
			stage = "--staged"
		}
		content, err := exec.Command("git", "diff", stage, "--numstat").Output()

		//tbsp: exit if error
		if err != nil {
			pterm.Error.Println("Error T0:", err)
			return
		}

		out := strings.Fields(string(content))
		message, file, short, files, diffs, rulesErr := rules(out, ncomment, selectFlag)

		//tbsp: allow for better error handling
		if rulesErr != nil {
			pterm.Error.Println(rulesErr)
			return
		}

		fmt.Printf("%s(%s): %s\n", message, file, short)

		if full {
			fmt.Printf("\n\n")
			for f := range files {
				fmt.Println("-", files[f], "-", diffs[f], "changes")
			}
			username, usernameErr := exec.Command("git", "config", "user.name").Output()
			email, emailErr := exec.Command("git", "config", "user.email").Output()

			if usernameErr != nil {
				pterm.Error.Println("Error T2:", usernameErr)
				return
			}
			if emailErr != nil {
				pterm.Error.Println("Error T2:", emailErr)
				return
			}

			fmt.Printf("\nAuthored-by: %s <%s>\n", strings.Fields(string(username))[0], strings.Fields(string(email))[0])

			//tbsp: allow for `--co-author` flag, syntax = <name>:<email@email.com>
			if coauth != "" {
				var addr string
				if len(strings.Split(coauth, ":")) > 1 {
					addr = strings.Split(coauth, ":")[1]
				} else {
					addr = coauth + "@users.noreply.github.com"
				}
				fmt.Printf("\nCo-Authored-by: %s <%s>\n", strings.Split(coauth, ":")[0], addr)
			}
		}
		pterm.Success.Println("Command Successfully Executed")
	},
}

func rules(input []string, ncomment bool, selectFlag string) (message string, file string, short string, files []string, diffs []int, rulesErr error) {
	var adds int
	var dels int
	rulesErr = nil

	//tbsp: Add error handling if no changes
	if len(input) <= 1 {
		rulesErr = errors.New("Error T0: No differences detected.")
		return
	}

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

	//tbsp: add `--select` flag to choose file
	if selectFlag == "" {
		for n := range diffs {
			s, _ := strconv.Atoi(selected[0])
			if n == 0 {
				selected = []string{strconv.Itoa(diffs[n]), files[n]}
			} else if diffs[n] > s {
				selected = []string{strconv.Itoa(diffs[n]), files[n]}
			}
		}
	} else {
		for n := range files {
			if files[n] == selectFlag {
				selected = []string{strconv.Itoa(diffs[n]), files[n]}
			}
		}
	}

	if len(selected) < 2 {
		rulesErr = errors.New("Error T6: File not found.")
		return
	}

	prompt := promptui.Select{
		Label: fmt.Sprintf("Select type for %s", selected[1]),
		Items: []string{"feat", "fix", "docs", "style", "refactor",
			"test", "chore"},
	}

	_, message, err := prompt.Run()

	if err != nil {
		rulesErr = errors.New("Error T5: " + err.Error())
		return
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
		rulesErr = errors.New("Error T3: " + diffErr.Error())
		return
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
			pterm.Error.Println("Error T4:", shortErr)
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
				pterm.Error.Println("Error T4:", shortErr)
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
	generateCmd.Flags().BoolP("unstaged", "u", false, "generate message for all changed files")
	generateCmd.Flags().BoolP("no-comment", "n", false, "prompt user for short description")
	generateCmd.Flags().StringP("select", "s", "", "choose file to showcase in short commit message")
	generateCmd.Flags().StringP("co-author", "c", "", "add co-author to commit")
}
