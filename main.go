package main

import (
	"bufio"
	"fmt"
	"os"
	"zen_test/parsers"

	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
)

func main() {
	file, _ := os.Open("WORLD.ZEN")
	zen := parsers.NewBsZen()
	zen.Read(kaitai.NewStream(file), nil, zen)

	// Prints all waypoints to standard output
	wplist := zen.GetWaypointList()
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	for _, wp := range wplist {
		f.WriteString(wp + "\n")
	}

	// Prints all items to standard output
	itemlist := zen.GetItemList()
	for key, amount := range itemlist {
		f.WriteString(fmt.Sprintf("%s: %d\n", key, amount))
	}
}
