package cmd

import (
	"fmt"

	"github.com/matheus-foscarinid/envs/internal/manager"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [key]",
	Short: "Get a variable from the active environment",
	Long:  `Print the value of a key from the active .env file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: key is required")
			return
		}
		if err := manager.Get(args[0]); err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
