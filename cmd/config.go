package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print out the config options that are applied.",
	Run: func(cmd *cobra.Command, args []string) {
		create, _ := cmd.Flags().GetBool("create")

		if create {
			// create the file
			f, err := os.Create("tbsp.json")
			if err != nil {
				pterm.Error.Println(errors.New("Error T8: " + err.Error()))
				return
			}
			// close the file with defer
			defer f.Close()

			// do operations

			//write directly into file
			f.Write([]byte("{\"commentID\": \"tbsp: \"}"))
		} else {

			type Config struct {
				CommentID string
			}

			info, infoErr := ioutil.ReadFile("tbsp.json")
			fmt.Println(string(info))
			if infoErr != nil {
				pterm.Error.Println(errors.New("Error T8: " + infoErr.Error()))
				return
			}

			var obj Config
			infoErr = json.Unmarshal(info, &obj)
			if infoErr != nil {
				pterm.Error.Println(errors.New("Error T8: " + infoErr.Error()))
				return
			}
		}

		pterm.Success.Println("✨ Command Successfully Executed")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().BoolP("create", "c", false, "create a tablespoon config file")
}
