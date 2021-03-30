package crude

import (
	"reflect"
	"testing"
)

func TestJoin(t *testing.T) {
	table := []struct {
		separator string
		values    []string
		expected  string
	}{
		{",", []string{"a", "b", "c"}, "a,b,c"},
		{", ", []string{"c", "b", "a"}, "c, b, a"},
		{"", []string{"x", "y", "z"}, "xyz"},
	}
	for _, row := range table {
		got := join(row.separator, row.values)
		if got != row.expected {
			t.Errorf("got: %s, expected: %s", got, row.expected)
		}
	}
}

func TestRepeat(t *testing.T) {
	table := []struct {
		times    int
		value    string
		expected []string
	}{
		{0, "a", []string{}},
		{1, "b", []string{"b"}},
		{10, "c", []string{"c", "c", "c", "c", "c", "c", "c", "c", "c", "c"}},
	}
	for _, row := range table {
		got := repeat(row.times, row.value)
		if !reflect.DeepEqual(got, row.expected) {
			t.Errorf("got: %s, expected: %s", got, row.expected)
		}
	}
}
