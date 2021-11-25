package cmd

import (
	"fmt"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"os/exec"
)

// dateCmd represents the date command
var commitCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a commit message.",
	Run: func(cmd *cobra.Command, args []string) {
		input := "test: test!!"
		out, err := exec.Command(fmt.Sprintf("git commit -m %s", input)).Output()

		// if there is an error with our execution handle it here
		if err != nil {
			pterm.Error.Printf("%s", err)
		}

		output := string(out[:])
		fmt.Println(output)
		pterm.Success.Println("Command Successfully Executed")
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
