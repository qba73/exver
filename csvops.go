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