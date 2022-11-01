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

package pasc

// // A TabulaRecta holds a tabula recta.
// type tabulaRecta struct {
// 	// this is used here only for key lookups, hence no corresponding "strict" option at this time
// 	ptAlphabet string
// 	// CtAlphabet  string
// 	keyAlphabet string

// 	key string

// 	autokeyer autokeyFunc
// 	tableau   map[rune]*masc.SimpleCipher
// }

// // // Encipher a plaintext rune with a given key alphabet rune.
// // // Encipher will return (-1, false) if the key rune is invalid.
// // // Encipher will return (-1, true) if the key rune is valid but the message rune is not.
// // // Encipher will otherwise return the transcoded rune as the first argument and true as the second.
// // func (tr *tabulaRecta) encipher(r rune, k rune) (rune, bool) {
// // 	c, ok := tr.pt2ct[k]
// // 	if !ok {
// // 		return (-1), false
// // 	}

// // 	if o, ok := c[r]; ok {
// // 		return o, true
// // 	}
// // 	return (-1), true
// // }

// // // Decipher a ciphertext rune with a given key alphabet rune.
// // // Decipher will return (-1, false) if the key rune is invalid.
// // // Decipher will return (-1, true) if the key rune is valid but the message rune is not.
// // // Decipher will otherwise return the transcoded rune as the first argument and true as the second.
// // func (tr *tabulaRecta) decipher(r rune, k rune) (rune, bool) {
// // 	c, ok := tr.ct2pt[k]
// // 	if !ok {
// // 		return (-1), false
// // 	}

// // 	if o, ok := c[r]; ok {
// // 		return o, true
// // 	}
// // 	return (-1), true
// // }
