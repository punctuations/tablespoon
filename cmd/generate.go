package cmd

import (
	"fmt"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"os/exec"
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

		var diffs []string

		for i := range out {
			deletions, _ := strconv.Atoi(out[i])
			if (i+1)%3 != 0 {
				if (i+1)%2 == 0 {
					additions, _ := strconv.Atoi(out[i-1])
					changes := deletions + additions
					fmt.Println(deletions)
					fmt.Println(additions)
					fmt.Println(additions, " + ", deletions, " = ", changes)
					diffs = append(diffs, []string{strconv.Itoa(changes), out[i+1]}...)
				}
				//fmt.Println(outInt)
			} else {
				fmt.Println(out[i], " - Not outputted")
			}
		}

		pterm.Success.Println(diffs)
		//format, _ := cmd.Flags().GetString("format")
		//letters := pterm.NewLettersFromString(time.Now().Format(format))
		//pterm.DefaultBigText.WithLetters(letters).Render()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
