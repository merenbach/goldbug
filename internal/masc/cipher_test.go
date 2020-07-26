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

import (
	"fmt"
	"testing"
)

const defaultMonoalphabeticAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func runMonoalphabeticEncipherTest(t *testing.T, input string, expected string, c *MonoalphabeticSubstitutionCipher) {
	output, err := c.EncipherString(input)
	if err != nil {
		t.Error("Encryption failed:", err)
	}
	if string(output) != expected {
		t.Errorf("ciphertext %q for input %q was incorrect; wanted %q", output, input, expected)
	}
}

func runMonoalphabeticDecipherTest(t *testing.T, input string, expected string, c *MonoalphabeticSubstitutionCipher) {
	output, err := c.DecipherString(input)
	if err != nil {
		t.Error("Decryption failed:", err)
	}
	if string(output) != expected {
		t.Errorf("plaintext %q for input %q was incorrect; wanted: %q", output, input, expected)
	}
}

func TestKeywordCipher(t *testing.T) {
	encipherTables := []struct {
		alphabet string
		input    string
		expected string
		keyword  string
		strict   bool
	}{
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "CRHHL, WLQHG!", "KANGAROO", false},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "CRHHLWLQHG", "KANGAROO", true},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "GDKKN, WNRKC!", "Q", false},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "GDKKNWNRKC", "Q", true},
	}
	decipherTables := []struct {
		alphabet string
		input    string
		expected string
		keyword  string
		strict   bool
	}{
		{defaultMonoalphabeticAlphabet, "CRHHL, WLQHG!", "HELLO, WORLD!", "KANGAROO", false},
		{defaultMonoalphabeticAlphabet, "CRHHL, WLQHG!", "HELLOWORLD", "KANGAROO", true},
		{defaultMonoalphabeticAlphabet, "GDKKN, WNRKC!", "HELLO, WORLD!", "Q", false},
		{defaultMonoalphabeticAlphabet, "GDKKN, WNRKC!", "HELLOWORLD", "Q", true},
	}

	for _, table := range encipherTables {
		c, err := NewKeywordCipher(table.alphabet, table.keyword)
		if err != nil {
			t.Error(err)
		}
		c.Strict = table.strict
		runMonoalphabeticEncipherTest(t, table.input, table.expected, c)
	}
	for _, table := range decipherTables {
		c, err := NewKeywordCipher(table.alphabet, table.keyword)
		if err != nil {
			t.Error(err)
		}
		c.Strict = table.strict
		runMonoalphabeticDecipherTest(t, table.input, table.expected, c)
	}
}

func TestAffineCipher(t *testing.T) {
	encipherTables := []struct {
		alphabet string
		input    string
		expected string
		a        int
		b        int
		strict   bool
	}{
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "AFCCX, BXSCY!", 7, 3, false},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "AFCCXBXSCY", 7, 3, true},
		{defaultMonoalphabeticAlphabet, "AFFINE CIPHER", "IHHWVC SWFRCP", 5, 8, false},
		{defaultMonoalphabeticAlphabet, "AFFINE CIPHER", "IHHWVCSWFRCP", 5, 8, true},
	}
	decipherTables := []struct {
		alphabet string
		input    string
		expected string
		a        int
		b        int
		strict   bool
	}{
		{defaultMonoalphabeticAlphabet, "AFCCX, BXSCY!", "HELLO, WORLD!", 7, 3, false},
		{defaultMonoalphabeticAlphabet, "AFCCX, BXSCY!", "HELLOWORLD", 7, 3, true},
		{defaultMonoalphabeticAlphabet, "IHHWVC SWFRCP", "AFFINE CIPHER", 5, 8, false},
		{defaultMonoalphabeticAlphabet, "IHHWVC SWFRCP", "AFFINECIPHER", 5, 8, true},
	}

	for _, table := range encipherTables {
		c, err := NewAffineCipher(table.alphabet, table.a, table.b)
		if err != nil {
			t.Error(err)
		}
		c.Strict = table.strict
		runMonoalphabeticEncipherTest(t, table.input, table.expected, c)
	}
	for _, table := range decipherTables {
		c, err := NewAffineCipher(table.alphabet, table.a, table.b)
		if err != nil {
			t.Error(err)
		}
		c.Strict = table.strict
		runMonoalphabeticDecipherTest(t, table.input, table.expected, c)
	}
}

func TestAtbashCipher(t *testing.T) {
	encipherTables := []struct {
		alphabet string
		input    string
		expected string
		strict   bool
	}{
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "SVOOL, DLIOW!", false},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "SVOOLDLIOW", true},
		{defaultMonoalphabeticAlphabet, "ATBASH CIPHER", "ZGYZHS XRKSVI", false},
		{defaultMonoalphabeticAlphabet, "ATBASH CIPHER", "ZGYZHSXRKSVI", true},
	}
	decipherTables := []struct {
		alphabet string
		input    string
		expected string
		strict   bool
	}{
		{defaultMonoalphabeticAlphabet, "SVOOL, DLIOW!", "HELLO, WORLD!", false},
		{defaultMonoalphabeticAlphabet, "SVOOL, DLIOW!", "HELLOWORLD", true},
		{defaultMonoalphabeticAlphabet, "ZGYZHS XRKSVI", "ATBASH CIPHER", false},
		{defaultMonoalphabeticAlphabet, "ZGYZHS XRKSVI", "ATBASHCIPHER", true},
	}

	for _, table := range encipherTables {
		c, err := NewAtbashCipher(table.alphabet)
		if err != nil {
			t.Error(err)
		}
		c.Strict = table.strict
		runMonoalphabeticEncipherTest(t, table.input, table.expected, c)
	}
	for _, table := range decipherTables {
		c, err := NewAtbashCipher(table.alphabet)
		if err != nil {
			t.Error(err)
		}
		c.Strict = table.strict
		runMonoalphabeticDecipherTest(t, table.input, table.expected, c)
	}
}

func TestCaesarCipher(t *testing.T) {
	encipherTables := []struct {
		alphabet string
		input    string
		expected string
		b        int
		strict   bool
	}{
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "HELLO, WORLD!", 26, false},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "HELLOWORLD", 26, true},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "KHOOR, ZRUOG!", 3, false},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "KHOORZRUOG", 3, true},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "YVCCF, NFICU!", 17, false},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "YVCCFNFICU", 17, true},
	}
	decipherTables := []struct {
		alphabet string
		input    string
		expected string
		b        int
		strict   bool
	}{
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "HELLO, WORLD!", 26, false},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "HELLOWORLD", 26, true},
		{defaultMonoalphabeticAlphabet, "KHOOR, ZRUOG!", "HELLO, WORLD!", 3, false},
		{defaultMonoalphabeticAlphabet, "KHOOR, ZRUOG!", "HELLOWORLD", 3, true},
		{defaultMonoalphabeticAlphabet, "YVCCF, NFICU!", "HELLO, WORLD!", 17, false},
		{defaultMonoalphabeticAlphabet, "YVCCF, NFICU!", "HELLOWORLD", 17, true},
	}

	for _, table := range encipherTables {
		c, err := NewCaesarCipher(table.alphabet, table.b)
		if err != nil {
			t.Error(err)
		}
		c.Strict = table.strict
		runMonoalphabeticEncipherTest(t, table.input, table.expected, c)
	}
	for _, table := range decipherTables {
		c, err := NewCaesarCipher(table.alphabet, table.b)
		if err != nil {
			t.Error(err)
		}
		c.Strict = table.strict
		runMonoalphabeticDecipherTest(t, table.input, table.expected, c)
	}
}

func TestDecimationCipher(t *testing.T) {
	encipherTables := []struct {
		alphabet string
		input    string
		expected string
		a        int
		strict   bool
	}{
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "HELLO, WORLD!", 1, false},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "HELLOWORLD", 1, true},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "XCZZU, YUPZV!", 7, false},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "XCZZUYUPZV", 7, true},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "TWPPM, EMJPX!", 25, false},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "TWPPMEMJPX", 25, true},
	}
	decipherTables := []struct {
		alphabet string
		input    string
		expected string
		a        int
		strict   bool
	}{
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "HELLO, WORLD!", 1, false},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "HELLOWORLD", 1, true},
		{defaultMonoalphabeticAlphabet, "XCZZU, YUPZV!", "HELLO, WORLD!", 7, false},
		{defaultMonoalphabeticAlphabet, "XCZZU, YUPZV!", "HELLOWORLD", 7, true},
		{defaultMonoalphabeticAlphabet, "TWPPM, EMJPX!", "HELLO, WORLD!", 25, false},
		{defaultMonoalphabeticAlphabet, "TWPPM, EMJPX!", "HELLOWORLD", 25, true},
	}

	for _, table := range encipherTables {
		c, err := NewDecimationCipher(table.alphabet, table.a)
		if err != nil {
			t.Error(err)
		}
		c.Strict = table.strict
		runMonoalphabeticEncipherTest(t, table.input, table.expected, c)
	}
	for _, table := range decipherTables {
		c, err := NewDecimationCipher(table.alphabet, table.a)
		if err != nil {
			t.Error(err)
		}
		c.Strict = table.strict
		runMonoalphabeticDecipherTest(t, table.input, table.expected, c)
	}
}

func TestRot13Cipher(t *testing.T) {
	encipherTables := []struct {
		alphabet string
		input    string
		expected string
		strict   bool
	}{
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "URYYB, JBEYQ!", false},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "URYYBJBEYQ", true},
		{defaultMonoalphabeticAlphabet, "URYYB, JBEYQ!", "HELLO, WORLD!", false},
		{defaultMonoalphabeticAlphabet, "URYYB, JBEYQ!", "HELLOWORLD", true},
	}
	decipherTables := []struct {
		alphabet string
		input    string
		expected string
		strict   bool
	}{
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "URYYB, JBEYQ!", false},
		{defaultMonoalphabeticAlphabet, "HELLO, WORLD!", "URYYBJBEYQ", true},
		{defaultMonoalphabeticAlphabet, "URYYB, JBEYQ!", "HELLO, WORLD!", false},
		{defaultMonoalphabeticAlphabet, "URYYB, JBEYQ!", "HELLOWORLD", true},
	}

	for _, table := range encipherTables {
		c, err := NewRot13Cipher(table.alphabet)
		if err != nil {
			t.Error(err)
		}
		c.Strict = table.strict
		runMonoalphabeticEncipherTest(t, table.input, table.expected, c)
	}
	for _, table := range decipherTables {
		c, err := NewRot13Cipher(table.alphabet)
		if err != nil {
			t.Error(err)
		}
		c.Strict = table.strict
		runMonoalphabeticDecipherTest(t, table.input, table.expected, c)
	}
}

func ExampleNewAtbashCipher() {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	c, err := NewAtbashCipher(alphabet)
	if err != nil {
		panic(err)
	}
	fmt.Println(c.Printable())
	// Output:
	// PT: ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// CT: ZYXWVUTSRQPONMLKJIHGFEDCBA
}

func ExampleNewRot13Cipher() {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	c, err := NewRot13Cipher(alphabet)
	if err != nil {
		panic(err)
	}
	fmt.Println(c.Printable())
	// Output:
	// PT: ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// CT: NOPQRSTUVWXYZABCDEFGHIJKLM
}
