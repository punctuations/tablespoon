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
		content, err := exec.Command("git", "diff", "--numstat").Output()
		if err != nil {
			pterm.Error.Println(err)
		}

		out := strings.Fields(string(content))
		message, file, verb, adds, dels := rules(out)

		pterm.Success.Printf("%s(%s): %s\n", message, file, verb)

		pterm.Success.Println(file, " + ", adds, " - ", dels)
		//format, _ := cmd.Flags().GetString("format")
		//letters := pterm.NewLettersFromString(time.Now().Format(format))
		//pterm.DefaultBigText.WithLetters(letters).Render()
	},
}

func rules(input []string) (message string, file string, verb string, adds int, dels int) {
	var diffs []int
	var files []string

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
			//fmt.Println(outInt)
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

	switch {
	case adds-dels == 0 && message == "fix":
		verb = "fix"
	case adds-dels <= 5 && message == "fix":
		verb = "add"
	case adds-dels >= -5 && message == "fix":
		verb = "remove"
	case message == "fix":
		verb = "fix"
	case adds-dels > 5:
		verb = "add"
	case adds-dels < -5:
		verb = "remove"

	default:
		verb = "change"
	}

	file = selected[1]

	return
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
