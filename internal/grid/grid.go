// Copyright 2020 Andrew Merenbach
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package grid

import (
	"sort"
	"strings"
)

// A cell holds a rune alongside table coordinates.
type cell struct {
	Row  int
	Col  int
	Rune rune
}

// A Grid is a slice of cells.
type Grid []cell

// FillByRow populates cell contents by row.
func (g Grid) FillByRow(s string) {
	g.sortByCol()
	g.fill(s)
	g.sortByRow()
}

// FillByCol populates cell contents by column.
func (g Grid) FillByCol(s string) {
	g.sortByRow()
	g.fill(s)
	g.sortByCol()
}

// ReadByRow concatenates cell contents by row.
func (g Grid) ReadByRow() string {
	g2 := make(Grid, len(g))
	copy(g2, g)
	g2.sortByRow()
	return g2.contents()
}

// ReadByCol concatenates cell contents by column.
func (g Grid) ReadByCol() string {
	g2 := make(Grid, len(g))
	copy(g2, g)
	g2.sortByCol()
	return g2.contents()
}

// Printable version of this grid.
func (g Grid) Printable() string {
	g2 := make(Grid, len(g))
	copy(g2, g)

	g2.sortByRow()

	var currentRow int
	var currentCol int

	var out strings.Builder
	for _, c := range g2 {
		if c.Row > currentRow {
			currentRow = c.Row
			currentCol = 0
			out.WriteRune('\n')
		}

		for c.Col > currentCol {
			out.WriteRune(' ')
			currentCol++
		}
		currentCol++

		out.WriteRune(c.Rune)
	}

	return out.String()
}

// Contents of this grid in the current order.
func (g Grid) contents() string {
	var b strings.Builder
	for _, c := range g {
		b.WriteRune(c.Rune)
	}
	return b.String()
}

// Fill a grid with runes.
func (g Grid) fill(s string) {
	for i, r := range []rune(s)[:len(g)] {
		g[i].Rune = r
	}
}

// SortByRow sorts a grid by row.
func (g Grid) sortByRow() {
	// SliceStable works here with the assumption that cells are already sorted by column.
	// A stable slice isn't required so much as having column values being used as a tiebreaker.
	// One functional equivalent (in this context) that works with sort.Slice() would be:
	//
	//    return (g[i].Row == g[j].Row && g[i].Col < g[j].Col) || g[i].Row < g[j].Row
	//
	sort.SliceStable(g, func(i, j int) bool {
		return g[i].Row < g[j].Row
	})
}

// SortByCol sorts a grid by column.
func (g Grid) sortByCol() {
	// SliceStable works here with the assumption that cells are already sorted by row.
	// A stable slice isn't required so much as having row values being used as a tiebreaker.
	// One functional equivalent (in this context) that works with sort.Slice() would be:
	//
	//    return (g[i].Col == g[j].Col && g[i].Row < g[j].Row) || g[i].Col < g[j].Col
	//
	sort.SliceStable(g, func(i, j int) bool {
		return g[i].Col < g[j].Col
	})
}
