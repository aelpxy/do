package cmd

import (
	"fmt"

	"github.com/aelpxy/dbctl/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:        config.CmdName,
	Short:      config.CmdShortDescription,
	Long:       config.CmdLongDescription,
	SuggestFor: []string{"db"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		} else {
			fmt.Println("Unknown command. Please use 'dbctl --help' to see available commands.")
		}
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the current dbctl version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s \n", config.Version)
	},
}

func Execute() {
	rootCmd.AddCommand(versionCmd)

	rootCmd.Execute()
}
