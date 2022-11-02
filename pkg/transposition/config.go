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

package transposition

// A Config struct for a cipher.
type Config struct {
	myszkowski bool
	keys       [][]int
}

// adapted from: https://www.sohamkamani.com/golang/options-pattern/

type ConfigOption func(*Config)

func WithMyszkowski() ConfigOption {
	return func(c *Config) {
		c.myszkowski = true
	}
}

func WithStringKey(s string) ConfigOption {
	return func(c *Config) {
		ii := lexorder([]rune(s))
		c.keys = append(c.keys, ii)
	}
}

func WithIntegerKey(i []int) ConfigOption {
	return func(c *Config) {
		c.keys = append(c.keys, i)
	}
}

func NewConfig(opts ...ConfigOption) *Config {
	c := &Config{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}
