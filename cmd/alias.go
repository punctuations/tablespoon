package cmd

import (
	"github.com/spf13/cobra"
)

// alias represents the config command
var alias = &cobra.Command{
	Use:   "alias",
	Short: "creates an alias for tablespoon.",
	Run: func(cmd *cobra.Command, args []string) {
		// add checking for diff shells
		// reminder: bash has function keyword in front
		println("tbsp () {{\n  $TBSP_CMD=$(\n    echo tablespoon $@\n  ) && eval $TBSP_CMD\n}}")
	},
}

func init() {
	rootCmd.AddCommand(alias)

	// add flag to allow for custom alias name
}
