// Iteration 2 assumes input is ascii only for benchmark comparison with deepdive solution
// This statemachine implementation benchmarked 20% faster than the deepdive generics map implementation
// 1395ns vs 1692ns, with the tradeoff being 3x more lines of code: SLOC 123 vs 49
// 
package piglatin

import (
  "strings"
)

type state int

const (
  START state = iota
  LEADING_CONSONANTS
  PREV_Q
  LEADING_X
  LEADING_Y
)

func Sentence(sentence string) string {
  words := strings.Split(strings.TrimSpace(sentence), " ")
  if len(words) == 0 {
    return ""
  }
  word := words[0]

  var output strings.Builder
  output.WriteString(translate(word))

  for _, word := range words[1:] {
    if len(word) == 0 {
      continue
    }
    output.WriteString(" ")
    output.WriteString(translate(word))
  }

  return output.String()
}

func translate(word string) string {
  var output strings.Builder
  st := START

  bytes := []byte(word)

  for pos := 0; pos < len(bytes); pos++ {
    curr := bytes[pos]

    switch st {
    case START:
      if isVowel(curr) {
        output.WriteString(word)
        output.WriteString("ay")
        return output.String()
      } else if curr == 'x' {
        st = LEADING_X
      } else if curr == 'y' {
        st = LEADING_Y
      } else if curr == 'q' {
        st = PREV_Q
      } else {
        st = LEADING_CONSONANTS
      }
    case LEADING_CONSONANTS:
      if curr == 'y' {
        if pos < len(bytes) {
          output.WriteString(word[pos:])
        }
        output.WriteString(word[:pos])
        output.WriteString("ay")
        return output.String()
      } else if curr == 'q' {
        st = PREV_Q
      } else if isVowel(curr) {
        output.WriteString(word[pos:])
        output.WriteString(word[:pos])
        output.WriteString("ay")
        return output.String()
      }
    case LEADING_X:
      if curr == 'r' {
        output.WriteString(word)
        output.WriteString("ay")
        return output.String()
      }
      st = LEADING_CONSONANTS
      pos -= 1;
    case LEADING_Y:
      if curr == 't' {
        output.WriteString(word)
        output.WriteString("ay")
        return output.String()
      }
      st = LEADING_CONSONANTS
      pos -= 1;
    case PREV_Q:
      if curr == 'u' {
        if pos < len(bytes)-1 {
          output.WriteString(word[pos+1:])
        }
        output.WriteString(word[:pos+1])
        output.WriteString("ay")
        return output.String()
      }
      st = LEADING_CONSONANTS
      pos -= 1
    }
  } 

  return word
}

var vowels = []byte{'a','e','i','o','u'}
func isVowel(b byte) bool {
  for _, v := range vowels {
    if b == v {
      return true
    }
  }
  return false
}

func isConsonant(b byte) bool {
  return !isVowel(b)
}
