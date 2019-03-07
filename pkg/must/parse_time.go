package must

import (
	"time"
)

// ParseTime parses time with layout(s) and panics on all errors
func ParseTime(value string, layouts ...string) time.Time {
	if len(layouts) == 0 {
		layouts = []string{time.RFC3339, time.RFC3339Nano, "2006-01-02", "2006-01-02 15:04:05"}
	}
	var (
		t   time.Time
		err error
	)
	for _, layout := range layouts {
		t, err = time.Parse(layout, value)
		if err == nil {
			break
		}
	}
	if err != nil {
		panic(err)
	}
	return t
}
