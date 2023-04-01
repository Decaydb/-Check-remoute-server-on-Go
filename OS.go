package main

import (
	"fmt"
	"runtime"
)

var OS string

func main() {
	switch runtime.GOOS {
	case "windows":
		OS = "windows"
	case "darwin":
		OS = "darwin"
	case "linux":
		OS = "linux"
	default:
		panic(fmt.Sprintf("Неизвестная OS %s", runtime.GOOS))
	}
	fmt.Println(OS)
}
