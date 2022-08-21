package timeparser_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/timeparser"
)

func TestParse(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		input          string
		expectedResult time.Time
		expectedError  string
	}{
		{
			scenario:      "rfc3339 - error",
			input:         "2020-01-02T03-",
			expectedError: `parsing time "2020-01-02T03-" as "2006-01-02T15:04:05Z07:00": cannot parse "-" as ":"`,
		},
		{
			scenario:       "rfc3339 - success",
			input:          "2020-01-02T03:04:05.000Z",
			expectedResult: time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC),
		},
		{
			scenario:      "ymd - error",
			input:         "2020-01-0",
			expectedError: `parsing time "2020-01-0" as "2006-01-02": cannot parse "0" as "02"`,
		},
		{
			scenario:       "ymd - success",
			input:          "2020-01-02",
			expectedResult: time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			result, err := timeparser.Parse(tc.input)

			assert.Equal(t, tc.expectedResult, result)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestParsePeriod(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario      string
		from          string
		to            string
		expectedStart time.Time
		expectedEnd   time.Time
		expectedError string
	}{
		{
			scenario:      "invalid from",
			from:          "foobar",
			expectedError: `parsing time "foobar" as "2006-01-02": cannot parse "foobar" as "2006"`,
		},
		{
			scenario:      "from is empty",
			expectedError: `parsing time "" as "2006-01-02": cannot parse "" as "2006"`,
		},
		{
			scenario:      "invalid to",
			from:          "2020-01-02T03:04:05Z",
			to:            "foobar",
			expectedError: `parsing time "foobar" as "2006-01-02": cannot parse "foobar" as "2006"`,
		},
		{
			scenario:      "to is empty",
			from:          "2020-01-02T03:04:05Z",
			expectedError: `parsing time "" as "2006-01-02": cannot parse "" as "2006"`,
		},
		{
			scenario:      "from and to are not empty and valid",
			from:          "2020-01-02T03:04:05Z",
			to:            "2020-02-02T03:04:05Z",
			expectedStart: time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC),
			expectedEnd:   time.Date(2020, 2, 2, 3, 4, 5, 0, time.UTC),
		},
		{
			scenario:      "from and to are not empty and invalid",
			from:          "2020-02-02T03:04:05Z",
			to:            "2020-01-02T03:04:05Z",
			expectedError: `invalid time period`,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			start, end, err := timeparser.ParsePeriod(tc.from, tc.to)

			assert.Equal(t, tc.expectedStart, start)
			assert.Equal(t, tc.expectedEnd, end)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestTimePtr(t *testing.T) {
	t.Parallel()

	ts := time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC)

	assert.Equal(t, &ts, timeparser.TimePtr(ts))
}
