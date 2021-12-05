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
package cmd

import (
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/renato0307/canivete/cmd/datetime"
	"github.com/renato0307/canivete/cmd/finance"
	"github.com/renato0307/canivete/cmd/internet"
	"github.com/renato0307/canivete/cmd/programming"
	"github.com/renato0307/canivete/pkg/iostreams"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "canivete",
	Short: "Utility functions you'll use for life",
	Long: heredoc.Doc(`
		canivete is a CLI to support you everyday, making your like simpler.

		Here you can find utility tools to:
		. Calculate compound interests
		. Generate UUIDs or nanoid
		. Etcetera

		Isn't that great?
	`),
	Version: "0.0.9",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.canivete.yaml)")

	iostreams := iostreams.IOStreams{
		ErrOut: os.Stderr,
		In:     os.Stdin,
		Out:    os.Stdout,
	}

	rootCmd.AddCommand(datetime.NewDatetimeCmd(iostreams))
	rootCmd.AddCommand(internet.NewInternetCmd(iostreams))
	rootCmd.AddCommand(finance.NewFinanceCmd(iostreams))
	rootCmd.AddCommand(programming.NewProgrammingCmd(iostreams))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".canivete" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".canivete")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
