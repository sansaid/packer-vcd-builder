package vcdbuilder

import (
	"context"
	"github.com/vmware/go-vcloud-director/v2/govcd"
	"github.com/hashicorp/packer/helper/multistep"
	"github.com/hashicorp/packer/packer"
)

type StepVMCreate struct {
	ParentVapp		string
	VDC				string
	VappTemplateUrl	string
}

func (s *StepVMCreate) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
	vcdClient := state.Get("vcdClient").(govcd.VCDClient)

	ui.Say("Creating new catalog reference")
	vcdCatalog := govcd.NewCatalog(vcdClient.Client)

	ui.Say("Fetching vApp template")
	vappTemplate, err := vcdCatalog.GetVappTemplateByHref(s.VAppTemplateUrl)

	stateError(err, state)

	vmName := generateRandomVmName() // TODO: needs to be implemented
	
	ui.Say("Creating new vApp from template")
	// Implement VM creation: 

	return multistep.ActionContinue
}

func (s *StepVMCreate) Cleanup(state multistep.StateBag) {
	// No cleanup needed in vApp template query stage
	return
}