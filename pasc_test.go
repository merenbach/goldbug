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
	"fmt"
	"testing"
)

const defaultPolyalphabeticAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func runPolyalphabeticReciprocalTests(t *testing.T, plaintext, ciphertext string, c VigenereFamilyCipher, strict bool) {
	encrypted := c.Encipher(plaintext, strict)
	decrypted := c.Decipher(ciphertext, strict)
	if string(encrypted) != ciphertext {
		t.Errorf("ciphertext %q was incorrect; wanted %q", encrypted, ciphertext)
	}
	if string(decrypted) != plaintext {
		t.Errorf("plaintext %q was incorrect; wanted: %q", decrypted, plaintext)
	}
}

func TestVigenereCipher(t *testing.T) {
	tables := []struct {
		alphabet    string
		plaintext   string
		ciphertext  string
		countersign string
		strict      bool
	}{
		{"", "", "", "", false},
		{defaultPolyalphabeticAlphabet, "", "", "OCEANOGRAPHYWHAT", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "VGPLB, KUILS!", "OCEANOGRAPHYWHAT", false},
		// {defaultPolyalphabeticAlphabet, "LJOOF, WFEOI!", "HELLO, WORLD!", "KANGAROO", false},
		// {defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "CRHHLWLQHG", "KANGROOO", true},
		// {defaultPolyalphabeticAlphabet, "LJOOF, WFEOI!", "HELLOWORLD", "KANGAROO", true},
	}
	for _, table := range tables {
		c := MakeVigenereCipher(table.countersign, table.alphabet)
		runPolyalphabeticReciprocalTests(t, table.plaintext, table.ciphertext, c, table.strict)
	}
}

func TestVigenereTextAutoclaveCipher(t *testing.T) {
	tables := []struct {
		alphabet    string
		plaintext   string
		ciphertext  string
		countersign string
		strict      bool
	}{
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "VGPLB, DSCWR!", "OCEANOGRAPHYWHAT"[:5], false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "XLPWZ, KKFCO!", "Q", false},
		// {defaultPolyalphabeticAlphabet, "LJOOF, WFEOI!", "HELLO, WORLD!", "KANGAROO", false},
		// {defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "CRHHLWLQHG", "KANGROOO", true},
		// {defaultPolyalphabeticAlphabet, "LJOOF, WFEOI!", "HELLOWORLD", "KANGAROO", true},
	}
	for _, table := range tables {
		c := MakeVigenereTextAutoclaveCipher(table.countersign, table.alphabet)
		runPolyalphabeticReciprocalTests(t, table.plaintext, table.ciphertext, c, table.strict)
	}
}

func TestVigenereKeyAutoclaveCipher(t *testing.T) {
	tables := []struct {
		alphabet    string
		plaintext   string
		ciphertext  string
		countersign string
		strict      bool
	}{
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "VGPLB, RUGWE!", "OCEANOGRAPHYWHAT"[:5], false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "XBMXL, HVMXA!", "Q", false},
		// {defaultPolyalphabeticAlphabet, "LJOOF, WFEOI!", "HELLO, WORLD!", "KANGAROO", false},
		// {defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "CRHHLWLQHG", "KANGROOO", true},
		// {defaultPolyalphabeticAlphabet, "LJOOF, WFEOI!", "HELLOWORLD", "KANGAROO", true},
	}
	for _, table := range tables {
		c := MakeVigenereKeyAutoclaveCipher(table.countersign, table.alphabet)
		runPolyalphabeticReciprocalTests(t, table.plaintext, table.ciphertext, c, table.strict)
	}
}

func TestBeaufortCipher(t *testing.T) {
	tables := []struct {
		alphabet    string
		plaintext   string
		ciphertext  string
		countersign string
		strict      bool
	}{
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "HYTPZ, SSAPM!", "OCEANOGRAPHYWHAT", false},
		// {defaultPolyalphabeticAlphabet, "LJOOF, WFEOI!", "HELLO, WORLD!", "KANGAROO", false},
		// {defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "CRHHLWLQHG", "KANGROOO", true},
		// {defaultPolyalphabeticAlphabet, "LJOOF, WFEOI!", "HELLOWORLD", "KANGAROO", true},
	}
	for _, table := range tables {
		c := MakeBeaufortCipher(table.countersign, table.alphabet)
		runPolyalphabeticReciprocalTests(t, table.plaintext, table.ciphertext, c, table.strict)
	}
}

func TestGronsfeldCipher(t *testing.T) {
	tables := []struct {
		alphabet    string
		plaintext   string
		ciphertext  string
		countersign string
		strict      bool
	}{
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "JHMOQ, YRSOF!", "23132", false},
		// {defaultPolyalphabeticAlphabet, "LJOOF, WFEOI!", "HELLO, WORLD!", "KANGAROO", false},
		// {defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "CRHHLWLQHG", "KANGROOO", true},
		// {defaultPolyalphabeticAlphabet, "LJOOF, WFEOI!", "HELLOWORLD", "KANGAROO", true},
	}
	for _, table := range tables {
		c := MakeGronsfeldCipher(table.countersign, table.alphabet)
		runPolyalphabeticReciprocalTests(t, table.plaintext, table.ciphertext, c, table.strict)
	}
}

func TestTrithemiusCipher(t *testing.T) {
	tables := []struct {
		alphabet   string
		plaintext  string
		ciphertext string
		strict     bool
	}{
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "HFNOS, BUYTM!", false},
		// {defaultPolyalphabeticAlphabet, "LJOOF, WFEOI!", "HELLO, WORLD!", "KANGAROO", false},
		// {defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "CRHHLWLQHG", "KANGROOO", true},
		// {defaultPolyalphabeticAlphabet, "LJOOF, WFEOI!", "HELLOWORLD", "KANGAROO", true},
	}
	for _, table := range tables {
		c := MakeTrithemiusCipher(table.alphabet)
		runPolyalphabeticReciprocalTests(t, table.plaintext, table.ciphertext, c, table.strict)
	}
}

func TestVariantBeaufortCipher(t *testing.T) {
	tables := []struct {
		alphabet    string
		plaintext   string
		ciphertext  string
		countersign string
		strict      bool
	}{
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "TCHLB, IIALO!", "OCEANOGRAPHYWHAT", false},
		// {defaultPolyalphabeticAlphabet, "LJOOF, WFEOI!", "HELLO, WORLD!", "KANGAROO", false},
		// {defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "CRHHLWLQHG", "KANGROOO", true},
		// {defaultPolyalphabeticAlphabet, "LJOOF, WFEOI!", "HELLOWORLD", "KANGAROO", true},
	}
	for _, table := range tables {
		c := MakeVariantBeaufortCipher(table.countersign, table.alphabet)
		runPolyalphabeticReciprocalTests(t, table.plaintext, table.ciphertext, c, table.strict)
	}
}

func TestDellaPortaCipher(t *testing.T) {
	tables := []struct {
		alphabet    string
		plaintext   string
		ciphertext  string
		countersign string
		strict      bool
	}{
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "OSNYI, CLJYX!", "OCEANOGRAPHYWHAT", false},
		// {defaultPolyalphabeticAlphabet, "LJOOF, WFEOI!", "HELLO, WORLD!", "KANGAROO", false},
		// {defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "CRHHLWLQHG", "KANGROOO", true},
		// {defaultPolyalphabeticAlphabet, "LJOOF, WFEOI!", "HELLOWORLD", "KANGAROO", true},
	}
	for _, table := range tables {
		c := MakeDellaPortaCipher(table.countersign, table.alphabet)
		runPolyalphabeticReciprocalTests(t, table.plaintext, table.ciphertext, c, table.strict)
	}
}

func TestReverseString(t *testing.T) {
	table := map[string]string{
		"hello": "olleh",
		"world": "dlrow",
	}

	for k, v := range table {
		if o := reverseString(k); o != v {
			t.Errorf("Reverse of string %q was %q; expected %q", k, o, v)
		}
	}
}

func ExampleVigenereCipher() {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	c := MakeVigenereCipher("", alphabet)
	fmt.Println(c)
	// Output:
	//     A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
	//   +----------------------------------------------------
	// A | A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
	// B | B C D E F G H I J K L M N O P Q R S T U V W X Y Z A
	// C | C D E F G H I J K L M N O P Q R S T U V W X Y Z A B
	// D | D E F G H I J K L M N O P Q R S T U V W X Y Z A B C
	// E | E F G H I J K L M N O P Q R S T U V W X Y Z A B C D
	// F | F G H I J K L M N O P Q R S T U V W X Y Z A B C D E
	// G | G H I J K L M N O P Q R S T U V W X Y Z A B C D E F
	// H | H I J K L M N O P Q R S T U V W X Y Z A B C D E F G
	// I | I J K L M N O P Q R S T U V W X Y Z A B C D E F G H
	// J | J K L M N O P Q R S T U V W X Y Z A B C D E F G H I
	// K | K L M N O P Q R S T U V W X Y Z A B C D E F G H I J
	// L | L M N O P Q R S T U V W X Y Z A B C D E F G H I J K
	// M | M N O P Q R S T U V W X Y Z A B C D E F G H I J K L
	// N | N O P Q R S T U V W X Y Z A B C D E F G H I J K L M
	// O | O P Q R S T U V W X Y Z A B C D E F G H I J K L M N
	// P | P Q R S T U V W X Y Z A B C D E F G H I J K L M N O
	// Q | Q R S T U V W X Y Z A B C D E F G H I J K L M N O P
	// R | R S T U V W X Y Z A B C D E F G H I J K L M N O P Q
	// S | S T U V W X Y Z A B C D E F G H I J K L M N O P Q R
	// T | T U V W X Y Z A B C D E F G H I J K L M N O P Q R S
	// U | U V W X Y Z A B C D E F G H I J K L M N O P Q R S T
	// V | V W X Y Z A B C D E F G H I J K L M N O P Q R S T U
	// W | W X Y Z A B C D E F G H I J K L M N O P Q R S T U V
	// X | X Y Z A B C D E F G H I J K L M N O P Q R S T U V W
	// Y | Y Z A B C D E F G H I J K L M N O P Q R S T U V W X
	// Z | Z A B C D E F G H I J K L M N O P Q R S T U V W X Y
}

func ExampleBeaufortCipher() {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	c := MakeBeaufortCipher("", alphabet)
	fmt.Println(c)
	// Output:
	//     A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
	//   +----------------------------------------------------
	// Z | Z Y X W V U T S R Q P O N M L K J I H G F E D C B A
	// Y | Y X W V U T S R Q P O N M L K J I H G F E D C B A Z
	// X | X W V U T S R Q P O N M L K J I H G F E D C B A Z Y
	// W | W V U T S R Q P O N M L K J I H G F E D C B A Z Y X
	// V | V U T S R Q P O N M L K J I H G F E D C B A Z Y X W
	// U | U T S R Q P O N M L K J I H G F E D C B A Z Y X W V
	// T | T S R Q P O N M L K J I H G F E D C B A Z Y X W V U
	// S | S R Q P O N M L K J I H G F E D C B A Z Y X W V U T
	// R | R Q P O N M L K J I H G F E D C B A Z Y X W V U T S
	// Q | Q P O N M L K J I H G F E D C B A Z Y X W V U T S R
	// P | P O N M L K J I H G F E D C B A Z Y X W V U T S R Q
	// O | O N M L K J I H G F E D C B A Z Y X W V U T S R Q P
	// N | N M L K J I H G F E D C B A Z Y X W V U T S R Q P O
	// M | M L K J I H G F E D C B A Z Y X W V U T S R Q P O N
	// L | L K J I H G F E D C B A Z Y X W V U T S R Q P O N M
	// K | K J I H G F E D C B A Z Y X W V U T S R Q P O N M L
	// J | J I H G F E D C B A Z Y X W V U T S R Q P O N M L K
	// I | I H G F E D C B A Z Y X W V U T S R Q P O N M L K J
	// H | H G F E D C B A Z Y X W V U T S R Q P O N M L K J I
	// G | G F E D C B A Z Y X W V U T S R Q P O N M L K J I H
	// F | F E D C B A Z Y X W V U T S R Q P O N M L K J I H G
	// E | E D C B A Z Y X W V U T S R Q P O N M L K J I H G F
	// D | D C B A Z Y X W V U T S R Q P O N M L K J I H G F E
	// C | C B A Z Y X W V U T S R Q P O N M L K J I H G F E D
	// B | B A Z Y X W V U T S R Q P O N M L K J I H G F E D C
	// A | A Z Y X W V U T S R Q P O N M L K J I H G F E D C B
}

func ExampleGronsfeldCipher() {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	c := MakeGronsfeldCipher("", alphabet)
	fmt.Println(c)
	// Output:
	//     A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
	//   +----------------------------------------------------
	// 0 | A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
	// 1 | B C D E F G H I J K L M N O P Q R S T U V W X Y Z A
	// 2 | C D E F G H I J K L M N O P Q R S T U V W X Y Z A B
	// 3 | D E F G H I J K L M N O P Q R S T U V W X Y Z A B C
	// 4 | E F G H I J K L M N O P Q R S T U V W X Y Z A B C D
	// 5 | F G H I J K L M N O P Q R S T U V W X Y Z A B C D E
	// 6 | G H I J K L M N O P Q R S T U V W X Y Z A B C D E F
	// 7 | H I J K L M N O P Q R S T U V W X Y Z A B C D E F G
	// 8 | I J K L M N O P Q R S T U V W X Y Z A B C D E F G H
	// 9 | J K L M N O P Q R S T U V W X Y Z A B C D E F G H I
}

func ExampleTrithemiusCipher() {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	c := MakeTrithemiusCipher(alphabet)
	fmt.Println(c)
	// Output:
	//     A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
	//   +----------------------------------------------------
	// A | A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
	// B | B C D E F G H I J K L M N O P Q R S T U V W X Y Z A
	// C | C D E F G H I J K L M N O P Q R S T U V W X Y Z A B
	// D | D E F G H I J K L M N O P Q R S T U V W X Y Z A B C
	// E | E F G H I J K L M N O P Q R S T U V W X Y Z A B C D
	// F | F G H I J K L M N O P Q R S T U V W X Y Z A B C D E
	// G | G H I J K L M N O P Q R S T U V W X Y Z A B C D E F
	// H | H I J K L M N O P Q R S T U V W X Y Z A B C D E F G
	// I | I J K L M N O P Q R S T U V W X Y Z A B C D E F G H
	// J | J K L M N O P Q R S T U V W X Y Z A B C D E F G H I
	// K | K L M N O P Q R S T U V W X Y Z A B C D E F G H I J
	// L | L M N O P Q R S T U V W X Y Z A B C D E F G H I J K
	// M | M N O P Q R S T U V W X Y Z A B C D E F G H I J K L
	// N | N O P Q R S T U V W X Y Z A B C D E F G H I J K L M
	// O | O P Q R S T U V W X Y Z A B C D E F G H I J K L M N
	// P | P Q R S T U V W X Y Z A B C D E F G H I J K L M N O
	// Q | Q R S T U V W X Y Z A B C D E F G H I J K L M N O P
	// R | R S T U V W X Y Z A B C D E F G H I J K L M N O P Q
	// S | S T U V W X Y Z A B C D E F G H I J K L M N O P Q R
	// T | T U V W X Y Z A B C D E F G H I J K L M N O P Q R S
	// U | U V W X Y Z A B C D E F G H I J K L M N O P Q R S T
	// V | V W X Y Z A B C D E F G H I J K L M N O P Q R S T U
	// W | W X Y Z A B C D E F G H I J K L M N O P Q R S T U V
	// X | X Y Z A B C D E F G H I J K L M N O P Q R S T U V W
	// Y | Y Z A B C D E F G H I J K L M N O P Q R S T U V W X
	// Z | Z A B C D E F G H I J K L M N O P Q R S T U V W X Y
}

func ExampleVariantBeaufortCipher() {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	c := MakeVariantBeaufortCipher("", alphabet)
	fmt.Println(c)
	// Output:
	//     Z Y X W V U T S R Q P O N M L K J I H G F E D C B A
	//   +----------------------------------------------------
	// A | Z Y X W V U T S R Q P O N M L K J I H G F E D C B A
	// B | Y X W V U T S R Q P O N M L K J I H G F E D C B A Z
	// C | X W V U T S R Q P O N M L K J I H G F E D C B A Z Y
	// D | W V U T S R Q P O N M L K J I H G F E D C B A Z Y X
	// E | V U T S R Q P O N M L K J I H G F E D C B A Z Y X W
	// F | U T S R Q P O N M L K J I H G F E D C B A Z Y X W V
	// G | T S R Q P O N M L K J I H G F E D C B A Z Y X W V U
	// H | S R Q P O N M L K J I H G F E D C B A Z Y X W V U T
	// I | R Q P O N M L K J I H G F E D C B A Z Y X W V U T S
	// J | Q P O N M L K J I H G F E D C B A Z Y X W V U T S R
	// K | P O N M L K J I H G F E D C B A Z Y X W V U T S R Q
	// L | O N M L K J I H G F E D C B A Z Y X W V U T S R Q P
	// M | N M L K J I H G F E D C B A Z Y X W V U T S R Q P O
	// N | M L K J I H G F E D C B A Z Y X W V U T S R Q P O N
	// O | L K J I H G F E D C B A Z Y X W V U T S R Q P O N M
	// P | K J I H G F E D C B A Z Y X W V U T S R Q P O N M L
	// Q | J I H G F E D C B A Z Y X W V U T S R Q P O N M L K
	// R | I H G F E D C B A Z Y X W V U T S R Q P O N M L K J
	// S | H G F E D C B A Z Y X W V U T S R Q P O N M L K J I
	// T | G F E D C B A Z Y X W V U T S R Q P O N M L K J I H
	// U | F E D C B A Z Y X W V U T S R Q P O N M L K J I H G
	// V | E D C B A Z Y X W V U T S R Q P O N M L K J I H G F
	// W | D C B A Z Y X W V U T S R Q P O N M L K J I H G F E
	// X | C B A Z Y X W V U T S R Q P O N M L K J I H G F E D
	// Y | B A Z Y X W V U T S R Q P O N M L K J I H G F E D C
	// Z | A Z Y X W V U T S R Q P O N M L K J I H G F E D C B
}

func ExampleDellaPortaCipher() {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	c := MakeDellaPortaCipher("", alphabet)
	fmt.Println(c)
	// Output:
	//     A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
	//   +----------------------------------------------------
	// A | N O P Q R S T U V W X Y Z A B C D E F G H I J K L M
	// B | N O P Q R S T U V W X Y Z A B C D E F G H I J K L M
	// C | O P Q R S T U V W X Y Z N M A B C D E F G H I J K L
	// D | O P Q R S T U V W X Y Z N M A B C D E F G H I J K L
	// E | P Q R S T U V W X Y Z N O L M A B C D E F G H I J K
	// F | P Q R S T U V W X Y Z N O L M A B C D E F G H I J K
	// G | Q R S T U V W X Y Z N O P K L M A B C D E F G H I J
	// H | Q R S T U V W X Y Z N O P K L M A B C D E F G H I J
	// I | R S T U V W X Y Z N O P Q J K L M A B C D E F G H I
	// J | R S T U V W X Y Z N O P Q J K L M A B C D E F G H I
	// K | S T U V W X Y Z N O P Q R I J K L M A B C D E F G H
	// L | S T U V W X Y Z N O P Q R I J K L M A B C D E F G H
	// M | T U V W X Y Z N O P Q R S H I J K L M A B C D E F G
	// N | T U V W X Y Z N O P Q R S H I J K L M A B C D E F G
	// O | U V W X Y Z N O P Q R S T G H I J K L M A B C D E F
	// P | U V W X Y Z N O P Q R S T G H I J K L M A B C D E F
	// Q | V W X Y Z N O P Q R S T U F G H I J K L M A B C D E
	// R | V W X Y Z N O P Q R S T U F G H I J K L M A B C D E
	// S | W X Y Z N O P Q R S T U V E F G H I J K L M A B C D
	// T | W X Y Z N O P Q R S T U V E F G H I J K L M A B C D
	// U | X Y Z N O P Q R S T U V W D E F G H I J K L M A B C
	// V | X Y Z N O P Q R S T U V W D E F G H I J K L M A B C
	// W | Y Z N O P Q R S T U V W X C D E F G H I J K L M A B
	// X | Y Z N O P Q R S T U V W X C D E F G H I J K L M A B
	// Y | Z N O P Q R S T U V W X Y B C D E F G H I J K L M A
	// Z | Z N O P Q R S T U V W X Y B C D E F G H I J K L M A
}
