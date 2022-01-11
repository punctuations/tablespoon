package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"io/ioutil"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print out the config options that are applied.",
	Run: func(cmd *cobra.Command, args []string) {

		type Config struct {
			commentID string
		}

		// fix me -- nothing being printed from file??!?!?!
		info, infoErr := ioutil.ReadFile("tablespoon.json")
		fmt.Println(string(info))
		if infoErr != nil {
			pterm.Error.Println(errors.New("Error T8: " + infoErr.Error()))
			return
		}

		var payload Config
		infoErr = json.Unmarshal(info, &payload)
		if infoErr != nil {
			pterm.Error.Println(errors.New("Error T8: " + infoErr.Error()))
			return
		}

		fmt.Println(payload.commentID)

		pterm.Success.Println("Command Successfully Executed")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
