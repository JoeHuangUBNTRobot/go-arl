package main

import (
	"fmt"
	"strings"

	networkARL "github.com/JoeHuangUBNTRobot/go-arl"
)

func main() {
	response, err := networkARL.ReadARLFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	arlRecords, err := networkARL.ParseARLResponse(response)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("ARL records:\n%v\n", arlRecords)

	portMACmap := make(map[int]string)
	for _, arlrecord := range arlRecords {
		portMACmap[arlrecord.Port] = strings.ReplaceAll(arlrecord.MACAddress, ":", "")
	}

	// Check if port "1" and "3" exists in the map
	if _, ok := portMACmap[1]; !ok {
		// Port 1 is not present in the map, response the error
		fmt.Printf("Error: Port 1 is not in ARL table")
	} else if _, ok := portMACmap[3]; !ok {
		// Port 3 is not present in the map, response the error
		fmt.Printf("Error: Port 3 is not in ARL table")
	}
}
