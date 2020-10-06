package routes_test

import (
	"testing"

	. "github.com/gessnerfl/awsroutes/routes"
	"github.com/stretchr/testify/require"
)

func TestShouldReturnTrueWhenWhenServiceOfIPPrefixIsRelevantForRouting(t *testing.T) {
	sut := IPPrefix{
		IPPrefix: "ip-prefix",
		Service:  "AMAZON",
		Region:   "us-east-1",
	}

	require.True(t, sut.IsRelevantService())
}

func TestShouldReturnFalseWhenWhenServiceOfIPPrefixIsNotRelevantForRouting(t *testing.T) {
	sut := IPPrefix{
		IPPrefix: "ip-prefix",
		Service:  "foo",
		Region:   "us-east-1",
	}

	require.False(t, sut.IsRelevantService())
}
