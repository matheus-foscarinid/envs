/*
Copyright © 2026 Matheus Foscarini Dias <matheus.foscarinid@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/matheus-foscarinid/envs/internal/manager"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
