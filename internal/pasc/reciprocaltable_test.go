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

package pasc

// func TestReciprocalTable(t *testing.T) {
// 	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// 	tables := []struct {
// 		keyAlphabet string
// 		ptAlphabet  string
// 		ctAlphabet  string
// 		keyRune     string
// 		srcRune     string
// 		dstRune     string
// 	}{
// 		{alphabet, alphabet, alphabet, "A", "A", "A"},
// 		{alphabet, alphabet, alphabet, "B", "B", "C"},
// 		{alphabet, alphabet, alphabet, "K", "V", "F"},
// 		{alphabet, alphabet, alphabet, "O", "K", "Y"},
// 		{alphabet, alphabet, alphabet, "Y", "Y", "W"},
// 		{alphabet, alphabet, alphabet, "Z", "O", "N"},
// 		{alphabet, alphabet, alphabet, "Z", "Z", "Y"},
// 	}

// 	for _, table := range tables {
// 		keyRunes := []rune(table.keyAlphabet)
// 		ctAlphabets := make([]string, len(keyRunes))

// 		for i := range keyRunes {
// 			ctAlphabets[i] = stringutil.WrapString(table.ctAlphabet, i)
// 		}

// 		tr := ReciprocalTable{
// 			PtAlphabet:  table.ptAlphabet,
// 			KeyAlphabet: table.keyAlphabet,
// 			CtAlphabets: ctAlphabets,
// 		}

// 		dst, err := tr.Encipher(table.srcRune, table.keyRune, nil)
// 		if err != nil {
// 			t.Errorf("Expected encipherment key rune %s to be in key alphabet, but it was not", table.keyRune)
// 		}
// 		if dst != table.dstRune {
// 			t.Errorf("Expected E(p=%s,k=%s)=%s but got %s instead", table.srcRune, table.keyRune, table.dstRune, dst)
// 		}

// 		src, err := tr.Decipher(dst, table.keyRune, nil)
// 		if err != nil {
// 			t.Errorf("Expected decipherment key rune %s to be in key alphabet, but it was not", table.keyRune)
// 		}
// 		if src != table.srcRune {
// 			t.Errorf("Expected D(c=%s,k=%s)=%s but got %s instead", table.dstRune, table.keyRune, table.srcRune, src)
// 		}
// 	}
// }
