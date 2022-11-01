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
	"errors"
	"fmt"

	"github.com/merenbach/goldbug/internal/mathutil"
)

// HullDobell tests for compliance with the Hull-Dobell theorem.
// The error parameter, if set, will contain the first-found failing constraint.
// When c != 0, this test passing means that the cycle is equal to g.multiplier.
func hullDobell(m int, a int, c int) error {
	switch {
	case !mathutil.Coprime(m, c):
		return errors.New("modulus and increment should be relatively prime")
	case !mathutil.Regular(a-1, m):
		return errors.New("prime factors of modulus should also divide multiplier less one")
	case m%4 == 0 && (a-1)%4 != 0:
		return errors.New("if 4 divides modulus, 4 should divide multiplier less one")
	default:
		return nil
	}
}

// An LCG is a linear congruential generator.
// An LCG is a type of pseudo-random number generator (PRNG).
// An LCG may not be cryptographically secure.
type LCG struct {
	Modulus    int // m
	Multiplier int // a
	Increment  int // c
	Seed       int // X_0
}

// // Multiplicative RNGs have a zero increment.
// // Multiplicative RNGs are also called Lehmer RNGs and Park-Miller RNGs.
// func (g *LCG) multiplicative() bool {
// 	return g.Increment == 0
// }

// Mixed RNGs have a non-zero increment.
func (g *LCG) mixed() bool {
	return g.Increment != 0
}

// Validate settings for this generator.
func (g *LCG) validate() error {
	if g.Modulus <= 0 {
		return errors.New("modulus must be greater than zero")
	}
	if g.Multiplier <= 0 {
		return errors.New("multiplier must be greater than zero")
	}

	if !g.mixed() {
		return nil
	}
	return hullDobell(g.Modulus, g.Multiplier, g.Increment)
}

// Iterator across LCG values.
func (g *LCG) iterator() (func() int, error) {
	if err := g.validate(); err != nil {
		return nil, fmt.Errorf("could not validate LCG: %w", err)
	}

	state := g.Seed % g.Modulus
	return func() int {
		prev := state
		state = (state*g.Multiplier + g.Increment) % g.Modulus
		return prev
	}, nil
}

// Slice of LCG values.
func (g *LCG) Slice(n int) ([]int, error) {
	iter, err := g.iterator()
	if err != nil {
		return nil, err
	}
	return take(n, iter), nil
}
