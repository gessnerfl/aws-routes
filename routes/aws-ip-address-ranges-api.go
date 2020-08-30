package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const url = "https://ip-ranges.amazonaws.com/ip-ranges.json"

//NewAWSIPRangeAPI creates a new instance of AWSIPRangeAPI
func NewAWSIPRangeAPI() AWSIPRangeAPI {
	return &baseAWSIPRangeAPI{}
}

//AWSIPRangeAPI interface definition of the utility class to download and unmarshal AWSIPAddressRanges from the AWS API
type AWSIPRangeAPI interface {
	//Download downloads the list of IP address ranges of AWS from its API
	Download() ([]byte, error)
	//Unmarshal converts the slice of bytes from the AWS API to the internal representation
	Unmarshal(data []byte) (*AWSIPAddressRanges, error)
}

type baseAWSIPRangeAPI struct {
}

//Download implementation of the AWSIPRangeAPI interface
func (d *baseAWSIPRangeAPI) Download() ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.Bytes(), nil
}

//Unmarshal implementation of the AWSIPRangeAPI interface
func (d *baseAWSIPRangeAPI) Unmarshal(data []byte) (*AWSIPAddressRanges, error) {
	result := &AWSIPAddressRanges{}
	err := json.Unmarshal(data, result)
	return result, err
}
