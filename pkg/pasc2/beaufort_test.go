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

package pasc2

import (
	"testing"

	"github.com/merenbach/goldbug/internal/fixture"
)

func TestBeaufortCipher_Encipher(t *testing.T) {
	var tables []struct {
		Alphabet string
		Caseless bool
		Strict   bool
		Key      string

		Input  string
		Output string
	}

	fixture.Load(t, &tables)
	for i, table := range tables {
		t.Logf("Running test %d of %d...", i+1, len(tables))

		var params []ConfigOption
		if table.Alphabet != "" {
			params = append(params, WithAlphabet(table.Alphabet))
		}
		if table.Strict {
			params = append(params, WithStrict())
		}
		if table.Caseless {
			params = append(params, WithCaseless())
		}

		c, err := NewBeaufortCipher(table.Key, params...)
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

func TestBeaufortCipher_Decipher(t *testing.T) {
	var tables []struct {
		Alphabet string
		Caseless bool
		Strict   bool
		Key      string

		Input  string
		Output string
	}

	fixture.Load(t, &tables)
	for i, table := range tables {
		t.Logf("Running test %d of %d...", i+1, len(tables))

		var params []ConfigOption
		if table.Alphabet != "" {
			params = append(params, WithAlphabet(table.Alphabet))
		}
		if table.Strict {
			params = append(params, WithStrict())
		}
		if table.Caseless {
			params = append(params, WithCaseless())
		}

		c, err := NewBeaufortCipher(table.Key, params...)
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

func TestBeaufortCipher_Printable(t *testing.T) {
	c, err := NewBeaufortCipher("")
	if err != nil {
		t.Fatal("Error:", err)
	}

	tableau, err := c.Printable()
	if err != nil {
		t.Fatal("Error:", err)
	}
	fixture.Golden(t, []byte(tableau))
}
