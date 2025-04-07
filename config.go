package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/adrg/xdg"
	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Theme         ThemeName
	Border        BorderColorName
	Dot           Dot
	Printables    []PrintableInfo
	DisableColors bool
}
type RawConfig struct {
	Theme         string
	Border        string
	Dot           string
	Fields        []string
	DisableColors bool
}

func GetConfig() Config {
	var fields []PrintableInfo
	config := Config{
		"6-colors",
		"white",
		" ",
		fields,
		false,
	}
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
		if valid, border := ValidBorder(rawconfig.Border); valid {
			config.Border = border
		}
		if valid, dot := ValidDot(rawconfig.Dot); valid {
			config.Dot = dot
		}
		config.Printables = SetValidPrintables(rawconfig.Fields)
		config.DisableColors = rawconfig.DisableColors
		return config
	}
}

func SetValidPrintables(fields []string) []PrintableInfo {
	var printables []PrintableInfo
	for _, field := range fields {
		switch field {
		case "user":
			printables = append(printables, PrintableInfo{"  " + field + "   ", string(GetUserName())})
		case "host":
			printables = append(printables, PrintableInfo{"  " + field + "   ", string(GetHostName())})
		case "distro":
			printables = append(printables, PrintableInfo{"  " + field + " ", string(GetDistro())})
		case "kernel":
			printables = append(printables, PrintableInfo{"  " + field + " ", string(GetKernel())})
		case "uptime":
			printables = append(printables, PrintableInfo{"  " + field + " ", string(GetUptime())})
		case "shell":
			printables = append(printables, PrintableInfo{"  " + field + "  ", string(GetShell())})
		case "pkgs":
			printables = append(printables, PrintableInfo{"󰏖  " + field + "   ", strconv.Itoa(int(GetPackages()))})
		case "memory":
			memoryString := fmt.Sprintf("%d | %d MiB", GetUsedMemory(), GetTotalMemory())
			printables = append(printables, PrintableInfo{"󰍛  " + field + " ", memoryString})

		}
	}
	return printables
}

func parseConfig(in []byte) RawConfig {
	var v RawConfig
	if err := toml.Unmarshal(in, &v); err != nil {
		log.Fatal(err)
	}
	// for i := range v.Fields {
	// 	println(v.Fields[i])
	// }
	return RawConfig{v.Theme, v.Border, v.Dot, v.Fields, v.DisableColors}
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
