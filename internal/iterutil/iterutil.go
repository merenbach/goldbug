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

package iterutil

// Successors to a start value based on a function.
// Successors is intended in the spirit of std::iter::successors in Rust.
func Successors[T any](first T, succ func(T) T) func() T {
	state := first
	return func() T {
		prev := state
		state = succ(state)
		return prev
	}
}

// Take a slice of values from a generating function.
func Take[T any](n int, f func() T) []T {
	out := make([]T, n)
	for i := range out {
		out[i] = f()
	}
	return out
}
