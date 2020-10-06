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

func TestShouldReturnRemoveAsNameOfRemoveOperation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	awsIPAddressRanges := mocks.NewMockAwsIPAddressRanges(ctrl)
	nativeExecutor := mocks.NewMockNativeExecutor(ctrl)
	logger := mocks.NewMockFieldLogger(ctrl)

	sut := NewRemoveOperation(awsIPAddressRanges, nativeExecutor, logger)

	require.Equal(t, "remove", sut.Name())
}

func TestShouldRemoveAllRoutesThroughRemoveOperation(t *testing.T) {
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

	awsIPAddressRanges.EXPECT().Read().Times(1).Return(&testData, nil)
	nativeExecutor.EXPECT().Execute("route", "delete", "-net", "ip-prefix-1", "-interface", iface).Times(1)
	nativeExecutor.EXPECT().Execute("route", "delete", "-net", "ip-prefix-2", "-interface", iface).Times(1)

	sut := NewRemoveOperation(awsIPAddressRanges, nativeExecutor, logger)

	err := sut.Apply(iface)

	require.NoError(t, err)
}

func TestShouldSkipProcessingOfIrrelevantServicesForRemoveOperation(t *testing.T) {
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

	awsIPAddressRanges.EXPECT().Read().Times(1).Return(&testData, nil)
	nativeExecutor.EXPECT().Execute("route", "delete", "-net", "ip-prefix-2", "-interface", iface).Times(1)

	sut := NewRemoveOperation(awsIPAddressRanges, nativeExecutor, logger)

	err := sut.Apply(iface)

	require.NoError(t, err)
}

func TestShouldLogWarningWhenRouteCannotBeRemovedForRemoveOperation(t *testing.T) {
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

	awsIPAddressRanges.EXPECT().Read().Times(1).Return(&testData, nil)
	nativeExecutor.EXPECT().Execute("route", "delete", "-net", "ip-prefix-1", "-interface", iface).Times(1).Return("", expectedError)
	nativeExecutor.EXPECT().Execute("route", "delete", "-net", "ip-prefix-2", "-interface", iface).Times(1)
	logger.EXPECT().Warnf(gomock.Any(), gomock.Eq("ip-prefix-1"), gomock.Eq(iface), gomock.Eq(expectedError.Error())).Times(1)

	sut := NewRemoveOperation(awsIPAddressRanges, nativeExecutor, logger)

	err := sut.Apply(iface)

	require.NoError(t, err)
}

func TestShouldReturnErrorWhenIPAddressRangesCannotBeReadForRemoveOperation(t *testing.T) {
	expectedError := errors.New("test error")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	awsIPAddressRanges := mocks.NewMockAwsIPAddressRanges(ctrl)
	nativeExecutor := mocks.NewMockNativeExecutor(ctrl)
	logger := mocks.NewMockFieldLogger(ctrl)

	awsIPAddressRanges.EXPECT().Read().Times(1).Return(nil, expectedError)

	sut := NewRemoveOperation(awsIPAddressRanges, nativeExecutor, logger)

	err := sut.Apply(iface)

	require.Error(t, err)
	require.Equal(t, expectedError, err)
}
