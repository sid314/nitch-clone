package main

// By default palettes contain 16 colors. If a theme has any less, the remaining are set to white

func Print() {
	config := GetConfig()
	theme := GenerateTheme(config)
	info := GetInfo()
	switch config.Style {
	case "nitch":
		printNitch(theme, info)
	case "classic":
	}
}

func printNitch(theme Theme, info Info) {
	dot := theme.dot
	theme.border.Printf("╭───────────╮\n")
	theme.border.Printf("│ ")
	theme.colors[0].Printf("  user   ")
	theme.border.Printf("│ ")
	theme.colors[1].Printf("%s\n", info.username)
	theme.border.Printf("│ ")
	theme.colors[2].Printf("  host   ")
	theme.border.Printf("│ ")
	theme.colors[3].Printf("%s\n", info.hostname)
	theme.border.Printf("│ ")
	theme.colors[4].Printf("  distro ")
	theme.border.Printf("│ ")
	theme.colors[5].Printf("%s\n", info.distro)
	theme.border.Printf("│ ")
	theme.colors[6].Printf("󰌢  kernel ")
	theme.border.Printf("│ ")
	theme.colors[7].Printf("%s\n", info.kernel)
	theme.border.Printf("│ ")
	theme.colors[8].Printf("  uptime ")
	theme.border.Printf("│ ")
	theme.colors[9].Printf("%s\n", info.uptime)
	theme.border.Printf("│ ")
	theme.colors[10].Printf("  shell  ")
	theme.border.Printf("│ ")
	theme.colors[11].Printf("%s\n", info.shell)
	theme.border.Printf("│ ")
	theme.colors[12].Printf("󰏖  pkgs   ")
	theme.border.Printf("│ ")
	theme.colors[13].Printf("%d\n", info.packages)
	theme.border.Printf("│ ")
	theme.colors[14].Printf("󰍛  memory ")
	theme.border.Printf("│ ")
	theme.colors[15].Printf("%d | %d MiB \n", info.usedMemory, info.totalMemory)
	theme.border.Printf("├───────────┤ \n")
	theme.border.Printf("│ ")
	theme.colors[0].Printf("󰏘  colors ")
	theme.border.Printf("│ ")
	for i := 0; i < 16; i += 2 {
		theme.colors[i].Printf("%s ", dot)
	}
	theme.border.Printf("\n╰───────────╯\n")
}
