package davinci

import (
	"context"
	"fmt"
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
						"company_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"client_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"first_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"phone_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"created_by_customer_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"created_by_company_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"email_verified": {
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
						"customer_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"created_date": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"email_verified_date": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"sk_user_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_login": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"customer_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"customer_count": {
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
		return diag.FromErr(err)
	}
	if err := d.Set("customer_count", customers.CustomerCount); err != nil {
		return diag.FromErr(err)
	}
	// if err := d.Set("customers", customers.Customers); err != nil {
	//   return diag.FromErr(err)
	// }

	custs := make([]interface{}, 0, len(customers.Customers))
	for i, _ := range custs {

		custs = append(custs, map[string]interface{}{
			"email":                  customers.Customers[i].Email,
			"company_id":             customers.Customers[i].CompanyID,
			"client_id":              customers.Customers[i].ClientID,
			"first_name":             customers.Customers[i].FirstName,
			"last_name":              customers.Customers[i].LastName,
			"phone_number":           customers.Customers[i].PhoneNumber,
			"created_by_customer_id": customers.Customers[i].CreatedByCustomerID,
			"created_by_company_id":  customers.Customers[i].CreatedByCompanyID,
			"email_verified":         customers.Customers[i].EmailVerified,
			"companies":              customers.Customers[i].Companies,
			"status":                 customers.Customers[i].Status,
			"customer_type":          customers.Customers[i].CustomerType,
			"created_date":           customers.Customers[i].CustomerType,
			"email_verified_date":    customers.Customers[i].EmailVerifiedDate,
			"sk_user_id":             customers.Customers[i].SkUserID,
			"last_login":             customers.Customers[i].LastLogin,
			"customer_id":            customers.Customers[i].CustomerID,
		})
	}

	d.Set("customers", custs)

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	fmt.Printf("customer count is: %d", d.Get("customer_count"))
	return diags
}
