package main

import (
	"flag"
	"fmt"
	"runtime/debug"
)
var version = "unknown"
func main() {
	vf := flag.Bool("version", false, "version flag")
	bif := flag.Bool("info", false, "build info flag")
	sf := flag.Bool("settings", false, "settings flag")
	flag.Parse()
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("error: ReadBuildInfo")
		return
	}
	if version == "unknown" {
		version = buildInfo.Main.Version
	}
	if *vf {
		fmt.Printf("version: %s\n", version)
		return
	} else if *bif {
		main := buildInfo.Main
		fmt.Printf("go: %s\n", buildInfo.GoVersion)
		fmt.Printf("path: %s\n", main.Path)
		fmt.Printf("sum: %s\n", main.Sum)
		return
	} else if *sf {
		for _, setting := range buildInfo.Settings {
			fmt.Printf("%s: %s\n", setting.Key, setting.Value)
		}
	}
	fmt.Printf("This is sample command %s\n", version)
}
