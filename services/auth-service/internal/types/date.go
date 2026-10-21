package types

import (
	"strings"
	"time"
)

type DateOnly struct {
	time.Time
}

// UnmarshalJSON parses "YYYY-MM-DD" date strings from JSON
func (d *DateOnly) UnmarshalJSON(b []byte) error {
	// Remove surrounding quotes from the JSON string
	s := strings.Trim(string(b), `"`)
	
	if s == "null" || s == "" {
		return nil
	}

	parsed, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}

	d.Time = parsed
	return nil
}

// MarshalJSON formats the date back to "YYYY-MM-DD" when serializing to JSON
func (d DateOnly) MarshalJSON() ([]byte, error) {
	if d.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + d.Format("2006-01-02") + `"`), nil
}