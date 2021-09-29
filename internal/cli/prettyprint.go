package cli

import (
	"encoding/json"
	"fmt"
)

// PrettyString returns a human-readable string representation of a value.
func PrettyString(in interface{}) string {
	res, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		return fmt.Sprintf("%+v", in)
	}

	return string(res)
}
