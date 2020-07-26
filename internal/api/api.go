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

	"github.com/merenbach/gold-bug/internal/masc"
	"github.com/merenbach/gold-bug/internal/pasc"
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

	c, err := masc.NewAffineCipher(payload.Alphabet, payload.Multiplier, payload.Shift)
	if err != nil {
		return "", err
	}

	c.Strict = payload.Strict
	if payload.Reverse {
		return c.DecipherString(payload.Message)
	}
	return c.EncipherString(payload.Message)
}

// Atbash cipher processing
func Atbash(s string) (string, error) {
	var payload struct {
		mascBaseConfig
	}
	if err := json.Unmarshal([]byte(s), &payload); err != nil {
		return "", err
	}

	c, err := masc.NewAtbashCipher(payload.Alphabet)
	if err != nil {
		return "", err
	}

	c.Strict = payload.Strict
	if payload.Reverse {
		return c.DecipherString(payload.Message)
	}
	return c.EncipherString(payload.Message)
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

	c, err := masc.NewCaesarCipher(payload.Alphabet, payload.Shift)
	if err != nil {
		return "", err
	}

	c.Strict = payload.Strict
	if payload.Reverse {
		return c.DecipherString(payload.Message)
	}
	return c.EncipherString(payload.Message)
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

	c, err := masc.NewDecimationCipher(payload.Alphabet, payload.Multiplier)
	if err != nil {
		return "", err
	}

	c.Strict = payload.Strict
	if payload.Reverse {
		return c.DecipherString(payload.Message)
	}
	return c.EncipherString(payload.Message)
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
	c, err := masc.NewKeywordCipher(payload.Alphabet, payload.Keyword)
	if err != nil {
		return "", err
	}

	c.Strict = payload.Strict
	if payload.Reverse {
		return c.DecipherString(payload.Message)
	}
	return c.EncipherString(payload.Message)
}

// Rot13 cipher processing
func Rot13(s string) (string, error) {
	var payload struct {
		mascBaseConfig
	}
	if err := json.Unmarshal([]byte(s), &payload); err != nil {
		return "", err
	}

	c, err := masc.NewRot13Cipher(payload.Alphabet)
	if err != nil {
		return "", err
	}

	c.Strict = payload.Strict
	if payload.Reverse {
		return c.DecipherString(payload.Message)
	}
	return c.EncipherString(payload.Message)
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

	c, err := pasc.NewVigenereCipher(payload.Countersign, payload.Alphabet)
	if err != nil {
		return "", err
	}

	c.TextAutoclave = payload.TextAutoclave
	c.KeyAutoclave = payload.KeyAutoclave
	c.Strict = payload.Strict
	if payload.Reverse {
		return c.DecipherString(payload.Message), nil
	}
	return c.EncipherString(payload.Message), nil
}

// Beaufort cipher processing
func Beaufort(s string) (string, error) {
	var payload struct {
		pascBaseConfig
	}
	if err := json.Unmarshal([]byte(s), &payload); err != nil {
		return "", err
	}

	c, err := pasc.NewBeaufortCipher(payload.Countersign, payload.Alphabet)
	if err != nil {
		return "", err
	}

	c.Strict = payload.Strict
	if payload.Reverse {
		return c.DecipherString(payload.Message), nil
	}
	return c.EncipherString(payload.Message), nil
}

// DellaPorta cipher processing
func DellaPorta(s string) (string, error) {
	var payload struct {
		pascBaseConfig
	}
	if err := json.Unmarshal([]byte(s), &payload); err != nil {
		return "", err
	}

	c, err := pasc.NewDellaPortaCipher(payload.Countersign, payload.Alphabet)
	if err != nil {
		return "", err
	}

	c.Strict = payload.Strict
	if payload.Reverse {
		return c.DecipherString(payload.Message), nil
	}
	return c.EncipherString(payload.Message), nil
}

// Gronsfeld cipher processing
func Gronsfeld(s string) (string, error) {
	var payload struct {
		pascBaseConfig
	}
	if err := json.Unmarshal([]byte(s), &payload); err != nil {
		return "", err
	}

	c, err := pasc.NewGronsfeldCipher(payload.Countersign, payload.Alphabet)
	if err != nil {
		return "", err
	}

	c.Strict = payload.Strict
	if payload.Reverse {
		return c.DecipherString(payload.Message), nil
	}
	return c.EncipherString(payload.Message), nil
}

// Trithemius cipher processing
func Trithemius(s string) (string, error) {
	var payload struct {
		pascBaseConfig
	}
	if err := json.Unmarshal([]byte(s), &payload); err != nil {
		return "", err
	}

	c, err := pasc.NewTrithemiusCipher(payload.Alphabet)
	if err != nil {
		return "", err
	}

	c.Strict = payload.Strict
	if payload.Reverse {
		return c.DecipherString(payload.Message), nil
	}
	return c.EncipherString(payload.Message), nil
}

// VariantBeaufort cipher processing
func VariantBeaufort(s string) (string, error) {
	var payload struct {
		pascBaseConfig
	}
	if err := json.Unmarshal([]byte(s), &payload); err != nil {
		return "", err
	}

	c, err := pasc.NewVariantBeaufortCipher(payload.Countersign, payload.Alphabet)
	if err != nil {
		return "", err
	}

	c.Strict = payload.Strict
	if payload.Reverse {
		return c.DecipherString(payload.Message), nil
	}
	return c.EncipherString(payload.Message), nil
}
