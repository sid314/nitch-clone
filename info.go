package main

import (
	"log"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
)

type (
	Hostname   string
	Username   string
	Distroname string
	Kernel     string
	Shell      string
	Uptime     string
	Memory     int
	Packages   int
)

type Info struct {
	hostname    Hostname
	username    Username
	distro      Distroname
	kernel      Kernel
	shell       Shell
	uptime      Uptime
	totalMemory Memory
	usedMemory  Memory
	packages    Packages
}

func GetInfo() Info {
	info := Info{
		GetHostName(),
		GetUserName(),
		GetDistro(),
		GetKernel(),
		GetShell(), GetUptime(),
		GetTotalMemory(),
		GetUsedMemory(),
		GetPackages(),
	}
	return info
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

func GetDistro() Distroname {
	osReleaseBytes, error := os.ReadFile("/etc/os-release")

	if error != nil {
		log.Fatal(error)
	}

	osRelease := string(osReleaseBytes)
	return Distroname(SnipSnip("PRETTY_NAME=\"", "\"", osRelease))
}

func GetKernel() Kernel {
	kernelBytes, error := exec.Command("uname", "-r").Output()
	if error != nil {
		log.Fatal(error)
	}
	kernel := string(kernelBytes)
	kernel = strings.TrimSpace(kernel)
	return Kernel(kernel)
}

func GetUptime() Uptime {
	uptimeBytes, error := exec.Command("uptime", "-p").Output()
	if error != nil {
		log.Fatal(error)
	}

	uptime := string(uptimeBytes)
	uptime = strings.ReplaceAll(uptime, "minutes", "m")
	uptime = strings.ReplaceAll(uptime, "hours", "h")
	uptime = strings.ReplaceAll(uptime, "days", "d")
	uptime = strings.ReplaceAll(uptime, "minute", "m")
	uptime = strings.ReplaceAll(uptime, "hour", "h")
	uptime = strings.ReplaceAll(uptime, "day", "d")
	uptime = strings.TrimSpace(uptime)
	return Uptime(uptime)
}

func GetShell() Shell {
	shellPieces := strings.SplitAfter(os.Getenv("SHELL"), "/")
	return Shell(shellPieces[len(shellPieces)-1])
}

func getRawTotalMemory() Memory {
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

func getRawFreeMemory() Memory {
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
	return Memory(getRawTotalMemory() / 1024)
}

func GetPackages() Packages {
	packagesBytes, error := exec.Command("pacman", "-Q").Output()
	if error != nil {
		log.Fatal(error)
	}

	packages := string(packagesBytes)
	lines := strings.Count(packages, "\n")
	return Packages(lines)
}

func GetUsedMemory() Memory {
	rawFreeMemory := getRawTotalMemory() - getRawFreeMemory()
	return Memory(rawFreeMemory / 1024)
}
