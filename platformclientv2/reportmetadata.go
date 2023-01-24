package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Reportmetadata
type Reportmetadata struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Title
	Title *string `json:"title,omitempty"`

	// Description
	Description *string `json:"description,omitempty"`

	// Keywords
	Keywords *[]string `json:"keywords,omitempty"`

	// AvailableLocales
	AvailableLocales *[]string `json:"availableLocales,omitempty"`

	// Parameters
	Parameters *[]Parameter `json:"parameters,omitempty"`

	// ExampleUrl
	ExampleUrl *string `json:"exampleUrl,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Reportmetadata) SetField(field string, fieldValue interface{}) {
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

func (o Reportmetadata) MarshalJSON() ([]byte, error) {
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
	type Alias Reportmetadata
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Keywords *[]string `json:"keywords,omitempty"`
		
		AvailableLocales *[]string `json:"availableLocales,omitempty"`
		
		Parameters *[]Parameter `json:"parameters,omitempty"`
		
		ExampleUrl *string `json:"exampleUrl,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Title: o.Title,
		
		Description: o.Description,
		
		Keywords: o.Keywords,
		
		AvailableLocales: o.AvailableLocales,
		
		Parameters: o.Parameters,
		
		ExampleUrl: o.ExampleUrl,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Reportmetadata) UnmarshalJSON(b []byte) error {
	var ReportmetadataMap map[string]interface{}
	err := json.Unmarshal(b, &ReportmetadataMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ReportmetadataMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ReportmetadataMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Title, ok := ReportmetadataMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Description, ok := ReportmetadataMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Keywords, ok := ReportmetadataMap["keywords"].([]interface{}); ok {
		KeywordsString, _ := json.Marshal(Keywords)
		json.Unmarshal(KeywordsString, &o.Keywords)
	}
	
	if AvailableLocales, ok := ReportmetadataMap["availableLocales"].([]interface{}); ok {
		AvailableLocalesString, _ := json.Marshal(AvailableLocales)
		json.Unmarshal(AvailableLocalesString, &o.AvailableLocales)
	}
	
	if Parameters, ok := ReportmetadataMap["parameters"].([]interface{}); ok {
		ParametersString, _ := json.Marshal(Parameters)
		json.Unmarshal(ParametersString, &o.Parameters)
	}
	
	if ExampleUrl, ok := ReportmetadataMap["exampleUrl"].(string); ok {
		o.ExampleUrl = &ExampleUrl
	}
    
	if SelfUri, ok := ReportmetadataMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Reportmetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
