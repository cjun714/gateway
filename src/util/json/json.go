package json

import (
	"encoding/json"
	"util/log"
)

// ToJSON is to searialize obj to JSON byte[]
func ToJSON(obj interface{}) []byte {
	body, e := json.Marshal(obj)
	if e != nil {
		log.E(e.Error())
	}
	return body
}

// ToObj is to unsearialize obj from JSON byte[]
func ToObj(jsonBytes []byte, obj interface{}) error {
	e := json.Unmarshal([]byte(jsonBytes), obj)
	if e != nil {
		log.E(e.Error())
		return e
	}
	return nil
}
