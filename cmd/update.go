package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the cli tool",
	Run: func(cmd *cobra.Command, args []string) {
		shellPath := os.Getenv("SHELL")

		os := runtime.GOOS
		switch os {
		case "windows": // self-explanatory
			exec.Command("iwr instl.sh/punctuations/tablespoon/windows | iex")
		case "darwin": // macOS
			shell := strings.Split(shellPath, "/")[len(strings.Split(shellPath, "/"))-1]

			if shell == "bash" || shell == "sh" {
				exec.Command("curl -sSL instl.sh/punctuations/tablespoon/macos | sudo bash")
			} else {
				exec.Command("sudo curl -sSL instl.sh/punctuations/tablespoon/macos | sudo bash")
			}
		default: // everything else basically
			shell := strings.Split(shellPath, "/")[len(strings.Split(shellPath, "/"))-1]

			if shell == "bash" || shell == "sh" {
				exec.Command("curl -sSL instl.sh/punctuations/tablespoon/linux | sudo bash")
			} else {
				exec.Command("sudo curl -sSL instl.sh/punctuations/tablespoon/linux | sudo bash")
			}
		}

		pterm.Success.Println("✨ Command Successfully Executed")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
