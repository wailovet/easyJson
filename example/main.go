package main

import (
	"fmt"
	"github.com/wailovet/easyJson"
)

func main() {
	jsonObj, _ := easyJson.Decode(`{"hello":"world","www":5.15453, "qqqq":{"a":{"b":true},"c":null,"d":[{"d":"d","e":"e"}]},"rrr":["s","f"]}`)
	fmt.Print(jsonObj)
	jsonStr, _ := easyJson.Encode(jsonObj)
	fmt.Print(jsonStr)
}
