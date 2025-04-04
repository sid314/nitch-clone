package main

import (
	"fmt"
)

// By default palettes contain 16 colors. If a theme has any less, the remaining are set to white

func Print() {
	config := GetConfig()
	switch config.Style {
	case "nitch":
		printNitch(config.Theme)
	case "classic":
		// printClassic(config.theme)
	}
}

func printNitch(theme ThemeName) {
	info := GetInfo()
	dot := ""
	printFunctions := GeneratePrintFunctions(theme)
	fmt.Printf("╭───────────╮\n")
	fmt.Printf("│ %s   │ %s \n", printFunctions[0]("  user"), printFunctions[1](info.username))
	fmt.Printf("│ %s   │ %s \n", printFunctions[2]("  host"), printFunctions[3](info.hostname))
	fmt.Printf("│ %s │ %s \n", printFunctions[4]("  distro"), printFunctions[5](info.distro))
	fmt.Printf("│ %s │ %s \n", printFunctions[6]("󰌢  kernel"), printFunctions[7](info.kernel))
	fmt.Printf("│ %s │ %s \n", printFunctions[8]("  uptime"), printFunctions[9](info.uptime))
	fmt.Printf("│ %s  │ %s \n", printFunctions[10]("  shell"), printFunctions[11](info.shell))
	fmt.Printf("│ %s   │ %s \n", printFunctions[12]("󰏖  pkgs"), printFunctions[13](info.packages))
	fmt.Printf("│ %s │ %s | %s \n", printFunctions[14]("󰍛  memory"), printFunctions[15](info.usedMemory), printFunctions[15](info.totalMemory))
	fmt.Printf("├───────────┤ \n")
	fmt.Printf("│ %s │", printFunctions[0]("󰏘  colors"))
	for i := 0; i < 16; i += 2 {
		fmt.Printf(" %s", printFunctions[i](dot))
	}
	fmt.Println("")
	fmt.Printf("╰───────────╯\n")
}
