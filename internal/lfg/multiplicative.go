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

import "errors"

// A Multiplicative LFG is a lagged Fibonacci generator that uses multiplication for new elements.
// A Multiplicative LFG is a type of pseudo-random number generator (PRNG).
// A Multiplicative LFG may not be cryptographically secure.
// A Multiplicative LFG uses 1-indexed taps.
type Multiplicative struct {
	Modulus int
	Seed    []int
	Taps    []int
}

// Iterate across a multiplicative lagged Fibonacci generator (MLFG) sequence.
func (g *Multiplicative) iterate() (func() int, error) {
	// if err := g.validate(); err != nil {
	// 	return nil, err
	// }

	// Ensure at least one item in seed is odd.
	if !all(g.Seed, func(i int) bool {
		return i%2 != 0
	}) {
		return nil, errors.New("All MLFG seed values must be odd")
	}

	return iterateLagTable(g.Modulus, g.Seed, func(lagTable []int) int {
		e := 1
		for _, t := range g.Taps {
			e *= lagTable[t-1]
		}
		return e
	}), nil
}

// Slice of LCG values.
func (g *Multiplicative) Slice(n int) ([]int, error) {
	iter, err := g.iterate()
	if err != nil {
		return nil, err
	}
	return take(n, iter), nil
}
