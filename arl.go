package network

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type ARLRecord struct {
	Port       int
	VLAN       int
	MACAddress string
	AgeTime    int
	Flags      int
}

func ParseARLResponse(response string) ([]ARLRecord, error) {
	var records []ARLRecord
	lines := strings.Split(strings.TrimSpace(response), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) != 5 {
			return nil, fmt.Errorf("invalid ARL record: %s", line)
		}
		port, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, fmt.Errorf("invalid port number: %s", fields[0])
		}
		vlan, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, fmt.Errorf("invalid VLAN number: %s", fields[1])
		}
		ageTime, err := strconv.Atoi(fields[3])
		if err != nil {
			return nil, fmt.Errorf("invalid age time: %s", fields[3])
		}
		flags, err := strconv.Atoi(fields[4])
		if err != nil {
			return nil, fmt.Errorf("invalid flags: %s", fields[4])
		}
		record := ARLRecord{
			Port:       port,
			VLAN:       vlan,
			MACAddress: fields[2],
			AgeTime:    ageTime,
			Flags:      flags,
		}
		records = append(records, record)
	}
	return records, nil
}

func ReadARLFile() (string, error) {
	bytes, err := ioutil.ReadFile("/proc/switch_uext/switch0/arl")
	if err != nil {
		return "", fmt.Errorf("failed to read ARL file: %s", err)
	}
	return string(bytes), nil
}
