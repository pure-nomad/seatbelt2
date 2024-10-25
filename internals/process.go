package process

import (
	ps "github.com/mitchellh/go-ps"
)

func checkOS(OperatingSys string) string {
	return OperatingSys
}

func returnUnixProcess()  {
	ps.Processes()
}