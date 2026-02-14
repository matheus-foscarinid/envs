package cmd

import (
	"fmt"

	"github.com/matheus-foscarinid/envs/internal/manager"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new environment",
	// TODO: Add a longer description
	Long: `Initialize a new environment with a sample file and a config file.
		Creates a new .env.json that will be the config file for envs manager.
		Creates a new .env.sample that will be the template for the new environment, if it doesn't exist.
		Creates a new .env file based on the sample file, if it doesn't exist.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := manager.Init(); err != nil {
			fmt.Println("Error initializing envs:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
