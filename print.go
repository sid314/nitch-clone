package main

// By default palettes contain 16 colors. If a theme has any less, the remaining are set to white
type PrintableInfo struct {
	Field string
	Value string
}

func largestFieldLength(disableColors bool, printables []PrintableInfo) int {
	largestFieldLength := 0
	if !disableColors {
		largestFieldLength = 10
	}
	for _, printable := range printables {
		if l := len(printable.Field); l > largestFieldLength {
			largestFieldLength = l
		}
	}
	return largestFieldLength - 2
}

func Print() {
	config := GetConfig()
	theme := GenerateTheme(config)
	// info := GetInfo()
	printables := config.Printables
	disableColors := config.DisableColors
	length := largestFieldLength(config.DisableColors, printables)
	dot := theme.dot
	theme.border.Printf("╭")
	for range length {
		theme.border.Printf("─")
	}
	theme.border.Printf("╮\n")
	for i, j := 0, 0; i < len(printables); i++ {

		theme.border.Printf("│ ")
		theme.colors[j].Print(printables[i].Field)
		theme.border.Printf("│ ")
		theme.colors[j+1].Printf("%s\n", printables[i].Value)
		j += 2
	}
	if !disableColors {
		theme.border.Printf("├")
		for range length {
			theme.border.Printf("─")
		}
		theme.border.Printf("┤ \n")
		theme.border.Printf("│ ")
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
	theme.border.Printf("╰")
	for range length {
		theme.border.Printf("─")
	}
	theme.border.Printf("╯\n")
}
