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

package lfg

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

// All slice items must pass the provided test for this function to return true.
func all(ii []int, f func(int) bool) bool {
	for _, i := range ii {
		if !f(i) {
			return false
		}
	}
	return true
}

func iterateLagTable(m int, seed []int, f func([]int) int) func() int {
	// current set of elements, effectively a FIFO queue
	lagTable := make([]int, len(seed))
	copy(lagTable, seed)
	last := len(lagTable) - 1

	return func() int {
		o := f(lagTable) % m
		copy(lagTable, lagTable[1:])
		lagTable[last] = o
		return o
	}
}

// Take a slice of integers from a generating function.
func take(n int, f func() int) []int {
	out := make([]int, n)
	for i := range out {
		out[i] = f()
	}
	return out
}

// // An LFG is a lagged Fibonacci generator.
// // An LFG is a type of pseudo-random number generator (PRNG).
// // An LFG may not be cryptographically secure.
// // An LFG uses 1-indexed taps.
// type LFG struct {
// 	Modulus int
// 	Seed    []int
// 	Taps    []int
// }

// // Validate the settings for this LFG.
// func (g *LFG) validate() error {
// 	for _, t := range g.Taps {
// 		switch {
// 		case t < 1:
// 			return fmt.Errorf("Tap value %d must be greater than zero", t)
// 		case t > len(g.Seed):
// 			return fmt.Errorf("Tap value %d exceeds seed length %d", t, len(g.Seed))
// 		}
// 	}

// 	return nil
// }

// // Iterator across LCG values.
// func (g *LCG) iterator() (func() int, error) {
// 	if err := g.validate(); err != nil {
// 		return nil, err
// 	}

// 	state := g.Seed % g.Modulus
// 	return func() int {
// 		prev := state
// 		state = (state*g.Multiplier + g.Increment) % g.Modulus
// 		return prev
// 	}, nil
// }

// // Slice of LCG values.
// func (g *LCG) Slice(n int) ([]int, error) {
// 	iter, err := g.iterator()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return take(n, iter), nil
// }
