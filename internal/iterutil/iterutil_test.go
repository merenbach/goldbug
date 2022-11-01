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

package iterutil

import (
	"reflect"
	"testing"
)

func TestSuccessors(t *testing.T) {
	tables := []struct {
		expect []int
		first  int
		succ   func(int) int
	}{
		{[]int{0}, 0, func(i int) int { return i + 1 }},
		{[]int{1, 1, 1}, 1, func(i int) int { return i }},
		{[]int{1, 2, 3, 4, 5}, 1, func(i int) int { return i + 1 }},
		{[]int{1, 2, 4, 8, 16}, 1, func(i int) int { return 2 * i }},
	}

	for i, table := range tables {
		t.Logf("Running test %d of %d...", i+1, len(tables))

		g := Successors(table.first, table.succ)
		out := Take(len(table.expect), g)

		if !reflect.DeepEqual(out, table.expect) {
			t.Errorf("Expected %+v, got %+v", table.expect, out)
		}

	}
}
