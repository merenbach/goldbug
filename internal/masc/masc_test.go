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

package masc

import "testing"

func TestMonoalphabeticSubstitutionCipher(t *testing.T) {

	tables := []struct {
		pt       string
		ct       string
		expected string
	}{
		{"", "", "PT: \nCT: "},
		{"A", "A", "PT: A\nCT: A"},
		{"ABCDE", "DEFGH", "PT: ABCDE\nCT: DEFGH"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", "DEFGHIJKLMNOPQRSTUVWXYZABC", "PT: ABCDEFGHIJKLMNOPQRSTUVWXYZ\nCT: DEFGHIJKLMNOPQRSTUVWXYZABC"},
	}
	for _, table := range tables {
		c, err := New(table.pt, table.ct)
		if err != nil {
			t.Fatal(err)
		}

		// TODO: use example tests for this
		if output := c.Printable(); output != table.expected {
			t.Errorf("Tableau printout doesn't match for PT %q and CT %q. Received: %s; expected: %s", table.pt, table.ct, output, table.expected)
		}

		var output string

		output, err = c.EncipherString(table.pt)
		if err != nil {
			t.Error("Encryption failed:", err)
		}
		if output != table.ct {
			t.Errorf("Tableau Pt2Ct doesn't match for PT %q and CT %q. Received: %s", table.pt, table.ct, output)
		}

		output, err = c.DecipherString(table.ct)
		if err != nil {
			t.Error("Decryption failed:", err)
		}
		if output != table.pt {
			t.Errorf("Tableau Ct2Pt doesn't match for PT %q and CT %q. Received: %s", table.pt, table.ct, output)
		}
	}
}
