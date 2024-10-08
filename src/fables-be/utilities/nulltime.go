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
		"time":  nt.Time,
		"valid": nt.Valid,
	})
}

// Convert pq.NullTime to custom NullTime
func ToNullTime(pqTime pq.NullTime) NullTime {
	return NullTime{
		Time:  pqTime.Time,
		Valid: pqTime.Valid,
	}
}

// Value implements the driver.Valuer interface for NullTime
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
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
		nt.Time = v
	default:
		return fmt.Errorf("cannot convert %v to time.Time", value)
	}
	return nil
}
