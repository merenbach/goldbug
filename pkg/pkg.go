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
	"github.com/merenbach/goldbug/pkg/masc"
	"github.com/merenbach/goldbug/pkg/pasc"
	"github.com/merenbach/goldbug/pkg/railfence"
	"github.com/merenbach/goldbug/pkg/transposition"
)

// A Cipher tranforms strings through encipherment and decipherment.
type cipher interface {
	Encipher(string) (string, error)
	Decipher(string) (string, error)
}

// Validate against this interface
var _ cipher = &railfence.Cipher{}
var _ cipher = &masc.SimpleCipher{}
var _ cipher = &transposition.Cipher{}
var _ cipher = &pasc.TabulaRectaCipher{}
