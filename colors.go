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

func catppuccinToColor(namedColor catppuccin.Color) *color.Color {
	r, g, b, _ := namedColor.RGBA()
	R, G, B := int(r), int(g), int(b)
	color := color.RGB(R, G, B)
	return color
}

func GeneratePrintFunctions(theme ThemeName) PrintFunctions {
	var colors Palette
	switch theme {
	case "catppuccin-mocha", "catppuccin-latte", "catppuccin-frappe", "catppuccin-macchiato":
		colors = generateCatpuccinPalette(theme)
	case "6-colors":
		colors = generate6ColorPalette()
	case "6-colors-high-intensity":
		colors = generate6HighIntensityColorPalette()
	case "random-6-colors":
		colors = generateRandom6colorPalette()
	case "random-6-colors-high-intensity":
		colors = generateRandomHighIntensity6colorPalette()

	default:
		colors = generateGrayscalePalette()

	}
	var functions [16]func(a ...any) string
	for i := range colors {
		functions[i] = colors[i].SprintFunc()
	}
	return functions
}

func generateGrayscalePalette() Palette {
	var palette Palette
	for i := range palette {
		palette[i] = color.RGB(255, 255, 255)
	}
	return palette
}

func generate6ColorPalette() Palette {
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

func generate6HighIntensityColorPalette() Palette {
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

func generateRandom6colorPalette() Palette {
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

func generateRandomHighIntensity6colorPalette() Palette {
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
