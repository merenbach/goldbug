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

package variantbeaufort

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestCipher_Encipher(t *testing.T) {
	testdata, err := ioutil.ReadFile(filepath.Join("testdata", "cipher_encipher.json"))
	if err != nil {
		t.Fatal("Could not read testdata fixture:", err)
	}

	var tables []struct {
		Cipher

		Input  string
		Output string
	}
	if err := json.Unmarshal(testdata, &tables); err != nil {
		t.Fatal("Could not unmarshal testdata:", err)
	}

	for _, table := range tables {
		if out, err := table.Encipher(table.Input); err != nil {
			t.Error("Could not encipher:", err)
		} else if out != table.Output {
			t.Errorf("Expected %q to encipher to %q, but instead got %q", table.Input, table.Output, out)
		}
	}
}

func TestCipher_Decipher(t *testing.T) {
	testdata, err := ioutil.ReadFile(filepath.Join("testdata", "cipher_decipher.json"))
	if err != nil {
		t.Fatal("Could not read testdata fixture:", err)
	}

	var tables []struct {
		Cipher

		Input  string
		Output string
	}
	if err := json.Unmarshal(testdata, &tables); err != nil {
		t.Fatal("Could not unmarshal testdata:", err)
	}

	for _, table := range tables {
		if out, err := table.Decipher(table.Input); err != nil {
			t.Error("Could not decipher:", err)
		} else if out != table.Output {
			t.Errorf("Expected %q to decipher to %q, but instead got %q", table.Input, table.Output, out)
		}
	}
}

func ExampleCipher_tableau() {
	c := Cipher{}
	out, err := c.tableau()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(out)

	// Output:
	//     Z Y X W V U T S R Q P O N M L K J I H G F E D C B A
	//   +----------------------------------------------------
	// A | Z Y X W V U T S R Q P O N M L K J I H G F E D C B A
	// B | Y X W V U T S R Q P O N M L K J I H G F E D C B A Z
	// C | X W V U T S R Q P O N M L K J I H G F E D C B A Z Y
	// D | W V U T S R Q P O N M L K J I H G F E D C B A Z Y X
	// E | V U T S R Q P O N M L K J I H G F E D C B A Z Y X W
	// F | U T S R Q P O N M L K J I H G F E D C B A Z Y X W V
	// G | T S R Q P O N M L K J I H G F E D C B A Z Y X W V U
	// H | S R Q P O N M L K J I H G F E D C B A Z Y X W V U T
	// I | R Q P O N M L K J I H G F E D C B A Z Y X W V U T S
	// J | Q P O N M L K J I H G F E D C B A Z Y X W V U T S R
	// K | P O N M L K J I H G F E D C B A Z Y X W V U T S R Q
	// L | O N M L K J I H G F E D C B A Z Y X W V U T S R Q P
	// M | N M L K J I H G F E D C B A Z Y X W V U T S R Q P O
	// N | M L K J I H G F E D C B A Z Y X W V U T S R Q P O N
	// O | L K J I H G F E D C B A Z Y X W V U T S R Q P O N M
	// P | K J I H G F E D C B A Z Y X W V U T S R Q P O N M L
	// Q | J I H G F E D C B A Z Y X W V U T S R Q P O N M L K
	// R | I H G F E D C B A Z Y X W V U T S R Q P O N M L K J
	// S | H G F E D C B A Z Y X W V U T S R Q P O N M L K J I
	// T | G F E D C B A Z Y X W V U T S R Q P O N M L K J I H
	// U | F E D C B A Z Y X W V U T S R Q P O N M L K J I H G
	// V | E D C B A Z Y X W V U T S R Q P O N M L K J I H G F
	// W | D C B A Z Y X W V U T S R Q P O N M L K J I H G F E
	// X | C B A Z Y X W V U T S R Q P O N M L K J I H G F E D
	// Y | B A Z Y X W V U T S R Q P O N M L K J I H G F E D C
	// Z | A Z Y X W V U T S R Q P O N M L K J I H G F E D C B
}
