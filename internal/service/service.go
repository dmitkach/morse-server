package service

import (
	"morse-server/pkg/morse"
	"unicode"
)

func TextTypeSwitch(s string) string {
	var isMorse = true
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			isMorse = false
			break
		}
	}

	var res string
	if isMorse {
		res = morse.ToText(s)
	} else {
		res = morse.ToMorse(s)
	}

	return res
}
