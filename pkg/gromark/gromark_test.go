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

package gromark

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
	c := Cipher{
		Key:    "ENIGMA",
		Primer: "23452",
	}
	out, err := c.tableau()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(out)

	// Output:
	//     A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
	//   +----------------------------------------------------
	// 0 | A J R X E B K S Y G F P V I D O U M H Q W N C L T Z
	// 1 | J R X E B K S Y G F P V I D O U M H Q W N C L T Z A
	// 2 | R X E B K S Y G F P V I D O U M H Q W N C L T Z A J
	// 3 | X E B K S Y G F P V I D O U M H Q W N C L T Z A J R
	// 4 | E B K S Y G F P V I D O U M H Q W N C L T Z A J R X
	// 5 | B K S Y G F P V I D O U M H Q W N C L T Z A J R X E
	// 6 | K S Y G F P V I D O U M H Q W N C L T Z A J R X E B
	// 7 | S Y G F P V I D O U M H Q W N C L T Z A J R X E B K
	// 8 | Y G F P V I D O U M H Q W N C L T Z A J R X E B K S
	// 9 | G F P V I D O U M H Q W N C L T Z A J R X E B K S Y
}
