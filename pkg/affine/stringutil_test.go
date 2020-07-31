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

package affine

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestBackpermute(t *testing.T) {
	testdata, err := ioutil.ReadFile(filepath.Join("testdata", "backpermute.json"))
	if err != nil {
		t.Fatal("Could not read testdata fixture:", err)
	}

	var tables []struct {
		Input   string
		Output  string
		Indices []int
		Success bool
	}
	if err := json.Unmarshal(testdata, &tables); err != nil {
		t.Fatal("Could not unmarshal testdata:", err)
	}

	for i, table := range tables {
		t.Logf("Testing table %d of %d", i+1, len(tables))
		if out, err := backpermute(table.Input, table.Indices); err != nil && table.Success {
			t.Error("Unexpected backpermute failure:", err)
		} else if err == nil && !table.Success {
			t.Error("Unexpected backpermute success")
		} else if string(out) != table.Output {
			t.Error("Received incorrect output:", out)
		}
	}
}
