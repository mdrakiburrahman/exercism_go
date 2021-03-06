package queenattack

import "testing"

// Arguments to CanQueenAttack are in algebraic notation.
// See http://en.wikipedia.org/wiki/Algebraic_notation_(chess)

var tests = []struct {
	w, b   string
	attack bool
	ok     bool
}{
	// Error cases
	{"a8", "b9", false, false}, // off board
	{"", "", false, false},
	{"here", "there", false, false}, // invalid
	{"b4", "b4", false, false},      // same square
	{"a0", "b1", false, false},
	{"g3", "i5", false, false},

	// Proper cases:
	// Same column
	{"b4", "b7", true, true}, // same file

	// Same row:
	{"e4", "b4", true, true}, // same rank

	// Diagonal?
	{"a1", "f6", true, true}, // common diagonals
	{"a6", "b7", true, true},
	{"d1", "f3", true, true},
	{"f1", "a6", true, true},
	{"a1", "h8", true, true},
	{"a8", "h1", true, true},

	// No hit:
	{"b3", "d7", false, true}, // no attack
	{"a1", "f8", false, true},
}

func TestCanQueenAttack(t *testing.T) {
	for _, test := range tests {
		switch attack, err := CanQueenAttack(test.w, test.b); {
		case err != nil:
			var _ error = err
			if test.ok {
				t.Fatalf("CanQueenAttack(%s, %s) returned error %q.  "+
					"Error not expected.",
					test.w, test.b, err)
			}
		case !test.ok:
			t.Fatalf("CanQueenAttack(%s, %s) = %t, %v.  Expected error.",
				test.w, test.b, attack, err)
		case attack != test.attack:
			t.Fatalf("CanQueenAttack(%s, %s) = %t, want %t.",
				test.w, test.b, attack, test.attack)
		}
	}
}

// Benchmark combined time for all test cases
func BenchmarkCanQueenAttack(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			CanQueenAttack(test.w, test.b)
		}
	}
}
