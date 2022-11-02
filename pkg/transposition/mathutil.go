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

package transposition

import "fmt"

// Abs returns the absolute value for an integer.
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// TODO: use this to make rail fence a special case of columnar transposition with Myszkowski
func zigzag(period int) []int {
	if period < 0 {
		panic(fmt.Sprintf("period must be nonnegative, but got: %d", period))
	}

	out := make([]int, period)
	for i := range out {
		n := i % period
		out[i] = min(n, period-n)
	}
	return out
}
