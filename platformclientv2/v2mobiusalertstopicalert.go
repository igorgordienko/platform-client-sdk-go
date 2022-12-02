package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2mobiusalertstopicalert
type V2mobiusalertstopicalert struct { 
	// Rule
	Rule *V2mobiusalertstopicalertruleproperties `json:"rule,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`


	// UserId
	UserId *string `json:"userId,omitempty"`


	// Notifications
	Notifications *[]V2mobiusalertstopicalertnotification `json:"notifications,omitempty"`


	// DateStart
	DateStart *time.Time `json:"dateStart,omitempty"`


	// DateEnd
	DateEnd *time.Time `json:"dateEnd,omitempty"`


	// Conditions
	Conditions *V2mobiusalertstopiccondition `json:"conditions,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]string `json:"additionalProperties,omitempty"`


	// Active
	Active *bool `json:"active,omitempty"`


	// Unread
	Unread *bool `json:"unread,omitempty"`


	// Muted
	Muted *bool `json:"muted,omitempty"`


	// Snoozed
	Snoozed *bool `json:"snoozed,omitempty"`


	// DateMutedUntil
	DateMutedUntil *time.Time `json:"dateMutedUntil,omitempty"`


	// DateSnoozedUntil
	DateSnoozedUntil *time.Time `json:"dateSnoozedUntil,omitempty"`


	// Action
	Action *string `json:"action,omitempty"`

}

func (o *V2mobiusalertstopicalert) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2mobiusalertstopicalert
	
	DateStart := new(string)
	if o.DateStart != nil {
		
		*DateStart = timeutil.Strftime(o.DateStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStart = nil
	}
	
	DateEnd := new(string)
	if o.DateEnd != nil {
		
		*DateEnd = timeutil.Strftime(o.DateEnd, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateEnd = nil
	}
	
	DateMutedUntil := new(string)
	if o.DateMutedUntil != nil {
		
		*DateMutedUntil = timeutil.Strftime(o.DateMutedUntil, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateMutedUntil = nil
	}
	
	DateSnoozedUntil := new(string)
	if o.DateSnoozedUntil != nil {
		
		*DateSnoozedUntil = timeutil.Strftime(o.DateSnoozedUntil, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateSnoozedUntil = nil
	}
	
	return json.Marshal(&struct { 
		Rule *V2mobiusalertstopicalertruleproperties `json:"rule,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		Notifications *[]V2mobiusalertstopicalertnotification `json:"notifications,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		
		DateEnd *string `json:"dateEnd,omitempty"`
		
		Conditions *V2mobiusalertstopiccondition `json:"conditions,omitempty"`
		
		AdditionalProperties *map[string]string `json:"additionalProperties,omitempty"`
		
		Active *bool `json:"active,omitempty"`
		
		Unread *bool `json:"unread,omitempty"`
		
		Muted *bool `json:"muted,omitempty"`
		
		Snoozed *bool `json:"snoozed,omitempty"`
		
		DateMutedUntil *string `json:"dateMutedUntil,omitempty"`
		
		DateSnoozedUntil *string `json:"dateSnoozedUntil,omitempty"`
		
		Action *string `json:"action,omitempty"`
		*Alias
	}{ 
		Rule: o.Rule,
		
		Id: o.Id,
		
		UserId: o.UserId,
		
		Notifications: o.Notifications,
		
		DateStart: DateStart,
		
		DateEnd: DateEnd,
		
		Conditions: o.Conditions,
		
		AdditionalProperties: o.AdditionalProperties,
		
		Active: o.Active,
		
		Unread: o.Unread,
		
		Muted: o.Muted,
		
		Snoozed: o.Snoozed,
		
		DateMutedUntil: DateMutedUntil,
		
		DateSnoozedUntil: DateSnoozedUntil,
		
		Action: o.Action,
		Alias:    (*Alias)(o),
	})
}

func (o *V2mobiusalertstopicalert) UnmarshalJSON(b []byte) error {
	var V2mobiusalertstopicalertMap map[string]interface{}
	err := json.Unmarshal(b, &V2mobiusalertstopicalertMap)
	if err != nil {
		return err
	}
	
	if Rule, ok := V2mobiusalertstopicalertMap["rule"].(map[string]interface{}); ok {
		RuleString, _ := json.Marshal(Rule)
		json.Unmarshal(RuleString, &o.Rule)
	}
	
	if Id, ok := V2mobiusalertstopicalertMap["id"].(map[string]interface{}); ok {
		IdString, _ := json.Marshal(Id)
		json.Unmarshal(IdString, &o.Id)
	}
	
	if UserId, ok := V2mobiusalertstopicalertMap["userId"].(map[string]interface{}); ok {
		UserIdString, _ := json.Marshal(UserId)
		json.Unmarshal(UserIdString, &o.UserId)
	}
	
	if Notifications, ok := V2mobiusalertstopicalertMap["notifications"].([]interface{}); ok {
		NotificationsString, _ := json.Marshal(Notifications)
		json.Unmarshal(NotificationsString, &o.Notifications)
	}
	
	if dateStartString, ok := V2mobiusalertstopicalertMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartString)
		o.DateStart = &DateStart
	}
	
	if dateEndString, ok := V2mobiusalertstopicalertMap["dateEnd"].(string); ok {
		DateEnd, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateEndString)
		o.DateEnd = &DateEnd
	}
	
	if Conditions, ok := V2mobiusalertstopicalertMap["conditions"].(map[string]interface{}); ok {
		ConditionsString, _ := json.Marshal(Conditions)
		json.Unmarshal(ConditionsString, &o.Conditions)
	}
	
	if AdditionalProperties, ok := V2mobiusalertstopicalertMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	
	if Active, ok := V2mobiusalertstopicalertMap["active"].(bool); ok {
		o.Active = &Active
	}
    
	if Unread, ok := V2mobiusalertstopicalertMap["unread"].(bool); ok {
		o.Unread = &Unread
	}
    
	if Muted, ok := V2mobiusalertstopicalertMap["muted"].(bool); ok {
		o.Muted = &Muted
	}
    
	if Snoozed, ok := V2mobiusalertstopicalertMap["snoozed"].(bool); ok {
		o.Snoozed = &Snoozed
	}
    
	if dateMutedUntilString, ok := V2mobiusalertstopicalertMap["dateMutedUntil"].(string); ok {
		DateMutedUntil, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateMutedUntilString)
		o.DateMutedUntil = &DateMutedUntil
	}
	
	if dateSnoozedUntilString, ok := V2mobiusalertstopicalertMap["dateSnoozedUntil"].(string); ok {
		DateSnoozedUntil, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateSnoozedUntilString)
		o.DateSnoozedUntil = &DateSnoozedUntil
	}
	
	if Action, ok := V2mobiusalertstopicalertMap["action"].(string); ok {
		o.Action = &Action
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *V2mobiusalertstopicalert) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
