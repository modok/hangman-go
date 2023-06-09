package util

import (
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	clear := make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear") //MacOS
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen for " + runtime.GOOS)
	}
}
