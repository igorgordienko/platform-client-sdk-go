package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Replacementterm
type Replacementterm struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// ExistingValue
	ExistingValue *string `json:"existingValue,omitempty"`


	// UpdatedValue
	UpdatedValue *string `json:"updatedValue,omitempty"`

}

// String returns a JSON representation of the model
func (o *Replacementterm) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
