package operations

import "os/exec"

//NativeExecutor interface definition for a wrapper to run native executables
type NativeExecutor interface {
	//Execute executes the given command with the provided arguments
	Execute(command string, args ...string) (string, error)
}

//NewNativeExecutor creates a new instance of a NativeExecutor
func NewNativeExecutor() NativeExecutor {
	return &baseNativeExecutor{}
}

type baseNativeExecutor struct{}

func (b *baseNativeExecutor) Execute(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)

	out, err := cmd.Output()

	return string(out), err
}
