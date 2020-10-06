package operations_test

import (
	"testing"

	. "github.com/gessnerfl/awsroutes/operations"
	"github.com/stretchr/testify/require"
)

func TestShouldSuccessfullyExecuteCommandWithArguments(t *testing.T) {
	sut := NewNativeExecutor()

	result, err := sut.Execute("echo", "-n", "test")

	require.NoError(t, err)
	require.Equal(t, "test", result)
}

func TestShouldReturnErrorWhenCommandDoesNotExist(t *testing.T) {
	sut := NewNativeExecutor()

	_, err := sut.Execute("invalidCommand", "-n")

	require.Error(t, err)
}

func TestShouldReturnErrorWhenCommandArgumentIsNotValid(t *testing.T) {
	sut := NewNativeExecutor()

	_, err := sut.Execute("ls", "-alF", "/invalid/path")

	require.Error(t, err)
}
