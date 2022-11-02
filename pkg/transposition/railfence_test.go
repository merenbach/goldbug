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

	"github.com/merenbach/goldbug/internal/fixture"
)

func TestRailFenceCipher_Encipher(t *testing.T) {
	var tables []struct {
		Rails int

		Input  string
		Output string
	}

	fixture.Load(t, &tables)
	for _, table := range tables {
		c, err := NewRailFenceCipher(table.Rails)
		if err != nil {
			t.Fatal("Could not create cipher:", err)
		}
		if out, err := c.Encipher(table.Input); err != nil {
			t.Error("Could not encipher:", err)
		} else if out != table.Output {
			t.Errorf("Expected %q to encipher to %q, but instead got %q", table.Input, table.Output, out)
		}
	}
}

func TestRailFenceCipher_Decipher(t *testing.T) {
	var tables []struct {
		Rails int

		Input  string
		Output string
	}

	fixture.Load(t, &tables)
	for _, table := range tables {
		c, err := NewRailFenceCipher(table.Rails)
		if err != nil {
			t.Fatal("Could not create cipher:", err)
		}
		if out, err := c.Decipher(table.Input); err != nil {
			t.Error("Could not decipher:", err)
		} else if out != table.Output {
			t.Errorf("Expected %q to decipher to %q, but instead got %q", table.Input, table.Output, out)
		}
	}
}

func TestRailFenceCipherReversibility(t *testing.T) {
	base := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	baseRunes := []rune(base)

	var counter int
	for msglen := 1; msglen < len(baseRunes); msglen++ {
		for rows := 1; rows < len(baseRunes); rows++ {
			counter++

			msg := string(baseRunes[:msglen])
			t.Logf("Iteration %d (rows = %d, message length = %d, message = %s)", counter, rows, msglen, msg)

			c, err := NewRailFenceCipher(rows)
			if err != nil {
				t.Error("Could not create new cipher:", err)
			}

			enciphered1, err := c.Encipher(msg)
			if err != nil {
				t.Error("Error:", err)
			}
			deciphered1, err := c.Decipher(enciphered1)
			if err != nil {
				t.Error("Error:", err)
			}
			if deciphered1 != msg {
				t.Errorf("Expected encipherment-then-decipherment of %q to be %q but got %q", enciphered1, msg, deciphered1)
			}

			enciphered2, err := c.Decipher(msg)
			if err != nil {
				t.Error("Error:", err)
			}
			deciphered2, err := c.Encipher(enciphered2)
			if err != nil {
				t.Error("Error:", err)
			}
			if deciphered2 != msg {
				t.Errorf("Expected decipherment-then-encipherment of %q to be %q but got %q", enciphered1, msg, deciphered1)
			}
		}
	}
}

// func ExampleCipher_enciphermentGrid() {
// 	c := Cipher{Rows: 3}
// 	out, err := c.enciphermentGrid("WEAREDISCOVEREDFLEEATONCE")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// 	fmt.Println(out)

// 	// Output:
// 	// W   E   C   R   L   T   E
// 	//  E R D S O E E F E A O C
// 	//   A   I   V   D   E   N
// }

// func ExampleCipher_deciphermentGrid() {
// 	c := Cipher{Rows: 3}
// 	out, err := c.deciphermentGrid("WECRLTEERDSOEEFEAOCAIVDEN")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// 	fmt.Println(out)

// 	// Output:
// 	// W   E   C   R   L   T   E
// 	//  E R D S O E E F E A O C
// 	//   A   I   V   D   E   N
// }
