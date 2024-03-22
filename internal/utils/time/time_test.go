package time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseDecimalString(t *testing.T) {
	tests := []struct {
		name          string
		val           string
		unit          time.Duration
		expectedValue time.Duration
		shouldError   bool
	}{
		{
			val:           "200s",
			unit:          time.Second,
			expectedValue: 0,
			shouldError:   true,
		},
		{
			val:           "200",
			unit:          time.Second,
			expectedValue: 200 * time.Second,
			shouldError:   false,
		},
		{
			val:           "0",
			unit:          time.Second,
			expectedValue: 0,
			shouldError:   false,
		},
		{
			val:           "-90",
			unit:          time.Second,
			expectedValue: -90 * time.Second,
			shouldError:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.val, func(t *testing.T) {
			val, err := ParseDecimalString(test.val, test.unit)
			assert.Equal(t, test.expectedValue, val, "Wrong time value")
			if test.shouldError {
				assert.Error(t, err, "Should be an error")
			} else {
				assert.Equal(t, nil, err, "Should not be an error")
			}
		})
	}
}
