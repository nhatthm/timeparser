package timeparser

import (
	"errors"
	"time"
)

const (
	layoutRFC3339 = time.RFC3339
	layoutYMD     = "2006-01-02"
)

// ErrInvalidTimePeriod indicates that the time input is invalid.
var ErrInvalidTimePeriod = errors.New("invalid time period")

// Parse parses string into time.Time.
func Parse(s string) (time.Time, error) {
	if len(s) > 10 {
		return time.Parse(layoutRFC3339, s)
	}

	return time.Parse(layoutYMD, s)
}

// ParsePeriod returns a time period from given input.
func ParsePeriod(from, to string) (time.Time, time.Time, error) {
	s, err := Parse(from)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	e, err := Parse(to)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	if s.After(e) {
		return time.Time{}, time.Time{}, ErrInvalidTimePeriod
	}

	return s, e, nil
}

// TimePtr returns a pointer to the given time.Time.
func TimePtr(t time.Time) *time.Time {
	return &t
}
