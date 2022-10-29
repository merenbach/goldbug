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

package simple

import (
	"fmt"

	"github.com/merenbach/goldbug/internal/masc"
)

// A Cipher implements a simple cipher.
type Cipher struct {
	ptAlphabet string
	ctAlphabet string
	caseless   bool
	strict     bool

	*masc.Tableau
}

// adapted from: https://www.sohamkamani.com/golang/options-pattern/

type CipherOption func(*Cipher)

func WithStrict() CipherOption {
	return func(c *Cipher) {
		c.strict = true
	}
}

func WithCaseless() CipherOption {
	return func(c *Cipher) {
		c.caseless = true
	}
}

func WithPtAlphabet(s string) CipherOption {
	return func(c *Cipher) {
		c.ptAlphabet = s
	}
}

func WithCtAlphabet(s string) CipherOption {
	return func(c *Cipher) {
		c.ctAlphabet = s
	}
}

func NewCipher(opts ...CipherOption) (*Cipher, error) {
	c := &Cipher{
		ptAlphabet: masc.Alphabet,
		ctAlphabet: masc.Alphabet,
	}
	for _, opt := range opts {
		opt(c)
	}

	tableau, err := masc.NewTableau(
		masc.WithPtAlphabet(c.ptAlphabet),
		masc.WithCtAlphabet(c.ctAlphabet),
		masc.WithStrict(c.strict),
		masc.WithCaseless(c.caseless),
	)
	if err != nil {
		return nil, fmt.Errorf("could not create tableau: %w", err)
	}
	c.Tableau = tableau

	return c, nil
}
