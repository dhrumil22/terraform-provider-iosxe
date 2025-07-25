// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-restconf"
)

func NewSNMPServerUserResource() resource.Resource {
	return &SNMPServerUserResource{}
}

type SNMPServerUserResource struct {
	data *IosxeProviderData
}

func (r *SNMPServerUserResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmp_server_user"
}

func (r *SNMPServerUserResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the SNMP Server User configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the object.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"delete_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure behavior when deleting/destroying the resource. Either delete the entire object (YANG container) being managed, or only delete the individual resource attributes configured explicitly and leave everything else as-is. Default value is `all`.").AddStringEnumDescription("all", "attributes").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("all", "attributes"),
				},
			},
			"username": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the user").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"grpname": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Group to which the user belongs").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"v3_auth_algorithm": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use HMAC SHA/MD5 algorithm for authentication").AddStringEnumDescription("md5", "sha").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("md5", "sha"),
				},
			},
			"v3_auth_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication password for user").String,
				Required:            true,
			},
			"v3_auth_priv_aes_algorithm": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("128", "192", "256").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("128", "192", "256"),
				},
			},
			"v3_auth_priv_aes_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication password for user").String,
				Optional:            true,
			},
			"v3_auth_priv_aes_access_ipv6_acl": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify IPv6 Named Access-List").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 194),
				},
			},
			"v3_auth_priv_aes_access_standard_acl": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Standard IP Access-list allowing access with this community string").AddIntegerRangeDescription(1, 99).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 99),
				},
			},
			"v3_auth_priv_aes_access_acl_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Access-list name").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 183),
				},
			},
			"v3_auth_priv_des_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication password for user").String,
				Optional:            true,
			},
			"v3_auth_priv_des_access_ipv6_acl": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify IPv6 Named Access-List").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 194),
				},
			},
			"v3_auth_priv_des_access_standard_acl": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Standard IP Access-list allowing access with this community string").AddIntegerRangeDescription(1, 99).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 99),
				},
			},
			"v3_auth_priv_des_access_acl_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Access-list name").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 183),
				},
			},
			"v3_auth_priv_des3_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication password for user").String,
				Optional:            true,
			},
			"v3_auth_priv_des3_access_ipv6_acl": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify IPv6 Named Access-List").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 194),
				},
			},
			"v3_auth_priv_des3_access_standard_acl": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Standard IP Access-list allowing access with this community string").AddIntegerRangeDescription(1, 99).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 99),
				},
			},
			"v3_auth_priv_des3_access_acl_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Access-list name").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 183),
				},
			},
			"v3_auth_access_ipv6_acl": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify IPv6 Named Access-List").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 194),
				},
			},
			"v3_auth_access_standard_acl": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Standard IP Access-list allowing access with this community string").AddIntegerRangeDescription(1, 99).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 99),
				},
			},
			"v3_auth_access_acl_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Access-list name").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 183),
				},
			},
		},
	}
}

func (r *SNMPServerUserResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.data = req.ProviderData.(*IosxeProviderData)
}

func (r *SNMPServerUserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SNMPServerUser

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getPath()))

	device, ok := r.data.Devices[plan.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", plan.Device.ValueString()))
		return
	}

	if device.Managed {
		// Create object
		body := plan.toBody(ctx)

		emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
		tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

		if YangPatch {
			edits := []restconf.YangPatchEdit{restconf.NewYangPatchEdit("merge", plan.getPath(), restconf.Body{Str: body})}
			for _, i := range emptyLeafsDelete {
				edits = append(edits, restconf.NewYangPatchEdit("remove", i, restconf.Body{}))
			}
			_, err := device.Client.YangPatchData("", "1", "", edits)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object, got error: %s", err))
				return
			}
		} else {
			res, err := device.Client.PatchData(plan.getPathShort(), body)
			if len(res.Errors.Error) > 0 && res.Errors.Error[0].ErrorMessage == "patch to a nonexistent resource" {
				_, err = device.Client.PutData(plan.getPath(), body)
			}
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PATCH), got error: %s", err))
				return
			}
			for _, i := range emptyLeafsDelete {
				res, err := device.Client.DeleteData(i)
				if err != nil && res.StatusCode != 404 {
					resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
					return
				}
			}
		}
	}

	plan.Id = types.StringValue(plan.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getPath()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *SNMPServerUserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SNMPServerUser

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.ValueString()))

	device, ok := r.data.Devices[state.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", state.Device.ValueString()))
		return
	}

	if device.Managed {
		res, err := device.Client.GetData(state.Id.ValueString())
		if res.StatusCode == 404 {
			state = SNMPServerUser{Device: state.Device, Id: state.Id}
		} else {
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
				return
			}

			imp, diags := helpers.IsFlagImporting(ctx, req)
			if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
				return
			}

			// After `terraform import` we switch to a full read.
			if imp {
				state.getIdsFromPath()
				state.fromBody(ctx, res.Res)
			} else {
				state.updateFromBody(ctx, res.Res)
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *SNMPServerUserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SNMPServerUser

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	device, ok := r.data.Devices[plan.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", plan.Device.ValueString()))
		return
	}

	if device.Managed {
		body := plan.toBody(ctx)

		deletedItems := plan.getDeletedItems(ctx, state)
		tflog.Debug(ctx, fmt.Sprintf("Removed items to delete: %+v", deletedItems))

		emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
		tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

		if YangPatch {
			edits := []restconf.YangPatchEdit{restconf.NewYangPatchEdit("merge", plan.getPath(), restconf.Body{Str: body})}
			for _, i := range deletedItems {
				edits = append(edits, restconf.NewYangPatchEdit("remove", i, restconf.Body{}))
			}
			for _, i := range emptyLeafsDelete {
				edits = append(edits, restconf.NewYangPatchEdit("remove", i, restconf.Body{}))
			}
			_, err := device.Client.YangPatchData("", "1", "", edits)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
				return
			}
		} else {
			res, err := device.Client.PatchData(plan.getPathShort(), body)
			if len(res.Errors.Error) > 0 && res.Errors.Error[0].ErrorMessage == "patch to a nonexistent resource" {
				_, err = device.Client.PutData(plan.getPath(), body)
			}
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PATCH), got error: %s", err))
				return
			}
			for _, i := range deletedItems {
				res, err := device.Client.DeleteData(i)
				if err != nil && res.StatusCode != 404 {
					resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
					return
				}
			}
			for _, i := range emptyLeafsDelete {
				res, err := device.Client.DeleteData(i)
				if err != nil && res.StatusCode != 404 {
					resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
					return
				}
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *SNMPServerUserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SNMPServerUser

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	device, ok := r.data.Devices[state.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", state.Device.ValueString()))
		return
	}

	if device.Managed {
		deleteMode := "all"
		if state.DeleteMode.ValueString() == "all" {
			deleteMode = "all"
		} else if state.DeleteMode.ValueString() == "attributes" {
			deleteMode = "attributes"
		}

		if deleteMode == "all" {
			res, err := device.Client.DeleteData(state.Id.ValueString())
			if err != nil && res.StatusCode != 404 && res.StatusCode != 400 {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
				return
			}
		} else {
			deletePaths := state.getDeletePaths(ctx)
			tflog.Debug(ctx, fmt.Sprintf("Paths to delete: %+v", deletePaths))

			if YangPatch {
				edits := []restconf.YangPatchEdit{}
				for _, i := range deletePaths {
					edits = append(edits, restconf.NewYangPatchEdit("remove", i, restconf.Body{}))
				}
				_, err := device.Client.YangPatchData("", "1", "", edits)
				if err != nil {
					resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
					return
				}
			} else {
				for _, i := range deletePaths {
					res, err := device.Client.DeleteData(i)
					if err != nil && res.StatusCode != 404 {
						resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
						return
					}
				}
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *SNMPServerUserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}
