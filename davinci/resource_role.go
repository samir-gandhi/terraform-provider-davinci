package davinci

import (
	"context"
	// "math/big"
	// "strconv"
	// "time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/samir-gandhi/davinci-client-go"
)

type resourceRoleType struct {

}

// Role Resource schema
func (r resourceRoleType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Computed: true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Type: types.StringType,
						Required: true,
					},
					"companyId": {
						Type: types.StringType,
						Computed: true,
					},
				}),
			},
			"createdDate": {
				Computed: true,
				Type: types.Int64Type,
			},
		},
	}, nil
}

type resourceRole struct {
	p provider
}

// New resource instance
func (r resourceRoleType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceRole{
		p: *(p.(*provider)),
	}, nil
}

// Create a new resource
func (r resourceRole) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	if !r.p.configured {
		resp.Diagnostics.AddError(
				"Provider not configured",
				"The provider hasn't been configured before apply, likely because it depends on an unknown value from another resource. This leads to weird stuff happening, so we'd prefer if you didn't do that. Thanks!",
		)
		return
	}

	// Retrieve values from plan
	var plan RoleCreate
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	payload := davinci.RoleCreate{
		Name: plan.Name,
	}
	roleResp, err := r.p.client.CreateRole(&r.p.client.CompanyID, &payload)
	if err != nil {
			resp.Diagnostics.AddError(
				"Error creating role",
				"Could not create role,	unexpected error: " + err.Error(),
			)
		return
	}

	roleRespSchema := RoleCreateResponse{
		CreatedDate: types.Int64{Value: roleResp.CreatedDate},
		ID: 	
	}
	diags = resp.State.Set(ctx, roleRespSchema)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
			return
	}

}

// Read resource information
func (r resourceRole) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state RoleCreateResponse
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	roleID := state.ID.Name

}

// Update resource
func (r resourceRole) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
}

// Delete resource
func (r resourceRole) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
}

// Import resource
func (r resourceRole) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	// Save the import identifier in the id attribute
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}
