package main

import (
	"errors"
	"log"
	"os"

	"github.com/adrg/xdg"
	"github.com/goccy/go-yaml"
)

type Config struct {
	theme string
	style string
}

var defaultConfig = Config{"grayscale", "nitch"}

func GetConfig() Config {
	configPath := xdg.ConfigHome + "/nitch-clone/config.yml"
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
		theme string `yaml:"theme"`
		style string `yaml:"style"`
	}
	if err := yaml.Unmarshal(in, &v); err != nil {
		log.Fatal(err)
	}
	config := Config{v.theme, v.style}
	if validateConfig(config) {
		return config
	} else {
		return defaultConfig
	}
}

func validateConfig(config Config) bool {
	var styleIsValid bool
	var themeIsValid bool
	switch config.style {
	case "classic", "nitch":
		styleIsValid = true
	default:
		styleIsValid = false
	}
	switch config.theme {
	case "catppuccin-mocha", "catppuccin-frappe", "catppuccin-latte", "catppuccin-macchiato", "grayscale":
		themeIsValid = true
	default:
		themeIsValid = false
	}
	return styleIsValid && themeIsValid
}
