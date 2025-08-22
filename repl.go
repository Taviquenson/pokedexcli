package main

import (
	"strings"
)

func cleanInput(text string) []string {
	subStrs := strings.Fields(text)
	for i, str := range subStrs {
		subStrs[i] = strings.ToLower(str)
	}
	return subStrs
}
