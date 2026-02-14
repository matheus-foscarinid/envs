package cmd

import (
	"fmt"

	"github.com/matheus-foscarinid/envs/internal/manager"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Switch active environment",
	Long: `Switch active environment to the given name.
		Example:
			envs use dev
			envs use prod
			envs use staging
			envs use test
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: a valid environment name is required")
			return
		}

		if err := manager.Use(args[0]); err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
