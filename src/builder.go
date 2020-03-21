//go:generate mapstructure-to-hcl2 -type Config

package vcdbuilder

import (
	"context"
	"github.com/hashicorp/packer/common"
	"github.com/hashicorp/packer/packer"
	"github.com/hashicorp/packer/packer/plugin"
)

// BuilderID : unique ID for builder
const BuilderID = "sansaid.vcd"

// Config is a structure encapsulating the builder configuration available to the user in the Packer template
// for the vCD builder
// TODO: Fill in types and `mapstructure` tags for each
// TODO: Identify full config needed to build images
type Config struct {
	common.PackerConfig	`mapstructure:",squash"`
	Username
	Password
	OrgName
	BaseImage
	SourceCatalog
	VAppSize
	PublishToSameCatalog	bool
	PublishCatalogName		bool
}

type Builder struct {
	
}

func (b* Builder) ConfigSpec() hcldec.ObjectSpec { return b.config.FlatMaptructure().HCL2Spec() }