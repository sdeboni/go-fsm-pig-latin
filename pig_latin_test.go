package piglatin

import (
	"testing"
  "piglatin/statemachine"
  "piglatin/maplookup"
)

func TestPigLatin(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := statemachine.Sentence(tc.input); actual != tc.expected {
				t.Fatalf("Sentence(%q) = %q, want %q", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkSentenceStatemachine(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			statemachine.Sentence(tc.input)
		}
	}
}
func BenchmarkSentenceMaplookup(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			maplookup.Sentence(tc.input)
		}
	}
}

