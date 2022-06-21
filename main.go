package main

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"terraform-provider-davinci/davinci"
)

func main() {
	tfsdk.Serve(context.Background(), davinci.New, tfsdk.ServeOpts{
		Name: "davinci",
	})
}

