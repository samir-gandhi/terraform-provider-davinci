package davinci

import "github.com/hashicorp/terraform-plugin-framework/types"

type Role struct {
	ID struct {
		Name      types.String `tfsdk:"name"`
		CompanyID types.String `tfsdk:"companyId"`
	} `tfsdk:"_id"`
	CreatedDate types.Int64  `tfsdk:"createdDate"`
	Description types.String `tfsdk:"description"`
	Policy      []struct {
		Resource types.String `tfsdk:"resource"`
		Actions  []struct {
			Action types.String `tfsdk:"action"`
			Allow  types.Bool   `tfsdk:"allow"`
		} `tfsdk:"actions"`
	} `tfsdk:"policy"`
}

type RoleCreate struct {
	Name types.String `tfsdk:"name"`
}

type RoleCreateResponse struct {
	ID struct {
		Name      types.String `tfsdk:"name"`
		CompanyID types.String `tfsdk:"companyId"`
		} `tfsdk:"_id"`
	CreatedDate types.Int64 `tfsdk:"createdDate"`
}

type RoleUpdate struct {
	Description types.String `tfsdk:"description"`
	Policy      []struct {
		Resource types.String `tfsdk:"resource"`
		Actions  []struct {
			Action types.String `tfsdk:"action"`
			Allow  types.Bool   `tfsdk:"allow"`
		} `tfsdk:"actions"`
	} `tfsdk:"policy"`
}