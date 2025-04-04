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
	Hostname string
	Username string
	Distro   string
	Kernel   string
	Shell    string
	Uptime   string
	Memory   int
	Packages int
)

type Info struct {
	hostname    Hostname
	username    Username
	distro      Distro
	kernel      Kernel
	shell       Shell
	uptime      Uptime
	totalMemory Memory
	usedMemory  Memory
	packages    Packages
}

func GetInfo() Info {
	info := Info{
		getHostName(),
		getUserName(),
		getDistro(),
		getKernel(),
		getShell(), getUptime(),
		getTotalMemory(),
		getUsedMemory(),
		getPackages(),
	}
	return info
}

func getHostName() Hostname {
	hostname, error := os.Hostname()
	if error != nil {
		log.Fatal(error)
	}
	return Hostname(hostname)
}

func getUserName() Username {
	user, error := user.Current()
	if error != nil {
		log.Fatal(error)
	}

	return Username(user.Username)
}

func getDistro() Distro {
	osReleaseBytes, error := os.ReadFile("/etc/os-release")

	if error != nil {
		log.Fatal(error)
	}

	osRelease := string(osReleaseBytes)
	return Distro(SnipSnip("PRETTY_NAME=\"", "\"", osRelease))
}

func getKernel() Kernel {
	kernelBytes, error := exec.Command("uname", "-r").Output()
	if error != nil {
		log.Fatal(error)
	}
	kernel := string(kernelBytes)
	kernel = strings.TrimSpace(kernel)
	return Kernel(kernel)
}

func getUptime() Uptime {
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

func getShell() Shell {
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

func getTotalMemory() Memory {
	return Memory(getRawTotalMemory() / 1024)
}

func getPackages() Packages {
	packagesBytes, error := exec.Command("pacman", "-Q").Output()
	if error != nil {
		log.Fatal(error)
	}

	packages := string(packagesBytes)
	lines := strings.Count(packages, "\n")
	return Packages(lines)
}

func getUsedMemory() Memory {
	rawFreeMemory := getRawTotalMemory() - getRawFreeMemory()
	return Memory(rawFreeMemory / 1024)
}
