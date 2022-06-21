package davinci

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	dv "github.com/samir-gandhi/davinci-client-go"
)

func dataSourceCustomers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCustomersRead,
		Schema: map[string]*schema.Schema{
			"customers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"email": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"companyId": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"clientId": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"firstName": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lastName": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"phoneNumber": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"createdByCustomerId": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"createdByCompanyId": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"emailVerified": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"companies": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"customerType": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"createdDate": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"emailVerifiedDate": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"skUserId": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lastLogin": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"customerId": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"customerCount": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceCustomersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dv.Client)

	var diags diag.Diagnostics

	customers, err := c.GetCustomers(&c.CompanyID, nil)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Davinci client",
			Detail:   "Unable to auth user",
		})
	}
	if err := d.Set("customers", customers); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Davinci client",
			Detail:   "Unable to auth user",
		})
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Unable to create Davinci client",
		Detail:   "Unable to auth user",
	})
	return diags
}