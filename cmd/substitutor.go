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
	"time"

	log "github.com/Sirupsen/logrus"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
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

func substitutor(*cobra.Command, []string) {

	log.Debug("Launching substitutor")
	// UI initialization
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	// Grid

	c := widgets.NewParagraph()
	c.Text = "TestTestTestTest"
	c.Title = "Ciphertext"

	s := widgets.NewParagraph()
	s.Text = "TestTestTestTest"
	s.Title = "Solution"

	m := widgets.NewParagraph()
	m.Text = "TestTestTestTest"
	m.Title = "Mapping"

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

	tickerCount := 1
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
				ui.Render(grid)
			}
		case <-ticker:
			if tickerCount == 100 {
				return
			}
			ui.Render(grid)
			tickerCount++
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
