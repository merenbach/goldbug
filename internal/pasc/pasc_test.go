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

package pasc

import (
	"testing"
)

const defaultPolyalphabeticAlphabet = Alphabet

func runPolyalphabeticEncipherTest(t *testing.T, input string, expected string, c *VigenereFamilyCipher) {
	output := c.EncipherString(input)
	if output != expected {
		t.Errorf("ciphertext %q for input %q was incorrect; wanted %q", output, input, expected)
	}
}

func runPolyalphabeticDecipherTest(t *testing.T, input string, expected string, c *VigenereFamilyCipher) {
	output := c.DecipherString(input)
	if output != expected {
		t.Errorf("plaintext %q for input %q was incorrect; wanted: %q", output, input, expected)
	}
}

// func runPolyalphabeticReciprocalTests(t *testing.T, plaintext, ciphertext string, c VigenereFamilyCipher, strict bool) {
// 	encrypted := c.EncipherString(plaintext, strict)
// 	decrypted := c.DecipherString(ciphertext, strict)
// 	if string(encrypted) != ciphertext {
// 		t.Errorf("ciphertext %q was incorrect; wanted %q", encrypted, ciphertext)
// 	}
// 	if string(decrypted) != plaintext {
// 		t.Errorf("plaintext %q was incorrect; wanted: %q", decrypted, plaintext)
// 	}
// }
