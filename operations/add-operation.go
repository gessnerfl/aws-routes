package operations

import (
	"github.com/gessnerfl/awsroutes/routes"

	log "github.com/sirupsen/logrus"
)

//NewDefaultAddOperation creates a new instance of an Operation to Add Routes
func NewDefaultAddOperation() Operation {
	awsIPAddressRanges := routes.NewDefaultAwsIPAddressRanges()
	nativeExecutor := NewNativeExecutor()
	return NewAddOperation(awsIPAddressRanges, nativeExecutor, log.StandardLogger())
}

//NewAddOperation creates a new instance of an Operation to Add Routes using the given IPAddressRanges
func NewAddOperation(awsIPAddressRanges routes.AwsIPAddressRanges, nativeExecutor NativeExecutor, logger log.FieldLogger) Operation {
	return &addOperation{
		awsIPAddressRanges: awsIPAddressRanges,
		nativeExecutor:     nativeExecutor,
		logger:             logger,
	}
}

type addOperation struct {
	awsIPAddressRanges routes.AwsIPAddressRanges
	nativeExecutor     NativeExecutor
	logger             log.FieldLogger
}

//Name implementation of the Operation interface
func (add *addOperation) Name() string {
	return "add"
}

//Apply implementation of the Operation interface
func (add *addOperation) Apply(iface string) error {
	ipAddressRanges, err := add.awsIPAddressRanges.Update()
	if err != nil {
		return err
	}
	for _, ip := range ipAddressRanges.Prefixes {
		if ip.IsRelevantService() {
			_, err := add.nativeExecutor.Execute("route", "add", "-net", ip.IPPrefix, "-interface", iface)
			if err != nil {
				add.logger.Warnf("Failed to create route for IP Prefix %s to interface %s; message = '%s'", ip.IPPrefix, iface, err.Error())
			}
		}
	}
	return nil
}
