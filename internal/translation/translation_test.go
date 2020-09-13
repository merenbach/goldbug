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

package translation

import (
	"strings"
	"testing"

	"github.com/merenbach/goldbug/internal/fixture"
)

func TestTable_Get(t *testing.T) {
	var tables []struct {
		Table

		Input  string
		Output string
		Strict bool
	}

	fixture.Load(t, &tables)
	for _, table := range tables {
		out := strings.Map(func(r rune) rune {
			o, _ := table.Get(r, table.Strict, false)
			return o
		}, table.Input)
		if out != table.Output {
			t.Errorf("Expected output %q for input %q, but instead got %q", table.Output, table.Input, out)
		}
	}
}
