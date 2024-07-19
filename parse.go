package lldpxml

import "encoding/xml"

func Parse(data []byte) Lldp {
	var result Lldp
	err := xml.Unmarshal(data, &result)
	if err != nil {
		panic(err)
	}
	return result
}
