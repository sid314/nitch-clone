package main

import catppuccin "github.com/catppuccin/go"

func isCatppuccin(theme themeName) (bool, catppuccin.Flavor) {
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

func catpuccinPalette(theme themeName) palette {
	var flavour catppuccin.Flavor
	var palette palette
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
	palette = append(palette, catpuccinToColor(flavour.Sapphire()))
	palette = append(palette, catpuccinToColor(flavour.Lavender()))
	palette = append(palette, catpuccinToColor(flavour.Teal()))
	palette = append(palette, catpuccinToColor(flavour.Green()))
	palette = append(palette, catpuccinToColor(flavour.Pink()))
	palette = append(palette, catpuccinToColor(flavour.Rosewater()))
	palette = append(palette, catpuccinToColor(flavour.Mauve()))
	palette = append(palette, catpuccinToColor(flavour.Sky()))
	palette = append(palette, catpuccinToColor(flavour.Peach()))
	palette = append(palette, catpuccinToColor(flavour.Blue()))
	palette = append(palette, catpuccinToColor(flavour.Yellow()))
	palette = append(palette, catpuccinToColor(flavour.Red()))
	palette = append(palette, catpuccinToColor(flavour.Maroon()))
	return palette
}
