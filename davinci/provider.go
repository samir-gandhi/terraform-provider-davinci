package davinci

import (
  "context"

  "github.com/samir-gandhi/davinci-client-go"
  "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DAVINCI_USERNAME", nil),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("DAVINCI_PASSWORD", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"hashicups_coffees": dataSourceCoffees(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
  username := d.Get("username").(string)
  password := d.Get("password").(string)

  // Warning or errors can be collected in a slice type
  var diags diag.Diagnostics

  if (username != "") && (password != "") {
    c, err := davinci.NewClient(nil, &username, &password)
    if err != nil {
      return nil, diag.FromErr(err)
    }

    return c, diags
  }

  c, err := davinci.NewClient(nil, nil, nil)
  if err != nil {
    return nil, diag.FromErr(err)
  }

  return c, diags
}
