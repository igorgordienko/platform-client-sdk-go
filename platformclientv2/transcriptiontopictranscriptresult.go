package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Transcriptiontopictranscriptresult
type Transcriptiontopictranscriptresult struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// UtteranceId
	UtteranceId *string `json:"utteranceId,omitempty"`

	// IsFinal
	IsFinal *bool `json:"isFinal,omitempty"`

	// Channel
	Channel *string `json:"channel,omitempty"`

	// Alternatives
	Alternatives *[]Transcriptiontopictranscriptalternative `json:"alternatives,omitempty"`

	// AgentAssistantId
	AgentAssistantId *string `json:"agentAssistantId,omitempty"`

	// EngineId
	EngineId *string `json:"engineId,omitempty"`

	// Dialect
	Dialect *string `json:"dialect,omitempty"`

	// SpeechTextAnalyticsProgramId
	SpeechTextAnalyticsProgramId *string `json:"speechTextAnalyticsProgramId,omitempty"`

	// AgentAssistEnabled
	AgentAssistEnabled *bool `json:"agentAssistEnabled,omitempty"`

	// VoiceTranscriptionEnabled
	VoiceTranscriptionEnabled *bool `json:"voiceTranscriptionEnabled,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Transcriptiontopictranscriptresult) SetField(field string, fieldValue interface{}) {
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

func (o Transcriptiontopictranscriptresult) MarshalJSON() ([]byte, error) {
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
	type Alias Transcriptiontopictranscriptresult
	
	return json.Marshal(&struct { 
		UtteranceId *string `json:"utteranceId,omitempty"`
		
		IsFinal *bool `json:"isFinal,omitempty"`
		
		Channel *string `json:"channel,omitempty"`
		
		Alternatives *[]Transcriptiontopictranscriptalternative `json:"alternatives,omitempty"`
		
		AgentAssistantId *string `json:"agentAssistantId,omitempty"`
		
		EngineId *string `json:"engineId,omitempty"`
		
		Dialect *string `json:"dialect,omitempty"`
		
		SpeechTextAnalyticsProgramId *string `json:"speechTextAnalyticsProgramId,omitempty"`
		
		AgentAssistEnabled *bool `json:"agentAssistEnabled,omitempty"`
		
		VoiceTranscriptionEnabled *bool `json:"voiceTranscriptionEnabled,omitempty"`
		Alias
	}{ 
		UtteranceId: o.UtteranceId,
		
		IsFinal: o.IsFinal,
		
		Channel: o.Channel,
		
		Alternatives: o.Alternatives,
		
		AgentAssistantId: o.AgentAssistantId,
		
		EngineId: o.EngineId,
		
		Dialect: o.Dialect,
		
		SpeechTextAnalyticsProgramId: o.SpeechTextAnalyticsProgramId,
		
		AgentAssistEnabled: o.AgentAssistEnabled,
		
		VoiceTranscriptionEnabled: o.VoiceTranscriptionEnabled,
		Alias:    (Alias)(o),
	})
}

func (o *Transcriptiontopictranscriptresult) UnmarshalJSON(b []byte) error {
	var TranscriptiontopictranscriptresultMap map[string]interface{}
	err := json.Unmarshal(b, &TranscriptiontopictranscriptresultMap)
	if err != nil {
		return err
	}
	
	if UtteranceId, ok := TranscriptiontopictranscriptresultMap["utteranceId"].(string); ok {
		o.UtteranceId = &UtteranceId
	}
    
	if IsFinal, ok := TranscriptiontopictranscriptresultMap["isFinal"].(bool); ok {
		o.IsFinal = &IsFinal
	}
    
	if Channel, ok := TranscriptiontopictranscriptresultMap["channel"].(string); ok {
		o.Channel = &Channel
	}
    
	if Alternatives, ok := TranscriptiontopictranscriptresultMap["alternatives"].([]interface{}); ok {
		AlternativesString, _ := json.Marshal(Alternatives)
		json.Unmarshal(AlternativesString, &o.Alternatives)
	}
	
	if AgentAssistantId, ok := TranscriptiontopictranscriptresultMap["agentAssistantId"].(string); ok {
		o.AgentAssistantId = &AgentAssistantId
	}
    
	if EngineId, ok := TranscriptiontopictranscriptresultMap["engineId"].(string); ok {
		o.EngineId = &EngineId
	}
    
	if Dialect, ok := TranscriptiontopictranscriptresultMap["dialect"].(string); ok {
		o.Dialect = &Dialect
	}
    
	if SpeechTextAnalyticsProgramId, ok := TranscriptiontopictranscriptresultMap["speechTextAnalyticsProgramId"].(string); ok {
		o.SpeechTextAnalyticsProgramId = &SpeechTextAnalyticsProgramId
	}
    
	if AgentAssistEnabled, ok := TranscriptiontopictranscriptresultMap["agentAssistEnabled"].(bool); ok {
		o.AgentAssistEnabled = &AgentAssistEnabled
	}
    
	if VoiceTranscriptionEnabled, ok := TranscriptiontopictranscriptresultMap["voiceTranscriptionEnabled"].(bool); ok {
		o.VoiceTranscriptionEnabled = &VoiceTranscriptionEnabled
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Transcriptiontopictranscriptresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
