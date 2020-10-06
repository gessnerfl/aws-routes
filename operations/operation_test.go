package operations_test

import (
	"fmt"
	"testing"

	. "github.com/gessnerfl/awsroutes/operations"
	"github.com/stretchr/testify/require"
)

func TestShouldReturnOperationByName(t *testing.T) {
	names := []string{"add", "remove"}
	for _, name := range names {
		t.Run(fmt.Sprintf("TestShouldReturnOperationByName_%s", name), createTestFofTestShouldReturnOperationByName(name))
	}
}

func createTestFofTestShouldReturnOperationByName(name string) func(*testing.T) {
	return func(t *testing.T) {
		operation, err := SupportedOperations.ByName(name)

		require.NoError(t, err)
		require.NotNil(t, operation)
	}
}

func TestShouldReturnErrorWhenOperationDoesNotExistForTheGivenName(t *testing.T) {
	operation, err := SupportedOperations.ByName("invalid")

	require.Error(t, err)
	require.Nil(t, operation)
}
