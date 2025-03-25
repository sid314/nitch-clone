package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/benhoyt/goawk/interp"
)

func GetHostName() string {
	hostname, _ := os.Hostname()
	return hostname
}

func GetUserName() string {
	user, _ := user.Current()
	return user.Username
}

func GetOS() string {
	var builder strings.Builder
	osReleaseBytes, _ := os.ReadFile("/etc/os-release")
	reader := strings.NewReader(string(osReleaseBytes))
	err := interp.Exec("'/PRETTY_NAME/ {print $2}'", "=", reader, &builder)
	if err != nil {
		fmt.Println(err.Error())
	}
	return builder.String()
}
