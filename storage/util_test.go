package storage_test

import (
	"strings"
	"testing"

	. "github.com/gessnerfl/awsroutes/storage"
	"github.com/stretchr/testify/require"
)

func TestShouldReturnUserHomeWithoutTrailingSlash(t *testing.T) {
	result := GetUserHome()

	require.DirExists(t, result)
	require.Condition(t, func() bool { return !strings.HasSuffix(result, "/") })
	require.NotEqual(t, "/tmp", result)
}
