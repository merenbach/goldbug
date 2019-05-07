package main

import (
	"reflect"
	"testing"
)

func TestMirrorSequenceUnit(t *testing.T) {
	tables := []struct {
		start    int
		pivot    int
		count    int
		expected []int
	}{
		{
			start:    1,
			pivot:    5,
			count:    15,
			expected: []int{1, 2, 3, 4, 5, 4, 3, 2, 1, 2, 3, 4, 5, 4, 3},
		},
		{
			start:    5,
			pivot:    1,
			count:    15,
			expected: []int{5, 4, 3, 2, 1, 2, 3, 4, 5, 4, 3, 2, 1, 2, 3},
		},
		{
			start:    -5,
			pivot:    -1,
			count:    15,
			expected: []int{-5, -4, -3, -2, -1, -2, -3, -4, -5, -4, -3, -2, -1, -2, -3},
		},
		{
			start:    5,
			pivot:    -1,
			count:    15,
			expected: []int{5, 4, 3, 2, 1, 0, -1, 0, 1, 2, 3, 4, 5, 4, 3},
		},
		{
			start:    -1,
			pivot:    -5,
			count:    15,
			expected: []int{-1, -2, -3, -4, -5, -4, -3, -2, -1, -2, -3, -4, -5, -4, -3},
		},
		{
			start:    1,
			pivot:    -5,
			count:    15,
			expected: []int{1, 0, -1, -2, -3, -4, -5, -4, -3, -2, -1, 0, 1, 0, -1},
		},
		{
			start:    1,
			pivot:    1,
			count:    5,
			expected: []int{1, 1, 1, 1, 1},
		},
		{
			start:    -1,
			pivot:    -1,
			count:    5,
			expected: []int{-1, -1, -1, -1, -1},
		},
		{
			start:    0,
			pivot:    0,
			count:    5,
			expected: []int{0, 0, 0, 0, 0},
		},
	}
	for _, table := range tables {
		out := mirrorSequenceUnit(table.start, table.pivot, table.count)
		if !reflect.DeepEqual(out, table.expected) {
			t.Errorf("Expected %v, got %v", table.expected, out)
		}
	}
}

func TestMirrorSequenceGen(t *testing.T) {
	tables := []struct {
		start    int
		pivot    int
		count    int
		expected []int
	}{
		{
			start:    1,
			pivot:    5,
			count:    15,
			expected: []int{1, 2, 3, 4, 5, 4, 3, 2, 1, 2, 3, 4, 5, 4, 3},
		},
		{
			start:    5,
			pivot:    1,
			count:    15,
			expected: []int{5, 4, 3, 2, 1, 2, 3, 4, 5, 4, 3, 2, 1, 2, 3},
		},
		{
			start:    -5,
			pivot:    -1,
			count:    15,
			expected: []int{-5, -4, -3, -2, -1, -2, -3, -4, -5, -4, -3, -2, -1, -2, -3},
		},
		{
			start:    5,
			pivot:    -1,
			count:    15,
			expected: []int{5, 4, 3, 2, 1, 0, -1, 0, 1, 2, 3, 4, 5, 4, 3},
		},
		{
			start:    -1,
			pivot:    -5,
			count:    15,
			expected: []int{-1, -2, -3, -4, -5, -4, -3, -2, -1, -2, -3, -4, -5, -4, -3},
		},
		{
			start:    1,
			pivot:    -5,
			count:    15,
			expected: []int{1, 0, -1, -2, -3, -4, -5, -4, -3, -2, -1, 0, 1, 0, -1},
		},
		{
			start:    1,
			pivot:    1,
			count:    5,
			expected: []int{1, 1, 1, 1, 1},
		},
		{
			start:    -1,
			pivot:    -1,
			count:    5,
			expected: []int{-1, -1, -1, -1, -1},
		},
		{
			start:    0,
			pivot:    0,
			count:    5,
			expected: []int{0, 0, 0, 0, 0},
		},
	}
	for _, table := range tables {
		gen := mirrorSequenceGen(table.start, table.pivot)
		out := make([]int, table.count)
		for i := range out {
			out[i] = gen()
		}
		if !reflect.DeepEqual(out, table.expected) {
			t.Errorf("Expected %v, got %v", table.expected, out)
		}
	}
}
