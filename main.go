package main

import (
	"bufio"
	"os"
	"zen_test/parsers"

	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
)

func main() {
	file, _ := os.Open("WORLD.ZEN")
	zen := parsers.NewBsZen()
	zen.Read(kaitai.NewStream(file), nil, zen)
	wplist := zen.GetWaypointList()

	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	for _, wp := range wplist {
		f.WriteString(wp + "\n")
	}
}
