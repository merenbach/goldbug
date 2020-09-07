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
	"errors"
	"fmt"
)

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

	return func() int {
		o := f(lagTable) % m
		copy(lagTable, lagTable[1:])
		lagTable[len(lagTable)-1] = o
		return o
	}
}

// An LFG is a lagged Fibonacci generator.
// An LFG is a type of pseudo-random number generator (PRNG).
// An LFG may not be cryptographically secure.
// An LFG uses 1-indexed taps.
type LFG struct {
	Modulus int
	Seed    []int
	Taps    []int
}

// Validate the settings for this LFG.
func (g *LFG) validate() error {
	for _, t := range g.Taps {
		switch {
		case t < 1:
			return fmt.Errorf("Tap value %d must be greater than zero", t)
		case t > len(g.Seed):
			return fmt.Errorf("Tap value %d exceeds seed length %d", t, len(g.Seed))
		}
	}

	return nil
}

// IteratorA returns an additive lagged Fibonacci generator (ALFG) function.
func (g *LFG) IteratorA() (func() int, error) {
	if err := g.validate(); err != nil {
		return nil, err
	}

	// Ensure at least one item in seed is odd.
	if all(g.Seed, func(i int) bool {
		return i%2 == 0
	}) {
		return nil, errors.New("At least one ALFG seed value must be odd")
	}

	return iterateLagTable(g.Modulus, g.Seed, func(lagTable []int) int {
		e := 0
		for _, t := range g.Taps {
			e += lagTable[t-1]
		}
		return e
	}), nil
}

// IteratorM returns a multiplicative lagged Fibonacci generator (MLFG) function.
func (g *LFG) IteratorM() (func() int, error) {
	if err := g.validate(); err != nil {
		return nil, err
	}

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
