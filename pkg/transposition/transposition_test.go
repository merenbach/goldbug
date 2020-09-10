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
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"
)

func TestLexicalKey(t *testing.T) {
	testdata, err := ioutil.ReadFile(filepath.Join("testdata", "lexicalkey.json"))
	if err != nil {
		t.Fatal("Could not read testdata fixture:", err)
	}

	var tables []struct {
		Input      string
		Repetition bool
		Output     []int
	}
	if err := json.Unmarshal(testdata, &tables); err != nil {
		t.Fatal("Could not unmarshal testdata:", err)
	}

	for _, table := range tables {
		if out := lexicalKey(table.Input, table.Repetition); err != nil {
			t.Error("Error:", err)
		} else if !reflect.DeepEqual(out, table.Output) {
			t.Errorf("Expected %q to lexically yield %v, but instead got %v", table.Input, table.Output, out)
		}
	}
}

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