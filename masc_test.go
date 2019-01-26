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

func runMonoalphabeticReciprocalTests(t *testing.T, plaintext, ciphertext string, c SimpleSubstitutionCipher, strict bool) {
	encrypted := c.Encipher(plaintext, strict)
	decrypted := c.Decipher(ciphertext, strict)
	if string(encrypted) != ciphertext {
		t.Errorf("ciphertext %q was incorrect; wanted %q", encrypted, ciphertext)
	}
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
		c := SimpleSubstitutionCipher{table.pt, table.ct}
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
	tables := []struct {
		alphabet   string
		plaintext  string
		ciphertext string
		keyword    string
		strict     bool
	}{
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "CRHHL, WLQHG!", "KANGROOO", false},
		{defaultMonoalphabeticAlphabet, "LJOOF, WFEOI!", "HELLO, WORLD!", "KANGAROO", false},
		// {defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "CRHHLWLQHG", "KANGROOO", true},
		// {defaultMonoalphabeticAlphabet, "LJOOF, WFEOI!", "HELLOWORLD", "KANGAROO", true},
	}
	for _, table := range tables {
		c := MakeKeywordCipher(table.alphabet, table.keyword)
		runMonoalphabeticReciprocalTests(t, table.plaintext, table.ciphertext, c, table.strict)
	}
}

func TestAffineCipher(t *testing.T) {
	tables := []struct {
		alphabet   string
		plaintext  string
		ciphertext string
		a          int
		b          int
		strict     bool
	}{
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "AFCCX, BXSCY!", 7, 3, false},
		{defaultMonoalphabeticAlphabet, "IPQQJ, ZJCQA!", "HELLO, WORLD!", 7, 3, false},
		// {defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "AFCCXBXSCY", 7, 3, true},
		// {defaultMonoalphabeticAlphabet, "IPQQJ, ZJCQA!", "HELLOWORLD", 7, 3, true},
	}
	for _, table := range tables {
		c := MakeAffineCipher(table.alphabet, table.a, table.b)
		runMonoalphabeticReciprocalTests(t, table.plaintext, table.ciphertext, c, table.strict)

	}
}

func TestAtbashCipher(t *testing.T) {
	tables := []struct {
		alphabet   string
		plaintext  string
		ciphertext string
		strict     bool
	}{
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "SVOOL, DLIOW!", false},
		{defaultMonoalphabeticAlphabet, "SVOOL, DLIOW!", "HELLO, WORLD!", false},
		// {defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "SVOOLDLIOW", true},
		// {defaultMonoalphabeticAlphabet, "SVOOL, DLIOW!", "HELLOWORLD", true},
	}
	for _, table := range tables {
		c := MakeAtbashCipher(table.alphabet)
		runMonoalphabeticReciprocalTests(t, table.plaintext, table.ciphertext, c, table.strict)
	}
}

func TestCaesarCipher(t *testing.T) {
	tables := []struct {
		alphabet   string
		plaintext  string
		ciphertext string
		b          int
		strict     bool
	}{
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "KHOOR, ZRUOG!", 3, false},
		{defaultMonoalphabeticAlphabet, "EBIIL, TLOIA!", "HELLO, WORLD!", 3, false},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "YVCCF, NFICU!", 17, false},
		{defaultMonoalphabeticAlphabet, "QNUUX, FXAUM!", "HELLO, WORLD!", 17, false},
		// {defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "KHOORZRUOG", 3, true},
		// {defaultMonoalphabeticAlphabet, "EBIIL, TLOIA!", "HELLOWORLD", 3, true},
		// {defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "YVCCFNFICU", 17, true},
		// {defaultMonoalphabeticAlphabet, "QNUUX, FXAUM!", "HELLOWORLD", 17, true},
	}
	for _, table := range tables {
		c := MakeCaesarCipher(table.alphabet, table.b)
		runMonoalphabeticReciprocalTests(t, table.plaintext, table.ciphertext, c, table.strict)
	}
}

func TestDecimationCipher(t *testing.T) {
	tables := []struct {
		alphabet   string
		plaintext  string
		ciphertext string
		a          int
		strict     bool
	}{
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "XCZZU, YUPZV!", 7, false},
		{defaultMonoalphabeticAlphabet, "BIJJC, SCVJT!", "HELLO, WORLD!", 7, false},
		// {defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "XCZZUYUPZV", 7, true},
		// {defaultMonoalphabeticAlphabet, "BIJJC, SCVJT!", "HELLOWORLD", 7, true},
	}
	for _, table := range tables {
		c := MakeDecimationCipher(table.alphabet, table.a)
		runMonoalphabeticReciprocalTests(t, table.plaintext, table.ciphertext, c, table.strict)
	}
}

func TestRot13Cipher(t *testing.T) {
	tables := []struct {
		alphabet   string
		plaintext  string
		ciphertext string
		strict     bool
	}{
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "URYYB, JBEYQ!", false},
		{defaultMonoalphabeticAlphabet, "URYYB, JBEYQ!", "HELLO, WORLD!", false},
		// {defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "URYYBJBEYQ", true},
		// {defaultMonoalphabeticAlphabet, "URYYB, JBEYQ!", "HELLOWORLD", true},
	}
	for _, table := range tables {
		c := MakeRot13Cipher(table.alphabet)
		runMonoalphabeticReciprocalTests(t, table.plaintext, table.ciphertext, c, table.strict)
	}
}
