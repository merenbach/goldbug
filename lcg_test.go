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

package main

import (
	"testing"
)

func TestGCD(t *testing.T) {
	tables := []struct {
		a        uint
		b        uint
		expected uint
	}{
		{3, 5, 1},
		{15, 5, 5},
		{24, 16, 8},
		{2, 4, 2},
		{2, 22, 2},
		{6, 15, 3},
		{14, 28, 14},
	}
	for _, table := range tables {
		if out := gcd(table.a, table.b); out != table.expected {
			t.Errorf("expected GCD of %d and %d to be %d, but got %d instead", table.a, table.b, table.expected, out)
		}
	}
}

func TestCoprime(t *testing.T) {
	tables := []struct {
		a        uint
		b        uint
		expected bool
	}{
		{3, 5, true},
		{7, 20, true},
		{14, 15, true},
		{172, 17, true},
		{2, 4, false},
		{2, 22, false},
		{3, 15, false},
		{14, 28, false},
	}
	for _, table := range tables {
		if out := coprime(table.a, table.b); out != table.expected {
			if table.expected {
				t.Errorf("%d and %d were expected to be comprime, but were not", table.a, table.b)
			} else {
				t.Errorf("%d and %d were not expected to be comprime, but were", table.a, table.b)
			}
		}
	}
}

func TestRegular(t *testing.T) {
	tables := []struct {
		a        uint
		b        uint
		expected bool
	}{
		{98, 168, true},
		{24, 6, true},
		{6, 24, true},
		{12, 18, true},
		{3, 0, true},
		{1, 1, true},
		{168, 98, false},
		{168, 132, false},
		{132, 168, false},
		{2, 1, false},
	}
	for _, table := range tables {
		if out := regular(table.a, table.b); out != table.expected {
			if table.expected {
				t.Errorf("%d and %d were expected to be regular, but were not", table.a, table.b)
			} else {
				t.Errorf("%d and %d were not expected to be regular, but were", table.a, table.b)
			}
		}
	}
}

func TestLCG(t *testing.T) {
	// Some sequences for verification borrowed from: <https://www.mi.fu-berlin.de/inf/groups/ag-tech/teaching/2012_SS/L_19540_Modeling_and_Performance_Analysis_with_Simulation/06.pdf>
	tables := []struct {
		m          uint
		a          uint
		c          uint
		seed       uint
		hulldobell bool
		expected   []uint
	}{
		{
			m:          100,
			a:          17,
			c:          43,
			seed:       27,
			hulldobell: false,
			expected:   []uint{27, 2, 77, 52, 27},
		},
		{
			m:          64,
			a:          13,
			c:          0,
			seed:       1,
			hulldobell: false,
			expected:   []uint{1, 13, 41, 21, 17, 29, 57, 37, 33, 45, 9, 53, 49, 61, 25, 5, 1},
		},
		{
			m:          64,
			a:          13,
			c:          0,
			seed:       2,
			hulldobell: false,
			expected:   []uint{2, 26, 18, 42, 34, 58, 50, 10, 2},
		},
		{
			m:          64,
			a:          13,
			c:          0,
			seed:       3,
			hulldobell: false,
			expected:   []uint{3, 39, 59, 63, 51, 23, 43, 47, 35, 7, 27, 31, 19, 55, 11, 15, 3},
		},
		{
			m:          64,
			a:          13,
			c:          0,
			seed:       4,
			hulldobell: false,
			expected:   []uint{4, 52, 36, 20, 4},
		},
	}

	for _, table := range tables {
		lcg := LCG{
			Modulus:    table.m,
			Multiplier: table.a,
			Increment:  table.c,
			Seed:       table.seed,
		}
		err := lcg.HullDobell()
		if (err != nil) == table.hulldobell {
			if err != nil {
				t.Errorf("LCG %+v satisfies Hull-Dobell, contrary to expectations", lcg)
			} else {
				t.Errorf("LCG %+v fails Hull-Dobell, contrary to expectations: %s", lcg, err)
			}
		}

		iter := lcg.Iterator()
		for idx, e := range table.expected {
			if f := iter(); e != f {
				t.Errorf("expected item %d from LCG %+v to equal %d, but got %d instead", idx, lcg, e, f)
			}
		}
	}
}
