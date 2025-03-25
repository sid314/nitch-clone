package main

import (
	"fmt"
)

func main() {
	const (
		yellow = "\033[93m\b"
		red    = "\033[91m\b"
		green  = "\033[92m\b"
		blue   = "\033[94m\b"
		pink   = "\033[95m\b"
		teal   = "\033[96m\b"
		grey   = "\033[97m\b"
		black  = "\033[30m\b"
		stop   = "\033[0m\b"
	)
	hostname := GetHostName()
	username := GetUserName()
	distroname := GetDistro()
	kernel := GetKernel()
	uptime := GetUptime()
	shell := GetShell()
	packages := GetPackages()
	totalmem := GetTotalMemory()
	usedmem := GetUsedMemory()

	dot := ""
	fmt.Println("╭───────────╮")
	fmt.Println("│", teal, " ", stop, "user   │", red, username, stop)
	fmt.Println("│", teal, " ", stop, "hname  │", yellow, hostname, stop)
	fmt.Println("│", green, " ", stop, "distro │", green, distroname, stop)
	fmt.Print("│ ", green, " 󰌢 ", stop, "  kernel │  ", blue, kernel, stop)
	fmt.Print("│ ", blue, "  ", stop, "  uptime │  ", red, uptime, stop)
	fmt.Println("│", blue, " ", stop, "shell  │", yellow, shell, stop)
	fmt.Println("│", red, "󰏖 ", stop, "pkgs   │", green, packages, stop)
	fmt.Println("│", red, "󰍛 ", stop, "mem    │", blue, usedmem, "|", totalmem, "MiB", stop)
	fmt.Println("├───────────┤")
	fmt.Println("│ 󰏘  colors │", grey, dot, red, dot, yellow, dot, green, dot, teal, dot, blue, dot, pink, dot, black, dot, stop)
	fmt.Println("╰───────────╯")
}
