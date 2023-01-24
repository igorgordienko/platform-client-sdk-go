package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Nludomain
type Nludomain struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of the NLU domain.
	Name *string `json:"name,omitempty"`

	// Language - The language culture of the NLU domain, e.g. `en-us`, `de-de`.
	Language *string `json:"language,omitempty"`

	// DraftVersion - The draft version of that NLU domain.
	DraftVersion *Nludomainversion `json:"draftVersion,omitempty"`

	// LastPublishedVersion - The last published version of that NLU domain.
	LastPublishedVersion *Nludomainversion `json:"lastPublishedVersion,omitempty"`

	// DateCreated - The date when the NLU domain was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - The date when the NLU domain was updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// EngineVersion - The version of the NLU engine to use.
	EngineVersion *string `json:"engineVersion,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Nludomain) SetField(field string, fieldValue interface{}) {
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

func (o Nludomain) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified", }
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
	type Alias Nludomain
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		DraftVersion *Nludomainversion `json:"draftVersion,omitempty"`
		
		LastPublishedVersion *Nludomainversion `json:"lastPublishedVersion,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		EngineVersion *string `json:"engineVersion,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Language: o.Language,
		
		DraftVersion: o.DraftVersion,
		
		LastPublishedVersion: o.LastPublishedVersion,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		EngineVersion: o.EngineVersion,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Nludomain) UnmarshalJSON(b []byte) error {
	var NludomainMap map[string]interface{}
	err := json.Unmarshal(b, &NludomainMap)
	if err != nil {
		return err
	}
	
	if Id, ok := NludomainMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := NludomainMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Language, ok := NludomainMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if DraftVersion, ok := NludomainMap["draftVersion"].(map[string]interface{}); ok {
		DraftVersionString, _ := json.Marshal(DraftVersion)
		json.Unmarshal(DraftVersionString, &o.DraftVersion)
	}
	
	if LastPublishedVersion, ok := NludomainMap["lastPublishedVersion"].(map[string]interface{}); ok {
		LastPublishedVersionString, _ := json.Marshal(LastPublishedVersion)
		json.Unmarshal(LastPublishedVersionString, &o.LastPublishedVersion)
	}
	
	if dateCreatedString, ok := NludomainMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := NludomainMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if EngineVersion, ok := NludomainMap["engineVersion"].(string); ok {
		o.EngineVersion = &EngineVersion
	}
    
	if SelfUri, ok := NludomainMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Nludomain) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
