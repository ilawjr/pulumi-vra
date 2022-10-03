// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vra

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VMware vRA CloudAccountAzure data source.
//
// ## Example Usage
// ### S
//
// **Azure cloud account data source by its id:**
//
// This is an example of how to read the cloud account data source using its id.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vra/sdk/go/vra"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/schmidtw/pulumi-vra/sdk/go/vra"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vra.LookupCloudAccountAzure(ctx, &GetCloudAccountAzureArgs{
//				Id: pulumi.StringRef(_var.Vra_cloud_account_azure_id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// **Azure cloud account data source by its name:**
//
// This is an example of how to read the cloud account data source using its name.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vra/sdk/go/vra"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/schmidtw/pulumi-vra/sdk/go/vra"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vra.LookupCloudAccountAzure(ctx, &GetCloudAccountAzureArgs{
//				Name: pulumi.StringRef(_var.Vra_cloud_account_azure_name),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupCloudAccountAzure(ctx *pulumi.Context, args *LookupCloudAccountAzureArgs, opts ...pulumi.InvokeOption) (*LookupCloudAccountAzureResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupCloudAccountAzureResult
	err := ctx.Invoke("vra:index/getCloudAccountAzure:getCloudAccountAzure", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudAccountAzure.
type LookupCloudAccountAzureArgs struct {
	// The id of this Azure cloud account.
	Id *string `pulumi:"id"`
	// The name of this Azure cloud account.
	Name *string `pulumi:"name"`
	// A set of tag keys and optional values that were set on this resource.
	// example:[ { "key" : "vmware", "value": "provider" } ]
	Tags []GetCloudAccountAzureTag `pulumi:"tags"`
}

// A collection of values returned by getCloudAccountAzure.
type LookupCloudAccountAzureResult struct {
	// Azure Client Application ID.
	ApplicationId string `pulumi:"applicationId"`
	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt string `pulumi:"createdAt"`
	// A human-friendly description.
	Description string `pulumi:"description"`
	Id          string `pulumi:"id"`
	// HATEOAS of the entity.
	Links []GetCloudAccountAzureLink `pulumi:"links"`
	Name  string                     `pulumi:"name"`
	// The id of the organization this entity belongs to.
	OrgId string `pulumi:"orgId"`
	// Email of the user that owns the entity.
	Owner string `pulumi:"owner"`
	// A set of region names that are enabled for this account.
	Regions []string `pulumi:"regions"`
	// Azure Subscription ID.
	SubscriptionId string `pulumi:"subscriptionId"`
	// A set of tag keys and optional values that were set on this resource.
	// example:[ { "key" : "vmware", "value": "provider" } ]
	Tags []GetCloudAccountAzureTag `pulumi:"tags"`
	// Azure Tenant ID.
	TenantId string `pulumi:"tenantId"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt string `pulumi:"updatedAt"`
}

func LookupCloudAccountAzureOutput(ctx *pulumi.Context, args LookupCloudAccountAzureOutputArgs, opts ...pulumi.InvokeOption) LookupCloudAccountAzureResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCloudAccountAzureResult, error) {
			args := v.(LookupCloudAccountAzureArgs)
			r, err := LookupCloudAccountAzure(ctx, &args, opts...)
			var s LookupCloudAccountAzureResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCloudAccountAzureResultOutput)
}

// A collection of arguments for invoking getCloudAccountAzure.
type LookupCloudAccountAzureOutputArgs struct {
	// The id of this Azure cloud account.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The name of this Azure cloud account.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// A set of tag keys and optional values that were set on this resource.
	// example:[ { "key" : "vmware", "value": "provider" } ]
	Tags GetCloudAccountAzureTagArrayInput `pulumi:"tags"`
}

func (LookupCloudAccountAzureOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudAccountAzureArgs)(nil)).Elem()
}

// A collection of values returned by getCloudAccountAzure.
type LookupCloudAccountAzureResultOutput struct{ *pulumi.OutputState }

func (LookupCloudAccountAzureResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudAccountAzureResult)(nil)).Elem()
}

func (o LookupCloudAccountAzureResultOutput) ToLookupCloudAccountAzureResultOutput() LookupCloudAccountAzureResultOutput {
	return o
}

func (o LookupCloudAccountAzureResultOutput) ToLookupCloudAccountAzureResultOutputWithContext(ctx context.Context) LookupCloudAccountAzureResultOutput {
	return o
}

// Azure Client Application ID.
func (o LookupCloudAccountAzureResultOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountAzureResult) string { return v.ApplicationId }).(pulumi.StringOutput)
}

// Date when the entity was created. The date is in ISO 6801 and UTC.
func (o LookupCloudAccountAzureResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountAzureResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// A human-friendly description.
func (o LookupCloudAccountAzureResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountAzureResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupCloudAccountAzureResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountAzureResult) string { return v.Id }).(pulumi.StringOutput)
}

// HATEOAS of the entity.
func (o LookupCloudAccountAzureResultOutput) Links() GetCloudAccountAzureLinkArrayOutput {
	return o.ApplyT(func(v LookupCloudAccountAzureResult) []GetCloudAccountAzureLink { return v.Links }).(GetCloudAccountAzureLinkArrayOutput)
}

func (o LookupCloudAccountAzureResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountAzureResult) string { return v.Name }).(pulumi.StringOutput)
}

// The id of the organization this entity belongs to.
func (o LookupCloudAccountAzureResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountAzureResult) string { return v.OrgId }).(pulumi.StringOutput)
}

// Email of the user that owns the entity.
func (o LookupCloudAccountAzureResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountAzureResult) string { return v.Owner }).(pulumi.StringOutput)
}

// A set of region names that are enabled for this account.
func (o LookupCloudAccountAzureResultOutput) Regions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCloudAccountAzureResult) []string { return v.Regions }).(pulumi.StringArrayOutput)
}

// Azure Subscription ID.
func (o LookupCloudAccountAzureResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountAzureResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

// A set of tag keys and optional values that were set on this resource.
// example:[ { "key" : "vmware", "value": "provider" } ]
func (o LookupCloudAccountAzureResultOutput) Tags() GetCloudAccountAzureTagArrayOutput {
	return o.ApplyT(func(v LookupCloudAccountAzureResult) []GetCloudAccountAzureTag { return v.Tags }).(GetCloudAccountAzureTagArrayOutput)
}

// Azure Tenant ID.
func (o LookupCloudAccountAzureResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountAzureResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// Date when the entity was last updated. The date is ISO 8601 and UTC.
func (o LookupCloudAccountAzureResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountAzureResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudAccountAzureResultOutput{})
}