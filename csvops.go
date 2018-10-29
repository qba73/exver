package main

import "strings"

func fixDots(s string) string {
	s = strings.TrimSuffix(s, ".")
	s = strings.Replace(s, ".", "_", 1)
	s = strings.ToUpper(s)
	return s
}

func fixDash(s string) string {
	s = strings.Replace(s, "-", "_", 5)
	s = strings.ToUpper(s)
	return s
}

func fixSpaces(s string) string {
	s = strings.Replace(s, " ", "_", 5)
	s = strings.ToUpper(s)
	return s
}

func headerType(s string) int {
	var h int
	if strings.Contains(s, ".") {
		h = 1
	} else if strings.Contains(s, "-") {
		h = 2
	} else if strings.Contains(s, " ") {
		h = 3
	} else {
		h = 4
	}
	return h
}

func fixHeader(sx []string) []string {
	for i, v := range sx {
		h := headerType(v)
		switch h {
		case 1:
			sx[i] = fixDots(v)
		case 2:
			sx[i] = fixDash(v)
		case 3:
			sx[i] = fixSpaces(v)
		default:
			sx[i] = v
		}
	}
	return sx
}
