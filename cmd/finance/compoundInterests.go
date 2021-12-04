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
	"fmt"
	"math"

	"github.com/MakeNowJust/heredoc"
	"github.com/renato0307/canivete/pkg/iostreams"
	"github.com/spf13/cobra"
)

func NewCompoundInterestsCmd(iostreams iostreams.IOStreams) *cobra.Command {

	var compoundInterestsCmd = &cobra.Command{
		Use:   "compoundinterests",
		Short: "Calculates compound interests",
		Long: heredoc.Doc(`
			Calculates compound interests.
	
			The formula for compound interests is a = p*((1+r/n)^(n * t))
			With different periodic payments an extra is needed:
				a_series = m * (y/n) {[(1 + r/n)^(n * t) - 1] / (r/n)}
				total = a + a_series
	
			Where:
				a = the future value of the investment/loan, including interest
				p = the principal investment amount (the initial deposit or loan amount)
				r = the annual interest rate (decimal)
				n = the number of times that interest is compounded per unit t
				t = the time the money is invested or borrowed for
				m = the regular contribution
				y = regular contributions in the compounded period
		`),
		Example: heredoc.Doc(`
			canivete finance compoundinterests -t 10 -p 1000 -r 5 -n 1
			canivete finance compoundinterests -t 25 -p 15000 -r 5 -n 1 -m 400 -y 12
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			p, _ := getFlagIntAsFloat64(cmd, "invest-amount")
			n, _ := getFlagIntAsFloat64(cmd, "compound-periods")
			t, _ := getFlagIntAsFloat64(cmd, "time")
			m, _ := getFlagIntAsFloat64(cmd, "regular-contributions")
			y, _ := getFlagIntAsFloat64(cmd, "regular-contributions-period")
			rint, _ := cmd.Flags().GetFloat64("annual-interest-rate")
			r := rint / 100

			a := p * math.Pow(1+r/n, n*t)
			aseries := 0.0
			if m > 0 {
				if y == 0 {
					return fmt.Errorf("the regular-contributions-period cannot be zero")
				}
				aseries = m * (y / n) * ((math.Pow(1+r/n, n*t) - 1) / (r / n))
			}

			total := a + aseries
			_, err := fmt.Fprintln(iostreams.Out, math.Ceil(total*100)/100)

			return err
		},
	}

	compoundInterestsCmd.Flags().IntP(
		"invest-amount",
		"p",
		0,
		"the principal investment amount (the initial deposit or loan amount)")

	compoundInterestsCmd.Flags().Float64P(
		"annual-interest-rate",
		"r",
		0,
		"the annual interest rate (decimal, percentage)")

	compoundInterestsCmd.Flags().IntP(
		"compound-periods",
		"n",
		0,
		"number of times interest compounds, i.e. 12 = monthly, 4 = quarterly, 2 = semi-annually, 1 = annually")

	compoundInterestsCmd.Flags().IntP(
		"time",
		"t",
		0,
		"the time the money is invested or borrowed for (e.g. 10 years)")

	compoundInterestsCmd.Flags().IntP(
		"regular-contributions",
		"m",
		0,
		"regular contributions (additional money added to investment)")

	compoundInterestsCmd.Flags().IntP(
		"regular-contributions-period",
		"y",
		12,
		"regular contributions in the compounded period (e.g. 12 if every month in a year)")

	compoundInterestsCmd.MarkFlagRequired("invest-amount")
	compoundInterestsCmd.MarkFlagRequired("annual-interest-rate")
	compoundInterestsCmd.MarkFlagRequired("compound-periods")
	compoundInterestsCmd.MarkFlagRequired("time")

	return compoundInterestsCmd
}

func getFlagIntAsFloat64(cmd *cobra.Command, name string) (float64, error) {
	valueInt, err := cmd.Flags().GetInt(name)
	return float64(valueInt), err
}
