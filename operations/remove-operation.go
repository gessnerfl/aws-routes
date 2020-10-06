package operations

import (
	"github.com/gessnerfl/awsroutes/routes"

	log "github.com/sirupsen/logrus"
)

//NewDefaultRemoveOperation creates a new instance of an Operation to Remove Routes
func NewDefaultRemoveOperation() Operation {
	awsIPAddressRanges := routes.NewDefaultAwsIPAddressRanges()
	nativeExecutor := NewNativeExecutor()
	return NewRemoveOperation(awsIPAddressRanges, nativeExecutor, log.StandardLogger())
}

//NewRemoveOperation creates a new instance of an Operation to Remove Routes using the given IPAddressRanges
func NewRemoveOperation(awsIPAddressRanges routes.AwsIPAddressRanges, nativeExecutor NativeExecutor, logger log.FieldLogger) Operation {
	return &removeOperation{
		awsIPAddressRanges: awsIPAddressRanges,
		nativeExecutor:     nativeExecutor,
		logger:             logger,
	}
}

type removeOperation struct {
	awsIPAddressRanges routes.AwsIPAddressRanges
	nativeExecutor     NativeExecutor
	logger             log.FieldLogger
}

//Name implementation of the Operation interface
func (rem *removeOperation) Name() string {
	return "remove"
}

//Apply implementation of the Operation interface
func (rem *removeOperation) Apply(iface string) error {
	ipAddressRanges, err := rem.awsIPAddressRanges.Read()
	if err != nil {
		return err
	}
	for _, ip := range ipAddressRanges.Prefixes {
		if ip.IsRelevantService() {
			_, err := rem.nativeExecutor.Execute("route", "delete", "-net", ip.IPPrefix, "-interface", iface)
			if err != nil {
				rem.logger.Warnf("Failed to delete route for IP Prefix %s to interface %s; message = '%s'", ip.IPPrefix, iface, err.Error())
			}
		}
	}
	return nil
}
