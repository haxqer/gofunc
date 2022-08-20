package gofunc

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

func PrintJson(v interface{}) {
	jsonText, _ := json.Marshal(v)
	fmt.Printf("%s\n", jsonText)
}

func PrintXML(v interface{}) {
	xmlText, _ := xml.Marshal(v)
	fmt.Printf("%s\n", xmlText)
}