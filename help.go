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

type Hex string

func catpuccinToColor(namedColor catppuccin.Color) *color.Color {
	r, g, b, _ := namedColor.RGBA()
	R, G, B := int(r), int(g), int(b)
	color := color.RGB(R, G, B)
	return color
}

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

func Mirror(palette Palette) Palette {
	var newPalette Palette
	for i := range palette {
		newPalette = append(newPalette, palette[i])
		newPalette = append(newPalette, palette[i])
	}
	return newPalette
}

func Randomise(palette Palette) Palette {
	rand.Shuffle(len(palette), func(i, j int) {
		palette[i], palette[j] = palette[j], palette[i]
	})
	return palette
}

func HexToColor(hex Hex) *color.Color {
	clr, err := colorful.Hex(string(hex))
	if err != nil {
		log.Fatal(err)
	}
	r, g, b := clr.RGB255()
	return color.RGB(int(r), int(g), int(b))
}
