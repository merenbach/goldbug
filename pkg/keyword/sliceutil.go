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

package keyword

import "github.com/merenbach/goldbug/internal/stringutil"

// Key a string with prefix text.
// TODO: rename Prefix? or something to allow semantically for suffix counterpart?
func Transform(xs string, keyword string) (string, error) {
	return stringutil.Deduplicate(keyword + xs), nil
}

/*
func Transform[T any](xs []T, keyword []T) ([]T, error) {
	ys := keyword[:]
	ys = append(ys, xs)
	return stringutil.Deduplicate(ys)
}
*/
