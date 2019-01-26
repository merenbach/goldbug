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
	"errors"
)

// Modulus returns the remainder of a Euclidean division operation.
// This works around https://github.com/golang/go/issues/448
/*func Modulus(a, b int) int {
	return ((a % b) + b) % b
}*/

// // abs returns the absolute value of an integer.
// func abs(a int) int {
// 	if a < 0 {
// 		return -a
// 	}
// 	return a
// }

// gcd returns the greatest common divisor for the provided parameters.
func gcd(a, b uint) uint {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Coprime tests if two numbers `a` and `b` are relatively prime.
// The order of the parameters does not matter.
func coprime(a, b uint) bool {
	return gcd(a, b) == 1
}

// Regular tests if all prime factors of `a` also divide `b`.
// Note that the order of the parameters is important, as `b` may have additional prime factors.
// TODO: just return true if either a or b is zero?
func regular(a, b uint) bool {
	if a == 0 {
		panic("Parameter `a` must be nonzero.")
	}
	if b == 0 {
		return true
	}

	for b != 1 {
		b = gcd(a, b)
		a /= b
	}
	return a == 1
}

// An LCG is a linear congruential generator, a method of generating pseudo-random numbers.
// TODO: this may be better as a simple generator function, rather than a struct.
type LCG struct {
	Modulus    uint // m
	Multiplier uint // a
	Increment  uint // c
	Seed       uint // X_0
}

// HullDobell tests for compliance with the Hull-Dobell theorem.
// The error parameter, if set, will contain the first-found failing constraint.
// When c != 0, this test passing means that the cycle is equal to g.multiplier.
func (g LCG) HullDobell() error {
	switch {
	case !coprime(g.Modulus, g.Increment):
		return errors.New("multiplier and increment should be coprime")
	case !regular(g.Modulus, g.Multiplier-1):
		return errors.New("prime factors of modulus should also divide multiplier-minus-one")
	case g.Modulus%4 == 0 && (g.Multiplier-1)%4 != 0:
		return errors.New("if 4 divides modulus, 4 should divide multiplier-minus-one")
	default:
		return nil
	}
}

// Iterator returns a generator function to iterate over the values of the LCG.
func (g LCG) Iterator() func() uint {
	if g.Modulus == 0 {
		panic("modulus must be greater than zero")
	}
	if g.Multiplier == 0 {
		panic("multiplier must be greater than zero")
	}

	state := g.Seed % g.Modulus
	return func() uint {
		prev := state
		state = (state*g.Multiplier + g.Increment) % g.Modulus
		return prev
	}
}
