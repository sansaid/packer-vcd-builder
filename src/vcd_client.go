package vcdbuilder

import (
	"fmt"
	"net/url"
	"github.com/vmware/go-vcloud-director/v2/govcd"
)

type VCDConfig struct {
	Username	string
	Password	string
	Org			string
	Endpoint	string
	VDC			string
	Insecure	bool
}

func (v *VCDConfig) getClient() (*govcd.VCDClient, error) {
	u, err := url.ParseRequestURI(v.Endpoint)

	if err != nil {
		return nil, fmt.Errorf("Unable to parse URL: %s", err)
	}

	vcdClient := govcd.NewVCDClient(*u, v.Insecure)
	err = vcdClient.Authenticate(v.Username, v.Password, v.Org)

	if err != nil {
		return nil, fmt.Errorf("Unable to authenticate: %s", err)
	}

	return vcdClient, nil
}