package common

import (
	"bytes"
	"encoding/json"
)

// Encoding is what can be injected into a subcommand when you need encoding-related items
type Encoding struct{}

// ToJSON returns a JSON byte slice encoding data
func (e Encoding) ToJSON(data interface{}) []byte {
	b, _ := json.Marshal(data)
	return b
}

// ToJSONString returns a JSON string encoding data
func (e Encoding) ToJSONString(data interface{}) string {
	return string(e.ToJSON(data))
}

// ToFormattedString returns a JSON formatted string encoding data
func (e Encoding) ToFormattedString(data interface{}) string {
	b := e.ToJSON(data)

	var out bytes.Buffer
	json.Indent(&out, b, "", "    ")
	return out.String()

}
