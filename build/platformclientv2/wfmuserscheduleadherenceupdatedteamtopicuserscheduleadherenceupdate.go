package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmuserscheduleadherenceupdatedteamtopicuserscheduleadherenceupdate
type Wfmuserscheduleadherenceupdatedteamtopicuserscheduleadherenceupdate struct { 
	// User
	User *Wfmuserscheduleadherenceupdatedteamtopicuserreference `json:"user,omitempty"`


	// ManagementUnitId
	ManagementUnitId *string `json:"managementUnitId,omitempty"`


	// Team
	Team *Wfmuserscheduleadherenceupdatedteamtopicurireference `json:"team,omitempty"`


	// ScheduledActivityCategory
	ScheduledActivityCategory *string `json:"scheduledActivityCategory,omitempty"`


	// SystemPresence
	SystemPresence *string `json:"systemPresence,omitempty"`


	// OrganizationSecondaryPresenceId
	OrganizationSecondaryPresenceId *string `json:"organizationSecondaryPresenceId,omitempty"`


	// RoutingStatus
	RoutingStatus *string `json:"routingStatus,omitempty"`


	// ActualActivityCategory
	ActualActivityCategory *string `json:"actualActivityCategory,omitempty"`


	// IsOutOfOffice
	IsOutOfOffice *bool `json:"isOutOfOffice,omitempty"`


	// AdherenceState
	AdherenceState *string `json:"adherenceState,omitempty"`


	// Impact
	Impact *string `json:"impact,omitempty"`


	// AdherenceChangeTime
	AdherenceChangeTime *time.Time `json:"adherenceChangeTime,omitempty"`


	// PresenceUpdateTime
	PresenceUpdateTime *time.Time `json:"presenceUpdateTime,omitempty"`


	// ActiveQueues
	ActiveQueues *[]Wfmuserscheduleadherenceupdatedteamtopicqueuereference `json:"activeQueues,omitempty"`


	// ActiveQueuesModifiedTime
	ActiveQueuesModifiedTime *time.Time `json:"activeQueuesModifiedTime,omitempty"`


	// RemovedFromManagementUnit
	RemovedFromManagementUnit *bool `json:"removedFromManagementUnit,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmuserscheduleadherenceupdatedteamtopicuserscheduleadherenceupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
