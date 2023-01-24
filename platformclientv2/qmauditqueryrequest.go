package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Qmauditqueryrequest
type Qmauditqueryrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Interval - Date and time range of data to query. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`

	// Filters - List of filters for the query.
	Filters *[]Qualityauditqueryfilter `json:"filters,omitempty"`

	// Sort - Sort parameter for the query.
	Sort *[]Auditquerysort `json:"sort,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Qmauditqueryrequest) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Qmauditqueryrequest) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Qmauditqueryrequest
	
	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		Filters *[]Qualityauditqueryfilter `json:"filters,omitempty"`
		
		Sort *[]Auditquerysort `json:"sort,omitempty"`
		Alias
	}{ 
		Interval: o.Interval,
		
		Filters: o.Filters,
		
		Sort: o.Sort,
		Alias:    (Alias)(o),
	})
}

func (o *Qmauditqueryrequest) UnmarshalJSON(b []byte) error {
	var QmauditqueryrequestMap map[string]interface{}
	err := json.Unmarshal(b, &QmauditqueryrequestMap)
	if err != nil {
		return err
	}
	
	if Interval, ok := QmauditqueryrequestMap["interval"].(string); ok {
		o.Interval = &Interval
	}
    
	if Filters, ok := QmauditqueryrequestMap["filters"].([]interface{}); ok {
		FiltersString, _ := json.Marshal(Filters)
		json.Unmarshal(FiltersString, &o.Filters)
	}
	
	if Sort, ok := QmauditqueryrequestMap["sort"].([]interface{}); ok {
		SortString, _ := json.Marshal(Sort)
		json.Unmarshal(SortString, &o.Sort)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Qmauditqueryrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
