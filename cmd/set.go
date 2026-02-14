package cmd

import (
	"fmt"

	"github.com/matheus-foscarinid/envs/internal/manager"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "Set a variable in the active environment",
	Long:  `Set or update a key-value pair in the active .env file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Error: key and value are required")
			return
		}
		if err := manager.Set(args[0], args[1]); err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
