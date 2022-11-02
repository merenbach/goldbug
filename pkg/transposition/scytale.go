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

package transposition

import "fmt"

// NewScytaleCipher creates and returns a new scytale.
func NewScytaleCipher(turns int) (*Cipher, error) {
	// Prepare a scytale cipher.
	// N.b.: The scytale cipher is a special case of a columnar transposition cipher
	//       with a key equal to an ascending consecutive integer sequence
	//       as long as the number of turns.
	//       A sequence with all the same digit may also work, but may depend on a stable sort.
	if turns < 1 {
		return nil, fmt.Errorf("turns must be greater than zero, but got %d", turns)
	}

	seq := make([]int, turns)
	for i := range seq {
		seq[i] = i
	}

	opts := []ConfigOption{
		WithMyszkowski(),
		WithIntegerKey(seq),
	}
	return NewCipher(opts...)
}
