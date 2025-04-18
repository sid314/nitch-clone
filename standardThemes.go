package main

import "github.com/fatih/color"

func grayscalePalette() Palette {
	var palette Palette
	palette = append(palette, color.RGB(255, 255, 255))
	return palette
}

func classicPalette() Palette {
	var palette Palette
	palette = append(palette, color.New(color.FgBlue))
	palette = append(palette, color.New(color.FgRed))
	palette = append(palette, color.New(color.FgYellow))
	palette = append(palette, color.New(color.FgCyan))
	palette = append(palette, color.New(color.FgGreen))
	palette = append(palette, color.New(color.FgMagenta))
	return palette
}

func highIntensityPalette() Palette {
	var palette Palette
	palette = append(palette, color.New(color.FgHiBlue))
	palette = append(palette, color.New(color.FgHiRed))
	palette = append(palette, color.New(color.FgHiYellow))
	palette = append(palette, color.New(color.FgHiCyan))
	palette = append(palette, color.New(color.FgHiGreen))
	palette = append(palette, color.New(color.FgHiMagenta))
	return palette
}
