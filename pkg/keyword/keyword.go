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

package keyword

import (
	"github.com/merenbach/goldbug/pkg/masc2"
	"github.com/merenbach/goldbug/pkg/simple"
)

// NewCipher creates and returns a new cipher.
func NewCipher(keyword string, opts ...masc2.ConfigOption) (*simple.Cipher, error) {
	c := masc2.NewConfig(opts...)
	ctAlphabet, _ := Transform([]rune(c.Alphabet()), []rune(keyword))

	return simple.NewCipher(string(ctAlphabet), opts...)
}
