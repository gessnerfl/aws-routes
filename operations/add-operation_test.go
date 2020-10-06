package operations_test

import (
	"errors"
	"testing"

	"github.com/gessnerfl/awsroutes/mocks"
	. "github.com/gessnerfl/awsroutes/operations"
	"github.com/gessnerfl/awsroutes/routes"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

const (
	iface = "network-interface"
)

func TestShouldReturnAddAsNameOfAddOperation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	awsIPAddressRanges := mocks.NewMockAwsIPAddressRanges(ctrl)
	nativeExecutor := mocks.NewMockNativeExecutor(ctrl)
	logger := mocks.NewMockFieldLogger(ctrl)

	sut := NewAddOperation(awsIPAddressRanges, nativeExecutor, logger)

	require.Equal(t, "add", sut.Name())
}

func TestShouldAddAllRoutesThroughAddOperation(t *testing.T) {
	testData := routes.IPAddressRanges{
		Prefixes: []routes.IPPrefix{
			{
				IPPrefix: "ip-prefix-1",
				Service:  "AMAZON",
			},
			{
				IPPrefix: "ip-prefix-2",
				Service:  "AMAZON",
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	awsIPAddressRanges := mocks.NewMockAwsIPAddressRanges(ctrl)
	nativeExecutor := mocks.NewMockNativeExecutor(ctrl)
	logger := mocks.NewMockFieldLogger(ctrl)

	awsIPAddressRanges.EXPECT().Update().Times(1).Return(&testData, nil)
	nativeExecutor.EXPECT().Execute("route", "add", "-net", "ip-prefix-1", "-interface", iface).Times(1)
	nativeExecutor.EXPECT().Execute("route", "add", "-net", "ip-prefix-2", "-interface", iface).Times(1)

	sut := NewAddOperation(awsIPAddressRanges, nativeExecutor, logger)

	err := sut.Apply(iface)

	require.NoError(t, err)
}

func TestShouldSkipProcessingOfIrrelevantServicesForAddOperation(t *testing.T) {
	testData := routes.IPAddressRanges{
		Prefixes: []routes.IPPrefix{
			{
				IPPrefix: "ip-prefix-1",
				Service:  "OTHER",
			},
			{
				IPPrefix: "ip-prefix-2",
				Service:  "AMAZON",
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	awsIPAddressRanges := mocks.NewMockAwsIPAddressRanges(ctrl)
	nativeExecutor := mocks.NewMockNativeExecutor(ctrl)
	logger := mocks.NewMockFieldLogger(ctrl)

	awsIPAddressRanges.EXPECT().Update().Times(1).Return(&testData, nil)
	nativeExecutor.EXPECT().Execute("route", "add", "-net", "ip-prefix-2", "-interface", iface).Times(1)

	sut := NewAddOperation(awsIPAddressRanges, nativeExecutor, logger)

	err := sut.Apply(iface)

	require.NoError(t, err)
}

func TestShouldLogWarningWhenRouteCannotBeAddedForAddOperation(t *testing.T) {
	testData := routes.IPAddressRanges{
		Prefixes: []routes.IPPrefix{
			{
				IPPrefix: "ip-prefix-1",
				Service:  "AMAZON",
			},
			{
				IPPrefix: "ip-prefix-2",
				Service:  "AMAZON",
			},
		},
	}
	expectedError := errors.New("test error")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	awsIPAddressRanges := mocks.NewMockAwsIPAddressRanges(ctrl)
	nativeExecutor := mocks.NewMockNativeExecutor(ctrl)
	logger := mocks.NewMockFieldLogger(ctrl)

	awsIPAddressRanges.EXPECT().Update().Times(1).Return(&testData, nil)
	nativeExecutor.EXPECT().Execute("route", "add", "-net", "ip-prefix-1", "-interface", iface).Times(1).Return("", expectedError)
	nativeExecutor.EXPECT().Execute("route", "add", "-net", "ip-prefix-2", "-interface", iface).Times(1)
	logger.EXPECT().Warnf(gomock.Any(), gomock.Eq("ip-prefix-1"), gomock.Eq(iface), gomock.Eq(expectedError.Error())).Times(1)

	sut := NewAddOperation(awsIPAddressRanges, nativeExecutor, logger)

	err := sut.Apply(iface)

	require.NoError(t, err)
}

func TestShouldReturnErrorWhenIPAddressRangesCannotBeUpdatedForAddOperation(t *testing.T) {
	expectedError := errors.New("test error")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	awsIPAddressRanges := mocks.NewMockAwsIPAddressRanges(ctrl)
	nativeExecutor := mocks.NewMockNativeExecutor(ctrl)
	logger := mocks.NewMockFieldLogger(ctrl)

	awsIPAddressRanges.EXPECT().Update().Times(1).Return(nil, expectedError)

	sut := NewAddOperation(awsIPAddressRanges, nativeExecutor, logger)

	err := sut.Apply(iface)

	require.Error(t, err)
	require.Equal(t, expectedError, err)
}
