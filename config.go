package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/adrg/xdg"
	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/pflag"
)

type config struct {
	Theme  themeName
	Border borderColorName
	Dot    symbol
	// Printable info means field name, icon and values
	// Printables    []PrintableInfo
	Printables    printables
	DisableColors bool
	Slow          bool
	Symmetric     bool
	Random        bool
}

// This will be read directly from the config
type rawConfig struct {
	Theme         string
	Border        string
	Dot           string
	Fields        []string
	DisableColors bool
	Slow          bool
	Symmetric     bool
	Random        bool
}

var (
	defaultFields = []string{
		"user",
		"kernel",
		"host",
		"shell",
		"uptime",
		"pkgs",
		"distro",
	}
	validThemes = []string{
		"classic",
		"high-intensity",
		"catppuccin-mocha",
		"catppuccin-macchiato",
		"catppuccin-frappe",
		"catppuccin-latte",
		"kanagawa-dragon",
		"kanagawa-wave",
		"kanagawa-lotus",
		"grayscale",
	}
	validBorders = []string{
		"none",
		"theme",
		"white",
	}
	defaultConfig = config{
		Theme:         "classic",
		Border:        "white",
		Dot:           "",
		Printables:    setValidPrintables(defaultFields),
		DisableColors: false,
		Slow:          false,
		Symmetric:     true,
		Random:        true,
	}
)

func getConfig() *config {
	configPath := xdg.ConfigHome + "/nitch-clone/config.toml"
	configFile, err := os.ReadFile(configPath)
	if errors.Is(err, os.ErrNotExist) {
		return setOverrides(&defaultConfig, parseFlags())
	} else if err != nil {
		return &defaultConfig
	} else {
		return setOverrides(ParseConfig([]byte(configFile)), parseFlags())
	}
}

func setOverrides(file, flags *config) *config {
	if pflag.CommandLine.Changed("slow") {
		file.Slow = flags.Slow
	}
	if pflag.CommandLine.Changed("random") {
		file.Random = flags.Random
	}
	if pflag.CommandLine.Changed("symmetric") {
		file.Symmetric = flags.Symmetric
	}
	if pflag.CommandLine.Changed("disableColors") {
		file.DisableColors = flags.DisableColors
	}
	if pflag.CommandLine.Changed("theme") {
		file.Theme = flags.Theme
	}
	if pflag.CommandLine.Changed("border") {
		file.Border = flags.Border
	}
	if pflag.CommandLine.Changed("fields") {
		file.Printables = flags.Printables
	}
	if pflag.CommandLine.Changed("dot") {
		file.Dot = flags.Dot
	}
	return file
}

func parseFlags() *config {
	config := defaultConfig
	slow := pflag.BoolP("slow", "s", config.Slow, "print slowly")
	pflag.Lookup("slow").NoOptDefVal = "true"
	random := pflag.BoolP("random", "r", config.Random, "randomise colors")
	pflag.Lookup("random").NoOptDefVal = "true"
	symmetric := pflag.BoolP("symmetric", "S", config.Symmetric, "print fields and values in the same color")
	pflag.Lookup("symmetric").NoOptDefVal = "true"
	disableColors := pflag.BoolP("disableColors", "d", config.DisableColors, "disable last line")
	pflag.Lookup("disableColors").NoOptDefVal = "true"
	dotf := pflag.StringP("dot", "D", string(config.Dot), "symbol printed on the last line")
	themef := pflag.StringP("theme", "t", string(config.Theme), "theme")
	borderf := pflag.StringP("border", "b", string(config.Border), "border color")
	fields := pflag.StringSliceP("fields", "f", fieldsFromPrintableInfo(config.Printables), "fields that will be printed")
	pflag.Parse()
	config.Slow = *slow
	config.Random = *random
	config.Symmetric = *symmetric
	config.DisableColors = *disableColors
	if valid, dot := validDot(*dotf); valid {
		config.Dot = dot
	}
	if valid, theme := validTheme(*themef); valid {
		config.Theme = theme
	}
	if valid, border := validBorder(*borderf); valid {
		config.Border = border
	}
	printables := setValidPrintables(*fields)
	if len(printables) != 0 {
		config.Printables = printables
	}
	return &config
}

func setValidPrintables(fields []string) printables {
	// Printables are added in the order they appear in the config
	var printables printables
	for _, field := range fields {
		switch field {
		case "user":
			printables = append(printables, &printableInfo{"  " + field, string(getUserName())})
		case "host":
			printables = append(printables, &printableInfo{"  " + field, string(getHostName())})
		case "distro":
			printables = append(printables, &printableInfo{"  " + field, string(getDistro())})
		case "kernel":
			printables = append(printables, &printableInfo{"  " + field, string(getKernel())})
		case "uptime":
			printables = append(printables, &printableInfo{"  " + field, string(getUptime())})
		case "shell":
			printables = append(printables, &printableInfo{"  " + field, string(getShell())})
		case "de":
			printables = append(printables, &printableInfo{"  " + field, string(getCurrentDesktop())})
		case "term":
			printables = append(printables, &printableInfo{"  " + field, string(getTerminal())})
		case "pkgs":
			printables = append(printables, &printableInfo{"  " + field, strconv.Itoa(int(getPackages()))})
		case "memory":
			memoryString := fmt.Sprintf("%d | %d MiB", getUsedMemory(), getTotalMemory())
			printables = append(printables, &printableInfo{"  " + field, memoryString})

		}
	}
	return printables
}

func ParseConfig(in []byte) *config {
	var v rawConfig
	if err := toml.Unmarshal(in, &v); err != nil {
		log.Fatal(err)
	}
	rawconfig := rawConfig{
		Theme:         v.Theme,
		Border:        v.Border,
		Dot:           v.Dot,
		Fields:        v.Fields,
		DisableColors: v.DisableColors,
		Slow:          v.Slow,
		Symmetric:     v.Symmetric,
		Random:        v.Random,
	}
	config := defaultConfig
	if valid, theme := validTheme(rawconfig.Theme); valid {
		config.Theme = theme
	}
	if valid, border := validBorder(rawconfig.Border); valid {
		config.Border = border
	}
	if valid, dot := validDot(rawconfig.Dot); valid {
		config.Dot = dot
	}
	if len(rawconfig.Fields) != 0 {
		config.Printables = setValidPrintables(rawconfig.Fields)
	}
	config.DisableColors = rawconfig.DisableColors
	config.Slow = rawconfig.Slow
	config.Symmetric = rawconfig.Symmetric
	config.Random = rawconfig.Random
	return &config
}

func validTheme(theme string) (bool, themeName) {
	if slices.Contains(validThemes, theme) {
		return true, themeName(theme)
	}
	return false, ""
}

func validBorder(border string) (bool, borderColorName) {
	if slices.Contains(validBorders, border) {
		return true, borderColorName(border)
	}
	return false, ""
}

func validDot(dot string) (bool, symbol) {
	if strings.TrimSpace(dot) != "" {
		return true, symbol(dot)
	} else {
		return false, ""
	}
}
