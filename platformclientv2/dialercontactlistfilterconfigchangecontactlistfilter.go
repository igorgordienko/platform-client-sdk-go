package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercontactlistfilterconfigchangecontactlistfilter
type Dialercontactlistfilterconfigchangecontactlistfilter struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ContactList
	ContactList *Dialercontactlistfilterconfigchangeurireference `json:"contactList,omitempty"`

	// ContactListColumns - The list of contact list columns
	ContactListColumns *[]string `json:"contactListColumns,omitempty"`

	// Clauses
	Clauses *[]Dialercontactlistfilterconfigchangefilterclause `json:"clauses,omitempty"`

	// FilterType - Contact list filter type
	FilterType *string `json:"filterType,omitempty"`

	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The UI-visible name of the object
	Name *string `json:"name,omitempty"`

	// DateCreated - Creation time of the entity
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Last modified time of the entity
	DateModified *time.Time `json:"dateModified,omitempty"`

	// Version - Required for updates, must match the version number of the most recent update
	Version *int `json:"version,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Dialercontactlistfilterconfigchangecontactlistfilter) SetField(field string, fieldValue interface{}) {
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

func (o Dialercontactlistfilterconfigchangecontactlistfilter) MarshalJSON() ([]byte, error) {
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
	type Alias Dialercontactlistfilterconfigchangecontactlistfilter
	
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
		ContactList *Dialercontactlistfilterconfigchangeurireference `json:"contactList,omitempty"`
		
		ContactListColumns *[]string `json:"contactListColumns,omitempty"`
		
		Clauses *[]Dialercontactlistfilterconfigchangefilterclause `json:"clauses,omitempty"`
		
		FilterType *string `json:"filterType,omitempty"`
		
		AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *int `json:"version,omitempty"`
		Alias
	}{ 
		ContactList: o.ContactList,
		
		ContactListColumns: o.ContactListColumns,
		
		Clauses: o.Clauses,
		
		FilterType: o.FilterType,
		
		AdditionalProperties: o.AdditionalProperties,
		
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		Alias:    (Alias)(o),
	})
}

func (o *Dialercontactlistfilterconfigchangecontactlistfilter) UnmarshalJSON(b []byte) error {
	var DialercontactlistfilterconfigchangecontactlistfilterMap map[string]interface{}
	err := json.Unmarshal(b, &DialercontactlistfilterconfigchangecontactlistfilterMap)
	if err != nil {
		return err
	}
	
	if ContactList, ok := DialercontactlistfilterconfigchangecontactlistfilterMap["contactList"].(map[string]interface{}); ok {
		ContactListString, _ := json.Marshal(ContactList)
		json.Unmarshal(ContactListString, &o.ContactList)
	}
	
	if ContactListColumns, ok := DialercontactlistfilterconfigchangecontactlistfilterMap["contactListColumns"].([]interface{}); ok {
		ContactListColumnsString, _ := json.Marshal(ContactListColumns)
		json.Unmarshal(ContactListColumnsString, &o.ContactListColumns)
	}
	
	if Clauses, ok := DialercontactlistfilterconfigchangecontactlistfilterMap["clauses"].([]interface{}); ok {
		ClausesString, _ := json.Marshal(Clauses)
		json.Unmarshal(ClausesString, &o.Clauses)
	}
	
	if FilterType, ok := DialercontactlistfilterconfigchangecontactlistfilterMap["filterType"].(string); ok {
		o.FilterType = &FilterType
	}
    
	if AdditionalProperties, ok := DialercontactlistfilterconfigchangecontactlistfilterMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	
	if Id, ok := DialercontactlistfilterconfigchangecontactlistfilterMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DialercontactlistfilterconfigchangecontactlistfilterMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := DialercontactlistfilterconfigchangecontactlistfilterMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := DialercontactlistfilterconfigchangecontactlistfilterMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := DialercontactlistfilterconfigchangecontactlistfilterMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercontactlistfilterconfigchangecontactlistfilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
