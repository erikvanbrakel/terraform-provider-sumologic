package terraform_provider_sumologic

import (
	"./provider"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts {
		ProviderFunc: provider.Provider,
	})
}
