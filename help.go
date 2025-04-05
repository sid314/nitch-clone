package main

import (
	"fmt"
	"strings"
)

func SnipSnip(prefix string, suffix string, s string) string {
	_, aftercut, _ := strings.Cut(s, prefix)
	beforecut, _, _ := strings.Cut(aftercut, suffix)
	return beforecut
}

func PrintConfig(config Config) {
	fmt.Println(config.Theme)
	fmt.Println(config.Border)
	fmt.Println(config.Dot)
	fmt.Println(config.Printables)
}
