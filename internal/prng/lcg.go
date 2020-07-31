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

	"github.com/merenbach/goldbug/internal/mathutil"
)

// An LCG is a linear congruential generator.
// An LCG is a type of pseudo-random number generator (PRNG).
// An LCG may not be cryptographically secure.
type LCG struct {
	Modulus    int // m
	Multiplier int // a
	Increment  int // c
	Seed       int // X_0
}

// // Lehmer RNG validation.
// // TODO
// func (g *LCG) Lehmer() bool {
// 	return g.Increment == 0
// }

// Copy this LCG into a fresh duplicate with pristine state.
func (g *LCG) Copy() *LCG {
	return &LCG{
		Modulus:    g.Modulus,
		Multiplier: g.Multiplier,
		Increment:  g.Increment,
		Seed:       g.Seed,
	}
}

// HullDobell tests for compliance with the Hull-Dobell theorem.
// The error parameter, if set, will contain the first-found failing constraint.
// When c != 0, this test passing means that the cycle is equal to g.multiplier.
func (g *LCG) HullDobell() error {
	switch {
	case !mathutil.Coprime(g.Modulus, g.Increment):
		return errors.New("multiplier and increment should be coprime")
	case !mathutil.Regular(g.Multiplier-1, g.Modulus):
		return errors.New("prime factors of modulus should also divide multiplier-minus-one")
	case g.Modulus%4 == 0 && (g.Multiplier-1)%4 != 0:
		return errors.New("if 4 divides modulus, 4 should divide multiplier-minus-one")
	default:
		return nil
	}
}

// Iterator across LCG values.
func (g *LCG) Iterator() (func() int, error) {
	if g.Modulus == 0 {
		return nil, errors.New("modulus must be greater than zero")
	}
	if g.Multiplier == 0 {
		return nil, errors.New("multiplier must be greater than zero")
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
	iter, err := g.Iterator()
	if err != nil {
		return nil, err
	}

	out := make([]int, n)
	for i := range out {
		out[i] = iter()
		// state = (state*g.Multiplier + g.Increment) % g.Modulus
	}

	return out, nil
}
