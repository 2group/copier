package copier

import (
	"strings"
	"time"
)

const isoDateLayout = "2006-01-02"

func isDateField(fieldName string) bool {
	// Go struct field names: TestDate, BirthDate, StartDate, EndDate...
	// Also support *_date style if you ever map tags/names
	n := strings.ToLower(fieldName)
	return strings.HasSuffix(n, "date") || strings.HasSuffix(n, "_date")
}

func parseISODateString(s string) (*time.Time, error) {
	if s == "" {
		return nil, nil
	}
	t, err := time.Parse(isoDateLayout, s)
	if err != nil {
		return nil, err
	}
	tt := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	return &tt, nil
}

func formatISODate(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.In(time.UTC).Format(isoDateLayout)
}
