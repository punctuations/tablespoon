package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"os/exec"
	"runtime"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the cli tool",
	Run: func(cmd *cobra.Command, args []string) {

		os := runtime.GOOS
		switch os {
		case "windows": // self-explanatory
			exec.Command("iwr instl.sh/punctuations/tablespoon/windows | iex")
		case "darwin": // macOS
			exec.Command("curl -sSL instl.sh/punctuations/tablespoon/macos | bash")
		default: // everything else basically
			exec.Command("curl -sSL instl.sh/punctuations/tablespoon/linux | bash")
		}

		pterm.Success.Println("✨ Command Successfully Executed")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
