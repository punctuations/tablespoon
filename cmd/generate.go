package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os/exec"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a commit message.",
	Run: func(cmd *cobra.Command, args []string) {
		exec.Command("git", "diff", "--numstat", "--output='../tmp/log.txt'")

		content, err := ioutil.ReadFile("../tmp/log.txt")
		if err != nil {
			pterm.Error.Println(err)
		}

		pterm.Success.Println(string(content))
		//format, _ := cmd.Flags().GetString("format")
		//letters := pterm.NewLettersFromString(time.Now().Format(format))
		//pterm.DefaultBigText.WithLetters(letters).Render()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
