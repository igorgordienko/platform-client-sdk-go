package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowdivisionview
type Flowdivisionview struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The flow identifier
	Id *string `json:"id,omitempty"`

	// Name - The flow name
	Name *string `json:"name,omitempty"`

	// Division - The division to which this entity belongs.
	Division *Writabledivision `json:"division,omitempty"`

	// VarType
	VarType *string `json:"type,omitempty"`

	// Description - the flow description
	Description *string `json:"description,omitempty"`

	// InputSchema - json schema describing the inputs for the flow
	InputSchema *Jsonschemadocument `json:"inputSchema,omitempty"`

	// OutputSchema - json schema describing the outputs for the flow
	OutputSchema *Jsonschemadocument `json:"outputSchema,omitempty"`

	// SupportedLanguages - List of supported languages for the published version of the flow.
	SupportedLanguages *[]Supportedlanguage `json:"supportedLanguages,omitempty"`

	// PublishedVersion - published version information if there is a published version
	PublishedVersion *Flowversion `json:"publishedVersion,omitempty"`

	// DebugVersion - debug version information if there is a debug version
	DebugVersion *Flowversion `json:"debugVersion,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Flowdivisionview) SetField(field string, fieldValue interface{}) {
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

func (o Flowdivisionview) MarshalJSON() ([]byte, error) {
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
	type Alias Flowdivisionview
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Division *Writabledivision `json:"division,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		InputSchema *Jsonschemadocument `json:"inputSchema,omitempty"`
		
		OutputSchema *Jsonschemadocument `json:"outputSchema,omitempty"`
		
		SupportedLanguages *[]Supportedlanguage `json:"supportedLanguages,omitempty"`
		
		PublishedVersion *Flowversion `json:"publishedVersion,omitempty"`
		
		DebugVersion *Flowversion `json:"debugVersion,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Division: o.Division,
		
		VarType: o.VarType,
		
		Description: o.Description,
		
		InputSchema: o.InputSchema,
		
		OutputSchema: o.OutputSchema,
		
		SupportedLanguages: o.SupportedLanguages,
		
		PublishedVersion: o.PublishedVersion,
		
		DebugVersion: o.DebugVersion,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Flowdivisionview) UnmarshalJSON(b []byte) error {
	var FlowdivisionviewMap map[string]interface{}
	err := json.Unmarshal(b, &FlowdivisionviewMap)
	if err != nil {
		return err
	}
	
	if Id, ok := FlowdivisionviewMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := FlowdivisionviewMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Division, ok := FlowdivisionviewMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if VarType, ok := FlowdivisionviewMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Description, ok := FlowdivisionviewMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if InputSchema, ok := FlowdivisionviewMap["inputSchema"].(map[string]interface{}); ok {
		InputSchemaString, _ := json.Marshal(InputSchema)
		json.Unmarshal(InputSchemaString, &o.InputSchema)
	}
	
	if OutputSchema, ok := FlowdivisionviewMap["outputSchema"].(map[string]interface{}); ok {
		OutputSchemaString, _ := json.Marshal(OutputSchema)
		json.Unmarshal(OutputSchemaString, &o.OutputSchema)
	}
	
	if SupportedLanguages, ok := FlowdivisionviewMap["supportedLanguages"].([]interface{}); ok {
		SupportedLanguagesString, _ := json.Marshal(SupportedLanguages)
		json.Unmarshal(SupportedLanguagesString, &o.SupportedLanguages)
	}
	
	if PublishedVersion, ok := FlowdivisionviewMap["publishedVersion"].(map[string]interface{}); ok {
		PublishedVersionString, _ := json.Marshal(PublishedVersion)
		json.Unmarshal(PublishedVersionString, &o.PublishedVersion)
	}
	
	if DebugVersion, ok := FlowdivisionviewMap["debugVersion"].(map[string]interface{}); ok {
		DebugVersionString, _ := json.Marshal(DebugVersion)
		json.Unmarshal(DebugVersionString, &o.DebugVersion)
	}
	
	if SelfUri, ok := FlowdivisionviewMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Flowdivisionview) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
