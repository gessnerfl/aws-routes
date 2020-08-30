package routes_test

import (
	"io/ioutil"
	"testing"

	. "github.com/gessnerfl/awsroutes/routes"
	"github.com/stretchr/testify/assert"
)

func TestShouldDownloadFileFromTheAPI(t *testing.T) {
	api := NewAWSIPRangeAPI()

	data, err := api.Download()

	assert.Nil(t, err)
	assert.NotNil(t, data)
}

func TestShouldUnmarshalJsonMessage(t *testing.T) {
	api := NewAWSIPRangeAPI()
	expectedResult := AWSIPAddressRanges{
		SyncToken:  "1598568675",
		CreateDate: "2020-08-27-22-51-15",
		Prefixes: []AWSIPPrefix{
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

	result, err := api.Unmarshal(input)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, &expectedResult, result)
}

func TestShouldFailToUnmarshalJsonMessageWhenMessageIsNotAValidJSON(t *testing.T) {
	api := NewAWSIPRangeAPI()
	expectedResult := AWSIPAddressRanges{}
	input := []byte("inalid data")

	result, err := api.Unmarshal(input)

	assert.NotNil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, &expectedResult, result)
}
