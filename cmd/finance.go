package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// financeCmd represents the finance command
var financeCmd = &cobra.Command{
	Use:   "finance",
	Short: "Finance related tools",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Error: must also specify a command like compoundInterests, etc.")
	},
}

func init() {
	rootCmd.AddCommand(financeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// financeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// financeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
