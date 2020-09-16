// Copyright 2019 Andrew Merenbach
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

package mathutil

import "testing"

func TestGCD(t *testing.T) {
	tables := []struct {
		a        int
		b        int
		expected int
	}{
		{0, 0, 0},
		{1, 1, 1},
		{0, 1, 1},
		{0, 2, 2},
		{1, 0, 1},
		{2, 0, 2},
		{1, 2, 1},
		{2, 1, 1},
		{0, 10, 10},
		{10, 0, 10},
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
		a        int
		b        int
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
		if out := Coprime(table.a, table.b); out != table.expected {
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
		a        int
		b        int
		expected bool
	}{
		{168, 98, true},
		{6, 24, true},
		{24, 6, true},
		{49, 7, true},
		{7, 49, true},
		{7, 98, false},
		{98, 7, true},
		{18, 12, true},
		{0, 3, true},
		{0, 1, true},
		{1, 1, true},
		{70, 1, true},
		{98, 168, false},
		{132, 168, false},
		{168, 132, false},
		{1, 2, false},
	}
	for _, table := range tables {
		if out := Regular(table.a, table.b); out != table.expected {
			if table.expected {
				t.Errorf("%d and %d were expected to be regular, but were not", table.a, table.b)
			} else {
				t.Errorf("%d and %d were not expected to be regular, but were", table.a, table.b)
			}
		}
	}
}
