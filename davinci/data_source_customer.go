package davinci

import (
  "context"
  "strconv"

  dv "github.com/samir-gandhi/davinci-client-go"
  "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCustomers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCustomersRead,
		Schema: map[string]*schema.Schema{
			"customers": &schema.Schema{
				Type: schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"email": &schema.Schema{
							Type: schema.TypeString,
							Computed: true,
						},
						"companyId": &schema.Schema{
							Type: schema.TypeString,
							Computed: true,
						},
						"clientId": &schema.Schema{
							Type: schema.TypeString,
							Computed: true,
						},
						"firstName": &schema.Schema{
							Type: schema.TypeString,
							Computed: true,
						},
						"lastName": &schema.Schema{
							Type: schema.TypeString,
							Computed: true,
						},
						"phoneNumber": &schema.Schema{
							Type: schema.TypeString,
							Computed: true,
						},
						"createdByCustomerId": &schema.Schema{
							Type: schema.TypeString,
							Computed: true,
						},
						"createdByCompanyId": &schema.Schema{
							Type: schema.TypeString,
							Computed: true,
						},
						"emailVerified": &schema.Schema{
							Type: schema.TypeBool,
							Computed: true,
						},
						"companies": &schema.Schema{
							Type: schema.TypeList,
							Computed: true,
							// Elem: ,
						},
						"status": &schema.Schema{
							Type: schema.TypeString,
							Computed: true,
						},
						"customerType": &schema.Schema{
							Type: schema.TypeString,
							Computed: true,
						},
						"createdDate": &schema.Schema{
							Type: schema.TypeInt,
							Computed: true,
						},
						"emailVerifiedDate": &schema.Schema{
							Type: schema.TypeInt,
							Computed: true,
						},
						"skUserId": &schema.Schema{
							Type: schema.TypeString,
							Computed: true,
						},
						"lastLogin": &schema.Schema{
							Type: schema.TypeInt,
							Computed: true,
						},
						"customerId": &schema.Schema{
							Type: schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"customerCount": &schema.Schema{
				Type: schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceCustomersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diags
}