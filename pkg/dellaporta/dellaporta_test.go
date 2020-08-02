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

package dellaporta

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
	// A | N O P Q R S T U V W X Y Z A B C D E F G H I J K L M
	// B | N O P Q R S T U V W X Y Z A B C D E F G H I J K L M
	// C | O P Q R S T U V W X Y Z N M A B C D E F G H I J K L
	// D | O P Q R S T U V W X Y Z N M A B C D E F G H I J K L
	// E | P Q R S T U V W X Y Z N O L M A B C D E F G H I J K
	// F | P Q R S T U V W X Y Z N O L M A B C D E F G H I J K
	// G | Q R S T U V W X Y Z N O P K L M A B C D E F G H I J
	// H | Q R S T U V W X Y Z N O P K L M A B C D E F G H I J
	// I | R S T U V W X Y Z N O P Q J K L M A B C D E F G H I
	// J | R S T U V W X Y Z N O P Q J K L M A B C D E F G H I
	// K | S T U V W X Y Z N O P Q R I J K L M A B C D E F G H
	// L | S T U V W X Y Z N O P Q R I J K L M A B C D E F G H
	// M | T U V W X Y Z N O P Q R S H I J K L M A B C D E F G
	// N | T U V W X Y Z N O P Q R S H I J K L M A B C D E F G
	// O | U V W X Y Z N O P Q R S T G H I J K L M A B C D E F
	// P | U V W X Y Z N O P Q R S T G H I J K L M A B C D E F
	// Q | V W X Y Z N O P Q R S T U F G H I J K L M A B C D E
	// R | V W X Y Z N O P Q R S T U F G H I J K L M A B C D E
	// S | W X Y Z N O P Q R S T U V E F G H I J K L M A B C D
	// T | W X Y Z N O P Q R S T U V E F G H I J K L M A B C D
	// U | X Y Z N O P Q R S T U V W D E F G H I J K L M A B C
	// V | X Y Z N O P Q R S T U V W D E F G H I J K L M A B C
	// W | Y Z N O P Q R S T U V W X C D E F G H I J K L M A B
	// X | Y Z N O P Q R S T U V W X C D E F G H I J K L M A B
	// Y | Z N O P Q R S T U V W X Y B C D E F G H I J K L M A
	// Z | Z N O P Q R S T U V W X Y B C D E F G H I J K L M A
}
