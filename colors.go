package main

import (
	catppuccin "github.com/catppuccin/go"
	"github.com/fatih/color"
)

type (
	Palette []*color.Color
)

type (
	StyleName       string
	ThemeName       string
	BorderColorName string
	Dot             string
	Theme           struct {
		name      ThemeName
		colors    Palette
		border    *color.Color
		dot       Dot
		symmetric bool
		random    bool
	}
)

func Color(namedColor catppuccin.Color) *color.Color {
	r, g, b, _ := namedColor.RGBA()
	R, G, B := int(r), int(g), int(b)
	color := color.RGB(R, G, B)
	return color
}

func IsCatppuccin(theme ThemeName) (bool, catppuccin.Flavor) {
	switch theme {
	case "catppuccin-mocha", "catppuccin-mocha-asymmetric":
		return true, catppuccin.Mocha
	case "catppuccin-frappe", "catppuccin-frappe-asymmetric":
		return true, catppuccin.Frappe
	case "catppuccin-latte", "catppuccin-latte-asymmetric":
		return true, catppuccin.Latte
	case "catppuccin-macchiato", "catppuccin-macchiato-asymmetric":
		return true, catppuccin.Macchiato
	default:
		return false, nil
	}
}

func GeneratePalette(theme ThemeName) Palette {
	var colors Palette
	switch theme {
	case "catppuccin-mocha", "catppuccin-latte", "catppuccin-frappe", "catppuccin-macchiato":
		colors = catpuccinPalette(theme)
	case "classic":
		colors = sixColorPalette()
	case "high-intensity":
		colors = sixHighIntensityColorPalette()

	default:
		colors = grayscalePalette()

	}
	return colors
}

func GenerateTheme(config Config) Theme {
	var theme Theme
	theme.name = config.Theme
	// theme.colors = GeneratePalette(config.Theme)
	rawPalette := GeneratePalette(config.Theme)
	if config.Random {
		rawPalette = Randomise(rawPalette)
	}
	if config.Symmetric {
		rawPalette = Mirror(rawPalette)
	}
	theme.colors = wrap(len(config.Printables)*2, rawPalette)
	switch config.Border {
	case "white":
		theme.border = color.RGB(255, 255, 255)
	case "none":
		theme.border = color.RGB(0, 0, 0)
	case "theme":
		if isCatpuccin, flavour := IsCatppuccin(theme.name); isCatpuccin {
			theme.border = Color(flavour.Base())
		} else {
			switch config.Theme {
			case "grayscale":
				theme.border = color.RGB(255, 255, 255)
			default:
				theme.border = color.New(color.FgWhite)
			}
		}
	}

	theme.dot = config.Dot
	return theme
}

func grayscalePalette() Palette {
	var palette Palette
	palette = append(palette, color.RGB(255, 255, 255))
	return palette
}

func sixColorPalette() Palette {
	var palette Palette
	palette = append(palette, color.New(color.FgBlue))
	palette = append(palette, color.New(color.FgRed))
	palette = append(palette, color.New(color.FgYellow))
	palette = append(palette, color.New(color.FgCyan))
	palette = append(palette, color.New(color.FgGreen))
	palette = append(palette, color.New(color.FgMagenta))
	return palette
}

func sixHighIntensityColorPalette() Palette {
	var palette Palette
	palette = append(palette, color.New(color.FgHiBlue))
	palette = append(palette, color.New(color.FgHiRed))
	palette = append(palette, color.New(color.FgHiYellow))
	palette = append(palette, color.New(color.FgHiCyan))
	palette = append(palette, color.New(color.FgHiGreen))
	palette = append(palette, color.New(color.FgHiMagenta))
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
	palette = append(palette, Color(flavour.Sapphire()))
	palette = append(palette, Color(flavour.Lavender()))
	palette = append(palette, Color(flavour.Teal()))
	palette = append(palette, Color(flavour.Green()))
	palette = append(palette, Color(flavour.Pink()))
	palette = append(palette, Color(flavour.Rosewater()))
	palette = append(palette, Color(flavour.Mauve()))
	palette = append(palette, Color(flavour.Sky()))
	palette = append(palette, Color(flavour.Peach()))
	palette = append(palette, Color(flavour.Blue()))
	palette = append(palette, Color(flavour.Yellow()))
	palette = append(palette, Color(flavour.Red()))
	palette = append(palette, Color(flavour.Maroon()))
	return palette
}
