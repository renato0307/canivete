/*
Copyright © 2021 Renato Torres <renato.torres@pm.me>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package datetime

import (
	"fmt"
	"time"

	"github.com/MakeNowJust/heredoc"
	"github.com/renato0307/canivete/pkg/iostreams"
	"github.com/spf13/cobra"
)

func NewFromUnixCmd(iostreams iostreams.IOStreams) *cobra.Command {
	var fromUnixCmd = &cobra.Command{
		Use:   "fromunix",
		Short: "Converts a Unix timestamp to human friendly format",
		Long: heredoc.Doc(`
			Converts a Unix timestamp to human friendly format.

			The Unix timestamp is a way to track time as a running total of seconds.
			This count starts at the Unix Epoch on January 1st, 1970 at UTC.
		`),
		RunE: func(cmd *cobra.Command, args []string) error {

			unixTime, _ := cmd.Flags().GetInt64("value")
			t := time.Unix(unixTime, 0)
			strDate := t.UTC().Format(time.UnixDate)
			_, err := fmt.Fprintln(iostreams.Out, strDate)

			return err
		},
		Example: heredoc.Doc(`
			canivete datetime fromunix --value 1638964800
			canivete datetime fromunix -v 1638964800`),
	}

	fromUnixCmd.Flags().Int64P("value", "v", 0, "the unix timestamp")
	fromUnixCmd.MarkFlagRequired("value")

	return fromUnixCmd
}
