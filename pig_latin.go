package piglatin

import (
  "unicode"
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

  runes := []rune(word)

  for pos := 0; pos < len(runes); pos++ {
    curr := unicode.ToLower(runes[pos])

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
        if pos < len(runes) {
          for _, r := range runes[pos:] {
            output.WriteRune(r)
          }
        }
        for _, r := range runes[:pos] { 
          output.WriteRune(r)
        }
        output.WriteString("ay")
        return output.String()
      } else if curr == 'q' {
        st = PREV_Q
      } else if isVowel(curr) {
        for _, r := range runes[pos:] {
          output.WriteRune(r)
        }
        for _, r := range runes[:pos] { 
          output.WriteRune(r)
        }
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
        if pos < len(runes)-1 {
          for _, r := range runes[pos+1:] {
            output.WriteRune(r)
          }
        }
        for _, r := range runes[:pos+1] {
          output.WriteRune(r)
        }
        output.WriteString("ay")
        return output.String()
      }
      st = LEADING_CONSONANTS
      pos -= 1
    }
  } 

  return word
}

var vowels = []rune{'a','e','i','o','u'}
func isVowel(r rune) bool {
  for _, v := range vowels {
    if r == v {
      return true
    }
  }
  return false
}

func isConsonant(r rune) bool {
  return !isVowel(r)
}
