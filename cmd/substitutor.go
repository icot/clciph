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
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/icot/clciph/analysis"
	"github.com/spf13/cobra"
)

// substitutorCmd represents the substitutor command
var substitutorCmd = &cobra.Command{
	Use:   "substitutor",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func substitutor(cmd *cobra.Command, args []string) {

	log.Debug("subsittutor called with args: ", args[0])
	analysis := analysis.AnalyzeFile(args[0])

	log.Debug("Launching substitutor")
	// UI initialization
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	// Grid
	c := widgets.NewParagraph()
	c.Text = analysis.Ciphertext
	c.Title = "Ciphertext"

	s := widgets.NewParagraph()
	s.Text = analysis.Ciphertext
	s.Title = "Solution"

	m := widgets.NewList()
	m.Title = "Mapping"
	m.Rows = make([]string, len(analysis.Mapping))
	// Iterate over mapping
	for k, v := range analysis.Mapping {
		m.Rows = append(m.Rows, fmt.Sprintf("%s: %s", string(k), string(v)))
	}

	m.TextStyle = ui.NewStyle(ui.ColorYellow)

	grid := ui.NewGrid()
	termhWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termhWidth, termHeight)
	grid.Set(
		ui.NewCol(1.0/2, m),
		ui.NewCol(1.0/2,
			ui.NewRow(1.0/2, c),
			ui.NewRow(1.0/2, s),
		),
	)

	ui.Render(grid)

	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(50 * time.Millisecond).C
	previousKey := ""
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "j", "<Down>":
				m.ScrollDown()
			case "k", "<Up>":
				m.ScrollUp()
			case "<C-d>":
				m.ScrollHalfPageDown()
			case "<C-u>":
				m.ScrollHalfPageUp()
			case "<C-f>":
				m.ScrollPageDown()
			case "<C-b>":
				m.ScrollPageUp()
			case "g":
				if previousKey == "g" {
					m.ScrollTop()
				}
			case "<Home>":
				m.ScrollTop()
			case "G", "<End>":
				m.ScrollBottom()
			}

			if previousKey == "g" {
				previousKey = ""
			} else {
				previousKey = e.ID
			}
		case <-ticker:
			ui.Render(grid)
		}
	}

}

func init() {
	rootCmd.AddCommand(substitutorCmd)
	substitutorCmd.Run = substitutor

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// substitutorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// substitutorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
