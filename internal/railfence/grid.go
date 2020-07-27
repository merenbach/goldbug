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

package railfence

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

// A grid is a slice of cells.
type grid []cell

// Printable version of this grid.
func (g grid) printable() string {
	g2 := make(grid, len(g))
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

// Contents of this grid.
func (g grid) contents() string {
	var b strings.Builder
	for _, c := range g {
		b.WriteRune(c.Rune)
	}
	return b.String()
}

// Fill a grid with runes.
func (g grid) fill(s string) {
	for i, r := range []rune(s)[:len(g)] {
		g[i].Rune = r
	}
}

// Renumber a grid with the given rail count.
func (g grid) renumber(k int) {
	k--
	for i := range g {
		g[i].Col = i

		// Cycle length is 2*(rails - 1), or 2*k
		g[i].Row = k - abs(k-i%(2*k))
	}
}

// SortByRow sorts a grid by row.
func (g grid) sortByRow() {
	// SliceStable works here with the assumption that cells are sorted by default by column.
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
func (g grid) sortByCol() {
	sort.Slice(g, func(i, j int) bool {
		return g[i].Col < g[j].Col
	})
}
