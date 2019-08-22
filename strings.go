package fuzzies

import "strings"

type Fuzzies struct {
	IgnoreCase  bool // case insensitive operations
	MinDistance int  // minimum distance to be considered a match
}

// Default configuration for "fuzzies" operations
var Default = Fuzzies{IgnoreCase: true, MinDistance: 3}

// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
//
// Note that Index will try to find the closest match (i.e. the position that returns the minimum "distance" between
// subst and the slice of s).
func (f *Fuzzies) Index(s, substr string) int {
	if f.IgnoreCase {
		s = strings.ToLower(s)
		substr = strings.ToLower(substr)
	}

	ls := len(s)
	lsub := len(substr)

	if ls <= lsub {
		d := Levenshtein(s, substr)
		if d < f.MinDistance {
			return 0
		}

		return -1
	}

	pos := -1
	mind := f.MinDistance

	if mind > (lsub+1)/2 {
		mind = (lsub + 1) / 2
	}

	for i := 0; i < (ls - lsub + 1); i++ {
		d := Levenshtein(substr, s[i:i+lsub])
		if d < mind {
			mind = d
			pos = i
		}
	}

	return pos
}

// Contains reports whether substr is within s.
func (f *Fuzzies) Contains(s, substr string) bool {
	return f.Index(s, substr) >= 0
}

// Compare returns true if a and b are within the specified distance (they are considered equal)
func (f *Fuzzies) Compare(a, b string) bool {
	if f.IgnoreCase {
		a = strings.ToLower(a)
		b = strings.ToLower(b)
	}

	d := Levenshtein(a, b)
	return d < f.MinDistance
}

// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
//
// Note that Index will try to find the closest match (i.e. the position that returns the minimum "distance" between
// subst and the slice of s).
func Index(s, substr string) int {
	return Default.Index(s, substr)
}

// Contains reports whether substr is within s.
func Contains(s, substr string) bool {
	return Default.Contains(s, substr)
}

// Compare returns true if a and b are within the specified distance (they are considered equal)
func Compare(s, substr string) bool {
	return Default.Compare(s, substr)
}
