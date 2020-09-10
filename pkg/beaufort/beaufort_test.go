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

package beaufort

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

func ExampleCipher_Tableau() {
	c := Cipher{}
	out, err := c.Tableau()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(out)

	// Output:
	//     A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
	//
	// Z   Z Y X W V U T S R Q P O N M L K J I H G F E D C B A
	// Y   Y X W V U T S R Q P O N M L K J I H G F E D C B A Z
	// X   X W V U T S R Q P O N M L K J I H G F E D C B A Z Y
	// W   W V U T S R Q P O N M L K J I H G F E D C B A Z Y X
	// V   V U T S R Q P O N M L K J I H G F E D C B A Z Y X W
	// U   U T S R Q P O N M L K J I H G F E D C B A Z Y X W V
	// T   T S R Q P O N M L K J I H G F E D C B A Z Y X W V U
	// S   S R Q P O N M L K J I H G F E D C B A Z Y X W V U T
	// R   R Q P O N M L K J I H G F E D C B A Z Y X W V U T S
	// Q   Q P O N M L K J I H G F E D C B A Z Y X W V U T S R
	// P   P O N M L K J I H G F E D C B A Z Y X W V U T S R Q
	// O   O N M L K J I H G F E D C B A Z Y X W V U T S R Q P
	// N   N M L K J I H G F E D C B A Z Y X W V U T S R Q P O
	// M   M L K J I H G F E D C B A Z Y X W V U T S R Q P O N
	// L   L K J I H G F E D C B A Z Y X W V U T S R Q P O N M
	// K   K J I H G F E D C B A Z Y X W V U T S R Q P O N M L
	// J   J I H G F E D C B A Z Y X W V U T S R Q P O N M L K
	// I   I H G F E D C B A Z Y X W V U T S R Q P O N M L K J
	// H   H G F E D C B A Z Y X W V U T S R Q P O N M L K J I
	// G   G F E D C B A Z Y X W V U T S R Q P O N M L K J I H
	// F   F E D C B A Z Y X W V U T S R Q P O N M L K J I H G
	// E   E D C B A Z Y X W V U T S R Q P O N M L K J I H G F
	// D   D C B A Z Y X W V U T S R Q P O N M L K J I H G F E
	// C   C B A Z Y X W V U T S R Q P O N M L K J I H G F E D
	// B   B A Z Y X W V U T S R Q P O N M L K J I H G F E D C
	// A   A Z Y X W V U T S R Q P O N M L K J I H G F E D C B
}
