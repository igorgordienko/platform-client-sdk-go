package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Attendancestatus
type Attendancestatus struct { 
	// DateWorkday - the workday date of this attendance status. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateWorkday *time.Time `json:"dateWorkday,omitempty"`


	// AttendanceStatusType - the attendance status
	AttendanceStatusType *string `json:"attendanceStatusType,omitempty"`

}

// String returns a JSON representation of the model
func (o *Attendancestatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
