/*
Copyright © 2024 Ashot Karapetian

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"time"

	"pomo/app"
	"pomo/pomodoro"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pomo",
	Short: "Interactive Pomodoro Timer",
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, err := getRepo()
		if err != nil {
			return err
		}

		config := pomodoro.NewConfig(
			repo,
			viper.GetDuration("pomo"),
			viper.GetDuration("short"),
			viper.GetDuration("long"),
		)

		return rootAction(os.Stdout, config)
	},
}

func rootAction(out io.Writer, config *pomodoro.IntervalConfig) error {
	a, err := app.New(config)
	if err != nil {
		return err
	}

	return a.Run()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pomo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("db", "d", "pomo.db", "database file")
	rootCmd.Flags().DurationP("pomo", "p", 25*time.Minute, "pomodoro duration")
	rootCmd.Flags().DurationP("short", "s", 5*time.Minute, "short break duration")
	rootCmd.Flags().DurationP("long", "l", 15*time.Minute, "long break duration")

	viper.BindPFlag("db", rootCmd.Flags().Lookup("db"))
	viper.BindPFlag("pomo", rootCmd.Flags().Lookup("pomo"))
	viper.BindPFlag("short", rootCmd.Flags().Lookup("short"))
	viper.BindPFlag("long", rootCmd.Flags().Lookup("long"))
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Search config in home directory with name ".pScan" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".pScan")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
