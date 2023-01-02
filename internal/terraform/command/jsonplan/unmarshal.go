package jsonplan


import (
	"encoding/json"
)

func Unmarshal(txt []byte) plan {

	var p plan
	json.Unmarshal(txt, &p)

	return p
}
