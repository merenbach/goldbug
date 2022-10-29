// Copyright 2020 Andrew Merenbach
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

package pkg

import (
	"github.com/merenbach/goldbug/pkg/affine"
	"github.com/merenbach/goldbug/pkg/atbash"
	"github.com/merenbach/goldbug/pkg/beaufort"
	"github.com/merenbach/goldbug/pkg/caesar"
	"github.com/merenbach/goldbug/pkg/decimation"
	"github.com/merenbach/goldbug/pkg/dellaporta"
	"github.com/merenbach/goldbug/pkg/gromark"
	"github.com/merenbach/goldbug/pkg/gronsfeld"
	"github.com/merenbach/goldbug/pkg/keyword"
	"github.com/merenbach/goldbug/pkg/railfence"
	"github.com/merenbach/goldbug/pkg/rot13"
	"github.com/merenbach/goldbug/pkg/transposition"
	"github.com/merenbach/goldbug/pkg/trithemius"
	"github.com/merenbach/goldbug/pkg/variantbeaufort"
	"github.com/merenbach/goldbug/pkg/vigenere"
)

// A Cipher tranforms strings through encipherment and decipherment.
type cipher interface {
	Encipher(string) (string, error)
	Decipher(string) (string, error)
}

// Validate against this interface
var _ cipher = &affine.Cipher{}
var _ cipher = &atbash.Cipher{}
var _ cipher = &beaufort.Cipher{}
var _ cipher = &caesar.Cipher{}
var _ cipher = &decimation.Cipher{}
var _ cipher = &dellaporta.Cipher{}
var _ cipher = &gromark.Cipher{}
var _ cipher = &gronsfeld.Cipher{}
var _ cipher = &keyword.Cipher{}
var _ cipher = &railfence.Cipher{}
var _ cipher = &rot13.Cipher{}
var _ cipher = &transposition.Cipher{}
var _ cipher = &trithemius.Cipher{}
var _ cipher = &variantbeaufort.Cipher{}
var _ cipher = &vigenere.Cipher{}
