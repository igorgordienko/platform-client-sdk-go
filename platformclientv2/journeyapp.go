package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyapp
type Journeyapp struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - Name of the application (e.g. mybankingapp).
	Name *string `json:"name,omitempty"`

	// Namespace - Namespace of the application (e.g. com.genesys.bancodinero).
	Namespace *string `json:"namespace,omitempty"`

	// Version - Version of the application (e.g. 5.9.27).
	Version *string `json:"version,omitempty"`

	// BuildNumber - Build number of the application (e.g. 701).
	BuildNumber *string `json:"buildNumber,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeyapp) SetField(field string, fieldValue interface{}) {
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

func (o Journeyapp) MarshalJSON() ([]byte, error) {
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
	type Alias Journeyapp
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Namespace *string `json:"namespace,omitempty"`
		
		Version *string `json:"version,omitempty"`
		
		BuildNumber *string `json:"buildNumber,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		Namespace: o.Namespace,
		
		Version: o.Version,
		
		BuildNumber: o.BuildNumber,
		Alias:    (Alias)(o),
	})
}

func (o *Journeyapp) UnmarshalJSON(b []byte) error {
	var JourneyappMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyappMap)
	if err != nil {
		return err
	}
	
	if Name, ok := JourneyappMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Namespace, ok := JourneyappMap["namespace"].(string); ok {
		o.Namespace = &Namespace
	}
    
	if Version, ok := JourneyappMap["version"].(string); ok {
		o.Version = &Version
	}
    
	if BuildNumber, ok := JourneyappMap["buildNumber"].(string); ok {
		o.BuildNumber = &BuildNumber
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyapp) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
