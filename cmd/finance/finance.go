package finance

import (
	"fmt"

	"github.com/spf13/cobra"
)

var financeCmd = &cobra.Command{
	Use:   "finance",
	Short: "Finance related tools",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Error: must also specify a command like compoundinterests, etc.")
	},
}

func NewFinanceCmd() *cobra.Command {
	return financeCmd
}

func init() {
}
