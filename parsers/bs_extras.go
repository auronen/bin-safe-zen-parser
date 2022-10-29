package parsers

import (
	"strings"
)

func (zen *BsZen) GetWaypointList() (list []string) {
	list = make([]string, 0, 20)
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

func (zen *BsZen) GetItemList() map[string]int {
	items := make(map[string]int, 20)
	for i := 0; i < len(zen.OCWorld.Properties); i++ {
		if zen.OCWorld.Properties[i].Type == 1 { // string type 0x1
			objectID := zen.OCWorld.Properties[i].PropBody.(*BsZen_TypeString).ObjectId
			if strings.Contains(objectID, "oCItem:zCVob") {
				items[zen.OCWorld.Properties[i+6].PropBody.(*BsZen_TypeString).ObjectId] += 1
			}
		}
	}
	return items
}
