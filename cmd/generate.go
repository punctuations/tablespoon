package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"time"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a commit message.",
	Run: func(cmd *cobra.Command, args []string) {
		format, _ := cmd.Flags().GetString("format")
		letters := pterm.NewLettersFromString(time.Now().Format(format))
		pterm.DefaultBigText.WithLetters(letters).Render()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
