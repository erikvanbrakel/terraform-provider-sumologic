package main

import (

	"github.com/hashicorp/terraform/plugin"
	"github.com/erikvanbrakel/terraform-provider-sumologic/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}
