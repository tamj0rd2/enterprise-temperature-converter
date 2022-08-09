package acceptance

import (
	"github.com/alecthomas/assert/v2"
	"testing"
)

type Converter interface {
	FromFToC(i int) (int, error)
}

// TODO: JIRA TICKET: Convert from C to F
func TestConvertingTemperatures(t *testing.T) {
	var (
		expectedC = 0
	)

	// arrange
	var converter Converter
	converter = newCliConverterDriver()

	// act
	c, err := converter.FromFToC(32)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, expectedC, c)
}

type cliConverterDriver struct{}

func newCliConverterDriver() *cliConverterDriver {
	return &cliConverterDriver{}
}

func (c cliConverterDriver) FromFToC(i int) (int, error) {
	return 0, nil
}
