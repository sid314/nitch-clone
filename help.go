package main

import (
	"strings"
)

func SnipSnip(prefix string, suffix string, s string) string {
	_, aftercut, _ := strings.Cut(s, prefix)
	beforecut, _, _ := strings.Cut(aftercut, suffix)
	return beforecut
}
