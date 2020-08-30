package operations

//RemoveOperation implementation of Operation to remove routes
type RemoveOperation struct{}

//Name implementation of the Operation interface
func (rem *RemoveOperation) Name() string {
	return "add"
}

//Apply implementation of the Operation interface
func (rem *RemoveOperation) Apply(iface string) error {
	return nil
}
