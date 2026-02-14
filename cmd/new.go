package cmd

import (
	"fmt"

	"github.com/matheus-foscarinid/envs/internal/manager"
	"github.com/spf13/cobra"
)

var newEmpty bool

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new environment",
	Long: `Create a new environment based on sample file, using the param as name
	Example:
		envs new dev
		envs new prod
		envs new staging
		envs new test
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: a valid environment name is required")
			return
		}
		if err := manager.NewEnv(args[0], newEmpty); err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().BoolVarP(&newEmpty, "empty", "e", false, "ignore sample file and create an empty .env")
}
