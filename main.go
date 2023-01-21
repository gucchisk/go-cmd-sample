package main

import (
	"flag"
	"fmt"
	"runtime/debug"
)
func main() {
	vf := flag.Bool("version", false, "version flag")
	flag.Parse()
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("error: ReadBuildInfo")
		return
	}
	version := buildInfo.Main.Version
	if *vf {
		fmt.Printf("version: %s\n", version)
		return
	}
	fmt.Printf("This is go-cmd-sample %s\n", version)
}
