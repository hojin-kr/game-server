package data

import "flag"

var (
	serverPlaceVersion = flag.Int64("serverPlaceVersion", 2, "data place version")
	names              = []string{"name1", "name2", "name3"}
	address            = []string{"address1", "address2", "address3"}
)

func Get(clientPlaceVer int64) ([]string, []string) {
	if clientPlaceVer >= *serverPlaceVersion {
		return []string{}, []string{}
	}
	return names, address
}
