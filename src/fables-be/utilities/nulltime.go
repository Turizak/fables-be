package utilities

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/lib/pq"
)

type NullTime struct {
	Time  time.Time `json:"time"`
	Valid bool      `json:"valid"`
}

// Custom MarshalJSON method to handle pq.NullTime
func (nt NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte(`null`), nil
	}
	return json.Marshal(map[string]interface{}{
		"time":  nt.Time.UTC(), // Ensure the time is marshaled in UTC
		"valid": nt.Valid,
	})
}

// UnmarshalJSON handles JSON string format for NullTime
func (nt *NullTime) UnmarshalJSON(data []byte) error {
	var t string
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}

	// Try to parse the string into time.Time in UTC
	parsedTime, err := time.Parse(time.RFC3339, t)
	if err != nil {
		nt.Valid = false
		return nil // if parsing fails, set Valid to false
	}

	nt.Time = parsedTime.UTC() // Set parsed time to UTC
	nt.Valid = true
	return nil
}

// Convert pq.NullTime to custom NullTime
func ToNullTime(pqTime pq.NullTime) NullTime {
	return NullTime{
		Time:  pqTime.Time.UTC(), // Ensure conversion to UTC
		Valid: pqTime.Valid,
	}
}

// Value implements the driver.Valuer interface for NullTime
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time.UTC(), nil // Store the time in UTC in the database
}

// Scan implements the sql.Scanner interface for NullTime
func (nt *NullTime) Scan(value interface{}) error {
	if value == nil {
		nt.Time, nt.Valid = time.Time{}, false
		return nil
	}

	nt.Valid = true
	switch v := value.(type) {
	case time.Time:
		nt.Time = v.UTC() // Ensure scanned time is set to UTC
	default:
		return fmt.Errorf("cannot convert %v to time.Time", value)
	}
	return nil
}
