package must_test

import (
	"testing"
	"time"

	"github.com/dohernandez/form3-service/pkg/must"
	"github.com/stretchr/testify/assert"
)

func TestParseTime(t *testing.T) {
	assert.NotPanics(t, func() {
		assert.Equal(t, "2018-01-02T00:00:00Z", must.ParseTime("2018-01-02T00:00:00Z").Format(time.RFC3339))
	})

	assert.NotPanics(t, func() {
		assert.Equal(t, "2018-01-02T00:00:00Z", must.ParseTime("01/02/2018", "01/02/2006").Format(time.RFC3339))
	})

	assert.Panics(t, func() {
		must.ParseTime("wat")
	})
}
