package routes_test

import (
	"io/ioutil"
	"testing"

	. "github.com/gessnerfl/awsroutes/routes"
	"github.com/stretchr/testify/assert"
)

func TestShouldUnmarshalJsonMessage(t *testing.T) {
	unmarshaller := NewUnmarshaller()
	expectedResult := IPAddressRanges{
		SyncToken:  "1598568675",
		CreateDate: "2020-08-27-22-51-15",
		Prefixes: []IPPrefix{
			{
				IPPrefix: "35.180.0.0/16",
				Region:   "eu-west-3",
				Service:  "AMAZON",
			},
			{
				IPPrefix: "52.93.178.234/32",
				Region:   "us-west-1",
				Service:  "AMAZON",
			},
		},
	}
	input, err := ioutil.ReadFile("test-file.json")
	assert.Nil(t, err)
	assert.NotNil(t, input)

	result, err := unmarshaller.Unmarshal(input)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, &expectedResult, result)
}

func TestShouldFailToUnmarshalJsonMessageWhenMessageIsNotAValidJSON(t *testing.T) {
	unmarshaller := NewUnmarshaller()
	expectedResult := IPAddressRanges{}
	input := []byte("inalid data")

	result, err := unmarshaller.Unmarshal(input)

	assert.NotNil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, &expectedResult, result)
}
