package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Meteredevaluationassignment
type Meteredevaluationassignment struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EvaluationContextId
	EvaluationContextId *string `json:"evaluationContextId,omitempty"`

	// Evaluators
	Evaluators *[]User `json:"evaluators,omitempty"`

	// MaxNumberEvaluations
	MaxNumberEvaluations *int `json:"maxNumberEvaluations,omitempty"`

	// EvaluationForm
	EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`

	// AssignToActiveUser
	AssignToActiveUser *bool `json:"assignToActiveUser,omitempty"`

	// TimeInterval
	TimeInterval *Timeinterval `json:"timeInterval,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Meteredevaluationassignment) SetField(field string, fieldValue interface{}) {
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

func (o Meteredevaluationassignment) MarshalJSON() ([]byte, error) {
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
	type Alias Meteredevaluationassignment
	
	return json.Marshal(&struct { 
		EvaluationContextId *string `json:"evaluationContextId,omitempty"`
		
		Evaluators *[]User `json:"evaluators,omitempty"`
		
		MaxNumberEvaluations *int `json:"maxNumberEvaluations,omitempty"`
		
		EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`
		
		AssignToActiveUser *bool `json:"assignToActiveUser,omitempty"`
		
		TimeInterval *Timeinterval `json:"timeInterval,omitempty"`
		Alias
	}{ 
		EvaluationContextId: o.EvaluationContextId,
		
		Evaluators: o.Evaluators,
		
		MaxNumberEvaluations: o.MaxNumberEvaluations,
		
		EvaluationForm: o.EvaluationForm,
		
		AssignToActiveUser: o.AssignToActiveUser,
		
		TimeInterval: o.TimeInterval,
		Alias:    (Alias)(o),
	})
}

func (o *Meteredevaluationassignment) UnmarshalJSON(b []byte) error {
	var MeteredevaluationassignmentMap map[string]interface{}
	err := json.Unmarshal(b, &MeteredevaluationassignmentMap)
	if err != nil {
		return err
	}
	
	if EvaluationContextId, ok := MeteredevaluationassignmentMap["evaluationContextId"].(string); ok {
		o.EvaluationContextId = &EvaluationContextId
	}
    
	if Evaluators, ok := MeteredevaluationassignmentMap["evaluators"].([]interface{}); ok {
		EvaluatorsString, _ := json.Marshal(Evaluators)
		json.Unmarshal(EvaluatorsString, &o.Evaluators)
	}
	
	if MaxNumberEvaluations, ok := MeteredevaluationassignmentMap["maxNumberEvaluations"].(float64); ok {
		MaxNumberEvaluationsInt := int(MaxNumberEvaluations)
		o.MaxNumberEvaluations = &MaxNumberEvaluationsInt
	}
	
	if EvaluationForm, ok := MeteredevaluationassignmentMap["evaluationForm"].(map[string]interface{}); ok {
		EvaluationFormString, _ := json.Marshal(EvaluationForm)
		json.Unmarshal(EvaluationFormString, &o.EvaluationForm)
	}
	
	if AssignToActiveUser, ok := MeteredevaluationassignmentMap["assignToActiveUser"].(bool); ok {
		o.AssignToActiveUser = &AssignToActiveUser
	}
    
	if TimeInterval, ok := MeteredevaluationassignmentMap["timeInterval"].(map[string]interface{}); ok {
		TimeIntervalString, _ := json.Marshal(TimeInterval)
		json.Unmarshal(TimeIntervalString, &o.TimeInterval)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Meteredevaluationassignment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
