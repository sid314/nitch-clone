package main

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/adrg/xdg"
	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Theme  ThemeName
	Style  StyleName
	Border BorderColorName
	Dot    Dot
}
type RawConfig struct {
	Theme  string
	Style  string
	Border string
	Dot    string
}

func GetConfig() Config {
	config := Config{"6-colors", "nitch", "white", " "}
	configPath := xdg.ConfigHome + "/nitch-clone/config.toml"
	configFile, err := os.ReadFile(configPath)
	if errors.Is(err, os.ErrNotExist) {
		return config
	} else if err != nil {
		return config
	} else {
		rawconfig := parseConfig(configFile)
		if valid, theme := ValidTheme(rawconfig.Theme); valid {
			config.Theme = theme
		}
		if valid, style := ValidStyle(rawconfig.Style); valid {
			config.Style = style
		}
		if valid, border := ValidBorder(rawconfig.Border); valid {
			config.Border = border
		}
		if valid, dot := ValidDot(rawconfig.Dot); valid {
			config.Dot = dot
		}
		return config
	}
}

func parseConfig(in []byte) RawConfig {
	var v RawConfig
	if err := toml.Unmarshal(in, &v); err != nil {
		log.Fatal(err)
	}
	return RawConfig{v.Theme, v.Style, v.Border, v.Dot}
}

func ValidTheme(theme string) (bool, ThemeName) {
	switch theme {
	case
		"catppuccin-mocha",
		"catppuccin-frappe",
		"catppuccin-latte",
		"catppuccin-macchiato",
		"grayscale", "6-colors",
		"6-colors-high-intensity",
		"random-6-colors",
		"random-6-colors-high-intensity":
		return true, ThemeName(theme)
	default:
		return false, ""

	}
}

func ValidStyle(style string) (bool, StyleName) {
	switch style {
	case "nitch", "classic":
		return true, StyleName(style)
	default:
		return false, ""

	}
}

func ValidBorder(border string) (bool, BorderColorName) {
	switch border {
	case "none", "theme", "white":
		return true, BorderColorName(border)
	default:
		return false, ""

	}
}

func ValidDot(dot string) (bool, Dot) {
	if strings.TrimSpace(dot) != "" {
		return true, Dot(dot)
	} else {
		return false, ""
	}
}
