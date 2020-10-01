package routes

import (
	"github.com/gessnerfl/awsroutes/storage"
)

//StorageFilename the filename of the downloaded ip ranges in the local storage
const StorageFilename = "ip-ranges.json"

//AwsIPAddressRanges interface definition for access to the AWS IPAddressRanges
type AwsIPAddressRanges interface {
	//Update updates the current list of IPAddressRanges in the local storage and returns the current version
	Update() (*IPAddressRanges, error)
	//Get reads the latest version of IPAddressRanges from the local storage
	Read() (*IPAddressRanges, error)
}

//NewDefaultAwsIPAddressRanges creates a new instance of AwsIPAddressRanges with a storage in a subfolder of the user home directory
func NewDefaultAwsIPAddressRanges() (AwsIPAddressRanges, error) {
	restAPI := NewAwsIPRangeRestAPI()
	unmarshaller := NewUnmarshaller()
	storage, err := storage.NewStorageAtUserHome()
	return NewAwsIPAddressRanges(restAPI, unmarshaller, storage), err
}

//NewAwsIPAddressRanges create a new instance of AwsIPAddressRanges with the given AwsIPRangeRestAPI, Unmarshaller, and Storage
func NewAwsIPAddressRanges(restAPI AwsIPRangeRestAPI, unmarshaller Unmarshaller, storage storage.Storage) AwsIPAddressRanges {
	return &baseAwsIPAddressRanges{
		restAPI:      restAPI,
		unmarshaller: unmarshaller,
		storage:      storage,
	}
}

type baseAwsIPAddressRanges struct {
	restAPI      AwsIPRangeRestAPI
	unmarshaller Unmarshaller
	storage      storage.Storage
}

func (b *baseAwsIPAddressRanges) Update() (*IPAddressRanges, error) {
	data, err := b.restAPI.Download()
	if err != nil {
		return &IPAddressRanges{}, err
	}

	err = b.storage.Update(StorageFilename, data)
	if err != nil {
		return &IPAddressRanges{}, err
	}

	return b.unmarshaller.Unmarshal(data)
}

func (b *baseAwsIPAddressRanges) Read() (*IPAddressRanges, error) {
	data, err := b.storage.Read(StorageFilename)
	if err != nil {
		return &IPAddressRanges{}, err
	}

	return b.unmarshaller.Unmarshal(data)
}
