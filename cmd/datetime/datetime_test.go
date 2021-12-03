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
	"testing"

	"github.com/renato0307/canivete/pkg/iostreams"
	"github.com/stretchr/testify/assert"
)

func TestNewDatetimeCmd(t *testing.T) {
	// arrange
	iostreams, _, _, _ := iostreams.Test()
	cmd := NewDatetimeCmd(*iostreams)

	// act
	_, err := cmd.ExecuteC()

	// assert
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "must specify a subcommand")
	assert.Len(t, cmd.Commands(), 3)
}
