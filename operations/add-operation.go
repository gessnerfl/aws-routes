package operations

//AddOperation implementation of Operation to add routes
type AddOperation struct{}

//Name implementation of the Operation interface
func (add *AddOperation) Name() string {
	return "add"
}

//Apply implementation of the Operation interface
func (add *AddOperation) Apply(iface string) error {
	return nil
}
