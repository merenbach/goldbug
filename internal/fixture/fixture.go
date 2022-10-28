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

package fixture

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// Portions of this file adapted liberally from:
//   <https://bbengfort.github.io/programmer/2018/09/22/go-testing-notes.html>
//   <https://ieftimov.com/posts/testing-in-go-golden-files/>

var update = flag.Bool("update", false, "update .golden files")

// Load a JSON fixture.
func Load(t *testing.T, v interface{}) {
	path := filepath.Join("testdata", t.Name()+".json")
	data, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("Unable to load fixture %s: %v", path, err)
		return
	}

	if err := json.Unmarshal(data, v); err != nil {
		t.Fatal("Could not unmarshal testdata:", err)
	}
}

// Golden fixture loading.
func Golden(t *testing.T, actual []byte) {
	golden := filepath.Join("testdata", t.Name()+".golden")
	if *update {
		err := ioutil.WriteFile(golden, actual, 0644)
		if err != nil {
			t.Fatalf("Could not write .golden file %q: %v", golden, err)
		}
	}

	expected, err := ioutil.ReadFile(golden)
	if err != nil {
		t.Fatal("Could not read testdata fixture:", err)
	}

	if !bytes.Equal(actual, expected) {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
