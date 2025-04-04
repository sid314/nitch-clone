package main

import (
	"fmt"
	"strings"

	catppuccin "github.com/catppuccin/go"
	"github.com/fatih/color"
)

func SnipSnip(prefix string, suffix string, s string) string {
	_, aftercut, _ := strings.Cut(s, prefix)
	beforecut, _, _ := strings.Cut(aftercut, suffix)
	return beforecut
}

func catppuccinToColor(namedColor catppuccin.Color) *color.Color {
	r, g, b, _ := namedColor.RGBA()
	R, G, B := int(r), int(g), int(b)
	color := color.RGB(R, G, B)
	return color
}

func GeneratePrintFunctions(colors Palette) PrintFunctions {
	var functions [16]func(a ...any) string
	for i := range colors {
		functions[i] = colors[i].SprintFunc()
	}
	return functions
}

func GeneratePalette(theme ThemeName) Palette {
	switch theme {
	case "catppuccin-mocha", "catppuccin-latte", "catppuccin-frappe", "catppuccin-macchiato":
		return generateCatpuccinPalette(theme)
	default:
		return generateGrayscalePalette()

	}
}

func PrintConfig(config Config) {
	fmt.Println(config.Style)
	fmt.Println(config.Theme)
}

func generateGrayscalePalette() Palette {
	var palette Palette
	for i := range palette {
		palette[i] = color.RGB(255, 255, 255)
	}
	return palette
}

func generateCatpuccinPalette(theme ThemeName) Palette {
	var flavour catppuccin.Flavor
	var palette Palette
	switch theme {
	case "catppuccin-mocha":
		flavour = catppuccin.Mocha
	case "catppuccin-macchiato":
		flavour = catppuccin.Macchiato
	case "catppuccin-frappe":
		flavour = catppuccin.Frappe
	case "catppuccin-latte":
		flavour = catppuccin.Latte
	}
	palette[0] = catppuccinToColor(flavour.Sky())
	palette[1] = catppuccinToColor(flavour.Sapphire())
	palette[2] = catppuccinToColor(flavour.Lavender())
	palette[3] = catppuccinToColor(flavour.Mauve())
	palette[4] = catppuccinToColor(flavour.Blue())
	palette[5] = catppuccinToColor(flavour.Sky())
	palette[6] = catppuccinToColor(flavour.Peach())
	palette[7] = catppuccinToColor(flavour.Yellow())
	palette[8] = catppuccinToColor(flavour.Red())
	palette[9] = catppuccinToColor(flavour.Maroon())
	palette[10] = catppuccinToColor(flavour.Green())
	palette[11] = catppuccinToColor(flavour.Teal())
	palette[12] = catppuccinToColor(flavour.Rosewater())
	palette[13] = catppuccinToColor(flavour.Flamingo())
	palette[14] = catppuccinToColor(flavour.Pink())
	palette[15] = catppuccinToColor(flavour.Lavender())
	return palette
}
