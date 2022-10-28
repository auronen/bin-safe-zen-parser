package parsers

import (
	"fmt"
	"os"
	"strings"
)

func (zen *BsZen) GetWaypointList() (list []string) {
	list = make([]string, 0, 20)
	fmt.Fprintf(os.Stdout, "zen: %v\n", zen)
	fmt.Fprintf(os.Stdout, "OCWorld: %v\n", zen.OCWorld)
	fmt.Fprintf(os.Stdout, "OCWorld: %v\n", zen.OCWorld)
	fmt.Fprintf(os.Stdout, "len(zen.OCWorld.Properties: %v\n", len(zen.OCWorld.Properties))

	for i := 0; i < len(zen.OCWorld.Properties); i++ {
		if zen.OCWorld.Properties[i].Type == 1 { // string type 0x1
			objectID := zen.OCWorld.Properties[i].PropBody.(*BsZen_TypeString).ObjectId
			parts := strings.Split(objectID, " ")
			if len(parts) > 3 {
				if parts[1] == "zCWaypoint" {
					list = append(list, zen.OCWorld.Properties[i+2].PropBody.(*BsZen_TypeString).ObjectId)
				}
			}
		}
	}
	return list
}
