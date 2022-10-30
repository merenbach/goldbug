package sliceutil

import (
	"reflect"
	"testing"

	"github.com/merenbach/goldbug/internal/fixture"
)

func TestArgsort(t *testing.T) {
	var tables []struct {
		Input  []int
		Output []int
	}

	fixture.Load(t, &tables)
	for i, table := range tables {
		t.Logf("Testing table %d of %d", i+1, len(tables))
		if o := Argsort(table.Input); !reflect.DeepEqual(o, table.Output) {
			t.Errorf("Expected %+v and got %+v", table.Output, o)
		}
	}
}

func TestBackpermute(t *testing.T) {
	var tables []struct {
		Input   []any
		Output  []any
		Indices []int
		Failure bool
	}

	fixture.Load(t, &tables)
	for i, table := range tables {
		t.Logf("Testing table %d of %d", i+1, len(tables))
		if out, err := Backpermute(table.Input, table.Indices); err != nil && !table.Failure {
			t.Error("Unexpected backpermute failure:", err)
		} else if err == nil && table.Failure {
			t.Error("Unexpected backpermute success")
		} else if !reflect.DeepEqual(out, table.Output) {
			t.Error("Received incorrect output:", out)
		}
	}
}

func TestDeduplicate(t *testing.T) {
	table := map[string]string{
		"hello":       "helo",
		"world":       "world",
		"hello world": "helo wrd",
	}

	for k, v := range table {
		if o := Deduplicate([]rune(k)); !reflect.DeepEqual(o, []rune(v)) {
			t.Errorf("Deduplication of string %q was %q; expected %q", k, o, v)
		}
	}
}
