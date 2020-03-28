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
	
	ui.Say("Creating VM from vApp template")
	vapp := govcd.NewVapp(vcdClient.Client)
	state.Put("vapp", *vapp)

	vmTask, err := vapp.AddNewVM(vmName, vappTemplate, nil, true)

	stateError(err, state)
	state.Put("vmTask", vmTask)

	err = vmTask.WaitTaskCompletion()

	stateError(err, state)
	
	err = vmTask.Refresh()

	stateError(err, state)
	state.Put("vmTask", vmTask)
	
	vm, err := vcdClient.Client.FindVMByHref(vmTask.HREF)

	stateError(err, state)
	state.Put("vm", vm)

	vmPowerOnTask, err := vm.PowerOn()

	stateError(err, state)
	
	err = vmPowerOnTask.WaitTaskCompletion()

	stateError(err, state)

	err = vm.Refresh()

	stateError(err, state)
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
	
	vm.Refresh()

	if vm.IsDeployed() {
		vapp.RemoveVM(vm)
	}
	
	return
}