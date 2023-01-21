package main

import (
	"flag"
	"fmt"
	"runtime/debug"
)
func main() {
	vf := flag.Bool("version", false, "version flag")
	flag.Parse()
	
	if *vf {
		buildInfo, ok := debug.ReadBuildInfo()
		if !ok {
			fmt.Println("error: ReadBuildInfo")
			return
		}
		version := buildInfo.Main.Version
		fmt.Printf("version: %s\n", version)
		return
	}
	fmt.Println("This is go-cmd-sample")
}
