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
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"
)

func TestMakeMap(t *testing.T) {
	testdata, err := ioutil.ReadFile(filepath.Join("testdata", "makemap.json"))
	if err != nil {
		t.Fatal("Could not read testdata fixture:", err)
	}

	var tables []struct {
		Src string
		Dst string
		Del string
		Map map[rune]rune
	}
	if err := json.Unmarshal(testdata, &tables); err != nil {
		t.Fatal("Could not unmarshal testdata:", err)
	}

	for _, table := range tables {
		m, err := makeMap(table.Src, table.Dst, table.Del)
		if err != nil {
			t.Error("Error:", err)
		}
		if !reflect.DeepEqual(m, table.Map) {
			t.Errorf("Got %+v instead of %+v", m, table.Map)
		}
	}
}

func TestTranslate(t *testing.T) {
	testdata, err := ioutil.ReadFile(filepath.Join("testdata", "translate.json"))
	if err != nil {
		t.Fatal("Could not read testdata fixture:", err)
	}

	var tables []struct {
		Input  string
		Output string
		Strict bool
		Map    map[rune]rune
	}
	if err := json.Unmarshal(testdata, &tables); err != nil {
		t.Fatal("Could not unmarshal testdata:", err)
	}

	for _, table := range tables {
		o := translate(table.Input, table.Map, table.Strict)
		if o != table.Output {
			t.Errorf("Expected %q, but got %q instead", table.Output, o)
		}
	}
}
