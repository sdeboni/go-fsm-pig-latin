// Author bobahop, copied from exercism deepdive section
// Package piglatin is a small library for translating a sentence into Pig Latin.
package maplookup

import (
	"strings"
)

type void struct{}

var blank void
var vowels = map[byte]void{'a': blank, 'e': blank, 'i': blank, 'o': blank, 'u': blank}
var specials = map[string]void{"xr": blank, "yt": blank}
var vowels_y = map[byte]void{'a': blank, 'e': blank, 'i': blank, 'o': blank, 'u': blank, 'y': blank}

type container interface {
	byte | string
}

func contains[T container](values map[T]void, value T) bool {
	_, ok := values[value]
	return ok
}

// Sentence translates a sentence into Pig Latin.
func Sentence(phrase string) string {
	piggyfied := strings.Builder{}

	for _, word := range strings.Split(phrase, " ") {
		if piggyfied.Len() != 0 {
			piggyfied.WriteByte(' ')
		}

		if contains(vowels, word[0]) || contains(specials, word[:2]) {
			piggyfied.WriteString(word + "ay")
			continue
		}

		for pos := 1; pos < len(word); pos++ {
			letter := word[pos]
			if contains(vowels_y, letter) {
				if letter == 'u' && word[pos-1] == 'q' {
					pos++
				}
				piggyfied.WriteString(word[pos:] + word[:pos] + "ay")
				break
			}
		}
	}
	return piggyfied.String()
}
