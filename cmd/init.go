/*
Copyright © 2026 Matheus Foscarini Dias <matheus.foscarinid@gmail.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/matheus-foscarinid/envs/internal/config"
	"github.com/matheus-foscarinid/envs/internal/constants"
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
		fmt.Println("⛰️ Initializing envs...")
		fmt.Println("\n")

		// check if the config file exists
		if _, err := os.Stat(constants.ConfigFile); os.IsNotExist(err) {
			fmt.Println("Creating default config file...")
			createDefaultConfigFile()
		} else {
			fmt.Println("✅ Config file already exists, skipping...")
		}

		// check if the sample file exists
		if _, err := os.Stat(constants.SampleFile); os.IsNotExist(err) {
			fmt.Println("Creating default sample file...")
			os.Create(constants.SampleFile)
		} else {
			fmt.Println("✅ Sample file already exists, skipping...")
		}

		// create the env file
		if _, err := os.Stat(constants.DotEnvFile); os.IsNotExist(err) {
			fmt.Println("Creating default env file...")
			os.Create(constants.DotEnvFile)
		} else {
			fmt.Println("✅ Env file already exists, skipping...")
		}

		fmt.Println("\n")
		fmt.Println("⛰️ Envs initialized successfully!")
	},
}

func createDefaultConfigFile() error {
	file, err := os.Create(constants.ConfigFile)
	if err != nil {
		fmt.Println("Error creating config file:", err)
		return err
	}
	defer file.Close()

	// write the default config file
	config := config.Config{
		Version: 1,
		Active:  "default",
	}

	json.NewEncoder(file).Encode(config)
	return nil
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
