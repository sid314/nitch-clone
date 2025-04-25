package main

import (
	"time"

	"github.com/rivo/uniseg"
)

type printableInfo struct {
	Field string
	Value string
}
type printables []*printableInfo

func print() {
	config := getConfig()
	theme := generateTheme(config)
	printables := config.Printables
	disableColors := config.DisableColors
	length := largestFieldLength(config.DisableColors, printables)
	padPrintables(printables, length)
	length = length + 2
	dot := theme.dot
	var delay time.Duration
	if config.Slow {
		delay = time.Millisecond * 50
	} else {
		delay = 0
	}
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
			theme.colors[j+1].Print(string(value.Bytes()))

			time.Sleep(delay)
		}
		theme.colors[j+1].Println()
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
		colorsLabel := uniseg.NewGraphemes("  colors")
		for colorsLabel.Next() {
			theme.colors[0].Print(string(colorsLabel.Bytes()))
			time.Sleep(delay)
		}
		for range length - 10 {
			theme.border.Print(" ")
		}
		theme.border.Printf("│ ")
		time.Sleep(delay)
		for i := 0; i < len(theme.colors); i += 2 {
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

func padPrintables(printables printables, requiredLength int) {
	for i, printable := range printables {
		for uniseg.StringWidth(printable.Field) < requiredLength+1 {
			printable.Field = printable.Field + " "
		}
		printables[i].Field = printable.Field
	}
}
