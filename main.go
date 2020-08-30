package main

import (
	"fmt"
	"os"

	"github.com/gessnerfl/awsroutes/operations"
)

func main() {
	if len(os.Args) != 2 {
		printUsageAndExit()
	}
	operation, err := getOperation()
	if err != nil {
		printErrorWithUsageAndExit(err)
	}
	iface := os.Args[1]

	err = operation.Apply(iface)
	if err != nil {
		fmt.Printf("Failed ot apply %s operation; %s\n", operation.Name(), err.Error())
		os.Exit(-2)
	}
}

func getOperation() (operations.Operation, error) {
	op := os.Args[0]
	return operations.SupportedOperations.ByName(op)
}

func printErrorWithUsageAndExit(err error) {
	fmt.Printf("%s\n\n---\n\n", err.Error())
	printUsageAndExit()
}

func printUsageAndExit() {
	fmt.Println("Usage awsroutes add|remove <interface>")
	os.Exit(-1)
}
