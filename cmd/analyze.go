/*
Copyright © 2019 Ignacio Coterillo <Ignacio.Coterillo@cern.ch>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	log "github.com/Sirupsen/logrus"
	"github.com/icot/clciph/analysis"
	"github.com/icot/clciph/tools/file"
	"github.com/spf13/cobra"
)

// analyzeCmd represents the analyze command
var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
	analyzeCmd.Run = analyze

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// analyzeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// analyzeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func analyze(cmd *cobra.Command, args []string) {
	log.Debug("analyze called with arg: ", args[0])
	buf := file.Load(args[0])
	analysis := analysis.Analyze(buf)
	log.Info(string(analysis.Bytes))
}
