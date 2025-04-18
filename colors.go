package main

import (
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

func GeneratePalette(theme ThemeName) Palette {
	var colors Palette
	switch theme {
	case "catppuccin-mocha", "catppuccin-latte", "catppuccin-frappe", "catppuccin-macchiato":
		colors = catpuccinPalette(theme)
	case "classic":
		colors = classicPalette()
	case "high-intensity":
		colors = highIntensityPalette()
	case "kanagawa-wave", "kanagawa-lotus", "kanagawa-dragon":
		colors = kanagawaPalette(theme)
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
			theme.border = catpuccinToColor(flavour.Base())
		} else {
			switch config.Theme {
			case "grayscale":
				theme.border = color.RGB(255, 255, 255)
			case "kanagawa-wave":
				theme.border = HexToColor("#16161d")
			case "kanagawa-lotus":
				theme.border = HexToColor("#1f1f28")
			case "kanagawa-dragon":
				theme.border = HexToColor("#0d0c0c")
			default:
				theme.border = color.New(color.FgWhite)
			}
		}
	}

	theme.dot = config.Dot
	return theme
}
