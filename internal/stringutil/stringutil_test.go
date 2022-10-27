// Copyright 2018 Andrew Merenbach
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

package stringutil

import (
	"reflect"
	"testing"

	"github.com/merenbach/goldbug/internal/fixture"
)

func TestBackpermute(t *testing.T) {
	var tables []struct {
		Input   []rune
		Output  []rune
		Indices []int
		Success bool
	}

	fixture.Load(t, &tables)
	for i, table := range tables {
		t.Logf("Testing table %d of %d", i+1, len(tables))
		if out, err := Backpermute(table.Input, table.Indices); err != nil && table.Success {
			t.Error("Unexpected backpermute failure:", err)
		} else if err == nil && !table.Success {
			t.Error("Unexpected backpermute success")
		} else if reflect.DeepEqual(out, table.Output) {
			t.Error("Received incorrect output:", out)
		}
	}
}

func TestDeduplicate(t *testing.T) {
	table := map[string]string{
		"hello":       "helo",
		"world":       "world",
		"hello world": "helo wrd",
	}

	for k, v := range table {
		if o := Deduplicate(k); o != v {
			t.Errorf("Deduplication of string %q was %q; expected %q", k, o, v)
		}
	}
}

func TestIntersect(t *testing.T) {
	tables := [][]string{
		{"HELLO, WORLD!", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "HELLOWORLD"},
		{"world", "word", "word"},
		{"world", "hello", "ol"},
		{"hello", "world", "llo"},
	}

	for _, table := range tables {
		if o := intersect(table[0], table[1]); o != table[2] {
			t.Errorf("Intersection of string %q with charset %q was %q; expected %q", table[0], table[1], o, table[2])
		}
	}
}

/*func TestReverse(t *testing.T) {
	table := map[string]string{
		"hello": "olleh",
		"world": "dlrow",
	}

	for k, v := range table {
		if o := Reverse(k); o != v {
			t.Errorf("Reverse of string %q was %q; expected %q", k, o, v)
		}
	}
}*/

// func TestWrapString(t *testing.T) {
// 	tables := []struct {
// 		s        string
// 		i        int
// 		expected string
// 	}{
// 		{"hello", 3, "lohel"},
// 		{"hello world", 0, "hello world"},
// 		{"hello world", 11, "hello world"},
// 	}
// 	for _, table := range tables {
// 		if o := WrapString(table.s, table.i); o != table.expected {
// 			t.Errorf("Wrapping of string %q by %d places was %q; expected %q", table.s, table.i, o, table.expected)
// 		}
// 	}
// }

// func ExampleWrapString() {
// 	s := "HELLO,_WORLD!"
// 	for i := range []rune(s) {
// 		fmt.Println(WrapString(s, i))
// 	}
// 	// Output:
// 	// HELLO,_WORLD!
// 	// ELLO,_WORLD!H
// 	// LLO,_WORLD!HE
// 	// LO,_WORLD!HEL
// 	// O,_WORLD!HELL
// 	// ,_WORLD!HELLO
// 	// _WORLD!HELLO,
// 	// WORLD!HELLO,_
// 	// ORLD!HELLO,_W
// 	// RLD!HELLO,_WO
// 	// LD!HELLO,_WOR
// 	// D!HELLO,_WORL
// 	// !HELLO,_WORLD
// }

func TestChunk(t *testing.T) {
	tables := []struct {
		s         string
		size      int
		delimiter rune
		expected  string
	}{
		{"HELLOWORLD", 2, ' ', "HE LL OW OR LD"},
		{"HELLOWORLD", 3, ' ', "HEL LOW ORL DXX"},
		{"HELLOWORLD", 4, ' ', "HELL OWOR LDXX"},
	}

	for _, table := range tables {
		if o := chunk(table.s, table.size, table.delimiter); o != table.expected {
			t.Errorf("Chunking of string %q with size %d and delimiter %q was %q; expected %q", table.s, table.size, table.delimiter, o, table.expected)
		}
	}
}

func TestGroupString(t *testing.T) {
	tables := []struct {
		s        string
		size     int
		padding  rune
		expected []string
	}{
		{"HELLOWORLD", 2, 'X', []string{"HE", "LL", "OW", "OR", "LD"}},
		{"HELLOWORLD", 3, 'X', []string{"HEL", "LOW", "ORL", "DXX"}},
		{"HELLOWORLD", 4, 'X', []string{"HELL", "OWOR", "LDXX"}},
	}

	for _, table := range tables {
		if o := groupString(table.s, table.size, table.padding); !reflect.DeepEqual(o, table.expected) {
			t.Errorf("Grouping of string %q with size %d and padding %q was %q; expected %q", table.s, table.size, table.padding, o, table.expected)
		}
	}
}

func TestDiffToMod(t *testing.T) {
	tables := [][]int{
		{6, 3, 0},
		{6, 4, 2},
		{7, 3, 2},
		{10, 2, 0},
		{10, 3, 2},
		{10, 4, 2},
	}

	for _, table := range tables {
		if o := diffToMod(table[0], table[1]); o != table[2] {
			t.Errorf("Diff to mod of %d %% %d was %d; expected %d", table[0], table[1], o, table[2])
		}
	}
}
