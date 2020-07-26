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
)

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
		{
			m:          100,
			a:          17,
			c:          43,
			seed:       27,
			hulldobell: false,
			expected:   []int{27, 2, 77, 52, 27},
		},
		{
			m:          64,
			a:          13,
			c:          0,
			seed:       1,
			hulldobell: false,
			expected:   []int{1, 13, 41, 21, 17, 29, 57, 37, 33, 45, 9, 53, 49, 61, 25, 5, 1},
		},
		{
			m:          64,
			a:          13,
			c:          0,
			seed:       2,
			hulldobell: false,
			expected:   []int{2, 26, 18, 42, 34, 58, 50, 10, 2},
		},
		{
			m:          64,
			a:          13,
			c:          0,
			seed:       3,
			hulldobell: false,
			expected:   []int{3, 39, 59, 63, 51, 23, 43, 47, 35, 7, 27, 31, 19, 55, 11, 15, 3},
		},
		{
			m:          64,
			a:          13,
			c:          0,
			seed:       4,
			hulldobell: false,
			expected:   []int{4, 52, 36, 20, 4},
		},
	}

	var lcg LCG
	for _, table := range tables {
		lcg.Reset()
		lcg.Modulus = table.m
		lcg.Multiplier = table.a
		lcg.Increment = table.c
		lcg.Seed = table.seed

		err := lcg.HullDobell()
		if (err != nil) == table.hulldobell {
			if err != nil {
				t.Errorf("LCG %+v satisfies Hull-Dobell, contrary to expectations", lcg)
			} else {
				t.Errorf("LCG %+v fails Hull-Dobell, contrary to expectations: %s", lcg, err)
			}
		}

		for idx, e := range table.expected {
			if f := lcg.Next(); e != f {
				t.Errorf("expected item %d from LCG %+v to equal %d, but got %d instead", idx, lcg, e, f)
			}
		}

		if lcg.Counter() != len(table.expected) {
			t.Errorf("Expected %d outputs, but LCG counter is at %d", len(table.expected), lcg.Counter())
		}
	}
}

func TestLCGCopy(t *testing.T) {
	tables := []struct {
		m    int
		a    int
		c    int
		seed int
	}{
		{
			m:    100,
			a:    17,
			c:    43,
			seed: 27,
		},
		{
			m:    64,
			a:    13,
			c:    0,
			seed: 1,
		},
		{
			m:    64,
			a:    13,
			c:    0,
			seed: 2,
		},
		{
			m:    64,
			a:    13,
			c:    0,
			seed: 3,
		},
		{
			m:    64,
			a:    13,
			c:    0,
			seed: 4,
		},
	}

	for _, table := range tables {
		lcg := &LCG{
			Modulus:    table.m,
			Multiplier: table.a,
			Increment:  table.c,
			Seed:       table.seed,
		}

		_ = lcg.Next()

		cp := lcg.Copy()
		_ = cp.Next()

		if !reflect.DeepEqual(lcg, cp) {
			t.Errorf("Expected %+v to be equal to %+v, but it was not", lcg, cp)
		}
	}
}
