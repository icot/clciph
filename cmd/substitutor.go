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
	"sort"

	log "github.com/Sirupsen/logrus"
	"github.com/icot/clciph/analysis"
	"github.com/jroimartin/gocui"
	"github.com/spf13/cobra"
)

var solution *analysis.Analysis

// substitutorCmd represents the substitutor command
var substitutorCmd = &cobra.Command{
	Use:   "substitutor",
	Short: "Launch subsittutor interface",
	Long:  `Launches substitutor`,
}

var (
	viewArr = []string{"Mapping", "Ciphertext", "Messages", "Solution"}
	active  = 0
)

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func nextView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (active + 1) % len(viewArr)
	name := viewArr[nextIndex]

	out, err := g.View("Messages")
	if err != nil {
		return err
	}
	fmt.Fprintln(out, "Going from view "+v.Name()+" to "+name)

	if _, err := setCurrentViewOnTop(g, name); err != nil {
		return err
	}

	if nextIndex == 0 || nextIndex == 3 {
		g.Cursor = true
	} else {
		g.Cursor = false
	}

	active = nextIndex
	return nil
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("Mapping", 0, 0, maxX/2-1, maxY/2-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Mapping (editable)"
		v.Editable = false
		v.Wrap = true
		// Map display
		keys := make([]int, 0, len(solution.Mapping))
		for k := range solution.Mapping {
			keys = append(keys, int(k))
		}
		sort.Slice(keys, func(i int, j int) bool { return keys[i] < keys[j] })
		for key := range keys {
			fmt.Fprintln(v, fmt.Sprintf("%c: %s\n", key, string(solution.Mapping[byte(key)])))
		}

		if _, err = setCurrentViewOnTop(g, "Mapping"); err != nil {
			return err
		}
	}

	if v, err := g.SetView("Ciphertext", maxX/2-1, 0, maxX-1, maxY/2-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Ciphertext"
		v.Wrap = true
		v.Autoscroll = true
		fmt.Fprint(v, solution.Ciphertext)

	}
	if v, err := g.SetView("Messages", 0, maxY/2-1, maxX/2-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Messages"
		v.Wrap = true
		v.Autoscroll = true
		fmt.Fprint(v, "Press TAB to change current view")
	}
	if v, err := g.SetView("Solution", maxX/2, maxY/2, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Solution (editable)"
		v.Editable = false
		fmt.Fprint(v, solution.Ciphertext)
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func substitutor(cmd *cobra.Command, args []string) {

	log.Debug("subsittutor called with args: ", args[0])
	solution = new(analysis.Analysis)
	solution = analysis.AnalyzeFile(args[0])
	log.Debug(solution)

	log.Debug("Launching substitutor")
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	/*
		g.Update(func(g *gocui.Gui) error {
			v, err := g.View("Ciphertext")
			if err != nil {
				// handle error
			}
			v.Clear()
			fmt.Fprintln(v, analysis.Ciphertext)
			return nil
		})
	*/

	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorGreen

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
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
