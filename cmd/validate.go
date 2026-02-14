package cmd

import (
	"fmt"
	"os"

	"github.com/matheus-foscarinid/envs/internal/manager"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Check that all sample keys are set in .env",
	Long:  `Compares keys from .env.sample against the active .env file and reports any that are missing.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := manager.Validate(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
