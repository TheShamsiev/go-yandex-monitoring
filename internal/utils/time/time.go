package time

import (
	"strconv"
	"time"
)

func ParseDecimalString(s string, unit time.Duration) (time.Duration, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return time.Duration(n) * unit, nil
}
