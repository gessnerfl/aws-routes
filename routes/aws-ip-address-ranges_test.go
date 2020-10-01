package routes_test

import (
	"errors"
	"testing"

	"github.com/gessnerfl/awsroutes/mocks"
	. "github.com/gessnerfl/awsroutes/routes"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestShouldSuccessfullyUpdateAwsIPAddressRanges(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rawdata := []byte("rawdata")
	data := IPAddressRanges{
		SyncToken: "test",
	}

	storage := mocks.NewMockStorage(ctrl)
	restapi := mocks.NewMockAwsIPRangeRestAPI(ctrl)
	unmarshaller := mocks.NewMockUnmarshaller(ctrl)

	restapi.EXPECT().Download().Times(1).Return(rawdata, nil)
	storage.EXPECT().Update(StorageFilename, rawdata).Times(1)
	unmarshaller.EXPECT().Unmarshal(rawdata).Times(1).Return(&data, nil)

	sut := NewAwsIPAddressRanges(restapi, unmarshaller, storage)
	result, err := sut.Update()

	require.NoError(t, err)
	require.Equal(t, &data, result)
}

func TestShouldFailToUpdateIPAddressRangesWhenIPRangesCannotBeDownloadedFromAwsAPI(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rawdata := make([]byte, 0)
	expectedError := errors.New("download failed")

	storage := mocks.NewMockStorage(ctrl)
	restapi := mocks.NewMockAwsIPRangeRestAPI(ctrl)
	unmarshaller := mocks.NewMockUnmarshaller(ctrl)

	restapi.EXPECT().Download().Times(1).Return(rawdata, expectedError)

	sut := NewAwsIPAddressRanges(restapi, unmarshaller, storage)
	_, err := sut.Update()

	require.Error(t, err)
	require.Equal(t, expectedError, err)
}

func TestShouldFailToUpdateIPAddressRangesWhenUpdateOfLocalStorageReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rawdata := make([]byte, 0)
	expectedError := errors.New("update of local storage failed")

	storage := mocks.NewMockStorage(ctrl)
	restapi := mocks.NewMockAwsIPRangeRestAPI(ctrl)
	unmarshaller := mocks.NewMockUnmarshaller(ctrl)

	restapi.EXPECT().Download().Times(1).Return(rawdata, nil)
	storage.EXPECT().Update(StorageFilename, rawdata).Times(1).Return(expectedError)

	sut := NewAwsIPAddressRanges(restapi, unmarshaller, storage)
	_, err := sut.Update()

	require.Error(t, err)
	require.Equal(t, expectedError, err)
}

func TestShouldFailToUpdateIPAddressRangesWhenUnmarshallingReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rawdata := make([]byte, 0)
	data := IPAddressRanges{}
	expectedError := errors.New("unmarshalling failed")

	storage := mocks.NewMockStorage(ctrl)
	restapi := mocks.NewMockAwsIPRangeRestAPI(ctrl)
	unmarshaller := mocks.NewMockUnmarshaller(ctrl)

	restapi.EXPECT().Download().Times(1).Return(rawdata, nil)
	storage.EXPECT().Update(StorageFilename, rawdata).Times(1)
	unmarshaller.EXPECT().Unmarshal(rawdata).Times(1).Return(&data, expectedError)

	sut := NewAwsIPAddressRanges(restapi, unmarshaller, storage)
	result, err := sut.Update()

	require.Error(t, err)
	require.Equal(t, expectedError, err)
	require.Equal(t, &data, result)
}

func TestShouldReadCurrentPersistedIPAddressRangesFromLocalStorage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rawdata := []byte("rawdata")
	data := IPAddressRanges{
		SyncToken: "test",
	}

	storage := mocks.NewMockStorage(ctrl)
	restapi := mocks.NewMockAwsIPRangeRestAPI(ctrl)
	unmarshaller := mocks.NewMockUnmarshaller(ctrl)

	storage.EXPECT().Read(StorageFilename).Times(1).Return(rawdata, nil)
	unmarshaller.EXPECT().Unmarshal(rawdata).Times(1).Return(&data, nil)

	sut := NewAwsIPAddressRanges(restapi, unmarshaller, storage)
	result, err := sut.Read()

	require.NoError(t, err)
	require.Equal(t, &data, result)
}

func TestShouldFailToReadCurrentPersistedIPAddressRangesFromLocalStorageWhenLocalStorageReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rawdata := make([]byte, 0)
	data := IPAddressRanges{}
	expectedError := errors.New("read fromm local storage failed")

	storage := mocks.NewMockStorage(ctrl)
	restapi := mocks.NewMockAwsIPRangeRestAPI(ctrl)
	unmarshaller := mocks.NewMockUnmarshaller(ctrl)

	storage.EXPECT().Read(StorageFilename).Times(1).Return(rawdata, expectedError)

	sut := NewAwsIPAddressRanges(restapi, unmarshaller, storage)
	result, err := sut.Read()

	require.Error(t, err)
	require.Equal(t, expectedError, err)
	require.Equal(t, &data, result)
}

func TestShouldFailToReadCurrentPersistedIPAddressRangesFromLocalStorageWhenUnmarshallingFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rawdata := []byte("rawdata")
	data := IPAddressRanges{}
	expectedError := errors.New("unmarshalling failed")

	storage := mocks.NewMockStorage(ctrl)
	restapi := mocks.NewMockAwsIPRangeRestAPI(ctrl)
	unmarshaller := mocks.NewMockUnmarshaller(ctrl)

	storage.EXPECT().Read(StorageFilename).Times(1).Return(rawdata, nil)
	unmarshaller.EXPECT().Unmarshal(rawdata).Times(1).Return(&data, expectedError)

	sut := NewAwsIPAddressRanges(restapi, unmarshaller, storage)
	result, err := sut.Read()

	require.Error(t, err)
	require.Equal(t, expectedError, err)
	require.Equal(t, &data, result)
}
