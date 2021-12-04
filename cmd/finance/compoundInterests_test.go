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
package finance

import (
	"testing"

	"github.com/renato0307/canivete/pkg/iostreams"
	"github.com/stretchr/testify/assert"
)

func TestNewCompoundInterestsCmd(t *testing.T) {
	// arrange
	iostreams, _, out, _ := iostreams.Test()
	cmd := NewCompoundInterestsCmd(*iostreams)

	// act
	cmd.SetArgs([]string{
		"--time=10",
		"--invest-amount=1000",
		"--annual-interest-rate=5",
		"--compound-periods=1",
	})
	_, err := cmd.ExecuteC()

	// assert
	if err != nil {
		t.Fatal(err)
	}
	assert.Contains(t, out.String(), "1628.9")
}

func TestNewCompoundInterestsWithRegularContributionsCmd(t *testing.T) {
	// arrange
	iostreams, _, out, _ := iostreams.Test()
	cmd := NewCompoundInterestsCmd(*iostreams)

	// act
	cmd.SetArgs([]string{
		"-t=10",
		"-p=5000",
		"-r=5",
		"-n=12",
		"-m=100",
		"-y=12",
	})
	_, err := cmd.ExecuteC()

	// assert
	if err != nil {
		t.Fatal(err)
	}
	assert.Contains(t, out.String(), "23763.28")
}
