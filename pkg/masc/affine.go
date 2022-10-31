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

package masc

import (
	"fmt"

	"github.com/merenbach/goldbug/internal/sliceutil"
)

// NewAffineCipher creates and returns a new affine cipher.
func NewAffineCipher(slope int, intercept int, opts ...ConfigOption) (*SimpleCipher, error) {
	c := NewConfig(opts...)

	ctAlphabet, err := sliceutil.Affine([]rune(c.alphabet), slope, intercept)
	if err != nil {
		return nil, fmt.Errorf("could not transform alphabet: %w", err)
	}

	return NewSimpleCipher(string(ctAlphabet), opts...)
}
