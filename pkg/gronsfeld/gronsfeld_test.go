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

package gronsfeld

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
	//     A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
	//   +----------------------------------------------------
	// 0 | A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
	// 1 | B C D E F G H I J K L M N O P Q R S T U V W X Y Z A
	// 2 | C D E F G H I J K L M N O P Q R S T U V W X Y Z A B
	// 3 | D E F G H I J K L M N O P Q R S T U V W X Y Z A B C
	// 4 | E F G H I J K L M N O P Q R S T U V W X Y Z A B C D
	// 5 | F G H I J K L M N O P Q R S T U V W X Y Z A B C D E
	// 6 | G H I J K L M N O P Q R S T U V W X Y Z A B C D E F
	// 7 | H I J K L M N O P Q R S T U V W X Y Z A B C D E F G
	// 8 | I J K L M N O P Q R S T U V W X Y Z A B C D E F G H
	// 9 | J K L M N O P Q R S T U V W X Y Z A B C D E F G H I
}
