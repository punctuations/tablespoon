package cmd

import (
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

		var diffs []int
		var files []string
		var adds int
		var dels int

		for i := range out {
			deletions, _ := strconv.Atoi(out[i])
			pattern := regexp.MustCompile("0|[1-9][0-9]*")
			if (i+1)%3 != 0 {
				if i == 0 { // initialize all the additions + deletions and store
					del, _ := strconv.Atoi(out[i+1])
					adds = deletions // not actually deletions in this case
					dels = del
				}
				if i >= 1 {
					if pattern.MatchString(out[i-1]) {
						additions, _ := strconv.Atoi(out[i-1])
						changes := deletions + additions
						diffs = append(diffs, []int{changes}...)
						files = append(files, []string{out[i+1]}...)
						if additions+deletions > adds+dels {
							adds = additions
							dels = deletions
						}
					}
				}
				//fmt.Println(outInt)
			}
		}

		var selected []string

		for n := range diffs {
			s, _ := strconv.Atoi(selected[0])
			if n == 0 {
				selected = []string{strconv.Itoa(diffs[n]), files[n]}
			} else if diffs[n] > s {
				selected = []string{strconv.Itoa(diffs[n]), files[n]}
			}
		}

		pterm.Success.Println(selected, " + ", adds, " - ", dels)
		//format, _ := cmd.Flags().GetString("format")
		//letters := pterm.NewLettersFromString(time.Now().Format(format))
		//pterm.DefaultBigText.WithLetters(letters).Render()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
