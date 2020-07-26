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

package prng

import (
	"reflect"
	"testing"
)

/*
TODO: based on https://cboard.cprogramming.com/general-discussions/151582-lagged-fibonacci-generator.html
can we simplify LFG?

Based on the Wikipedia article on Lagged Fibonacci generator, a type of pseudorandom number generator:
Code:
?
1
f(n) = ( f(n-j) OP f(n-k) ) mod m
where j < k < n, m is usually a power of two, and OP is a binary operator (addition, subtraction, multiplication, or exclusive or, in practice).

"Before the first cycle", f(0 .. k-1) must be computed or set first.

In practice, you keep f(n-k) to f(n-1) stored in an array. Then, the next value is obtained via
Code:
?
1
f[ n mod k ] = f[ (n + k - j) mod k ] OP f[ (n + k - k) mod k ]
where the oldest value in the array is always replaced with the new value, and you only need an array with k entries.
*/

func TestALFG(t *testing.T) {
	tables := []struct {
		*LFG
		expected []int
	}{
		{
			LFG: &LFG{
				Modulus: 1000,
				Seed:    []int{1, 1},
				Taps:    []int{1, 2},
			},
			expected: []int{2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 597, 584, 181, 765, 946, 711},
		},
		{
			LFG: &LFG{
				Modulus: 10,
				Seed:    []int{4, 8, 9, 4, 6, 0, 8},
				Taps:    []int{1, 2},
			},
			expected: []int{2, 7, 3, 0, 6, 8, 0, 9, 0, 3, 6, 4, 8, 9, 9, 3, 9, 0, 2, 7},
		},
		{
			LFG: &LFG{
				Modulus: 10,
				Seed:    []int{4, 8, 9, 4, 6, 0, 8},
				Taps:    []int{3, 7},
			},
			expected: []int{7, 1, 7, 7, 5, 2, 3, 0, 7, 2, 4, 7, 7, 4, 6, 0, 7, 4, 8, 4},
		},
	}

	for _, table := range tables {
		iter, err := table.LFG.IteratorA()
		if err != nil {
			t.Fatalf("Could not create ALFG iterator: %v", err)
		}

		output := make([]int, len(table.expected))
		for i := range output {
			output[i] = iter()
		}

		if !reflect.DeepEqual(output, table.expected) {
			t.Errorf("Expected %v but got %v", table.expected, output)
		}
	}
}

func TestMLFG(t *testing.T) {
	tables := []struct {
		*LFG
		expected []int
	}{
		{
			LFG: &LFG{
				Modulus: 10,
				Seed:    []int{1, 1},
				Taps:    []int{1, 2},
			},
			expected: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			LFG: &LFG{
				Modulus: 10,
				Seed:    []int{7, 7, 7, 9, 3, 1, 1},
				Taps:    []int{1, 2},
			},
			expected: []int{9, 9, 3, 7, 3, 1, 9, 1, 7, 1, 1, 3, 9, 9, 7, 7, 1, 3, 7, 1},
		},
		{
			LFG: &LFG{
				Modulus: 10,
				Seed:    []int{7, 7, 7, 9, 3, 1, 1},
				Taps:    []int{3, 7},
			},
			expected: []int{7, 3, 9, 9, 9, 3, 9, 1, 9, 1, 3, 7, 7, 3, 3, 9, 3, 1, 3, 9},
		},
	}

	for _, table := range tables {
		iter, err := table.LFG.IteratorM()
		if err != nil {
			t.Fatalf("Could not create MLFG iterator: %v", err)
		}

		output := make([]int, len(table.expected))
		for i := range output {
			output[i] = iter()
		}

		if !reflect.DeepEqual(output, table.expected) {
			t.Errorf("Expected %v but got %v", table.expected, output)
		}
	}
}
