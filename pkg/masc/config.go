// Copyright 2022 Andrew Merenbach
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

const Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// A Config struct for a cipher.
type Config struct {
	alphabet string
	caseless bool
	strict   bool
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
	const (
		defaultAlphabet = Alphabet
	)

	c := &Config{alphabet: defaultAlphabet}
	for _, opt := range opts {
		opt(c)
	}
	return c
}
