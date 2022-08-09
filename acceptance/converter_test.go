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
	t.Run("Converging from F to C", func(t *testing.T) {
		var (
			expectedC = 0
		)

		// arrange
		var converter Converter
		converter = newCliConverterDriver(t)

		// act
		c, err := converter.FromFToC(32)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, expectedC, c)
	})
}
