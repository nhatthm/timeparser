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
func ParsePeriod(from, to string) (start *time.Time, end *time.Time, err error) {
	if from == "" && to == "" {
		return nil, nil, nil
	}

	if from != "" {
		s, err := Parse(from)
		if err != nil {
			return nil, nil, err
		}

		start = &s
	}

	if to != "" {
		e, err := Parse(to)
		if err != nil {
			return nil, nil, err
		}

		end = &e
	}

	if start != nil && end != nil && start.After(*end) {
		return nil, nil, ErrInvalidTimePeriod
	}

	return start, end, nil
}

// TimePtr returns a pointer to the given time.Time.
func TimePtr(t time.Time) *time.Time {
	return &t
}
