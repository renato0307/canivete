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
package programming

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/google/uuid"
	"github.com/renato0307/canivete/pkg/iostreams"
	"github.com/spf13/cobra"
)

func NewUuidCmd(stream iostreams.IOStreams) *cobra.Command {
	var uuidCmd = &cobra.Command{
		Use:   "uuid",
		Short: "Generates UUIDs (or GUIDs)",
		Long: heredoc.Doc(`
			UUID also known as GUID is a 16 byte or 128-bit number.
			It is meant to uniquely identify something.
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := fmt.Fprintln(stream.Out, uuid.New())

			return err
		},
	}

	return uuidCmd
}
