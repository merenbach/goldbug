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
	"fmt"
	"reflect"
	"testing"
)

func TestBackpermute(t *testing.T) {
	successTables := []struct {
		expected string
		s        string
		ii       []int
	}{
		{"eh", "hello", []int{1, 0}},
	}

	failTables := []struct {
		s  string
		ii []int
	}{
		{"hello", []int{1, 0, -1}},
		{"hello", []int{1, 0, 5}},
	}

	for _, table := range successTables {
		out, err := backpermute([]rune(table.s), table.ii)
		if err != nil {
			t.Fatal(err)
		}

		if string(out) != table.expected {
			t.Errorf("For backpermutation of %q with slice %v, expected output %q, but got %q instead", table.s, table.ii, table.expected, out)
		}
	}

	for _, table := range failTables {
		if _, err := backpermute([]rune(table.s), table.ii); err == nil {
			t.Errorf("Expected backpermutation of %q with slice %v to fail", table.s, table.ii)
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

func TestWrapString(t *testing.T) {
	tables := []struct {
		s        string
		i        int
		expected string
	}{
		{"hello", 3, "lohel"},
		{"hello world", 0, "hello world"},
		{"hello world", 11, "hello world"},
	}
	for _, table := range tables {
		if o := WrapString(table.s, table.i); o != table.expected {
			t.Errorf("Wrapping of string %q by %d places was %q; expected %q", table.s, table.i, o, table.expected)
		}
	}
}

func ExampleWrapString() {
	s := "HELLO,_WORLD!"
	for i := range []rune(s) {
		fmt.Println(WrapString(s, i))
	}
	// Output:
	// HELLO,_WORLD!
	// ELLO,_WORLD!H
	// LLO,_WORLD!HE
	// LO,_WORLD!HEL
	// O,_WORLD!HELL
	// ,_WORLD!HELLO
	// _WORLD!HELLO,
	// WORLD!HELLO,_
	// ORLD!HELLO,_W
	// RLD!HELLO,_WO
	// LD!HELLO,_WOR
	// D!HELLO,_WORL
	// !HELLO,_WORLD
}

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

func TestMakeTranslationTable(t *testing.T) {
	tables := []struct {
		a   string
		b   string
		c   string
		out map[rune]rune
	}{
		{
			a:   "ABC",
			b:   "DEF",
			c:   "",
			out: map[rune]rune{'A': 'D', 'B': 'E', 'C': 'F'},
		},
		{
			a:   "ABC",
			b:   "abc",
			c:   "",
			out: map[rune]rune{'A': 'a', 'B': 'b', 'C': 'c'},
		},
		{
			a:   "ABC",
			b:   "abc",
			c:   "DEF",
			out: map[rune]rune{'A': 'a', 'B': 'b', 'C': 'c', 'D': (-1), 'E': (-1), 'F': (-1)},
		},
	}
	for _, table := range tables {
		o, err := MakeTranslationTable(table.a, table.b, table.c)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(o, table.out) {
			t.Errorf("Got %+v instead of %+v", o, table.out)
		}
	}
}

func TestTranslate(t *testing.T) {
	tables := []struct {
		a      string
		b      string
		m      map[rune]rune
		strict bool
	}{
		{
			a:      "MAKE",
			b:      "LOVE",
			m:      map[rune]rune{'M': 'L', 'A': 'O', 'K': 'V'},
			strict: false,
		},
		{
			a:      "MAKE",
			b:      "LOVE",
			m:      map[rune]rune{'M': 'L', 'A': 'O', 'K': 'V', 'E': 'E'},
			strict: true,
		},
		{
			a:      "NOT",
			b:      "WAR",
			m:      map[rune]rune{'N': 'W', 'O': 'A', 'T': 'R'},
			strict: true,
		},
		{
			a:      "NOTARY",
			b:      "PUBLIC",
			m:      map[rune]rune{'N': 'P', 'O': 'U', 'T': 'B', 'A': 'L', 'R': 'I', 'Y': 'C'},
			strict: true,
		},
		{
			a:      "NOTARIES",
			b:      "PUBLIC",
			m:      map[rune]rune{'N': 'P', 'O': 'U', 'T': 'B', 'A': 'L', 'R': 'I', 'I': 'C', 'E': (-1), 'S': (-1)},
			strict: false,
		},
	}

	for _, table := range tables {
		o := Translate(table.a, table.m, table.strict)
		if o != table.b {
			t.Errorf("Expected %q, but got %q instead", table.b, o)
		}
	}
}
