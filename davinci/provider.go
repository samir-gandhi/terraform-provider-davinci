package davinci

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/samir-gandhi/davinci-client-go"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DAVINCI_USERNAME", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("DAVINCI_PASSWORD", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"davinci_customers": dataSourceCustomers(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Warning Message Summary",
		Detail:   "This is the detailed warning message from providerConfigure",
	})
	var c *davinci.Client
	
	// if (username != "") && (password != "") {
		// 	diags = append(diags, diag.Diagnostic{
			// 		Severity: diag.Error,
			// 		Summary:  "User or Password Not Provided",
			// 		Detail:   "Unable to auth user",
			// 	})
			// }
			
	if (username != "") && (password != "") {
		fmt.Printf("username is: %s",username)
		c, err := davinci.NewClient(nil, &username, &password)
		if err != nil {
      diags = append(diags, diag.Diagnostic{
        Severity: diag.Error,
        Summary:  "Unable to create Davinci client",
        Detail:   "Unable to auth user",
      })
      return nil, diags
		}

		return c, diags
	}

	// c, err := davinci.NewClient(nil, nil, nil)
	// if err != nil {
	// 	diags = append(diags, diag.Diagnostic{
	// 		Severity: diag.Error,
	// 		Summary:  "Unable to create Davinci client",
	// 		Detail:   "Unable to auth user",
	// 	})
	// }

	return c, diags
}
