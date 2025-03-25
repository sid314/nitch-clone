package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/benhoyt/goawk/interp"
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
	osrelease := GetOS()

	// osreleasebyte, _ := os.ReadFile("/etc/os-release")
	// osrelease := string(osreleasebyte)
	getPrettyNameRegex := regexp.MustCompile(`(?m)^PRETTY_NAME=\"?([^\"]*?)\"?$`)
	distroname := getPrettyNameRegex.FindStringSubmatch(osrelease)
	kernelbytes, _ := exec.Command("uname", "-r").Output()
	kernel := string(kernelbytes)
	uptimebytes, _ := exec.Command("uptime", "-p").Output()
	uptime := string(uptimebytes)
	uptime = strings.ReplaceAll(uptime, "minutes", "m")
	uptime = strings.ReplaceAll(uptime, "hours", "h")
	uptime = strings.ReplaceAll(uptime, "days", "d")
	uptime = strings.ReplaceAll(uptime, "minute", "m")
	uptime = strings.ReplaceAll(uptime, "hour", "h")
	uptime = strings.ReplaceAll(uptime, "day", "d")
	shellreader := strings.NewReader(os.Getenv("SHELL"))
	var builder strings.Builder
	err := interp.Exec("{print$NF}", "/", shellreader, &builder)
	if err != nil {
		fmt.Println(err.Error())
	}
	shell := builder.String()
	builder.Reset()
	// echo $SHELL |awk -F "/" '{print$NF}'
	packagesbytes, _ := exec.Command("sh", "-c", " pacman -Q|wc -l").Output()
	packages := string(packagesbytes)
	totalmembytesinkilobytesbytes, _ := exec.Command("sh", "-c", "awk '/MemTotal/ {print $2}' /proc/meminfo").Output()
	totalmembytesinkilobytes := string(totalmembytesinkilobytesbytes)
	totalmembytesinkilobytes = strings.TrimSuffix(totalmembytesinkilobytes, "\n")
	totalmeminkilobytes, _ := strconv.Atoi(totalmembytesinkilobytes)
	totalmeminmebibytes := totalmeminkilobytes / 1024
	freemembytesinkilobytesbytes, _ := exec.Command("sh", "-c", "awk '/MemAvailable/ {print $2}' /proc/meminfo").Output()
	freemembytesinkilobytes := string(freemembytesinkilobytesbytes)
	freemembytesinkilobytes = strings.TrimSuffix(freemembytesinkilobytes, "\n")
	freememinkilobytes, _ := strconv.Atoi(freemembytesinkilobytes)
	freememinmebibytes := freememinkilobytes / 1024
	usedmeminmebibytes := totalmeminmebibytes - freememinmebibytes
	dot := ""
	fmt.Println("╭───────────╮")
	fmt.Println("│", teal, " ", stop, "user   │", red, username, stop)
	fmt.Println("│", teal, " ", stop, "hname  │", yellow, hostname, stop)
	fmt.Println("│", green, " ", stop, "distro │", green, distroname[1], stop)
	fmt.Print("│ ", green, " 󰌢 ", stop, "  kernel │  ", blue, kernel, stop)
	fmt.Print("│ ", blue, "  ", stop, "  uptime │  ", red, uptime, stop)
	fmt.Print("│  ", blue, "  ", stop, " shell  │", "  ", yellow, shell, stop)
	fmt.Print("│ ", red, " 󰏖 ", stop, "  pkgs   │  ", green, packages, stop)
	fmt.Println("│", red, "󰍛 ", stop, "mem    │", blue, usedmeminmebibytes, "|", totalmeminmebibytes, "MiB", stop)
	fmt.Println("├───────────┤")
	fmt.Println("│ 󰏘  colors │", grey, dot, red, dot, yellow, dot, green, dot, teal, dot, blue, dot, pink, dot, black, dot, stop)
	fmt.Println("╰───────────╯")
}
