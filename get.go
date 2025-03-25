package main

import (
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
)

func GetHostName() string {
	hostname, _ := os.Hostname()
	return hostname
}

func GetUserName() string {
	user, _ := user.Current()
	return user.Username
}

func GetDistro() string {
	osReleaseBytes, _ := os.ReadFile("/etc/os-release")
	osrelease := string(osReleaseBytes)
	return SnipSnip("PRETTY_NAME=\"", "\"", osrelease)
}

func GetKernel() string {
	kernelbytes, _ := exec.Command("uname", "-r").Output()
	return string(kernelbytes)
}

func GetUptime() string {
	uptimebytes, _ := exec.Command("uptime", "-p").Output()
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
	meminfobytes, _ := os.ReadFile("/proc/meminfo")
	meminfostring := string(meminfobytes)
	totalmemorystring := SnipSnip("MemTotal:", " kB", meminfostring)
	totalmemorystring = strings.TrimSpace(totalmemorystring)
	totalrawmemory, _ := strconv.Atoi(totalmemorystring)
	return totalrawmemory
}

func getRawFreeMemory() int {
	meminfobytes, _ := os.ReadFile("/proc/meminfo")
	meminfostring := string(meminfobytes)
	totalmemorystring := SnipSnip("MemAvailable:", " kB", meminfostring)
	totalmemorystring = strings.TrimSpace(totalmemorystring)
	totalrawmemory, _ := strconv.Atoi(totalmemorystring)
	return totalrawmemory
}

func GetTotalMemory() int {
	return getRawTotalMemory() / 1024
}

func GetPackages() int {
	packagesbytes, _ := exec.Command("pacman", "-Q").Output()
	packages := string(packagesbytes)
	lines := strings.Count(packages, "\n")
	return lines
}

func GetUsedMemory() int {
	rawfreememory := getRawTotalMemory() - getRawFreeMemory()
	return rawfreememory / 1024
}
