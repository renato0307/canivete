package programming

import (
	"fmt"

	"github.com/spf13/cobra"
)

var programmingCmd = &cobra.Command{
	Use:   "programming",
	Short: "Programming tools",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Error: must also specify a command like uuid, etc.")
	},
}

func NewProgrammingCmd() *cobra.Command {
	return programmingCmd
}

func init() {
}
