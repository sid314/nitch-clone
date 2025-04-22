package main

import (
	"github.com/fatih/color"
)

type (
	palette []*color.Color
)

type (
	styleName       string
	themeName       string
	borderColorName string
	symbol          string
	theme           struct {
		name      themeName
		colors    palette
		border    *color.Color
		dot       symbol
		symmetric bool
		random    bool
	}
)

func generatePalette(theme themeName) palette {
	var colors palette
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

func generateTheme(config *config) *theme {
	var theme theme
	theme.name = config.Theme
	// theme.colors = GeneratePalette(config.Theme)
	rawPalette := generatePalette(config.Theme)
	if config.Random {
		rawPalette = randomise(rawPalette)
	}
	if config.Symmetric {
		rawPalette = mirror(rawPalette)
	}
	theme.colors = wrap(len(config.Printables)*2, rawPalette)
	switch config.Border {
	case "white":
		theme.border = color.RGB(255, 255, 255)
	case "none":
		theme.border = color.RGB(0, 0, 0)
	case "theme":
		if isCatpuccin, flavour := isCatppuccin(theme.name); isCatpuccin {
			theme.border = catpuccinToColor(flavour.Base())
		} else {
			switch config.Theme {
			case "grayscale":
				theme.border = color.RGB(255, 255, 255)
			case "kanagawa-wave":
				theme.border = hexToColor("#16161d")
			case "kanagawa-lotus":
				theme.border = hexToColor("#1f1f28")
			case "kanagawa-dragon":
				theme.border = hexToColor("#16161d")
			default:
				theme.border = color.New(color.FgWhite)
			}
		}
	}

	theme.dot = config.Dot
	return &theme
}
