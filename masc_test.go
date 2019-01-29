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

package main

import (
	"testing"
)

const defaultMonoalphabeticAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func runMonoalphabeticEncipherTest(t *testing.T, plaintext, ciphertext string, c *SimpleSubstitutionCipher, strict bool) {
	encrypted := c.Encipher(plaintext, strict)
	if string(encrypted) != ciphertext {
		t.Errorf("ciphertext %q was incorrect; wanted %q", encrypted, ciphertext)
	}
}

func runMonoalphabeticDecipherTest(t *testing.T, ciphertext string, plaintext string, c *SimpleSubstitutionCipher, strict bool) {
	decrypted := c.Decipher(ciphertext, strict)
	if string(decrypted) != plaintext {
		t.Errorf("plaintext %q was incorrect; wanted: %q", decrypted, plaintext)
	}
}

func TestSimpleSubstitutionCipher(t *testing.T) {

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
		c := NewSimpleSubstitutionCipher(table.pt, table.ct)
		if output := c.String(); output != table.expected {
			t.Errorf("Tableau printout doesn't match for PT %q and CT %q. Received: %s; expected: %s", table.pt, table.ct, output, table.expected)
		}
		if output := c.Encipher(table.pt, false); output != table.ct {
			t.Errorf("Tableau Pt2Ct doesn't match for PT %q and CT %q. Received: %s", table.pt, table.ct, output)
		}
		if output := c.Decipher(table.ct, false); output != table.pt {
			t.Errorf("Tableau Ct2Pt doesn't match for PT %q and CT %q. Received: %s", table.pt, table.ct, output)
		}
	}
}

func TestKeywordCipher(t *testing.T) {
	encipherTables := []struct {
		alphabet   string
		plaintext  string
		ciphertext string
		keyword    string
		strict     bool
	}{
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "CRHHL, WLQHG!", "KANGAROO", false},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "CRHHLWLQHG", "KANGAROO", true},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "GDKKN, WNRKC!", "Q", false},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "GDKKNWNRKC", "Q", true},
	}
	decipherTables := []struct {
		alphabet   string
		plaintext  string
		ciphertext string
		keyword    string
		strict     bool
	}{
		{defaultMonoalphabeticAlphabet, "CRHHL, WLQHG!", "HELLO, WORLD!", "KANGAROO", false},
		{defaultMonoalphabeticAlphabet, "CRHHL, WLQHG!", "HELLOWORLD", "KANGAROO", true},
		{defaultMonoalphabeticAlphabet, "GDKKN, WNRKC!", "HELLO, WORLD!", "Q", false},
		{defaultMonoalphabeticAlphabet, "GDKKN, WNRKC!", "HELLOWORLD", "Q", true},
	}

	for _, table := range encipherTables {
		c := NewKeywordCipher(table.alphabet, table.keyword)
		runMonoalphabeticEncipherTest(t, table.plaintext, table.ciphertext, c, table.strict)
	}
	for _, table := range decipherTables {
		c := NewKeywordCipher(table.alphabet, table.keyword)
		runMonoalphabeticDecipherTest(t, table.plaintext, table.ciphertext, c, table.strict)
	}
}

func TestAffineCipher(t *testing.T) {
	encipherTables := []struct {
		alphabet   string
		plaintext  string
		ciphertext string
		a          int
		b          int
		strict     bool
	}{
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "AFCCX, BXSCY!", 7, 3, false},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "AFCCXBXSCY", 7, 3, true},
		{defaultMonoalphabeticAlphabet, "AFFINE CIPHER", "IHHWVC SWFRCP", 5, 8, false},
		{defaultMonoalphabeticAlphabet, "AFFINE CIPHER", "IHHWVCSWFRCP", 5, 8, true},
	}
	decipherTables := []struct {
		alphabet   string
		plaintext  string
		ciphertext string
		a          int
		b          int
		strict     bool
	}{
		{defaultMonoalphabeticAlphabet, "AFCCX, BXSCY!", "HELLO, WORLD!", 7, 3, false},
		{defaultMonoalphabeticAlphabet, "AFCCX, BXSCY!", "HELLOWORLD", 7, 3, true},
		{defaultMonoalphabeticAlphabet, "IHHWVC SWFRCP", "AFFINE CIPHER", 5, 8, false},
		{defaultMonoalphabeticAlphabet, "IHHWVC SWFRCP", "AFFINECIPHER", 5, 8, true},
	}

	for _, table := range encipherTables {
		c := NewAffineCipher(table.alphabet, table.a, table.b)
		runMonoalphabeticEncipherTest(t, table.plaintext, table.ciphertext, c, table.strict)
	}
	for _, table := range decipherTables {
		c := NewAffineCipher(table.alphabet, table.a, table.b)
		runMonoalphabeticDecipherTest(t, table.plaintext, table.ciphertext, c, table.strict)
	}
}

func TestAtbashCipher(t *testing.T) {
	encipherTables := []struct {
		alphabet   string
		plaintext  string
		ciphertext string
		strict     bool
	}{
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "SVOOL, DLIOW!", false},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "SVOOLDLIOW", true},
		{defaultMonoalphabeticAlphabet, "ATBASH CIPHER", "ZGYZHS XRKSVI", false},
		{defaultMonoalphabeticAlphabet, "ATBASH CIPHER", "ZGYZHSXRKSVI", true},
	}
	decipherTables := []struct {
		alphabet   string
		plaintext  string
		ciphertext string
		strict     bool
	}{
		{defaultMonoalphabeticAlphabet, "SVOOL, DLIOW!", "HELLO, WORLD!", false},
		{defaultMonoalphabeticAlphabet, "SVOOL, DLIOW!", "HELLOWORLD", true},
		{defaultMonoalphabeticAlphabet, "ZGYZHS XRKSVI", "ATBASH CIPHER", false},
		{defaultMonoalphabeticAlphabet, "ZGYZHS XRKSVI", "ATBASHCIPHER", true},
	}

	for _, table := range encipherTables {
		c := NewAtbashCipher(table.alphabet)
		runMonoalphabeticEncipherTest(t, table.plaintext, table.ciphertext, c, table.strict)
	}
	for _, table := range decipherTables {
		c := NewAtbashCipher(table.alphabet)
		runMonoalphabeticDecipherTest(t, table.plaintext, table.ciphertext, c, table.strict)
	}
}

func TestCaesarCipher(t *testing.T) {
	encipherTables := []struct {
		alphabet   string
		plaintext  string
		ciphertext string
		b          int
		strict     bool
	}{
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "HELLO, WORLD!", 26, false},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "HELLOWORLD", 26, true},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "KHOOR, ZRUOG!", 3, false},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "KHOORZRUOG", 3, true},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "YVCCF, NFICU!", 17, false},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "YVCCFNFICU", 17, true},
	}
	decipherTables := []struct {
		alphabet   string
		plaintext  string
		ciphertext string
		b          int
		strict     bool
	}{
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "HELLO, WORLD!", 26, false},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "HELLOWORLD", 26, true},
		{defaultMonoalphabeticAlphabet, "KHOOR, ZRUOG!", "HELLO, WORLD!", 3, false},
		{defaultMonoalphabeticAlphabet, "KHOOR, ZRUOG!", "HELLOWORLD", 3, true},
		{defaultMonoalphabeticAlphabet, "YVCCF, NFICU!", "HELLO, WORLD!", 17, false},
		{defaultMonoalphabeticAlphabet, "YVCCF, NFICU!", "HELLOWORLD", 17, true},
	}

	for _, table := range encipherTables {
		c := NewCaesarCipher(table.alphabet, table.b)
		runMonoalphabeticEncipherTest(t, table.plaintext, table.ciphertext, c, table.strict)
	}
	for _, table := range decipherTables {
		c := NewCaesarCipher(table.alphabet, table.b)
		runMonoalphabeticDecipherTest(t, table.plaintext, table.ciphertext, c, table.strict)
	}
}

func TestDecimationCipher(t *testing.T) {
	encipherTables := []struct {
		alphabet   string
		plaintext  string
		ciphertext string
		a          int
		strict     bool
	}{
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "HELLO, WORLD!", 1, false},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "HELLOWORLD", 1, true},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "XCZZU, YUPZV!", 7, false},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "XCZZUYUPZV", 7, true},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "TWPPM, EMJPX!", 25, false},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "TWPPMEMJPX", 25, true},
	}
	decipherTables := []struct {
		alphabet   string
		plaintext  string
		ciphertext string
		a          int
		strict     bool
	}{
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "HELLO, WORLD!", 1, false},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "HELLOWORLD", 1, true},
		{defaultMonoalphabeticAlphabet, "XCZZU, YUPZV!", "HELLO, WORLD!", 7, false},
		{defaultMonoalphabeticAlphabet, "XCZZU, YUPZV!", "HELLOWORLD", 7, true},
		{defaultMonoalphabeticAlphabet, "TWPPM, EMJPX!", "HELLO, WORLD!", 25, false},
		{defaultMonoalphabeticAlphabet, "TWPPM, EMJPX!", "HELLOWORLD", 25, true},
	}

	for _, table := range encipherTables {
		c := NewDecimationCipher(table.alphabet, table.a)
		runMonoalphabeticEncipherTest(t, table.plaintext, table.ciphertext, c, table.strict)
	}
	for _, table := range decipherTables {
		c := NewDecimationCipher(table.alphabet, table.a)
		runMonoalphabeticDecipherTest(t, table.plaintext, table.ciphertext, c, table.strict)
	}
}

func TestRot13Cipher(t *testing.T) {
	encipherTables := []struct {
		alphabet   string
		plaintext  string
		ciphertext string
		strict     bool
	}{
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "URYYB, JBEYQ!", false},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "URYYBJBEYQ", true},
		{defaultMonoalphabeticAlphabet, "URYYB, JBEYQ!", "HELLO, WORLD!", false},
		{defaultMonoalphabeticAlphabet, "URYYB, JBEYQ!", "HELLOWORLD", true},
	}
	decipherTables := []struct {
		alphabet   string
		plaintext  string
		ciphertext string
		strict     bool
	}{
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "URYYB, JBEYQ!", false},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "URYYBJBEYQ", true},
		{defaultMonoalphabeticAlphabet, "URYYB, JBEYQ!", "HELLO, WORLD!", false},
		{defaultMonoalphabeticAlphabet, "URYYB, JBEYQ!", "HELLOWORLD", true},
	}

	for _, table := range encipherTables {
		c := NewRot13Cipher(table.alphabet)
		runMonoalphabeticEncipherTest(t, table.plaintext, table.ciphertext, c, table.strict)
	}
	for _, table := range decipherTables {
		c := NewRot13Cipher(table.alphabet)
		runMonoalphabeticDecipherTest(t, table.plaintext, table.ciphertext, c, table.strict)
	}
}
