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

package api

import (
	"encoding/json"
	"errors"

	"github.com/merenbach/goldbug/pkg/beaufort"
	"github.com/merenbach/goldbug/pkg/dellaporta"
	"github.com/merenbach/goldbug/pkg/gronsfeld"
	"github.com/merenbach/goldbug/pkg/trithemius"
	"github.com/merenbach/goldbug/pkg/variantbeaufort"
	"github.com/merenbach/goldbug/pkg/vigenere"
)

// MascBaseConfig is a base configuration for a monoalphabetic substitution cipher operation
type mascBaseConfig struct {
	Alphabet string `json:"alphabet"`
	Message  string `json:"message"`
	Reverse  bool   `json:"reverse"`
	Strict   bool   `json:"strict"`
}

// MascBaseConfig is a base configuration for a polyalphabetic substitution cipher operation
type pascBaseConfig struct {
	mascBaseConfig
	Countersign string `json:"countersign"`
}

// Affine cipher processing
func Affine(s string) (string, error) {
	var payload struct {
		mascBaseConfig
		Shift      int `json:"shift"`
		Multiplier int `json:"multiplier"`
	}
	if err := json.Unmarshal([]byte(s), &payload); err != nil {
		return "", err
	}

	c := affine.Cipher{
		Alphabet:  payload.Alphabet,
		Slope:     payload.Multiplier,
		Intercept: payload.Shift,
		Strict:    payload.Strict,
	}

	if payload.Reverse {
		return c.Decipher(payload.Message)
	}
	return c.Encipher(payload.Message)
}

// Atbash cipher processing
func Atbash(s string) (string, error) {
	var payload struct {
		mascBaseConfig
	}
	if err := json.Unmarshal([]byte(s), &payload); err != nil {
		return "", err
	}

	c := atbash.Cipher{
		Alphabet: payload.Alphabet,
		Strict:   payload.Strict,
	}

	if payload.Reverse {
		return c.Decipher(payload.Message)
	}
	return c.Encipher(payload.Message)
}

// Caesar cipher processing
func Caesar(s string) (string, error) {
	var payload struct {
		mascBaseConfig
		Shift int `json:"shift"`
	}
	if err := json.Unmarshal([]byte(s), &payload); err != nil {
		return "", err
	}

	c := caesar.Cipher{
		Alphabet: payload.Alphabet,
		Shift:    payload.Shift,
		Strict:   payload.Strict,
	}

	if payload.Reverse {
		return c.Decipher(payload.Message)
	}
	return c.Encipher(payload.Message)
}

// Decimation cipher processing
func Decimation(s string) (string, error) {
	var payload struct {
		mascBaseConfig
		Multiplier int `json:"multiplier"`
	}
	if err := json.Unmarshal([]byte(s), &payload); err != nil {
		return "", err
	}

	c := decimation.Cipher{
		Alphabet:   payload.Alphabet,
		Multiplier: payload.Multiplier,
		Strict:     payload.Strict,
	}

	if payload.Reverse {
		return c.Decipher(payload.Message)
	}
	return c.Encipher(payload.Message)
}

// Keyword cipher processing
func Keyword(s string) (string, error) {
	var payload struct {
		mascBaseConfig
		Keyword string `json:"keyword"`
	}
	if err := json.Unmarshal([]byte(s), &payload); err != nil {
		return "", err
	}
	c := keyword.Cipher{
		Alphabet: payload.Alphabet,
		Keyword:  payload.Keyword,
		Strict:   payload.Strict,
	}

	if payload.Reverse {
		return c.Decipher(payload.Message)
	}
	return c.Encipher(payload.Message)
}

// Rot13 cipher processing
func Rot13(s string) (string, error) {
	var payload struct {
		mascBaseConfig
	}
	if err := json.Unmarshal([]byte(s), &payload); err != nil {
		return "", err
	}

	c := rot13.Cipher{}

	if payload.Reverse {
		return c.Decipher(payload.Message)
	}
	return c.Encipher(payload.Message)
}

// Vigenere cipher processing
func Vigenere(s string) (string, error) {
	var payload struct {
		pascBaseConfig
		TextAutoclave bool `json:"textAutoclave"`
		KeyAutoclave  bool `json:"keyAutoclave"`
	}
	if err := json.Unmarshal([]byte(s), &payload); err != nil {
		return "", err
	}

	if payload.TextAutoclave && payload.KeyAutoclave {
		return "", errors.New("Text autoclave and key autoclave are mutually exclusive")
	}

	c := vigenere.Cipher{
		Alphabet: payload.Alphabet,
		Key:      payload.Countersign,
		Strict:   payload.Strict,
	}

	if payload.TextAutoclave {
		c.Autokey = vigenere.TextAutokey
	} else if payload.KeyAutoclave {
		c.Autokey = vigenere.KeyAutokey
	}

	if payload.Reverse {
		return c.Decipher(payload.Message)
	}
	return c.Encipher(payload.Message)
}

// Beaufort cipher processing
func Beaufort(s string) (string, error) {
	var payload struct {
		pascBaseConfig
	}
	if err := json.Unmarshal([]byte(s), &payload); err != nil {
		return "", err
	}

	c := beaufort.Cipher{
		Alphabet: payload.Alphabet,
		Key:      payload.Countersign,
		Strict:   payload.Strict,
	}

	if payload.Reverse {
		return c.Decipher(payload.Message)
	}
	return c.Encipher(payload.Message)
}

// DellaPorta cipher processing
func DellaPorta(s string) (string, error) {
	var payload struct {
		pascBaseConfig
	}
	if err := json.Unmarshal([]byte(s), &payload); err != nil {
		return "", err
	}

	c := dellaporta.Cipher{
		Alphabet: payload.Alphabet,
		Key:      payload.Countersign,
		Strict:   payload.Strict,
	}

	if payload.Reverse {
		return c.Decipher(payload.Message)
	}
	return c.Encipher(payload.Message)
}

// Gronsfeld cipher processing
func Gronsfeld(s string) (string, error) {
	var payload struct {
		pascBaseConfig
	}
	if err := json.Unmarshal([]byte(s), &payload); err != nil {
		return "", err
	}

	c := gronsfeld.Cipher{
		Alphabet: payload.Alphabet,
		Key:      payload.Countersign,
		Strict:   payload.Strict,
	}

	if payload.Reverse {
		return c.Decipher(payload.Message)
	}
	return c.Encipher(payload.Message)
}

// Trithemius cipher processing
func Trithemius(s string) (string, error) {
	var payload struct {
		pascBaseConfig
	}
	if err := json.Unmarshal([]byte(s), &payload); err != nil {
		return "", err
	}

	c := trithemius.Cipher{
		Alphabet: payload.Alphabet,
		Strict:   payload.Strict,
	}

	if payload.Reverse {
		return c.Decipher(payload.Message)
	}
	return c.Encipher(payload.Message)
}

// VariantBeaufort cipher processing
func VariantBeaufort(s string) (string, error) {
	var payload struct {
		pascBaseConfig
	}
	if err := json.Unmarshal([]byte(s), &payload); err != nil {
		return "", err
	}

	c := variantbeaufort.Cipher{
		Alphabet: payload.Alphabet,
		Key:      payload.Countersign,
		Strict:   payload.Strict,
	}

	if payload.Reverse {
		return c.Decipher(payload.Message)
	}
	return c.Encipher(payload.Message)
}
