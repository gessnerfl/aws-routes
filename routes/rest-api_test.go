package routes_test

import (
	"testing"

	. "github.com/gessnerfl/awsroutes/routes"
	"github.com/stretchr/testify/assert"
)

func TestShouldDownloadFileFromTheAPI(t *testing.T) {
	api := NewAwsIPRangeRestAPI()

	data, err := api.Download()

	assert.Nil(t, err)
	assert.NotNil(t, data)
}
