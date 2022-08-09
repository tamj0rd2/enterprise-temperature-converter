package acceptance

import (
	"bytes"
	"github.com/alecthomas/assert/v2"
	"os"
	"os/exec"
	"path"
	"strconv"
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
	converter = newCliConverterDriver(t)

	// act
	c, err := converter.FromFToC(32)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, expectedC, c)
}

type cliConverterDriver struct {
	binaryPath string
}

func newCliConverterDriver(tb testing.TB) *cliConverterDriver {
	dir, err := os.MkdirTemp("", "converter-*")
	assert.NoError(tb, err)
	tb.Cleanup(func() {
		_ = os.RemoveAll(dir)
	})
	binaryPath := path.Join(dir, "converter-cli")
	tb.Log(binaryPath)

	b, err := exec.Command("go", "build", "-o", binaryPath, "../cmd/converter-cli/...").CombinedOutput()
	output := string(b)
	assert.NoError(tb, err, output)

	return &cliConverterDriver{binaryPath: binaryPath}
}

func (d cliConverterDriver) FromFToC(i int) (int, error) {
	var buf bytes.Buffer

	cmd := exec.Command(d.binaryPath)
	cmd.Stdout = &buf

	if err := cmd.Run(); err != nil {
		panic(err)
	}

	output := buf.String()
	c, err := strconv.Atoi(output)
	if err != nil {
		panic(err)
	}
	
	return c, nil
}
