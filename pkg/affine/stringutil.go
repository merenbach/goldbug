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
	"unicode/utf8"

	"github.com/merenbach/goldbug/internal/mathutil"
	"github.com/merenbach/goldbug/internal/prng"
	"github.com/merenbach/goldbug/internal/sliceutil"
)

// Transform a string according to an affine equation.
func transform(s string, slope int, intercept int) (string, error) {
	if s == "" {
		return "", nil
	}

	m := utf8.RuneCountInString(s)

	for slope < 0 {
		slope += m
	}
	for intercept < 0 {
		intercept += m
	}

	if !mathutil.Coprime(m, slope) {
		return "", errors.New("slope and string length must be coprime")
	}

	lcg := &prng.LCG{
		Modulus:    m,
		Multiplier: 1,
		Increment:  slope,
		Seed:       intercept,
	}

	positions, err := lcg.Slice(m)
	if err != nil {
		return "", fmt.Errorf("couldn't initialize LCG: %w", err)
	}

	out, err := sliceutil.Backpermute([]rune(s), positions)
	if err != nil {
		return "", fmt.Errorf("couldn't backpermute input: %w", err)
	}

	return string(out), nil
}
