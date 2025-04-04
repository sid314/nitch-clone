package main

import (
	"math/rand"

	catppuccin "github.com/catppuccin/go"
	"github.com/fatih/color"
)

type (
	Palette        [16]*color.Color
	PrintFunctions [16]func(a ...any) string
)

type (
	StyleName string
	ThemeName string
	Theme     struct {
		name   ThemeName
		colors Palette
	}
)

func Color(namedColor catppuccin.Color) *color.Color {
	r, g, b, _ := namedColor.RGBA()
	R, G, B := int(r), int(g), int(b)
	color := color.RGB(R, G, B)
	return color
}

func GeneratePrintFunctions(theme ThemeName) PrintFunctions {
	var colors Palette
	switch theme {
	case "catppuccin-mocha", "catppuccin-latte", "catppuccin-frappe", "catppuccin-macchiato":
		colors = catpuccinPalette(theme)
	case "6-colors":
		colors = sixColorPalette()
	case "6-colors-high-intensity":
		colors = sixHighIntensityColorPalette()
	case "random-6-colors":
		colors = randomSixColorPalette()
	case "random-6-colors-high-intensity":
		colors = randomHighIntensitySixcolorPalette()

	default:
		colors = grayscalePalette()

	}
	var functions [16]func(a ...any) string
	for i := range colors {
		functions[i] = colors[i].SprintFunc()
	}
	return functions
}

func grayscalePalette() Palette {
	var palette Palette
	for i := range palette {
		palette[i] = color.RGB(255, 255, 255)
	}
	return palette
}

func sixColorPalette() Palette {
	var palette Palette
	palette[0] = color.New(color.FgBlue)
	palette[1] = color.New(color.FgBlue)
	palette[2] = color.New(color.FgRed)
	palette[3] = color.New(color.FgRed)
	palette[4] = color.New(color.FgYellow)
	palette[5] = color.New(color.FgYellow)
	palette[6] = color.New(color.FgCyan)
	palette[7] = color.New(color.FgCyan)
	palette[8] = color.New(color.FgGreen)
	palette[9] = color.New(color.FgGreen)
	palette[10] = color.New(color.FgMagenta)
	palette[11] = color.New(color.FgMagenta)
	palette[12] = color.New(color.FgRed)
	palette[13] = color.New(color.FgRed)
	palette[14] = color.New(color.FgYellow)
	palette[15] = color.New(color.FgYellow)
	return palette
}

func sixHighIntensityColorPalette() Palette {
	var palette Palette
	palette[0] = color.New(color.FgHiBlue)
	palette[1] = color.New(color.FgHiBlue)
	palette[2] = color.New(color.FgHiRed)
	palette[3] = color.New(color.FgHiRed)
	palette[4] = color.New(color.FgHiYellow)
	palette[5] = color.New(color.FgHiYellow)
	palette[6] = color.New(color.FgHiCyan)
	palette[7] = color.New(color.FgHiCyan)
	palette[8] = color.New(color.FgHiGreen)
	palette[9] = color.New(color.FgHiGreen)
	palette[10] = color.New(color.FgHiMagenta)
	palette[11] = color.New(color.FgHiMagenta)
	palette[12] = color.New(color.FgHiRed)
	palette[13] = color.New(color.FgHiRed)
	palette[14] = color.New(color.FgHiYellow)
	palette[15] = color.New(color.FgHiYellow)
	return palette
}

func randomSixColorPalette() Palette {
	colors := [16]*color.Color{
		color.New(color.FgBlue),
		color.New(color.FgBlue),
		color.New(color.FgRed),
		color.New(color.FgRed),
		color.New(color.FgYellow),
		color.New(color.FgYellow),
		color.New(color.FgCyan),
		color.New(color.FgCyan),
		color.New(color.FgGreen),
		color.New(color.FgGreen),
		color.New(color.FgMagenta),
		color.New(color.FgMagenta),
		color.New(color.FgYellow),
		color.New(color.FgYellow),
		color.New(color.FgRed),
		color.New(color.FgRed),
	}
	var palette Palette
	for i := range len(palette) {
		random := rand.Intn(16)
		palette[i] = colors[random]

	}
	return palette
}

func randomHighIntensitySixcolorPalette() Palette {
	colors := [16]*color.Color{
		color.New(color.FgHiBlue),
		color.New(color.FgHiBlue),
		color.New(color.FgHiRed),
		color.New(color.FgHiRed),
		color.New(color.FgHiYellow),
		color.New(color.FgHiYellow),
		color.New(color.FgHiCyan),
		color.New(color.FgHiCyan),
		color.New(color.FgHiGreen),
		color.New(color.FgHiGreen),
		color.New(color.FgHiMagenta),
		color.New(color.FgHiMagenta),
		color.New(color.FgHiYellow),
		color.New(color.FgHiYellow),
		color.New(color.FgHiRed),
		color.New(color.FgHiRed),
	}
	var palette Palette
	for i := range len(palette) {
		random := rand.Intn(16)
		palette[i] = colors[random]

	}
	return palette
}

func catpuccinPalette(theme ThemeName) Palette {
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
	palette[0] = Color(flavour.Sky())
	palette[1] = Color(flavour.Sapphire())
	palette[2] = Color(flavour.Lavender())
	palette[3] = Color(flavour.Mauve())
	palette[4] = Color(flavour.Blue())
	palette[5] = Color(flavour.Sky())
	palette[6] = Color(flavour.Peach())
	palette[7] = Color(flavour.Yellow())
	palette[8] = Color(flavour.Red())
	palette[9] = Color(flavour.Maroon())
	palette[10] = Color(flavour.Green())
	palette[11] = Color(flavour.Teal())
	palette[12] = Color(flavour.Rosewater())
	palette[13] = Color(flavour.Flamingo())
	palette[14] = Color(flavour.Pink())
	palette[15] = Color(flavour.Lavender())
	return palette
}
