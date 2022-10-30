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

package affine

import (
	"fmt"

	"github.com/merenbach/goldbug/internal/masc"
)

// A Config struct for a cipher.
type Config struct {
	alphabet string
	caseless bool
	strict   bool
}

func (c *Config) Alphabet() string {
	return c.alphabet
}

func (c *Config) Caseless() bool {
	return c.caseless
}

func (c *Config) Strict() bool {
	return c.strict
}

// A Cipher implements an affine cipher.
type Cipher struct {
	*Config
	*masc.Tableau
}

// adapted from: https://www.sohamkamani.com/golang/options-pattern/

type ConfigOption func(*Config)

func WithStrict() ConfigOption {
	return func(c *Config) {
		c.strict = true
	}
}

func WithCaseless() ConfigOption {
	return func(c *Config) {
		c.caseless = true
	}
}

func WithAlphabet(s string) ConfigOption {
	return func(c *Config) {
		c.alphabet = s
	}
}

func NewConfig(opts ...ConfigOption) *Config {
	c := &Config{alphabet: masc.Alphabet}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func NewCipher(slope int, intercept int, opts ...ConfigOption) (*Cipher, error) {
	c := NewConfig(opts...)

	ctAlphabet, err := Transform([]rune(c.alphabet), slope, intercept)
	if err != nil {
		return nil, fmt.Errorf("could not transform alphabet: %w", err)
	}

	tableau, err := masc.NewTableau(
		c.alphabet,
		string(ctAlphabet),
		c.strict,
		c.caseless,
	)
	if err != nil {
		return nil, fmt.Errorf("could not create tableau: %w", err)
	}

	return &Cipher{
		c,
		tableau,
	}, nil
}
