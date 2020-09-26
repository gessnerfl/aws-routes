package routes

import (
	"bytes"
	"net/http"
)

const url = "https://ip-ranges.amazonaws.com/ip-ranges.json"

//NewAWSIPRangeRestAPI creates a new instance of AWSIPRangeAPI
func NewAWSIPRangeRestAPI() AWSIPRangeRestAPI {
	return &baseAWSIPRangeRestAPI{}
}

//AWSIPRangeRestAPI interface definition of the utility class to download and unmarshal IPAddressRanges from the AWS API
type AWSIPRangeRestAPI interface {
	//Download downloads the list of IP address ranges of AWS from its API
	Download() ([]byte, error)
}

type baseAWSIPRangeRestAPI struct {
}

//Download implementation of the AWSIPRangeAPI interface
func (d *baseAWSIPRangeRestAPI) Download() ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.Bytes(), nil
}
