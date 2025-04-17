package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/adrg/xdg"
	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Theme  ThemeName
	Border BorderColorName
	Dot    Dot
	// Printable info means field name, icon and values
	Printables    []PrintableInfo
	DisableColors bool
	Slow          bool
	Symmetric     bool
	Random        bool
}

// This will be read directly from the config
type RawConfig struct {
	Theme         string
	Border        string
	Dot           string
	Fields        []string
	DisableColors bool
	Slow          bool
	Symmetric     bool
	Random        bool
}

func GetConfig() Config {
	var fields []PrintableInfo
	// The default config enables the six-colors theme
	// with white border and dot set to 
	// no fields, disableColors set to false and slow mode
	// set to false
	config := Config{
		Theme:         "6-colors",
		Border:        "white",
		Dot:           " ",
		Printables:    fields,
		DisableColors: false,
		Slow:          false,
		Symmetric:     true,
		Random:        true,
	}
	configPath := xdg.ConfigHome + "/nitch-clone/config.toml"
	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return config
	} else {
		rawconfig := parseConfig(configFile)
		// Whatever config option is defined in the config
		// and is valid get read and overrides the default config
		if valid, theme := ValidTheme(rawconfig.Theme); valid {
			config.Theme = theme
		}
		if valid, border := ValidBorder(rawconfig.Border); valid {
			config.Border = border
		}
		if valid, dot := ValidDot(rawconfig.Dot); valid {
			config.Dot = dot
		}
		config.Printables = SetValidPrintables(rawconfig.Fields)
		config.DisableColors = rawconfig.DisableColors
		config.Slow = rawconfig.Slow
		config.Symmetric = rawconfig.Symmetric
		config.Random = rawconfig.Random
		return config
	}
}

func SetValidPrintables(fields []string) []PrintableInfo {
	// Printables are added in the order they appear in the config
	var printables []PrintableInfo
	for _, field := range fields {
		switch field {
		case "user":
			printables = append(printables, PrintableInfo{"  " + field, string(GetUserName())})
		case "host":
			printables = append(printables, PrintableInfo{"  " + field, string(GetHostName())})
		case "distro":
			printables = append(printables, PrintableInfo{"  " + field, string(GetDistro())})
		case "kernel":
			printables = append(printables, PrintableInfo{"  " + field, string(GetKernel())})
		case "uptime":
			printables = append(printables, PrintableInfo{"  " + field, string(GetUptime())})
		case "shell":
			printables = append(printables, PrintableInfo{"  " + field, string(GetShell())})
		case "de":
			printables = append(printables, PrintableInfo{"  " + field, string(GetCurrentDesktop())})
		case "term":
			printables = append(printables, PrintableInfo{"  " + field, string(GetTerminal())})
		case "pkgs":
			printables = append(printables, PrintableInfo{"󰏖  " + field, strconv.Itoa(int(GetPackages()))})
		case "memory":
			memoryString := fmt.Sprintf("%d | %d MiB", GetUsedMemory(), GetTotalMemory())
			printables = append(printables, PrintableInfo{"󰍛  " + field, memoryString})

		}
	}
	return printables
}

func parseConfig(in []byte) RawConfig {
	var v RawConfig
	if err := toml.Unmarshal(in, &v); err != nil {
		log.Fatal(err)
	}
	return RawConfig{
		Theme:         v.Theme,
		Border:        v.Border,
		Dot:           v.Dot,
		Fields:        v.Fields,
		DisableColors: v.DisableColors,
		Slow:          v.Slow,
		Symmetric:     v.Symmetric,
		Random:        v.Random,
	}
}

func ValidTheme(theme string) (bool, ThemeName) {
	switch theme {
	case
		"catppuccin-mocha",
		"catppuccin-frappe",
		"catppuccin-latte",
		"catppuccin-macchiato",
		"grayscale", "6-colors":
		return true, ThemeName(theme)
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
