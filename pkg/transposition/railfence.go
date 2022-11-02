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

// NewRailFencecipher creates and returns a new rail fence cipher.
func NewRailFenceCipher(rails int) (*Cipher, error) {
	// Prepare a rail fence cipher.
	// N.b.: The rail fence cipher is a special case of a columnar transposition cipher
	//       with Myszkowski transposition and a key equal to a zigzag sequence
	//       that converts the row count into the appropriate period.
	if rails < 1 {
		return nil, fmt.Errorf("rails must be greater than zero, but got %d", rails)
	}
	period := max(1, 2*(rails-1))

	key := zigzag(period)
	opts := []ConfigOption{
		WithMyszkowski(),
		WithIntegerKey(key),
	}
	return NewCipher(opts...)
}
