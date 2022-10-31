// Copyright 2022 Andrew Merenbach
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

package simple

import (
	"fmt"
	"testing"

	"github.com/merenbach/goldbug/internal/fixture"
	"github.com/merenbach/goldbug/pkg/masc2"
)

func TestCipher_Encipher(t *testing.T) {
	var tables []struct {
		Alphabet   string
		Caseless   bool
		Strict     bool
		CtAlphabet string

		Input  string
		Output string
	}

	fixture.Load(t, &tables)
	for i, table := range tables {
		t.Logf("Running test %d of %d...", i+1, len(tables))

		var params []masc2.ConfigOption
		if table.Strict {
			params = append(params, masc2.WithStrict())
		}
		if table.Caseless {
			params = append(params, masc2.WithCaseless())
		}

		c, err := NewCipher(table.CtAlphabet, params...)
		if err != nil {
			t.Error("Could not create cipher:", err)
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
		Alphabet   string
		Caseless   bool
		Strict     bool
		Keyword    string
		CtAlphabet string

		Input  string
		Output string
	}

	fixture.Load(t, &tables)
	for i, table := range tables {
		t.Logf("Running test %d of %d...", i+1, len(tables))

		var params []masc2.ConfigOption
		if table.Strict {
			params = append(params, masc2.WithStrict())
		}
		if table.Caseless {
			params = append(params, masc2.WithCaseless())
		}

		c, err := NewCipher(table.CtAlphabet, params...)
		if err != nil {
			t.Error("Could not create cipher:", err)
		}

		if out, err := c.Decipher(table.Input); err != nil {
			t.Error("Could not decipher:", err)
		} else if out != table.Output {
			t.Errorf("Expected %q to decipher to %q, but instead got %q", table.Input, table.Output, out)
		}
	}
}

func ExampleCipher_Tableau() {
	c, err := NewCipher("MLSDEFPTJCARNUVWYXOGQKIZHB")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(c)

	// Output:
	// PT: ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// CT: MLSDEFPTJCARNUVWYXOGQKIZHB
}