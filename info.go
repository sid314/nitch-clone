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
	host     string
	username string
	distro   string
	kernel   string
	shell    string
	uptime   string
	desktop  string
	term     string
	memory   int
	pkgs     int
)

func getCurrentDesktop() desktop {
	return desktop(os.Getenv("XDG_CURRENT_DESKTOP"))
}

func getTerminal() term {
	// First tries to read $TERM_PROGRAM
	// If that returns an emty string,
	// then tries to read $TERM
	// This is very inconsistent and does not work for all terminals
	terminal := os.Getenv("TERM_PROGRAM")
	terminal = strings.TrimSpace(terminal)
	if terminal == "" {
		return term(os.Getenv("TERM"))
	}
	return term(terminal)
}

func getHostName() host {
	hostname, error := os.Hostname()
	if error != nil {
		log.Fatal(error)
	}
	return host(hostname)
}

func getUserName() username {
	user, error := user.Current()
	if error != nil {
		log.Fatal(error)
	}

	return username(user.Username)
}

// linux-specific
func getDistro() distro {
	osReleaseBytes, error := os.ReadFile("/etc/os-release")

	if error != nil {
		log.Fatal(error)
	}

	osRelease := string(osReleaseBytes)
	return distro(snipSnip("PRETTY_NAME=\"", "\"", osRelease))
}

func getKernel() kernel {
	u := unix.Utsname{}
	err := unix.Uname(&u)
	if err != nil {
		log.Fatal(err)
	}
	var buffer bytes.Buffer
	buffer.Write(u.Release[:])
	return kernel(u.Release[:bytes.IndexByte(u.Release[:], 0)])
}

func getUptime() uptime {
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
	return uptime(builder.String())
}

func getShell() shell {
	shellPieces := strings.SplitAfter(os.Getenv("SHELL"), "/")
	return shell(shellPieces[len(shellPieces)-1])
}

func getRawTotalMemory() memory {
	memInfoBytes, error := os.ReadFile("/proc/meminfo")
	if error != nil {
		log.Fatal(error)
	}

	meminfostring := string(memInfoBytes)
	totalMemoryString := snipSnip("MemTotal:", " kB", meminfostring)
	totalMemoryString = strings.TrimSpace(totalMemoryString)
	totalRawMemory, error := strconv.Atoi(totalMemoryString)
	if error != nil {
		log.Fatal(error)
	}

	return memory(totalRawMemory)
}

func getRawFreeMemory() memory {
	memInfoBytes, error := os.ReadFile("/proc/meminfo")
	if error != nil {
		log.Fatal(error)
	}

	memInfoString := string(memInfoBytes)
	totalmemorystring := snipSnip("MemAvailable:", " kB", memInfoString)
	totalmemorystring = strings.TrimSpace(totalmemorystring)
	totalRawMemory, error := strconv.Atoi(totalmemorystring)
	if error != nil {
		log.Fatal(error)
	}

	return memory(totalRawMemory)
}

func getTotalMemory() memory {
	return memory(getRawTotalMemory() / 1024)
}

// To add a new package manager first add
// its name to packageManagers
// then add a case for it and add
// whatever command it uses to list all installed
// packages ON SEPARATE LINES
func getPackages() pkgs {
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
	return pkgs(lines)
}

func getUsedMemory() memory {
	rawFreeMemory := getRawTotalMemory() - getRawFreeMemory()
	return memory(rawFreeMemory / 1024)
}
