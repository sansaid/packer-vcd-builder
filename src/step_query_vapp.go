package vcdbuilder

import (
	"context"
	"github.com/vmware/go-vcloud-director/v2/govcd"
	"github.com/hashicorp/packer/helper/multistep"
	"github.com/hashicorp/packer/packer"
)

type StepVappTemplate struct {
	VAppTemplateUrl string
}

func (s *StepVappTemplate) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
	vcdClient := state.Get("vcdClient").(govcd.VCDClient)

	ui.Say("Creating new catalog reference")
	vcdCatalog := govcd.NewCatalog(vcdClient)

	ui.Say("Fetching vApp template")
	vappTemplate, err := vcdCatalog.GetVappTemplateByHref(s.VAppTemplateUrl)

	stateError(err, state)

	state.Put("vappTemplate", vappTemplate)

	return multistep.ActionContinue
}

func (s *StepVappTemplate) Cleanup(state multistep.StateBag) {
	// No cleanup needed in vApp template query stage
	return
}