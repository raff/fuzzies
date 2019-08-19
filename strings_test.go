package fuzzies

import "testing"

var compareTests = []struct {
	a, b string
	c    bool
}{
	{"", "", true},
	{"", " ", true},
	{" ", "", true},
	{"a", "", true},
	{"", "a", true},
	{"abc", "abc", true},
	{"abc", "adc", true},
	{"adc", "abc", true},
	{"hello", "heloa", true},
	{"hello", "haloa", false},
	{"hello", "there", false},
}

func TestCompare(t *testing.T) {
	for _, tt := range compareTests {
		cmp := Compare(tt.a, tt.b)
		if cmp != tt.c {
			t.Errorf(`Compare(%q, %q) = %v`, tt.a, tt.b, cmp)
		}
	}
}

var containsTests = []struct {
	a, b string
	c    bool
}{
	{"", "", true},
	{"", " ", true},
	{" ", "", true},
	{"a", "", true},
	{"", "a", true},
	{"  abc  ", "abc", true},
	{"abcdefgh", "adc", true},
	{"dfgabcgh", "adc", true},
	{"say hello to me", "hello", true},
	{"say heloa to me", "hello", true},
	{"say haloa to me", "hello", false},
	{"here and there", "hello", false},
}

func TestContains(t *testing.T) {
	for _, tt := range compareTests {
		cmp := Contains(tt.a, tt.b)
		if cmp != tt.c {
			t.Errorf(`Contains(%q, %q) = %v`, tt.a, tt.b, cmp)
		}
	}
}

var indexTests = []struct {
	a, b string
	i    int
}{
	{"", "", 0},
	{"", " ", 0},
	{" ", "", 0},
	{"a", "", 0},
	{"", "a", 0},
	{"  abc  ", "abc", 2},
	{"abcdefgh", "adc", 0},
	{"dfgabcgh", "adc", 3},
	{"say hello to me", "hello", 4},
	{"say heloa to me", "hello", 3}, // should be 4, but the distance between ` helo`/hello and heloa/hello is the same
	{"say haloa to me", "hello", -1},
	{"here and there", "hello", -1},
}

func TestIndex(t *testing.T) {
	for _, tt := range indexTests {
		cmp := Index(tt.a, tt.b)
		if cmp != tt.i {
			t.Errorf(`Index(%q, %q) = %v`, tt.a, tt.b, cmp)
		}
	}
}
