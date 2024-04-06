// Пакет stringutil_test содержит тесты для пакета stringutil.
package stringutil_test

import (
	stringutil "awesomeProject"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello", "olleH"},
		{"world", "dlrow"},
		{"", ""},
	}

	for _, test := range tests {
		if output := stringutil.Reverse(test.input); output != test.expected {
			t.Errorf("Reverse(%q) = %q, ожидалось %q", test.input, output, test.expected)
		}
	}
}

func TestLength(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"Hello", 5},
		{"world", 5},
		{"", 0},
		{"Привет", 6},
	}

	for _, test := range tests {
		if output := stringutil.Length(test.input); output != test.expected {
			t.Errorf("Length(%q) = %d, ожидалось %d", test.input, output, test.expected)
		}
	}
}
