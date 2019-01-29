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

func runPolyalphabeticEncipherTest(t *testing.T, input string, expected string, c *VigenereFamilyCipher, strict bool) {
	output := c.Encipher(input, strict)
	if string(output) != expected {
		t.Errorf("ciphertext %q for input %q was incorrect; wanted %q", output, input, expected)
	}
}
func runPolyalphabeticDecipherTest(t *testing.T, input string, expected string, c *VigenereFamilyCipher, strict bool) {
	output := c.Decipher(input, strict)
	if string(output) != expected {
		t.Errorf("plaintext %q for input %q was incorrect; wanted: %q", output, input, expected)
	}
}

// func runPolyalphabeticReciprocalTests(t *testing.T, plaintext, ciphertext string, c VigenereFamilyCipher, strict bool) {
// 	encrypted := c.Encipher(plaintext, strict)
// 	decrypted := c.Decipher(ciphertext, strict)
// 	if string(encrypted) != ciphertext {
// 		t.Errorf("ciphertext %q was incorrect; wanted %q", encrypted, ciphertext)
// 	}
// 	if string(decrypted) != plaintext {
// 		t.Errorf("plaintext %q was incorrect; wanted: %q", decrypted, plaintext)
// 	}
// }

func TestVigenereCipher(t *testing.T) {
	encipherTables := []struct {
		alphabet    string
		input       string
		expected    string
		countersign string
		strict      bool
	}{
		{"", "", "", "", false},
		{defaultPolyalphabeticAlphabet, "", "", "OCEANOGRAPHYWHAT", false},
		{defaultPolyalphabeticAlphabet, "", "", "OCEANOGRAPHYWHAT", true},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "HELLO, WORLD!", "A", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "HELLOWORLD", "A", true},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "VGPLB, KUILS!", "OCEANOGRAPHYWHAT", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "VGPLBKUILS", "OCEANOGRAPHYWHAT", true},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "XUBBE, MEHBT!", "Q", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "XUBBEMEHBT", "Q", true},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "REYRO, NCFVD!", "KANGAROO", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "REYRONCFVD", "KANGAROO", true},
	}
	decipherTables := []struct {
		alphabet    string
		input       string
		expected    string
		countersign string
		strict      bool
	}{
		{"", "", "", "", false},
		{defaultPolyalphabeticAlphabet, "", "", "OCEANOGRAPHYWHAT", false},
		{defaultPolyalphabeticAlphabet, "", "", "OCEANOGRAPHYWHAT", true},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "HELLO, WORLD!", "A", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "HELLOWORLD", "A", true},
		{defaultPolyalphabeticAlphabet, "VGPLB, KUILS!", "HELLO, WORLD!", "OCEANOGRAPHYWHAT", false},
		{defaultPolyalphabeticAlphabet, "VGPLB, KUILS!", "HELLOWORLD", "OCEANOGRAPHYWHAT", true},
		{defaultPolyalphabeticAlphabet, "XUBBE, MEHBT!", "HELLO, WORLD!", "Q", false},
		{defaultPolyalphabeticAlphabet, "XUBBE, MEHBT!", "HELLOWORLD", "Q", true},
		{defaultPolyalphabeticAlphabet, "REYRO, NCFVD!", "HELLO, WORLD!", "KANGAROO", false},
		{defaultPolyalphabeticAlphabet, "REYRO, NCFVD!", "HELLOWORLD", "KANGAROO", true},
	}

	for _, table := range encipherTables {
		c := NewVigenereCipher(table.countersign, table.alphabet)
		runPolyalphabeticEncipherTest(t, table.input, table.expected, c, table.strict)
	}
	for _, table := range decipherTables {
		c := NewVigenereCipher(table.countersign, table.alphabet)
		runPolyalphabeticDecipherTest(t, table.input, table.expected, c, table.strict)
	}
}

func TestVigenereTextAutoclaveCipher(t *testing.T) {
	encipherTables := []struct {
		alphabet    string
		input       string
		expected    string
		countersign string
		strict      bool
	}{
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "VGPLB, KUILS!", "OCEANOGRAPHYWHAT", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "VGPLBKUILS", "OCEANOGRAPHYWHAT", true},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "HLPWZ, KKFCO!", "A", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "HLPWZKKFCO", "A", true},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "XLPWZ, KKFCO!", "Q", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "XLPWZKKFCO", "Q", true},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "REYRO, NCFSH!", "KANGAROO", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "REYRONCFSH", "KANGAROO", true},
	}
	decipherTables := []struct {
		alphabet    string
		input       string
		expected    string
		countersign string
		strict      bool
	}{
		{defaultPolyalphabeticAlphabet, "VGPLB, KUILS!", "HELLO, WORLD!", "OCEANOGRAPHYWHAT", false},
		{defaultPolyalphabeticAlphabet, "VGPLB, KUILS!", "HELLOWORLD", "OCEANOGRAPHYWHAT", true},
		{defaultPolyalphabeticAlphabet, "HLPWZ, KKFCO!", "HELLO, WORLD!", "A", false},
		{defaultPolyalphabeticAlphabet, "HLPWZ, KKFCO!", "HELLOWORLD", "A", true},
		{defaultPolyalphabeticAlphabet, "XLPWZ, KKFCO!", "HELLO, WORLD!", "Q", false},
		{defaultPolyalphabeticAlphabet, "XLPWZ, KKFCO!", "HELLOWORLD", "Q", true},
		{defaultPolyalphabeticAlphabet, "REYRO, NCFSH!", "HELLO, WORLD!", "KANGAROO", false},
		{defaultPolyalphabeticAlphabet, "REYRO, NCFSH!", "HELLOWORLD", "KANGAROO", true},
	}

	for _, table := range encipherTables {
		c := NewVigenereTextAutoclaveCipher(table.countersign, table.alphabet)
		runPolyalphabeticEncipherTest(t, table.input, table.expected, c, table.strict)
	}
	for _, table := range decipherTables {
		c := NewVigenereTextAutoclaveCipher(table.countersign, table.alphabet)
		runPolyalphabeticDecipherTest(t, table.input, table.expected, c, table.strict)
	}
}

func TestVigenereKeyAutoclaveCipher(t *testing.T) {
	encipherTables := []struct {
		alphabet    string
		input       string
		expected    string
		countersign string
		strict      bool
	}{
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "VGPLB, KUILS!", "OCEANOGRAPHYWHAT", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "VGPLBKUILS", "OCEANOGRAPHYWHAT", true},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "HLWHV, RFWHK!", "A", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "HLWHVRFWHK", "A", true},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "XBMXL, HVMXA!", "Q", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "XBMXLHVMXA", "Q", true},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "REYRO, NCFCH!", "KANGAROO", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "REYRONCFCH", "KANGAROO", true},
	}
	decipherTables := []struct {
		alphabet    string
		input       string
		expected    string
		countersign string
		strict      bool
	}{
		{defaultPolyalphabeticAlphabet, "VGPLB, KUILS!", "HELLO, WORLD!", "OCEANOGRAPHYWHAT", false},
		{defaultPolyalphabeticAlphabet, "VGPLB, KUILS!", "HELLOWORLD", "OCEANOGRAPHYWHAT", true},
		{defaultPolyalphabeticAlphabet, "HLWHV, RFWHK!", "HELLO, WORLD!", "A", false},
		{defaultPolyalphabeticAlphabet, "HLWHV, RFWHK!", "HELLOWORLD", "A", true},
		{defaultPolyalphabeticAlphabet, "XBMXL, HVMXA!", "HELLO, WORLD!", "Q", false},
		{defaultPolyalphabeticAlphabet, "XBMXL, HVMXA!", "HELLOWORLD", "Q", true},
		{defaultPolyalphabeticAlphabet, "REYRO, NCFCH!", "HELLO, WORLD!", "KANGAROO", false},
		{defaultPolyalphabeticAlphabet, "REYRO, NCFCH!", "HELLOWORLD", "KANGAROO", true},
	}

	for _, table := range encipherTables {
		c := NewVigenereKeyAutoclaveCipher(table.countersign, table.alphabet)
		runPolyalphabeticEncipherTest(t, table.input, table.expected, c, table.strict)
	}
	for _, table := range decipherTables {
		c := NewVigenereKeyAutoclaveCipher(table.countersign, table.alphabet)
		runPolyalphabeticDecipherTest(t, table.input, table.expected, c, table.strict)
	}
}

func TestBeaufortCipher(t *testing.T) {
	encipherTables := []struct {
		alphabet    string
		input       string
		expected    string
		countersign string
		strict      bool
	}{
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "HYTPZ, SSAPM!", "OCEANOGRAPHYWHAT", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "HYTPZSSAPM", "OCEANOGRAPHYWHAT", true},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "JMFFC, UCZFN!", "Q", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "JMFFCUCZFN", "Q", true},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "DWCVM, VAXZX!", "KANGAROO", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "DWCVMVAXZX", "KANGAROO", true},
	}
	decipherTables := []struct {
		alphabet    string
		input       string
		expected    string
		countersign string
		strict      bool
	}{
		{defaultPolyalphabeticAlphabet, "HYTPZ, SSAPM!", "HELLO, WORLD!", "OCEANOGRAPHYWHAT", false},
		{defaultPolyalphabeticAlphabet, "HYTPZ, SSAPM!", "HELLOWORLD", "OCEANOGRAPHYWHAT", true},
		{defaultPolyalphabeticAlphabet, "JMFFC, UCZFN!", "HELLO, WORLD!", "Q", false},
		{defaultPolyalphabeticAlphabet, "JMFFC, UCZFN!", "HELLOWORLD", "Q", true},
		{defaultPolyalphabeticAlphabet, "DWCVM, VAXZX!", "HELLO, WORLD!", "KANGAROO", false},
		{defaultPolyalphabeticAlphabet, "DWCVM, VAXZX!", "HELLOWORLD", "KANGAROO", true},
	}

	for _, table := range encipherTables {
		c := NewBeaufortCipher(table.countersign, table.alphabet)
		runPolyalphabeticEncipherTest(t, table.input, table.expected, c, table.strict)
	}

	for _, table := range decipherTables {
		c := NewBeaufortCipher(table.countersign, table.alphabet)
		runPolyalphabeticDecipherTest(t, table.input, table.expected, c, table.strict)
	}
}

func TestGronsfeldCipher(t *testing.T) {
	encipherTables := []struct {
		alphabet    string
		input       string
		expected    string
		countersign string
		strict      bool
	}{
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "HELLO, WORLD!", "0", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "HELLOWORLD", "0", true},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "JHMOQ, YRSOF!", "23132", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "JHMOQYRSOF", "23132", true},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "KMUNX, WPRNG!", "389290102394957", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "KMUNXWPRNG", "389290102394957", true},
	}
	decipherTables := []struct {
		alphabet    string
		input       string
		expected    string
		countersign string
		strict      bool
	}{
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "HELLO, WORLD!", "0", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "HELLOWORLD", "0", true},
		{defaultPolyalphabeticAlphabet, "JHMOQ, YRSOF!", "HELLO, WORLD!", "23132", false},
		{defaultPolyalphabeticAlphabet, "JHMOQ, YRSOF!", "HELLOWORLD", "23132", true},
		{defaultPolyalphabeticAlphabet, "KMUNX, WPRNG!", "HELLO, WORLD!", "389290102394957", false},
		{defaultPolyalphabeticAlphabet, "KMUNX, WPRNG!", "HELLOWORLD", "389290102394957", true},
	}

	for _, table := range encipherTables {
		c := NewGronsfeldCipher(table.countersign, table.alphabet)
		runPolyalphabeticEncipherTest(t, table.input, table.expected, c, table.strict)
	}
	for _, table := range decipherTables {
		c := NewGronsfeldCipher(table.countersign, table.alphabet)
		runPolyalphabeticDecipherTest(t, table.input, table.expected, c, table.strict)
	}
}

func TestTrithemiusCipher(t *testing.T) {
	encipherTables := []struct {
		alphabet string
		input    string
		expected string
		strict   bool
	}{
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "HFNOS, BUYTM!", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "HFNOSBUYTM", true},
	}
	decipherTables := []struct {
		alphabet string
		input    string
		expected string
		strict   bool
	}{
		{defaultPolyalphabeticAlphabet, "HFNOS, BUYTM!", "HELLO, WORLD!", false},
		{defaultPolyalphabeticAlphabet, "HFNOS, BUYTM!", "HELLOWORLD", true},
	}

	for _, table := range encipherTables {
		c := NewTrithemiusCipher(table.alphabet)
		runPolyalphabeticEncipherTest(t, table.input, table.expected, c, table.strict)
	}
	for _, table := range decipherTables {
		c := NewTrithemiusCipher(table.alphabet)
		runPolyalphabeticDecipherTest(t, table.input, table.expected, c, table.strict)
	}
}

func TestVariantBeaufortCipher(t *testing.T) {
	encipherTables := []struct {
		alphabet    string
		input       string
		expected    string
		countersign string
		strict      bool
	}{
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "TCHLB, IIALO!", "OCEANOGRAPHYWHAT", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "TCHLBIIALO", "OCEANOGRAPHYWHAT", true},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "HELLO, WORLD!", "A", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "HELLOWORLD", "A", true},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "ROVVY, GYBVN!", "Q", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "ROVVYGYBVN", "Q", true},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "XEYFO, FADBD!", "KANGAROO", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "XEYFOFADBD", "KANGAROO", true},
	}
	decipherTables := []struct {
		alphabet    string
		input       string
		expected    string
		countersign string
		strict      bool
	}{
		{defaultPolyalphabeticAlphabet, "TCHLB, IIALO!", "HELLO, WORLD!", "OCEANOGRAPHYWHAT", false},
		{defaultPolyalphabeticAlphabet, "TCHLB, IIALO!", "HELLOWORLD", "OCEANOGRAPHYWHAT", true},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "HELLO, WORLD!", "A", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "HELLOWORLD", "A", true},
		{defaultPolyalphabeticAlphabet, "ROVVY, GYBVN!", "HELLO, WORLD!", "Q", false},
		{defaultPolyalphabeticAlphabet, "ROVVY, GYBVN!", "HELLOWORLD", "Q", true},
		{defaultPolyalphabeticAlphabet, "XEYFO, FADBD!", "HELLO, WORLD!", "KANGAROO", false},
		{defaultPolyalphabeticAlphabet, "XEYFO, FADBD!", "HELLOWORLD", "KANGAROO", true},
	}

	for _, table := range encipherTables {
		c := NewVariantBeaufortCipher(table.countersign, table.alphabet)
		runPolyalphabeticEncipherTest(t, table.input, table.expected, c, table.strict)
	}
	for _, table := range decipherTables {
		c := NewVariantBeaufortCipher(table.countersign, table.alphabet)
		runPolyalphabeticDecipherTest(t, table.input, table.expected, c, table.strict)
	}
}

func TestDellaPortaCipher(t *testing.T) {
	encipherTables := []struct {
		alphabet    string
		input       string
		expected    string
		countersign string
		strict      bool
	}{
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "OSNYI, CLJYX!", "OCEANOGRAPHYWHAT", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "OSNYICLJYX", "OCEANOGRAPHYWHAT", true},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "PZTTG, BGJTY!", "Q", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "PZTTGBGJTY", "Q", true},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "ZRROB, BHKQQ!", "KANGAROO", false},
		{defaultPolyalphabeticAlphabet, "HELLO, WORLD!", "ZRROBBHKQQ", "KANGAROO", true},
	}
	decipherTables := []struct {
		alphabet    string
		input       string
		expected    string
		countersign string
		strict      bool
	}{
		{defaultPolyalphabeticAlphabet, "OSNYI, CLJYX!", "HELLO, WORLD!", "OCEANOGRAPHYWHAT", false},
		{defaultPolyalphabeticAlphabet, "OSNYI, CLJYX!", "HELLOWORLD", "OCEANOGRAPHYWHAT", true},
		{defaultPolyalphabeticAlphabet, "PZTTG, BGJTY!", "HELLO, WORLD!", "Q", false},
		{defaultPolyalphabeticAlphabet, "PZTTG, BGJTY!", "HELLOWORLD", "Q", true},
		{defaultPolyalphabeticAlphabet, "ZRROB, BHKQQ!", "HELLO, WORLD!", "KANGAROO", false},
		{defaultPolyalphabeticAlphabet, "ZRROB, BHKQQ!", "HELLOWORLD", "KANGAROO", true},
	}
	for _, table := range encipherTables {
		c := NewDellaPortaCipher(table.countersign, table.alphabet)
		runPolyalphabeticEncipherTest(t, table.input, table.expected, c, table.strict)
	}
	for _, table := range decipherTables {
		c := NewDellaPortaCipher(table.countersign, table.alphabet)
		runPolyalphabeticDecipherTest(t, table.input, table.expected, c, table.strict)
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
	c := NewVigenereCipher("", alphabet)
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
	c := NewBeaufortCipher("", alphabet)
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
	c := NewGronsfeldCipher("", alphabet)
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
	c := NewTrithemiusCipher(alphabet)
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
	c := NewVariantBeaufortCipher("", alphabet)
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
	c := NewDellaPortaCipher("", alphabet)
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
