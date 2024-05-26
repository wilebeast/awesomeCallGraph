package main

import "encoding/json"

func marshal2String(v interface{}) string {
	bytes, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return ""
	}
	return string(bytes)
}

func unmarshal2map(in string) map[string]interface{} {
	out := make(map[string]interface{})
	err := json.Unmarshal([]byte(in), &out)
	if err != nil {
		return nil
	}
	return out
}
