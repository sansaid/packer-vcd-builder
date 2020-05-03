package vcdbuilder

import (
	"context"
	"github.com/vmware/go-vcloud-director/v2/govcd"
	"github.com/hashicorp/packer/helper/multistep"
	"github.com/hashicorp/packer/packer"
)

type StepVMCreate struct {
	BaseVappTemplateUrl	string
}

func (s *StepVMCreate) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
	vcdClient := state.Get("vcdClient").(govcd.VCDClient)

	ui.Say("Fetching vApp template")
	vcdCatalog := govcd.NewCatalog(vcdClient.Client)
	vappTemplate, err := vcdCatalog.GetVappTemplateByHref(s.BaseVappTemplateUrl)

	if err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}
	
	ui.Say("Creating VM from vApp template")
	vapp := govcd.NewVapp(vcdClient.Client)
	state.Put("vapp", *vapp)

	vmName, err := generateRandomVmName()

	if err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}
	
	vmTask, err := vapp.AddNewVM(vmName, vappTemplate, nil, true)

	if err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}

	state.Put("vmTask", vmTask)

	if err = vmTask.WaitTaskCompletion(); err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}

	if err = vmTask.Refresh(); err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}

	state.Put("vmTask", vmTask)
	
	ui.Say("Powering on newly created VM")
	vm, err := vcdClient.Client.FindVMByHref(vmTask.HREF)

	if err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}

	state.Put("vm", vm)

	vmPowerOnTask, err := vm.PowerOn()

	if err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}

	if err = vmPowerOnTask.WaitTaskCompletion(); err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}

	if err = vm.Refresh(); err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}

	state.Put("vm", vm)
	
	return multistep.ActionContinue
}

func (s *StepVMCreate) Cleanup(state multistep.StateBag) {
	completedTaskStatuses := []string{"success", "error", "aborted"}

	ui := state.Get("ui").(packer.Ui)
	vcdClient := state.Get("vcdClient").(govcd.VCDClient)
	vmTask := state.Get("vmTask").(govcd.Task)
	vapp := state.Get("vapp").(govcd.VApp)
	vm := state.Get("vm").(govcd.VM)

	vmTask.Refresh()

	for _, status := completedTaskStatuses {
		if vmTask.Status != status {
			vmTask.CancelTask()
			return 
		}
	}
	
	vapp.Delete()
	
	return
}