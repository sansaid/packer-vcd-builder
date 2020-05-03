//go:generate mapstructure-to-hcl2 -type Config

package vcdbuilder

import (
	"context"
	"github.com/hashicorp/packer/common"
	"github.com/hashicorp/packer/packer"
	"github.com/hashicorp/packer/packer/plugin"
	"github.com/hashicorp/packer/helper/multistep"
	"github.com/vmware/go-vcloud-director/v2/types/v56"
)

// BuilderID : unique ID for builder
const BuilderID = "sansaid.vcd"

// Config is a structure encapsulating the builder configuration available to the user in the Packer template
// for the vCD builder
// TODO: Fill in types and `mapstructure` tags for each
// TODO: Identify full config needed to build images
type Config struct {
	common.PackerConfig	`mapstructure:",squash"`
	Username				string
	Password				string
	OrgName					string
	VAppTemplateHref		string
	PublishToSameCatalog	bool
	PublishCatalogName		bool
	AcceptAllEulas			bool // (should default to true)
}

type Builder struct {
	config Config
	runner multistep.Runner	
}

func (b *Builder) ConfigSpec() hcldec.ObjectSpec { return b.config.FlatMaptructure().HCL2Spec() }

func (b *Builder) Prepare(raws ...interface{}) ([]string, []string, error) {

}

func (b *Builder) Run(ctx context.Context, ui packer.Ui, hook packer.Hook) (packer.Artifact, error) {
	// Creating statebag
	state := new(multistep.BasicStateBag)

	// Initialising statebag
	state.Put("config", &b.config)
	state.Put("hook", hook)
	state.Put("ui", ui)

	/* TODO: 
	1. Query vApp Template to be used as base - vApp template must have only one VM
	>> vappTemplate := queryVappTemplate(vcdClient, <TODO>, <TODO>) - https://github.com/vmware/go-vcloud-director/blob/master/govcd/catalog.go

	2. Ensure vApp template must only have one VM (NOTE: not sure if necessary or will be handled better downstream?)
	>> len(vappTemplate.VAppTemplate.Children.VM) == 1

	3. Randomly generate VM name
	>> builderName := <TODO>

	4. Create VM using vAppTemplate (ensure vAppTemplate has sshd)

	5. Power on VM

	6. Provision VM

	7. Power off VM

	8. Remove VM
	*/
}