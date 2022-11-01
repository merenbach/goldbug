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

package prng

import (
	"reflect"
	"testing"

	"github.com/merenbach/goldbug/internal/iterutil"
)

func TestHullDobell(t *testing.T) {
	tables := []struct {
		expectSuccess bool
		m             int
		a             int
		c             int
	}{
		{true, 1, 1, 0},
		{true, 26, 1, 7},
		{true, 26, 1, 3},
		{true, 26, 1, 1},
		{false, 0, 1, 0},
		{false, 2, 1, 0},
		{false, 26, 1, 13},
		{false, 26, 2, 7},
		{false, 28, 3, 7},
	}

	for i, table := range tables {
		t.Logf("Running test %d of %d...", i+1, len(tables))
		err := hullDobell(table.m, table.a, table.c)
		if table.expectSuccess && err != nil {
			t.Error("Expected success and got failure:", err)
		} else if !table.expectSuccess && err == nil {
			t.Error("Expected failure and got success")
		}
	}
}

func TestLCG(t *testing.T) {
	// Some sequences for verification borrowed from: <https://www.mi.fu-berlin.de/inf/groups/ag-tech/teaching/2012_SS/L_19540_Modeling_and_Performance_Analysis_with_Simulation/06.pdf>
	tables := []struct {
		m          int
		a          int
		c          int
		seed       int
		hulldobell bool
		expected   []int
	}{
		// {
		// 	m:        100,
		// 	a:        17,
		// 	c:        43,
		// 	seed:     27,
		// 	expected: []int{27, 2, 77, 52, 27},
		// },
		{
			m:        64,
			a:        13,
			c:        0,
			seed:     1,
			expected: []int{1, 13, 41, 21, 17, 29, 57, 37, 33, 45, 9, 53, 49, 61, 25, 5, 1},
		},
		{
			m:        64,
			a:        13,
			c:        0,
			seed:     2,
			expected: []int{2, 26, 18, 42, 34, 58, 50, 10, 2},
		},
		{
			m:        64,
			a:        13,
			c:        0,
			seed:     3,
			expected: []int{3, 39, 59, 63, 51, 23, 43, 47, 35, 7, 27, 31, 19, 55, 11, 15, 3},
		},
		{
			m:        64,
			a:        13,
			c:        0,
			seed:     4,
			expected: []int{4, 52, 36, 20, 4},
		},
	}

	var lcg LCG
	for _, table := range tables {
		lcg.Modulus = table.m
		lcg.Multiplier = table.a
		lcg.Increment = table.c
		lcg.Seed = table.seed

		if g, err := lcg.Iterator(); err != nil {
			t.Errorf("Error for LCG %+v: %+v", lcg, err)
		} else {
			out := iterutil.Take(len(table.expected), g)
			if !reflect.DeepEqual(out, table.expected) {
				t.Errorf("expected LCG %+v to produce values %+v, but got %+v instead", lcg, table.expected, out)
			}
		}
	}
}
