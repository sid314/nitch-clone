package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	"time"

	"golang.org/x/sys/unix"
)

type (
	Hostname string
	Username string
	Distro   string
	Kernel   string
	Shell    string
	Uptime   string
	Desktop  string
	Terminal string
	Memory   int
	Packages int
)

func GetCurrentDesktop() Desktop {
	return Desktop(os.Getenv("XDG_CURRENT_DESKTOP"))
}

func GetTerminal() Terminal {
	// First tries to read $TERM_PROGRAM
	// If that returns an emty string,
	// then tries to read $TERM
	// This is very inconsistent and does not work for all terminals
	terminal := os.Getenv("TERM_PROGRAM")
	terminal = strings.TrimSpace(terminal)
	if terminal == "" {
		return Terminal(os.Getenv("TERM"))
	}
	return Terminal(terminal)
}

func GetHostName() Hostname {
	hostname, error := os.Hostname()
	if error != nil {
		log.Fatal(error)
	}
	return Hostname(hostname)
}

func GetUserName() Username {
	user, error := user.Current()
	if error != nil {
		log.Fatal(error)
	}

	return Username(user.Username)
}

// linux-specific
func GetDistro() Distro {
	osReleaseBytes, error := os.ReadFile("/etc/os-release")

	if error != nil {
		log.Fatal(error)
	}

	osRelease := string(osReleaseBytes)
	return Distro(SnipSnip("PRETTY_NAME=\"", "\"", osRelease))
}

func GetKernel() Kernel {
	u := unix.Utsname{}
	err := unix.Uname(&u)
	if err != nil {
		log.Fatal(err)
	}
	var buffer bytes.Buffer
	buffer.Write(u.Release[:])
	return Kernel(u.Release[:bytes.IndexByte(u.Release[:], 0)])
}

func GetUptime() Uptime {
	sysinfo := &unix.Sysinfo_t{}
	if err := unix.Sysinfo(sysinfo); err != nil {
		log.Fatal(err)
	}
	uptimeDuration := time.Duration(sysinfo.Uptime * int64(time.Second))
	seconds := int(uptimeDuration.Seconds())
	days := seconds / 86400
	seconds = seconds % 86400
	hours := seconds / 3600
	seconds = seconds % 3600
	minutes := seconds / 60
	seconds = seconds % 60

	var builder strings.Builder
	if days > 0 {
		builder.WriteString(fmt.Sprintf("%dd ", days))
	}
	if hours > 0 {
		builder.WriteString(fmt.Sprintf("%dh ", hours))
	}
	if minutes > 0 {
		builder.WriteString(fmt.Sprintf("%dm ", minutes))
	}
	if seconds > 0 {
		builder.WriteString(fmt.Sprintf("%ds ", seconds))
	}
	return Uptime(builder.String())
}

func GetShell() Shell {
	shellPieces := strings.SplitAfter(os.Getenv("SHELL"), "/")
	return Shell(shellPieces[len(shellPieces)-1])
}

func GetRawTotalMemory() Memory {
	memInfoBytes, error := os.ReadFile("/proc/meminfo")
	if error != nil {
		log.Fatal(error)
	}

	meminfostring := string(memInfoBytes)
	totalMemoryString := SnipSnip("MemTotal:", " kB", meminfostring)
	totalMemoryString = strings.TrimSpace(totalMemoryString)
	totalRawMemory, error := strconv.Atoi(totalMemoryString)
	if error != nil {
		log.Fatal(error)
	}

	return Memory(totalRawMemory)
}

func GetRawFreeMemory() Memory {
	memInfoBytes, error := os.ReadFile("/proc/meminfo")
	if error != nil {
		log.Fatal(error)
	}

	memInfoString := string(memInfoBytes)
	totalmemorystring := SnipSnip("MemAvailable:", " kB", memInfoString)
	totalmemorystring = strings.TrimSpace(totalmemorystring)
	totalRawMemory, error := strconv.Atoi(totalmemorystring)
	if error != nil {
		log.Fatal(error)
	}

	return Memory(totalRawMemory)
}

func GetTotalMemory() Memory {
	return Memory(GetRawTotalMemory() / 1024)
}

// To add a new package manager first add
// its name to packageManagers
// then add a case for it and add
// whatever command it uses to list all installed
// packages ON SEPARATE LINES
func GetPackages() Packages {
	var command *exec.Cmd
	packageManagers := []string{"pacman", "dnf", "rpm", "apt"}
	var packageManager string
	for i := range packageManagers {
		maybePackageManager, _ := exec.Command("which", packageManagers[i]).Output()
		if string(maybePackageManager) != "" {
			packageManager = packageManagers[i]
		}

	}
	switch packageManager {
	case "pacman":
		command = exec.Command("pacman", "-Q")
	case "apt":
		command = exec.Command("apt", "list", "--installed")
	case "dnf":
		command = exec.Command("dnf", "list", "--installed")
	case "rpm":
		command = exec.Command("rpm", "-qa")
	default:
		command = exec.Command("echo", "boo")
	}
	packagesBytes, error := command.Output()
	if error != nil {
		log.Fatal(error)
	}

	packages := string(packagesBytes)
	lines := strings.Count(packages, "\n")
	return Packages(lines)
}

func GetUsedMemory() Memory {
	rawFreeMemory := GetRawTotalMemory() - GetRawFreeMemory()
	return Memory(rawFreeMemory / 1024)
}
