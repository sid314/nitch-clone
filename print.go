package main

import "fmt"

func SimplePrint() {
	info := GetInfo()
	// dot := ""
	fmt.Printf("╭───────────╮\n")
	fmt.Printf("│   user   │ %s \n", info.username)
	fmt.Printf("│   host   │ %s \n", info.hostname)
	fmt.Printf("│   distro │ %s \n", info.distro)
	fmt.Printf("│ 󰌢  kernel │ %s \n", info.kernel)
	fmt.Printf("│   uptime │ %s \n", info.uptime)
	fmt.Printf("│   shell  │ %s \n", info.shell)
	fmt.Printf("│ 󰏖  pkgs   │ %d \n", info.packages)
	fmt.Printf("│ 󰍛  memory │ %d | %d \n", info.usedMemory, info.totalMemory)
	fmt.Printf("├───────────┤ \n")
	fmt.Printf("│ 󰏘  colors │         \n")
	fmt.Printf("╰───────────╯\n")
	// fmt.Println("│", teal, " ", stop, "hname  │", colornames[1], hostname, stop)
	// fmt.Println("│", green, " ", stop, "distro │", colornames[2], distroname, stop)
	// fmt.Println("│", green, "󰌢 ", stop, "kernel │", colornames[3], kernel, stop)
	// fmt.Println("│", blue, " ", stop, "uptime │", colornames[4], uptime, stop)
	// fmt.Println("│", blue, " ", stop, "shell  │", colornames[5], shell, stop)
	// fmt.Println("│", red, "󰏖 ", stop, "pkgs   │", colornames[6], packages, stop)
	// fmt.Println("│", red, "󰍛 ", stop, "mem    │", colornames[7], usedmem, "|", totalmem, "MiB", stop)
	// fmt.Println("├───────────┤")
	// fmt.Println("│ 󰏘  colors │", grey, dot, red, dot, yellow, dot, green, dot, teal, dot, blue, dot, pink, dot, black, dot, stop)
	// fmt.Println("╰───────────╯")
}
