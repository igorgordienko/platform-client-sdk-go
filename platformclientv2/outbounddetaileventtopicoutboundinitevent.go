package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Outbounddetaileventtopicoutboundinitevent
type Outbounddetaileventtopicoutboundinitevent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EventTime
	EventTime *int `json:"eventTime,omitempty"`

	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`

	// ParticipantId
	ParticipantId *string `json:"participantId,omitempty"`

	// SessionId
	SessionId *string `json:"sessionId,omitempty"`

	// MediaType
	MediaType *string `json:"mediaType,omitempty"`

	// Provider
	Provider *string `json:"provider,omitempty"`

	// Direction
	Direction *string `json:"direction,omitempty"`

	// Ani
	Ani *string `json:"ani,omitempty"`

	// Dnis
	Dnis *string `json:"dnis,omitempty"`

	// AddressTo
	AddressTo *string `json:"addressTo,omitempty"`

	// AddressFrom
	AddressFrom *string `json:"addressFrom,omitempty"`

	// Subject
	Subject *string `json:"subject,omitempty"`

	// MessageType
	MessageType *string `json:"messageType,omitempty"`

	// OutboundCampaignId
	OutboundCampaignId *string `json:"outboundCampaignId,omitempty"`

	// DivisionId
	DivisionId *string `json:"divisionId,omitempty"`

	// OutboundContactListId
	OutboundContactListId *string `json:"outboundContactListId,omitempty"`

	// OutboundContactId
	OutboundContactId *string `json:"outboundContactId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Outbounddetaileventtopicoutboundinitevent) SetField(field string, fieldValue interface{}) {
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

func (o Outbounddetaileventtopicoutboundinitevent) MarshalJSON() ([]byte, error) {
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
	type Alias Outbounddetaileventtopicoutboundinitevent
	
	return json.Marshal(&struct { 
		EventTime *int `json:"eventTime,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		ParticipantId *string `json:"participantId,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		Ani *string `json:"ani,omitempty"`
		
		Dnis *string `json:"dnis,omitempty"`
		
		AddressTo *string `json:"addressTo,omitempty"`
		
		AddressFrom *string `json:"addressFrom,omitempty"`
		
		Subject *string `json:"subject,omitempty"`
		
		MessageType *string `json:"messageType,omitempty"`
		
		OutboundCampaignId *string `json:"outboundCampaignId,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		OutboundContactListId *string `json:"outboundContactListId,omitempty"`
		
		OutboundContactId *string `json:"outboundContactId,omitempty"`
		Alias
	}{ 
		EventTime: o.EventTime,
		
		ConversationId: o.ConversationId,
		
		ParticipantId: o.ParticipantId,
		
		SessionId: o.SessionId,
		
		MediaType: o.MediaType,
		
		Provider: o.Provider,
		
		Direction: o.Direction,
		
		Ani: o.Ani,
		
		Dnis: o.Dnis,
		
		AddressTo: o.AddressTo,
		
		AddressFrom: o.AddressFrom,
		
		Subject: o.Subject,
		
		MessageType: o.MessageType,
		
		OutboundCampaignId: o.OutboundCampaignId,
		
		DivisionId: o.DivisionId,
		
		OutboundContactListId: o.OutboundContactListId,
		
		OutboundContactId: o.OutboundContactId,
		Alias:    (Alias)(o),
	})
}

func (o *Outbounddetaileventtopicoutboundinitevent) UnmarshalJSON(b []byte) error {
	var OutbounddetaileventtopicoutboundiniteventMap map[string]interface{}
	err := json.Unmarshal(b, &OutbounddetaileventtopicoutboundiniteventMap)
	if err != nil {
		return err
	}
	
	if EventTime, ok := OutbounddetaileventtopicoutboundiniteventMap["eventTime"].(float64); ok {
		EventTimeInt := int(EventTime)
		o.EventTime = &EventTimeInt
	}
	
	if ConversationId, ok := OutbounddetaileventtopicoutboundiniteventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if ParticipantId, ok := OutbounddetaileventtopicoutboundiniteventMap["participantId"].(string); ok {
		o.ParticipantId = &ParticipantId
	}
    
	if SessionId, ok := OutbounddetaileventtopicoutboundiniteventMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if MediaType, ok := OutbounddetaileventtopicoutboundiniteventMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if Provider, ok := OutbounddetaileventtopicoutboundiniteventMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if Direction, ok := OutbounddetaileventtopicoutboundiniteventMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if Ani, ok := OutbounddetaileventtopicoutboundiniteventMap["ani"].(string); ok {
		o.Ani = &Ani
	}
    
	if Dnis, ok := OutbounddetaileventtopicoutboundiniteventMap["dnis"].(string); ok {
		o.Dnis = &Dnis
	}
    
	if AddressTo, ok := OutbounddetaileventtopicoutboundiniteventMap["addressTo"].(string); ok {
		o.AddressTo = &AddressTo
	}
    
	if AddressFrom, ok := OutbounddetaileventtopicoutboundiniteventMap["addressFrom"].(string); ok {
		o.AddressFrom = &AddressFrom
	}
    
	if Subject, ok := OutbounddetaileventtopicoutboundiniteventMap["subject"].(string); ok {
		o.Subject = &Subject
	}
    
	if MessageType, ok := OutbounddetaileventtopicoutboundiniteventMap["messageType"].(string); ok {
		o.MessageType = &MessageType
	}
    
	if OutboundCampaignId, ok := OutbounddetaileventtopicoutboundiniteventMap["outboundCampaignId"].(string); ok {
		o.OutboundCampaignId = &OutboundCampaignId
	}
    
	if DivisionId, ok := OutbounddetaileventtopicoutboundiniteventMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
    
	if OutboundContactListId, ok := OutbounddetaileventtopicoutboundiniteventMap["outboundContactListId"].(string); ok {
		o.OutboundContactListId = &OutboundContactListId
	}
    
	if OutboundContactId, ok := OutbounddetaileventtopicoutboundiniteventMap["outboundContactId"].(string); ok {
		o.OutboundContactId = &OutboundContactId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Outbounddetaileventtopicoutboundinitevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
