// Copyright 2020 Andrew Merenbach
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

package railfence

import (
	"fmt"
	"testing"
)

func TestCipher(t *testing.T) {
	tables := []struct {
		ciphertext string
		plaintext  string
		rails      int
	}{
		{
			ciphertext: "WECRLTEERDSOEEFEAOCAIVDEN",
			plaintext:  "WEAREDISCOVEREDFLEEATONCE",
			rails:      3,
		},
		{
			ciphertext: "IA_EZS_ELYLK_UZERLIPL",
			plaintext:  "I_REALLY_LIKE_PUZZLES",
			rails:      3,
		},
		{
			ciphertext: "MEMTNGTXETEOIHQZ",
			plaintext:  "MEETMETONIGHTQXZ",
			rails:      2,
		},
		{
			ciphertext: "MMNTETEOIHQZETGX",
			plaintext:  "MEETMETONIGHTQXZ",
			rails:      3,
		},
		{
			ciphertext: "AALUHNHSEDFYMNAGIGIHAOFZ",
			plaintext:  "AMANLAUGHINGHISHEADOFFYZ",
			rails:      2,
		},
		{
			ciphertext: "HELLO, WORLD!",
			plaintext:  "HELLO, WORLD!",
			rails:      1,
		},
	}

	for _, table := range tables {
		c := Cipher{Rails: table.rails}
		if out := c.Decipher(table.ciphertext); out != table.plaintext {
			t.Errorf("Expected %q to decipher to %q, but instead got %q", table.ciphertext, table.plaintext, out)
		}

		if out := c.Encipher(table.plaintext); out != table.ciphertext {
			t.Errorf("Expected %q to encipher to %q, but instead got %q", table.plaintext, table.ciphertext, out)
		}
	}
}

func TestCipherReversibility(t *testing.T) {
	base := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	baseRunes := []rune(base)

	var counter int
	for msglen := 1; msglen < len(baseRunes); msglen++ {
		for rails := 1; rails < len(baseRunes); rails++ {
			counter++

			msg := string(baseRunes[:msglen])
			t.Logf("Iteration %d (rails = %d, message length = %d, message = %s)", counter, rails, msglen, msg)

			c := Cipher{Rails: rails}

			enciphered1 := c.Encipher(msg)
			deciphered1 := c.Decipher(enciphered1)
			if deciphered1 != msg {
				t.Errorf("Expected encipherment-then-decipherment of %q to be %q but got %q", enciphered1, msg, deciphered1)
			}

			enciphered2 := c.Decipher(msg)
			deciphered2 := c.Encipher(enciphered2)
			if deciphered2 != msg {
				t.Errorf("Expected decipherment-then-encipherment of %q to be %q but got %q", enciphered1, msg, deciphered1)
			}
		}
	}
}

func ExampleCipher_EnciphermentGrid() {
	c := Cipher{Rails: 3}
	fmt.Println(c.EnciphermentGrid("WEAREDISCOVEREDFLEEATONCE"))

	// Output:
	// W   E   C   R   L   T   E
	//  E R D S O E E F E A O C
	//   A   I   V   D   E   N
}

func ExampleCipher_DeciphermentGrid() {
	c := Cipher{Rails: 3}
	fmt.Println(c.DeciphermentGrid("WECRLTEERDSOEEFEAOCAIVDEN"))

	// Output:
	// W   E   C   R   L   T   E
	//  E R D S O E E F E A O C
	//   A   I   V   D   E   N
}
