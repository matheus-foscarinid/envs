package cmd

import (
	"fmt"

	"github.com/matheus-foscarinid/envs/internal/manager"
	"github.com/spf13/cobra"
)

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Show the active environment",
	Long: `Show the active environment.
		Example:
			envs current
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := manager.Current(); err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(currentCmd)
}
