package acceptance

import (
	"github.com/alecthomas/assert/v2"
	"testing"
)

type Converter interface {
	FromFToC(f float64) (float64, error)
}

// TODO: JIRA TICKET: Convert from C to F
func TestConvertingTemperatures(t *testing.T) {
	t.Run("Converting from F to C", func(t *testing.T) {
		// arrange
		var converter Converter
		converter = newCliConverterDriver(t)

		// act
		c, err := converter.FromFToC(32)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 0, c)
	})

	//t.Run("Converting from F to C with another input", func(t *testing.T) {
	//	// arrange
	//	var converter Converter
	//	converter = newCliConverterDriver(t)
	//
	//	// act
	//	c, err := converter.FromFToC(90)
	//
	//	// assert
	//	assert.NoError(t, err)
	//	assert.Equal(t, 32.2222, c)
	//})
}
