package finance

import (
	"fmt"
	"math"

	"github.com/spf13/cobra"
)

var compoundInterestsCmd = &cobra.Command{
	Use:   "compoundinterests",
	Short: "Calculates compound interests",
	Long: `
Calculates compound interests.

The formula for compound interests is A = P * ((1 + r/n) ^ (n * t))

Where:

A = the future value of the investment/loan, including interest
P = the principal investment amount (the initial deposit or loan amount)
r = the annual interest rate (decimal)
n = the number of times that interest is compounded per unit t
t = the time the money is invested or borrowed for
`,
	Run: func(cmd *cobra.Command, args []string) {
		investAmount, _ := cmd.Flags().GetInt("invest-amount")
		annualInterestRate, _ := cmd.Flags().GetFloat64("annual-interest-rate")
		compoundPeriods, _ := cmd.Flags().GetInt("compound-periods")
		time, _ := cmd.Flags().GetInt("time")

		part1 := (1 + ((annualInterestRate / 100) / float64(compoundPeriods)))
		part2 := float64(compoundPeriods * time)
		result := float64(investAmount) * math.Pow(part1, part2)
		roundedResult := math.Ceil(result*100) / 100

		fmt.Println(roundedResult)
	},
}

func init() {
	financeCmd.AddCommand(compoundInterestsCmd)

	compoundInterestsCmd.Flags().IntP("invest-amount", "p", 0, "the principal investment amount (the initial deposit or loan amount)")
	compoundInterestsCmd.Flags().Float64P("annual-interest-rate", "r", 0.0, "the annual interest rate (decimal, percentage)")
	compoundInterestsCmd.Flags().IntP("compound-periods", "n", 0, "the number of times that interest is compounded per unit t (e.g. 12 if monthly)")
	compoundInterestsCmd.Flags().IntP("time", "t", 0, "the time the money is invested or borrowed for (e.g. 10 years)")

	compoundInterestsCmd.MarkFlagRequired("invest-amount")
	compoundInterestsCmd.MarkFlagRequired("annual-interest-rate")
	compoundInterestsCmd.MarkFlagRequired("compound-periods")
	compoundInterestsCmd.MarkFlagRequired("time")
}
