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

func TestDellaPortaCipher_Encipher(t *testing.T) {
	var tables []struct {
		Alphabet string
		Caseless bool
		Strict   bool
		Key      string

		Input  string
		Output string
	}

	fixture.Load(t, &tables)
	for _, table := range tables {
		var params []ConfigOption
		if table.Alphabet != "" {
			params = append(params, WithPtAlphabet(table.Alphabet))
		}
		if table.Strict {
			params = append(params, WithStrict())
		}
		if table.Caseless {
			params = append(params, WithCaseless())
		}

		c, err := NewDellaPortaCipher(table.Key, params...)
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

func TestDellaPortaCipher_Decipher(t *testing.T) {
	var tables []struct {
		Alphabet string
		Caseless bool
		Strict   bool
		Key      string

		Input  string
		Output string
	}

	fixture.Load(t, &tables)
	for _, table := range tables {
		var params []ConfigOption
		if table.Alphabet != "" {
			params = append(params, WithPtAlphabet(table.Alphabet))
		}
		if table.Strict {
			params = append(params, WithStrict())
		}
		if table.Caseless {
			params = append(params, WithCaseless())
		}

		c, err := NewDellaPortaCipher(table.Key, params...)
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

func TestDellaPortaCipher_Printable(t *testing.T) {
	c, err := NewDellaPortaCipher("")
	if err != nil {
		t.Fatal("Error:", err)
	}

	tableau, err := c.Printable()
	if err != nil {
		t.Fatal("Error:", err)
	}
	fixture.Golden(t, []byte(tableau))
}
