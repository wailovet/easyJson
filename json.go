package easyJson

import "encoding/json"

func Decode(str string) (interface{}, error) {
	b := []byte(str)
	var f interface{}
	err := json.Unmarshal(b, &f)
	return f, err
}

func Encode(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	return string(b), err
}
