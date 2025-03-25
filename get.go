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

	osrelease := string(osReleaseBytes)
	return SnipSnip("PRETTY_NAME=\"", "\"", osrelease)
}

func GetKernel() string {
	kernelbytes, error := exec.Command("uname", "-r").Output()
	if error != nil {
		log.Fatal(error)
	}

	return string(kernelbytes)
}

func GetUptime() string {
	uptimebytes, error := exec.Command("uptime", "-p").Output()
	if error != nil {
		log.Fatal(error)
	}

	uptime := string(uptimebytes)
	uptime = strings.ReplaceAll(uptime, "minutes", "m")
	uptime = strings.ReplaceAll(uptime, "hours", "h")
	uptime = strings.ReplaceAll(uptime, "days", "d")
	uptime = strings.ReplaceAll(uptime, "minute", "m")
	uptime = strings.ReplaceAll(uptime, "hour", "h")
	uptime = strings.ReplaceAll(uptime, "day", "d")
	return uptime
}

func GetShell() string {
	shellpieces := strings.SplitAfter(os.Getenv("SHELL"), "/")
	return shellpieces[len(shellpieces)-1]
}

func getRawTotalMemory() int {
	meminfobytes, error := os.ReadFile("/proc/meminfo")
	if error != nil {
		log.Fatal(error)
	}

	meminfostring := string(meminfobytes)
	totalmemorystring := SnipSnip("MemTotal:", " kB", meminfostring)
	totalmemorystring = strings.TrimSpace(totalmemorystring)
	totalrawmemory, error := strconv.Atoi(totalmemorystring)
	if error != nil {
		log.Fatal(error)
	}

	return totalrawmemory
}

func getRawFreeMemory() int {
	meminfobytes, error := os.ReadFile("/proc/meminfo")
	if error != nil {
		log.Fatal(error)
	}

	meminfostring := string(meminfobytes)
	totalmemorystring := SnipSnip("MemAvailable:", " kB", meminfostring)
	totalmemorystring = strings.TrimSpace(totalmemorystring)
	totalrawmemory, error := strconv.Atoi(totalmemorystring)
	if error != nil {
		log.Fatal(error)
	}

	return totalrawmemory
}

func GetTotalMemory() int {
	return getRawTotalMemory() / 1024
}

func GetPackages() int {
	packagesbytes, error := exec.Command("pacman", "-Q").Output()
	if error != nil {
		log.Fatal(error)
	}

	packages := string(packagesbytes)
	lines := strings.Count(packages, "\n")
	return lines
}

func GetUsedMemory() int {
	rawfreememory := getRawTotalMemory() - getRawFreeMemory()
	return rawfreememory / 1024
}
