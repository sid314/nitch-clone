package main

import (
	"log"
	"math/rand"
	"strings"

	catppuccin "github.com/catppuccin/go"
	"github.com/fatih/color"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/rivo/uniseg"
)

type hex string

func catpuccinToColor(namedColor catppuccin.Color) *color.Color {
	r, g, b, _ := namedColor.RGBA()
	R, G, B := int(r), int(g), int(b)
	color := color.RGB(R, G, B)
	return color
}

func largestFieldLength(disableColors bool, printables printables) int {
	largestFieldLength := 0
	if !disableColors {
		largestFieldLength = 9
	}
	for _, printable := range printables {
		if l := uniseg.StringWidth(printable.Field); l > largestFieldLength {
			largestFieldLength = l
		}
	}
	return largestFieldLength
}

func snipSnip(prefix string, suffix string, s string) string {
	_, aftercut, _ := strings.Cut(s, prefix)
	beforecut, _, _ := strings.Cut(aftercut, suffix)
	return beforecut
}

func wrap(fieldsNumber int, colors palette) palette {
	var newPalette palette
	for len(newPalette) <= fieldsNumber+1 {
		newPalette = append(newPalette, colors...)
	}
	return newPalette
}

func mirror(colors palette) palette {
	var newPalette palette
	for i := range colors {
		newPalette = append(newPalette, colors[i])
		newPalette = append(newPalette, colors[i])
	}
	return newPalette
}

func randomise(colors palette) palette {
	rand.Shuffle(len(colors), func(i, j int) {
		colors[i], colors[j] = colors[j], colors[i]
	})
	return colors
}

func hexToColor(hex hex) *color.Color {
	clr, err := colorful.Hex(string(hex))
	if err != nil {
		log.Fatal(err)
	}
	r, g, b := clr.RGB255()
	return color.RGB(int(r), int(g), int(b))
}

func fieldsFromPrintableInfo(printables printables) []string {
	var fields []string
	for i := range printables {
		fields = append(fields, printables[i].Field)
	}
	return fields
}
