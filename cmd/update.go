package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"os/exec"
	"runtime"
	"strings"
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
			shellOut, shellErr := exec.Command("echo $SHELL").Output()
			if shellErr != nil {
				pterm.Error.Println("500: Unable to identify shell; ", shellErr)
				return
			}
			shell := strings.Split(string(shellOut), "/")[len(strings.Split(string(shellOut), "/"))-1]
			if shell == "bash" || shell == "sh" {
				exec.Command("curl -sSL instl.sh/punctuations/tablespoon/macos | sudo bash")
			} else {
				exec.Command("sudo curl -sSL instl.sh/punctuations/tablespoon/macos | sudo bash")
			}
		default: // everything else basically
			shellOut, shellErr := exec.Command("echo $SHELL").Output()
			if shellErr != nil {
				pterm.Error.Println("500: Unable to identify shell; ", shellErr)
				return
			}
			shell := strings.Split(string(shellOut), "/")[len(strings.Split(string(shellOut), "/"))-1]
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
