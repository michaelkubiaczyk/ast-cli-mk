//go:build integration

package integration

import (
	"testing"

	"github.com/checkmarx/ast-cli/internal/params"
)

const (
	maskCommand           = "mask"
	resultsFileValue      = "data/package.json"
	resultFileNonExisting = "data/package.jso"
)

func TestMaskSecrets(t *testing.T) {
	executeCmdNilAssertion(
		t,
		"Remediating kics result",
		utilsCommand,
		maskCommand,
		flag(params.ChatResultFile),
		resultsFileValue,
	)
}

func TestFailedMaskSecrets(t *testing.T) {
	args := []string{
		utilsCommand,
		maskCommand,
		flag(params.ChatResultFile),
		resultFileNonExisting,
	}
	err, _ := executeCommand(t, args...)
	assertError(t, err, "Error opening file : open data/package.jso: no such file or directory")
}
