package main

import (
	"time"

	"github.com/rivo/uniseg"
)

type PrintableInfo struct {
	Field string
	Value string
}

func largestFieldLength(disableColors bool, printables []PrintableInfo) int {
	largestFieldLength := 0
	if !disableColors {
		largestFieldLength = 8
	}
	for _, printable := range printables {
		if l := uniseg.StringWidth(printable.Field); l > largestFieldLength {
			largestFieldLength = l
		}
	}
	return largestFieldLength
}

func paddedPrintables(printables []PrintableInfo, requiredLength int) []PrintableInfo {
	for i, printable := range printables {
		for uniseg.StringWidth(printable.Field) < requiredLength+1 {
			printable.Field = printable.Field + " "
		}
		printables[i].Field = printable.Field
	}
	return printables
}

func fastPrint() {
	config := GetConfig()
	theme := GenerateTheme(config)
	printables := config.Printables
	disableColors := config.DisableColors
	length := largestFieldLength(config.DisableColors, printables)
	printables = paddedPrintables(printables, length)
	length = length + 2
	dot := theme.dot
	theme.border.Printf("  ╭")
	for range length {
		theme.border.Printf("─")
	}
	theme.border.Printf("╮\n")
	for i, j := 0, 0; i < len(printables); i++ {
		theme.border.Printf("  │ ")
		theme.colors[j].Print(printables[i].Field)
		theme.border.Printf("│ ")
		theme.colors[j+1].Printf("%s\n", printables[i].Value)
		if j >= 14 {
			j = -2
		}
		j += 2
	}
	if !disableColors {
		theme.border.Printf("  ├")
		for range length {
			theme.border.Printf("─")
		}
		theme.border.Printf("┤ \n")
		theme.border.Printf("  │ ")
		theme.colors[0].Printf("󰏘 colors")
		for range length - 9 {
			theme.border.Print(" ")
		}
		theme.border.Printf("│ ")
		for i := 0; i < 16; i += 2 {
			theme.colors[i].Printf("%s ", dot)
		}
		theme.border.Print("\n")
	}
	theme.border.Printf("  ╰")
	for range length {
		theme.border.Printf("─")
	}
	theme.border.Printf("╯\n")
}

func slowPrint() {
	config := GetConfig()
	theme := GenerateTheme(config)
	printables := config.Printables
	disableColors := config.DisableColors
	length := largestFieldLength(config.DisableColors, printables)
	printables = paddedPrintables(printables, length)
	length = length + 2
	dot := theme.dot
	delay := time.Millisecond * 50
	theme.border.Printf("  ╭")
	time.Sleep(delay)
	for range length {
		theme.border.Printf("─")
		time.Sleep(delay)
	}
	theme.border.Printf("╮\n")
	time.Sleep(delay)
	for i, j := 0, 0; i < len(printables); i++ {
		theme.border.Printf("  │ ")
		time.Sleep(delay)
		// theme.colors[j].Print(printables[i].Field)
		field := uniseg.NewGraphemes(printables[i].Field)
		for field.Next() {
			theme.colors[j].Print(string(field.Bytes()))
			time.Sleep(delay)
		}
		time.Sleep(delay)
		theme.border.Printf("│ ")
		time.Sleep(delay)
		value := uniseg.NewGraphemes(printables[i].Value)
		for value.Next() {
			theme.colors[j].Print(string(value.Bytes()))

			time.Sleep(delay)
		}
		theme.colors[j+1].Println()
		// time.Sleep(delay)
		if j >= 14 {
			j = -2
		}
		j += 2
	}
	if !disableColors {
		theme.border.Printf("  ├")
		time.Sleep(delay)
		for range length {
			theme.border.Printf("─")
			time.Sleep(delay)
		}
		theme.border.Printf("┤ \n")
		time.Sleep(delay)
		time.Sleep(delay)
		theme.border.Printf("  │ ")
		time.Sleep(delay)
		// theme.colors[0].Printf("󰏘 colors")
		colorsLabel := uniseg.NewGraphemes("󰏘 colors")
		for colorsLabel.Next() {
			theme.colors[0].Print(string(colorsLabel.Bytes()))
			time.Sleep(delay)
		}
		for range length - 9 {
			theme.border.Print(" ")
		}
		theme.border.Printf("│ ")
		time.Sleep(delay)
		for i := 0; i < 16; i += 2 {
			theme.colors[i].Printf("%s ", dot)
			time.Sleep(delay)
		}
		theme.border.Print("\n")
		time.Sleep(delay)
	}
	theme.border.Printf("  ╰")
	time.Sleep(delay)
	for range length {
		theme.border.Printf("─")
		time.Sleep(delay)
	}
	theme.border.Printf("╯\n")
}
