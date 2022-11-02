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

package transposition

import (
	"testing"
)

func TestScytaleyyCipher(t *testing.T) {
	tables := []struct {
		ciphertext string
		plaintext  string
		turns      int
	}{
		{
			plaintext:  "Iamhurtverybadlyhelp",
			ciphertext: "Iryyatbhmvaehedlurlp",
			turns:      5,
		},
		{
			ciphertext: "Tolhxaejzquyumdipocsgkobvreorwtnhfe",
			plaintext:  "Thequickbrownfoxjumpsoverthelazydog",
			turns:      14,
		},
		// {
		// 	ciphertext: "TWLQHTHIAUTHELCAXEMLERXRELIEIEETNAGSTATTNTIKHEOANEEIRBGPSGEC",
		// 	plaintext:  "THEMEETINGWILLTAKEPLACEINTHESQUAREATEIGHTXXXXXXXXXXXXXXXXXXX",
		// 	turns:      6,
		// },
		{
			ciphertext: "HENTEIDTLAEAPMRCMUAK",
			plaintext:  "HELPMEIAMUNDERATTACK",
			turns:      5,
		},
		// {
		// 	ciphertext: "AALUHNHSEDFYMNAGIGIHAOFZ",
		// 	plaintext:  "AMANLAUGHINGHISHEADOFFYZ",
		// 	turns:      2,
		// },
		// {
		// 	ciphertext: "HELLO, WORLD!",
		// 	plaintext:  "HELLO, WORLD!",
		// 	turns:      1,
		// },
	}

	for i, table := range tables {
		t.Logf("Running test %d of %d...", i+1, len(tables))

		c, err := NewScytaleCipher(table.turns)
		if err != nil {
			t.Fatal("Could not create cipher:", err)
		}

		if out, err := c.Decipher(table.ciphertext); err != nil {
			t.Error("Could not decipher:", err)
		} else if out != table.plaintext {
			t.Errorf("Expected %q to decipher to %q, but instead got %q", table.ciphertext, table.plaintext, out)
		}

		if out, err := c.Encipher(table.plaintext); err != nil {
			t.Error("Could not encipher:", err)
		} else if out != table.ciphertext {
			t.Errorf("Expected %q to encipher to %q, but instead got %q", table.plaintext, table.ciphertext, out)
		}
	}
}

// func TestCipherReversibility(t *testing.T) {
// 	base := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
// 	baseRunes := []rune(base)

// 	var counter int
// 	for msglen := 1; msglen < len(baseRunes); msglen++ {
// 		for turns := 1; turns < len(baseRunes); turns++ {
// 			counter++

// 			msg := string(baseRunes[:msglen])
// 			t.Logf("Iteration %d (turns = %d, message length = %d, message = %s)", counter, turns, msglen, msg)

// 			c := Cipher{Turns: turns}

// 			enciphered1 := c.Encipher(msg)
// 			deciphered1 := c.Decipher(enciphered1)
// 			if deciphered1 != msg {
// 				t.Errorf("Expected encipherment-then-decipherment of %q to be %q but got %q", enciphered1, msg, deciphered1)
// 			}

// 			enciphered2 := c.Decipher(msg)
// 			deciphered2 := c.Encipher(enciphered2)
// 			if deciphered2 != msg {
// 				t.Errorf("Expected decipherment-then-encipherment of %q to be %q but got %q", enciphered1, msg, deciphered1)
// 			}
// 		}
// 	}
// }

// // func ExampleCipher_EnciphermentGrid() {
// // 	c := Cipher{Turns: 3}
// // 	fmt.Println(c.EnciphermentGrid("WEAREDISCOVEREDFLEEATONCE"))

// // 	// Output:
// // 	// W   E   C   R   L   T   E
// // 	//  E R D S O E E F E A O C
// // 	//   A   I   V   D   E   N
// // }

// // func ExampleCipher_DeciphermentGrid() {
// // 	c := Cipher{Turns: 3}
// // 	fmt.Println(c.DeciphermentGrid("WECRLTEERDSOEEFEAOCAIVDEN"))

// // 	// Output:
// // 	// W   E   C   R   L   T   E
// // 	//  E R D S O E E F E A O C
// // 	//   A   I   V   D   E   N
// // }
