package vcdbuilder

import (
	"context"
	"github.com/vmware/go-vcloud-director/v2/govcd"
	"github.com/hashicorp/packer/helper/multistep"
	"github.com/hashicorp/packer/packer"
)

type StepCreateVappTemplate struct {
	// Properties
}

func (s *StepCreateVappTemplate) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
	vcdClient := state.Get("vcdClient").(govcd.VCDClient)

	
}

func (s *StepCreateVappTemplate) Cleanup(state multistep.StateBag) {

}