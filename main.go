package main

import (
	"runtime"

	"github.com/ezParth/structure/bashCommand"
	windowsCommand "github.com/ezParth/structure/windowCommand"
)

func main() {
	if runtime.GOOS == "windows" {
		windowsCommand.MakeFile()
	} else {
		bashCommand.MakeFile()
	}
}
