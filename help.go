package main

import (
	"strings"

	"github.com/rivo/uniseg"
)

func largestFieldLength(disableColors bool, printables []PrintableInfo) int {
	largestFieldLength := 0
	if !disableColors {
		largestFieldLength = 8
	}
	for _, printable := range printables {
		if l := uniseg.StringWidth(printable.Field); l > largestFieldLength {
			largestFieldLength = l
		}
	}
	return largestFieldLength
}

func SnipSnip(prefix string, suffix string, s string) string {
	_, aftercut, _ := strings.Cut(s, prefix)
	beforecut, _, _ := strings.Cut(aftercut, suffix)
	return beforecut
}

func wrap(fieldsNumber int, palette Palette) Palette {
	var newPalette Palette
	for len(newPalette) <= fieldsNumber+1 {
		newPalette = append(newPalette, palette...)
	}
	return newPalette
}
