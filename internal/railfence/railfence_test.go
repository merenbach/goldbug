package railfence

import (
	"testing"
)

func TestCipher(t *testing.T) {
	tables := []struct {
		ciphertext string
		plaintext  string
		rails      int
	}{
		{
			ciphertext: "WECRLTEERDSOEEFEAOCAIVDEN",
			plaintext:  "WEAREDISCOVEREDFLEEATONCE",
			rails:      3,
		},
		{
			ciphertext: "IA_EZS_ELYLK_UZERLIPL",
			plaintext:  "I_REALLY_LIKE_PUZZLES",
			rails:      3,
		},
		{
			ciphertext: "MEMTNGTXETEOIHQZ",
			plaintext:  "MEETMETONIGHTQXZ",
			rails:      2,
		},
		{
			ciphertext: "MMNTETEOIHQZETGX",
			plaintext:  "MEETMETONIGHTQXZ",
			rails:      3,
		},
		{
			ciphertext: "AALUHNHSEDFYMNAGIGIHAOFZ",
			plaintext:  "AMANLAUGHINGHISHEADOFFYZ",
			rails:      2,
		},
		{
			ciphertext: "HELLO, WORLD!",
			plaintext:  "HELLO, WORLD!",
			rails:      1,
		},
	}

	for _, table := range tables {
		c := Cipher{Rails: table.rails}
		if out := c.Decipher(table.ciphertext); out != table.plaintext {
			t.Errorf("Expected %q to decipher to %q, but instead got %q", table.ciphertext, table.plaintext, out)
		}

		if out := c.Encipher(table.plaintext); out != table.ciphertext {
			t.Errorf("Expected %q to encipher to %q, but instead got %q", table.plaintext, table.ciphertext, out)
		}
	}
}

func TestCipherReversibility(t *testing.T) {
	base := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	baseRunes := []rune(base)

	var counter int
	for msglen := 1; msglen < len(baseRunes); msglen++ {
		for rails := 1; rails < len(baseRunes); rails++ {
			counter++

			msg := string(baseRunes[:msglen])
			t.Logf("Iteration %d (rails = %d, message length = %d, message = %s)", counter, rails, msglen, msg)

			c := Cipher{Rails: rails}

			enciphered1 := c.Encipher(msg)
			deciphered1 := c.Decipher(enciphered1)
			if deciphered1 != msg {
				t.Errorf("Expected encipherment-then-decipherment of %q to be %q but got %q", enciphered1, msg, deciphered1)
			}

			enciphered2 := c.Decipher(msg)
			deciphered2 := c.Encipher(enciphered2)
			if deciphered2 != msg {
				t.Errorf("Expected decipherment-then-encipherment of %q to be %q but got %q", enciphered1, msg, deciphered1)
			}
		}
	}
}
