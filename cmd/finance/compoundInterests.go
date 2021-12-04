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

type compoundInterestsOutput struct {
	FutureValue float64 `json:"futureValue"`
}

const flagInvestAmount = "invest-amount"
const flagCompoundPeriods = "compound-periods"
const flagTime = "time"
const flagRegularContributions = "regular-contributions"
const flagRegularContributionsPeriod = "regular-contributions-period"
const flagAnnualInterestRate = "annual-interest-rate"

func NewCompoundInterestsCmd(iostreams iostreams.IOStreams) *cobra.Command {

	// Command definition
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
			p, _ := getFlagIntAsFloat64(cmd, flagInvestAmount)
			n, _ := getFlagIntAsFloat64(cmd, flagCompoundPeriods)
			t, _ := getFlagIntAsFloat64(cmd, flagTime)
			m, _ := getFlagIntAsFloat64(cmd, flagRegularContributions)
			y, _ := getFlagIntAsFloat64(cmd, flagRegularContributionsPeriod)
			rint, _ := cmd.Flags().GetFloat64(flagAnnualInterestRate)

			if m > 0 && y == 0 {
				return fmt.Errorf("the regular-contributions-period cannot be zero")
			}

			output := run(p, n, t, m, y, rint)
			err := iostreams.PrintOutput(output)

			return err
		},
	}

	// Command flags
	compoundInterestsCmd.Flags().IntP(
		flagInvestAmount,
		"p",
		0,
		"the principal investment amount (the initial deposit or loan amount)")
	compoundInterestsCmd.MarkFlagRequired(flagInvestAmount)

	compoundInterestsCmd.Flags().Float64P(
		flagAnnualInterestRate,
		"r",
		0,
		"the annual interest rate (decimal, percentage)")
	compoundInterestsCmd.MarkFlagRequired(flagAnnualInterestRate)

	compoundInterestsCmd.Flags().IntP(
		flagCompoundPeriods,
		"n",
		0,
		"number of times interest compounds, i.e. 12 = monthly, 4 = quarterly, 2 = semi-annually, 1 = annually")
	compoundInterestsCmd.MarkFlagRequired(flagCompoundPeriods)

	compoundInterestsCmd.Flags().IntP(
		flagTime,
		"t",
		0,
		"the time the money is invested or borrowed for (e.g. 10 years)")
	compoundInterestsCmd.MarkFlagRequired(flagTime)

	compoundInterestsCmd.Flags().IntP(
		flagRegularContributions,
		"m",
		0,
		"regular contributions (additional money added to investment)")

	compoundInterestsCmd.Flags().IntP(
		flagRegularContributionsPeriod,
		"y",
		12,
		"regular contributions in the compounded period (e.g. 12 if every month in a year)")

	return compoundInterestsCmd
}

func run(p, n, t, m, y, rInt float64) compoundInterestsOutput {
	r := rInt / 100

	// base calculation
	a := p * math.Pow(1+r/n, n*t)

	// calculation for regular contributions
	aseries := 0.0
	if m > 0 {
		aseries = m * (y / n) * ((math.Pow(1+r/n, n*t) - 1) / (r / n))
	}

	// total
	total := roundTwoDecimalPlaces(a + aseries)

	return compoundInterestsOutput{FutureValue: total}
}

func getFlagIntAsFloat64(cmd *cobra.Command, name string) (float64, error) {
	valueInt, err := cmd.Flags().GetInt(name)
	return float64(valueInt), err
}

func roundTwoDecimalPlaces(value float64) float64 {
	return math.Ceil(value*100) / 100
}
