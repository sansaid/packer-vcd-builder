package vcdbuilder

import (
	"github.com/hashicorp/packer/helper/multistep"
)

func stateError(err error, state multistep.StateBag) {
	if err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}
}

func generateRandomVmName() string {
	return
}