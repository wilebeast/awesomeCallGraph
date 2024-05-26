package main

import "encoding/json"

func marshal2String(v interface{}) string {
	bytes, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return ""
	}
	return string(bytes)
}
