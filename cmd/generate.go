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

		for i := range out {
			outInt, _ := strconv.Atoi(out[i])
			if outInt%3 != 0 {
				fmt.Println(out[i])
			}
			fmt.Println(false)
		}

		pterm.Success.Println(out[0])
		//format, _ := cmd.Flags().GetString("format")
		//letters := pterm.NewLettersFromString(time.Now().Format(format))
		//pterm.DefaultBigText.WithLetters(letters).Render()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
