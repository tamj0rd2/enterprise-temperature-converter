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

func (d cliConverterDriver) FromFToC(f float64) (float64, error) {
	var buf bytes.Buffer

	cmd := exec.Command(d.binaryPath)
	cmd.Stdout = &buf

	if err := cmd.Run(); err != nil {
		panic(err)
	}

	output := buf.String()
	c, err := strconv.ParseFloat(output, 64)
	if err != nil {
		panic(err)
	}

	return c, nil
}
