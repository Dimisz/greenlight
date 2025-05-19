package data

import (
	"fmt"
	"strconv"
)

// custom runtype with underlying type of int32
// implemented to demostrate fine-tuning JSON Marshalling
type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	// needs to be surrounded in double quotes to be a valid JSON
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}
