package main

import (
	"fmt"
	"github.com/mitchty/sys/sys"
)

func main() {
	test := sys.NewInfo()
	fmt.Println(test)
	fmt.Printf("buildinfo: %s %s %s\n", GitHash, GitTag, GitBuildDate)
}
