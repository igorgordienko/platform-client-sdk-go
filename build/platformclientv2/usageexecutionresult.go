package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Usageexecutionresult
type Usageexecutionresult struct { 
	// ExecutionId - The id of the query execution
	ExecutionId *string `json:"executionId,omitempty"`


	// ResultsUri - URI where the query results can be retrieved
	ResultsUri *string `json:"resultsUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Usageexecutionresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
