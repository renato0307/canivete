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
