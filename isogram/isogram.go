// Package isogram determines if a word or phrase is an isogram.
package isogram

import "strings"

// IsIsogram checks whether given word isogram or not
func IsIsogram(word string) bool {
	m := make(map[string]int)
	for _, r := range word {
		s := string(r)
		s = strings.ToLower(s)
		if s == "-" || s == "_" || s == " " {
			continue
		}
		if val, ok := m[s]; ok {
			m[s] = val + 1
		} else {
			m[s] = 1
		}
		if m[s] > 1 {
			return false
		}
	}
	return true
}
