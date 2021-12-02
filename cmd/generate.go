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
			outInt, _ := strconv.Atoi(out[i])
			if (i+1)%3 != 0 {
				if (i+1)%2 == 0 {
					prevInt, _ := strconv.Atoi(out[i-1])
					changes := outInt + prevInt
					diffs = append(diffs, []string{string(rune(changes)), out[i+1]}...)
				}

				fmt.Println(outInt)
			}
			fmt.Println(false)
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
