package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"os/exec"
	"regexp"
	"sort"
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

		for i := range out {
			deletions, _ := strconv.Atoi(out[i])
			pattern := regexp.MustCompile("0|[1-9][0-9]*")
			if (i+1)%3 != 0 {
				if i >= 1 {
					if pattern.MatchString(out[i-1]) {
						additions, _ := strconv.Atoi(out[i-1])
						changes := deletions + additions
						diffs = append(diffs, []int{changes}...)
						files = append(files, []string{out[i+1]}...)
					}
				}
				//fmt.Println(outInt)
			}
		}

		sort.Ints(diffs[:])

		pterm.Success.Println(diffs)
		//format, _ := cmd.Flags().GetString("format")
		//letters := pterm.NewLettersFromString(time.Now().Format(format))
		//pterm.DefaultBigText.WithLetters(letters).Render()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
