package vcdbuilder

import (
	"context"
	"net/url"
	"github.com/vmware/go-vcloud-director/v2/govcd"
	"github.com/hashicorp/packer/helper/multistep"
	"github.com/hashicorp/packer/packer"
)

type StepVCDConfigure struct {
	Username	string
	Password	string
	Org			string
	Endpoint	string
	Insecure	bool
}

func (s *StepVCDConfig) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)

	ui.Say("Parsing endpoint URL")
	u, err := url.ParseRequestURI(v.Endpoint)

	if err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}

	ui.Say("Authenticating with vCD server")
	vcdClient := govcd.NewVCDClient(*u, v.Insecure)
	
	if err = vcdClient.Authenticate(v.Username, v.Password, v.Org); err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}

	ui.Say("Authentication successful")
	state.Put("vcdClient", vcdClient)
	
	return multistep.ActionContinue
}

func (s *StepVCDConfigure) Cleanup(state multistep.StateBag) {
	// No cleanup needed in authenticate step
	return
}