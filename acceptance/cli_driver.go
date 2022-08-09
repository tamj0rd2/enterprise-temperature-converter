package acceptance

import (
	"bytes"
	"fmt"
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
	tb.Helper()
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

func (d cliConverterDriver) FromFToC(inputF float64) (float64, error) {
	var stdoutBuf bytes.Buffer
	var stderrBuf bytes.Buffer

	cmd := exec.Command(d.binaryPath)
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	w, err := cmd.StdinPipe()
	if err != nil {
		return 0, err
	}

	if err := cmd.Start(); err != nil {
		return 0, fmt.Errorf("failed to run cli - %w: %s", err, stderrBuf.String())
	}

	if _, err := fmt.Fprintf(w, "%g\n", inputF); err != nil {
		return 0, fmt.Errorf("failed to write to stdin - %w: %s", err, stderrBuf.String())
	}

	if err := cmd.Wait(); err != nil {
		return 0, fmt.Errorf("failed to run cli - %w: %s", err, stderrBuf.String())
	}

	output := stdoutBuf.String()
	c, err := strconv.ParseFloat(output, 64)
	if err != nil {
		panic(err)
	}

	return c, nil
}
