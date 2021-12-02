package programming

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generates UUIDs (or GUIDs)",
	Long: heredoc.Doc(`
		UUID also known as GUID is a 16 byte or 128-bit number.
		It is meant to uniquely identify something.
	`),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(uuid.New())
	},
}

func init() {
	programmingCmd.AddCommand(uuidCmd)
}
