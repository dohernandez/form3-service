package internal

import (
	"regexp"
	"time"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/dohernandez/form3-service/pkg/test/feature"
)

// RestContext struct to hold necessary dependencies
type RestContext struct {
	*feature.RestContext
}

// RegisterRestContext parser the response
func RegisterRestContext(s *godog.Suite, upstream *feature.RestContext) *RestContext {
	c := RestContext{
		RestContext: upstream,
	}

	// AfterStep is used to sleep the test before check into the database,
	// otherwise the corresponding message hasn't been dispatched when the test run check against the table
	// therefore the projection is not found.
	s.AfterStep(func(step *gherkin.Step, e error) {
		if e != nil {
			return
		}

		re := regexp.MustCompile(`^I request REST endpoint with method "([^"]*)" and path "([^"]*)".*$`)
		if re.Match([]byte(step.Text)) {
			time.Sleep(1 * time.Second)
		}
	})

	return &c
}
