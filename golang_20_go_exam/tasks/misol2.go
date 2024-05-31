package tasks

import (
 "errors"
 "strings"
 "unicode"
)

func TextToTitle(text string) (string, error) {
 if len(text) == 0 {
  return "", errors.New("matn uzunligi 0 ga teng")
 }

 hasLetter := false
 for _, r := range text {
  if unicode.IsLetter(r) {
   hasLetter = true
   break
  }
 }
 if !hasLetter {
  return "", errors.New("matnda umuman harf yo'q")
 }

 return strings.Title(strings.ToLower(text)), nil
}