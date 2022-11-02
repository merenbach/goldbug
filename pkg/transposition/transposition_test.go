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

// func TestLexicalKey(t *testing.T) {
// 	var tables []struct {
// 		Input      string
// 		Repetition bool
// 		Output     []int
// 	}

// 	fixture.Load(t, &tables)
// 	for _, table := range tables {
// 		out := lexicalKey(table.Input, table.Repetition)
// 		if !reflect.DeepEqual(out, table.Output) {
// 			t.Errorf("Expected %q to lexically yield %v, but instead got %v", table.Input, table.Output, out)
// 		}
// 	}
// }

func TestCipher_Encipher(t *testing.T) {
	var tables []struct {
		Myszkowski bool
		Keys       []string

		Input  string
		Output string
	}

	fixture.Load(t, &tables)
	for i, table := range tables {
		t.Logf("Running test %d of %d...", i+1, len(tables))

		var params []ConfigOption
		if table.Myszkowski {
			params = append(params, WithMyszkowski())
		}
		for _, key := range table.Keys {
			params = append(params, WithStringKey(key))
		}

		c, err := NewCipher(params...)
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

func TestCipher_Decipher(t *testing.T) {
	var tables []struct {
		Myszkowski bool
		Keys       []string

		Input  string
		Output string
	}

	fixture.Load(t, &tables)
	for i, table := range tables {
		t.Logf("Running test %d of %d...", i+1, len(tables))

		var params []ConfigOption
		if table.Myszkowski {
			params = append(params, WithMyszkowski())
		}
		for _, key := range table.Keys {
			params = append(params, WithStringKey(key))
		}

		c, err := NewCipher(params...)
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
