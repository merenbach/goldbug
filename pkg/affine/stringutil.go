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

package affine

import (
	"errors"
	"fmt"
	"log"
	"unicode/utf8"

	"github.com/merenbach/goldbug/internal/mathutil"
	"github.com/merenbach/goldbug/internal/prng"
)

// Backpermute a string based on a slice of index values.
// Backpermute will return [E E O H L O] for inputs [H E L L O] and [1 1 4 0 2 4]
// Backpermute will return an error if the transform function returns any invalid string index values.
func backpermute(s string, ii []int) (string, error) {
	out := make([]rune, len(ii))

	rr := []rune(s)
	for n, i := range ii {
		if i < 0 || i >= len(rr) {
			return "", fmt.Errorf("Index %d out of bounds of interval [0, %d)", i, len(rr))
		}
		out[n] = rr[i] // or rr[ii[n]]
	}

	return string(out), nil
}

// Transform a string according to an affine equation.
func transform(s string, slope int, intercept int) (string, error) {
	m := utf8.RuneCountInString(s)

	for slope < 0 {
		slope += m
	}
	for intercept < 0 {
		intercept += m
	}

	if !mathutil.Coprime(m, slope) {
		return "", errors.New("Slope and string length must be coprime")
	}

	lcg := &prng.LCG{
		Modulus:    m,
		Multiplier: 1,
		Increment:  slope,
		Seed:       intercept,
	}

	positions, err := lcg.Slice(m)
	if err != nil {
		log.Println("Couldn't initialize LCG")
		return "", err
	}

	out, err := backpermute(s, positions)
	if err != nil {
		log.Println("Couldn't backpermute input")
		return "", err
	}

	return out, nil
}
