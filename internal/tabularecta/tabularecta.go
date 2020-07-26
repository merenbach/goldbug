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

package tabularecta

import (
	"errors"
	"fmt"
	"strings"

	"github.com/merenbach/gold-bug/internal/stringutil"
)

// TabulaRecta holds a tabula recta.
type TabulaRecta struct {
	keyAlphabet string
	ptAlphabet  string
	ctAlphabets []string

	pt2ct map[rune]map[rune]rune // map[key]=>map[pt]=>ct
	ct2pt map[rune]map[rune]rune // map[key]=>map[ct]=>pt
}

// New tabula recta.
func New(columnHeaders string, rowHeaders string, rows []string) (*TabulaRecta, error) {
	pt2ct := make(map[rune]map[rune]rune)
	ct2pt := make(map[rune]map[rune]rune)

	keyRunes := []rune(rowHeaders)
	if len(keyRunes) != len(rowHeaders) {
		return nil, errors.New("Row headers must have same rune length as rows slice")
	}

	for i, r := range keyRunes {
		ptCipher, err := stringutil.MakeTranslationTable(columnHeaders, rows[i], "")
		if err != nil {
			return nil, err
		}

		ctCipher, err := stringutil.MakeTranslationTable(rows[i], columnHeaders, "")
		if err != nil {
			return nil, err
		}

		pt2ct[r], ct2pt[r] = ptCipher, ctCipher
	}

	return &TabulaRecta{
		ptAlphabet:  columnHeaders,
		keyAlphabet: rowHeaders,
		ctAlphabets: rows,
		pt2ct:       pt2ct,
		ct2pt:       ct2pt,
	}, nil
}

func (tr *TabulaRecta) String() string {
	return fmt.Sprintf("%+v", map[string]interface{}{
		"k":  tr.keyAlphabet,
		"pt": tr.ptAlphabet,
		"ct": tr.ctAlphabets,
	})
}

// Printable representation of this tabula recta.
func (tr *TabulaRecta) Printable() string {
	var out strings.Builder
	var err error
	formatForPrinting := func(s string) string {
		spl := strings.Split(s, "")
		return strings.Join(spl, " ")
	}
	if _, err = out.WriteString("    " + formatForPrinting(tr.ptAlphabet) + "\n  +"); err != nil {
		return ""
	}
	for range tr.ptAlphabet {
		if _, err = out.WriteRune('-'); err != nil {
			return ""
		}
		if _, err = out.WriteRune('-'); err != nil {
			return ""
		}
	}
	for i, r := range tr.keyAlphabet {
		ctAlpha := fmt.Sprintf("\n%c | %s", r, formatForPrinting(tr.ctAlphabets[i]))
		if _, err = out.WriteString(ctAlpha); err != nil {
			return ""
		}
	}
	return out.String()
}

// Encipher a plaintext rune with a given key alphabet rune.
// Encipher will return (-1, false) if the key rune is invalid.
// Encipher will return (-1, true) if the key rune is valid but the message rune is not.
// Encipher will otherwise return the transcoded rune as the first argument and true as the second.
func (tr *TabulaRecta) Encipher(r rune, k rune) (rune, bool) {
	c, ok := tr.pt2ct[k]
	if !ok {
		return (-1), false
	}

	if o, ok := c[r]; ok {
		return o, true
	}
	return (-1), true
}

// Decipher a ciphertext rune with a given key alphabet rune.
// Decipher will return (-1, false) if the key rune is invalid.
// Decipher will return (-1, true) if the key rune is valid but the message rune is not.
// Decipher will otherwise return the transcoded rune as the first argument and true as the second.
func (tr *TabulaRecta) Decipher(r rune, k rune) (rune, bool) {
	c, ok := tr.ct2pt[k]
	if !ok {
		return (-1), false
	}

	if o, ok := c[r]; ok {
		return o, true
	}
	return (-1), true
}
