package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/johncming/terraform-provider-mydns/mydns"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: mydns.Provider,
	})
}
