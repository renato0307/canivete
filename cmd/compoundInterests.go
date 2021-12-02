/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"math"

	"github.com/spf13/cobra"
)

// compoundInterestsCmd represents the compoundInterests command
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// compoundInterestsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	compoundInterestsCmd.Flags().IntP("invest-amount", "p", 0, "the principal investment amount (the initial deposit or loan amount)")
	compoundInterestsCmd.Flags().Float64P("annual-interest-rate", "r", 0.0, "the annual interest rate (decimal, percentage)")
	compoundInterestsCmd.Flags().IntP("compound-periods", "n", 0, "the number of times that interest is compounded per unit t (e.g. 12 if monthly)")
	compoundInterestsCmd.Flags().IntP("time", "t", 0, "the time the money is invested or borrowed for (e.g. 10 years)")

	compoundInterestsCmd.MarkFlagRequired("invest-amount")
	compoundInterestsCmd.MarkFlagRequired("annual-interest-rate")
	compoundInterestsCmd.MarkFlagRequired("compound-periods")
	compoundInterestsCmd.MarkFlagRequired("time")
}
