// Copyright 2019 Andrew Merenbach
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

package lfg

import (
	"reflect"
	"testing"

	"github.com/merenbach/goldbug/internal/fixture"
	"github.com/merenbach/goldbug/internal/iterutil"
)

func TestMultiplicative_Slice(t *testing.T) {
	var tables []struct {
		Modulus int
		Seed    []int
		Taps    []int

		Output []int
	}

	fixture.Load(t, &tables)
	for i, table := range tables {
		t.Logf("Running test %d of %d...", i+1, len(tables))

		if g, err := Multiplicative(table.Modulus, table.Seed, table.Taps); err != nil {
			t.Error("Error:", err)
		} else {
			out := iterutil.Take(len(table.Output), g)
			if !reflect.DeepEqual(out, table.Output) {
				t.Errorf("Expected output %v but got %v", out, table.Output)
			}
		}
	}
}
