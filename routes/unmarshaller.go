package routes

import "encoding/json"

//Unmarshaller interface definition of a unmarshaller for IPAddressRanges received from the REST API
type Unmarshaller interface {
	//Unmarshal unmarshalls the provided bytes into the struct type IPAddressRanges
	Unmarshal(data []byte) (*IPAddressRanges, error)
}

//NewUnmarshaller creates a new instance of Unmarshaller
func NewUnmarshaller() Unmarshaller {
	return &baseUnmarshaller{}
}

type baseUnmarshaller struct{}

func (u *baseUnmarshaller) Unmarshal(data []byte) (*IPAddressRanges, error) {
	result := &IPAddressRanges{}
	err := json.Unmarshal(data, result)
	return result, err
}
