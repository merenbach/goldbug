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

// GCD returns the greatest common divisor for the provided parameters.
// http://anh.cs.luc.edu/331/notes/xgcd.pdf
func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// XGCD calculates the extended GCD of two numbers.
// http://anh.cs.luc.edu/331/notes/xgcd.pdf
func XGCD(a int, b int) (int, int, int) {
	prevx, x := 1, 0
	prevy, y := 0, 1
	for b != 0 {
		q := a / b
		x, prevx = prevx-q*x, x
		y, prevy = prevy-q*y, y
		a, b = b, a%b
	}
	return a, prevx, prevy
}

// ModInv returns the modular multiplicative inverse of two numbers.
// ModInv will always return a non-negative number if a modular multiplicative inverse is found.
// ModInv returns (-1) if no modular multiplicative inverse is found.
func ModInv(a int, m int) int {
	g, x, _ := XGCD(a, m)
	if g != 1 {
		return (-1)
	}
	for x < 0 {
		x += m
	}
	return x % m
}

// Coprime tests if two numbers `a` and `b` are relatively prime.
// The order of the parameters does not matter.
func Coprime(a int, b int) bool {
	return gcd(a, b) == 1
}

// Regular tests if all prime factors of `b` also divide `a`.
// Note that the order of the parameters is important, as `a` may have additional prime factors.
// TODO: just return true if either a or b is zero?
func Regular(a int, b int) bool {
	if b == 0 {
		panic("Parameter `b` must be nonzero.")
	}

	for a != 1 {
		a = gcd(a, b)
		b /= a
	}
	return b == 1
}

// Modulo returns the remainder of a Euclidean division operation.
// This works around https://github.com/golang/go/issues/448
// func Modulo(a int, b int) int {
// 	return ((a % b) + b) % b
// }
