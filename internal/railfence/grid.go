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

func (g grid) String() string {
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
	sort.SliceStable(g, func(i, j int) bool {
		return g[i].Row < g[j].Row
	})
}

// SortByCol sorts a grid by column.
func (g grid) sortByCol() {
	sort.SliceStable(g, func(i, j int) bool {
		return g[i].Col < g[j].Col
	})
}
