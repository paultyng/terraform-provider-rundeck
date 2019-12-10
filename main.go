package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/terraform-providers/terraform-provider-rundeck/rundeck"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: rundeck.Provider})
}
