/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>
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
	"calc/scanner"
	"calc/stack"
	"calc/version"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "calc",
	Short: "Math expression evaluator",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("calc version ", version.Number)

		input, _ := cmd.Flags().GetString("input")
		if input == "" {
			fmt.Println("Math expression input is required, run \"calc -h\" for detail info")
			return
		}

		nScanner := scanner.New(
			scanner.WithQuote(input),
			scanner.WithStack(stack.New()),
		)

		if err := nScanner.DoScan(); err != nil {
			fmt.Println(err.Error())
		}

		fmt.Printf("Result of \"%s\" is: %s\n", input, nScanner.DoSolve())
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("input", "i", "", "Math expression input, addition and subtraction only (i.e:\"5+2-1+-3\")")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
