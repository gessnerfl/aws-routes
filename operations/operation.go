package operations

import "fmt"

//Operation interface definition of a operation of the awsroutes CLI
type Operation interface {
	Name() string
	Apply(iface string) error
}

//Operations a slice of Operations
type Operations []Operation

//ByName searches and returns the Operation of the given name from the underlying slice or an error when no operation exists for the given name
func (ops Operations) ByName(name string) (Operation, error) {
	for _, this := range ops {
		if this.Name() == name {
			return this, nil
		}
	}
	return nil, fmt.Errorf("%s is not a supported operation", name)
}

//SupportedOperations slice of supported operations of the awsroutes cli
var SupportedOperations = Operations{NewDefaultAddOperation(), &RemoveOperation{}}
