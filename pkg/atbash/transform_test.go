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

package atbash

import (
	"reflect"
	"testing"

	"github.com/merenbach/goldbug/internal/fixture"
)

func TestTransform(t *testing.T) {
	var tables []struct {
		Input  []int
		Output []int
	}

	fixture.Load(t, &tables)
	for i, table := range tables {
		t.Logf("Running test %d of %d...", i+1, len(tables))

		out, err := Transform(table.Input)
		if err != nil {
			t.Error("Could not complete transformation:", err)
		}

		if !reflect.DeepEqual(out, table.Output) {
			t.Errorf("Expected %+v to transform to %v, but instead got %v", table.Input, table.Output, out)
		}
	}
}
