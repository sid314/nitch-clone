package main

import (
	"errors"
	"log"
	"os"

	"github.com/adrg/xdg"
	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Theme ThemeName
	Style StyleName
}

var defaultConfig = Config{"6-colors", "nitch"}

func GetConfig() Config {
	configPath := xdg.ConfigHome + "/nitch-clone/config.toml"
	configFile, err := os.ReadFile(configPath)
	if errors.Is(err, os.ErrNotExist) {
		return defaultConfig
	} else if err != nil {
		return defaultConfig
	} else {
		return parseConfig(configFile)
	}
}

func parseConfig(in []byte) Config {
	var v struct {
		Theme string
		Style string
	}
	if err := toml.Unmarshal(in, &v); err != nil {
		log.Fatal(err)
	}
	config := Config{ThemeName(v.Theme), StyleName(v.Style)}
	if validateConfig(config) {
		return config
	} else {
		return defaultConfig
	}
}

func validateConfig(config Config) bool {
	var styleIsValid bool
	var themeIsValid bool
	switch config.Style {
	case "classic", "nitch":
		styleIsValid = true
	default:
		styleIsValid = false
	}
	switch config.Theme {
	case
		"catppuccin-mocha",
		"catppuccin-frappe",
		"catppuccin-latte",
		"catppuccin-macchiato",
		"grayscale", "6-colors",
		"6-colors-high-intensity",
		"random-6-colors",
		"random-6-colors-high-intensity":
		themeIsValid = true
	default:
		themeIsValid = false
	}
	return styleIsValid && themeIsValid
}
