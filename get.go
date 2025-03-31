package main

import (
	"log"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
)

func GetHostName() string {
	hostname, error := os.Hostname()
	if error != nil {
		log.Fatal(error)
	}

	return hostname
}

func GetUserName() string {
	user, error := user.Current()
	if error != nil {
		log.Fatal(error)
	}

	return user.Username
}

func GetDistro() string {
	osReleaseBytes, error := os.ReadFile("/etc/os-release")

	if error != nil {
		log.Fatal(error)
	}

	osRelease := string(osReleaseBytes)
	return SnipSnip("PRETTY_NAME=\"", "\"", osRelease)
}

func GetKernel() string {
	kernelBytes, error := exec.Command("uname", "-r").Output()
	if error != nil {
		log.Fatal(error)
	}

	return string(kernelBytes)
}

func GetUptime() string {
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
	return uptime
}

func GetShell() string {
	shellPieces := strings.SplitAfter(os.Getenv("SHELL"), "/")
	return shellPieces[len(shellPieces)-1]
}

func getRawTotalMemory() int {
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

	return totalRawMemory
}

func getRawFreeMemory() int {
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

	return totalRawMemory
}

func GetTotalMemory() int {
	return getRawTotalMemory() / 1024
}

func GetPackages() int {
	packagesBytes, error := exec.Command("pacman", "-Q").Output()
	if error != nil {
		log.Fatal(error)
	}

	packages := string(packagesBytes)
	lines := strings.Count(packages, "\n")
	return lines
}

func GetUsedMemory() int {
	rawFreeMemory := getRawTotalMemory() - getRawFreeMemory()
	return rawFreeMemory / 1024
}
