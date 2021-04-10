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

func TestToString(t *testing.T) {
	table := []struct {
		value    interface{}
		expected string
	}{
		{0, "0"},
		{3.14, "3.14"},
		{true, "true"},
		{"a", "a"},
	}
	for _, row := range table {
		got := toString(row.value)
		if !reflect.DeepEqual(got, row.expected) {
			t.Errorf("got: %s, expected: %s", got, row.expected)
		}
	}
}

func TestToStrings(t *testing.T) {
	table := []struct {
		values   []interface{}
		expected []string
	}{
		{[]interface{}{0, 3.14, true, "a"}, []string{"0", "3.14", "true", "a"}},
	}
	for _, row := range table {
		got := toStrings(row.values)
		if !reflect.DeepEqual(got, row.expected) {
			t.Errorf("got: %s, expected: %s", got, row.expected)
		}
	}
}

func TestWrap(t *testing.T) {
	table := []struct {
		left     string
		right    string
		values   []string
		expected []string
	}{
		{"", "", []string{"a"}, []string{"a"}},
		{"", "", []string{"a", "b"}, []string{"a", "b"}},
		{"x", "", []string{"a", "b"}, []string{"xa", "xb"}},
		{"x", "y", []string{"xay", "xby"}, []string{"xxayy", "xxbyy"}},
	}
	for _, row := range table {
		got := wrap(row.left, row.right, row.values)
		if !reflect.DeepEqual(got, row.expected) {
			t.Errorf("got: %s, expected: %s", got, row.expected)
		}
	}
}

func TestZip(t *testing.T) {
	table := []struct {
		separator   string
		leftValues  []string
		rightValues []string
		expected    []string
	}{
		{"", []string{"a"}, []string{}, nil},
		{"", []string{"a"}, []string{"b"}, []string{"ab"}},
		{"-", []string{"a", "b"}, []string{"c", "d"}, []string{"a-c", "b-d"}},
	}
	for _, row := range table {
		got := zip(row.separator, row.leftValues, row.rightValues)
		if !reflect.DeepEqual(got, row.expected) {
			t.Errorf("got: %s, expected: %s", got, row.expected)
		}
	}
}
