// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudaccount

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VMware vRA cloudaccount.Azure data source.
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
//	"github.com/pulumi/pulumi-vra/sdk/go/vra/cloudaccount"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vra/sdk/go/vra/cloudaccount"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudaccount.LookupAzure(ctx, &cloudaccount.LookupAzureArgs{
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
//	"github.com/pulumi/pulumi-vra/sdk/go/vra/cloudaccount"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vra/sdk/go/vra/cloudaccount"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudaccount.LookupAzure(ctx, &cloudaccount.LookupAzureArgs{
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
func LookupAzure(ctx *pulumi.Context, args *LookupAzureArgs, opts ...pulumi.InvokeOption) (*LookupAzureResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupAzureResult
	err := ctx.Invoke("vra:cloudaccount/getAzure:getAzure", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAzure.
type LookupAzureArgs struct {
	// The id of this Azure cloud account.
	Id *string `pulumi:"id"`
	// The name of this Azure cloud account.
	Name *string `pulumi:"name"`
	// A set of tag keys and optional values that were set on this resource.
	// example:[ { "key" : "vmware", "value": "provider" } ]
	Tags []GetAzureTag `pulumi:"tags"`
}

// A collection of values returned by getAzure.
type LookupAzureResult struct {
	// Azure Client Application ID.
	ApplicationId string `pulumi:"applicationId"`
	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt string `pulumi:"createdAt"`
	// A human-friendly description.
	Description string `pulumi:"description"`
	Id          string `pulumi:"id"`
	// HATEOAS of the entity.
	Links []GetAzureLink `pulumi:"links"`
	Name  string         `pulumi:"name"`
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
	Tags []GetAzureTag `pulumi:"tags"`
	// Azure Tenant ID.
	TenantId string `pulumi:"tenantId"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt string `pulumi:"updatedAt"`
}

func LookupAzureOutput(ctx *pulumi.Context, args LookupAzureOutputArgs, opts ...pulumi.InvokeOption) LookupAzureResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAzureResult, error) {
			args := v.(LookupAzureArgs)
			r, err := LookupAzure(ctx, &args, opts...)
			var s LookupAzureResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAzureResultOutput)
}

// A collection of arguments for invoking getAzure.
type LookupAzureOutputArgs struct {
	// The id of this Azure cloud account.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The name of this Azure cloud account.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// A set of tag keys and optional values that were set on this resource.
	// example:[ { "key" : "vmware", "value": "provider" } ]
	Tags GetAzureTagArrayInput `pulumi:"tags"`
}

func (LookupAzureOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureArgs)(nil)).Elem()
}

// A collection of values returned by getAzure.
type LookupAzureResultOutput struct{ *pulumi.OutputState }

func (LookupAzureResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureResult)(nil)).Elem()
}

func (o LookupAzureResultOutput) ToLookupAzureResultOutput() LookupAzureResultOutput {
	return o
}

func (o LookupAzureResultOutput) ToLookupAzureResultOutputWithContext(ctx context.Context) LookupAzureResultOutput {
	return o
}

// Azure Client Application ID.
func (o LookupAzureResultOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureResult) string { return v.ApplicationId }).(pulumi.StringOutput)
}

// Date when the entity was created. The date is in ISO 6801 and UTC.
func (o LookupAzureResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// A human-friendly description.
func (o LookupAzureResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupAzureResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureResult) string { return v.Id }).(pulumi.StringOutput)
}

// HATEOAS of the entity.
func (o LookupAzureResultOutput) Links() GetAzureLinkArrayOutput {
	return o.ApplyT(func(v LookupAzureResult) []GetAzureLink { return v.Links }).(GetAzureLinkArrayOutput)
}

func (o LookupAzureResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureResult) string { return v.Name }).(pulumi.StringOutput)
}

// The id of the organization this entity belongs to.
func (o LookupAzureResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureResult) string { return v.OrgId }).(pulumi.StringOutput)
}

// Email of the user that owns the entity.
func (o LookupAzureResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureResult) string { return v.Owner }).(pulumi.StringOutput)
}

// A set of region names that are enabled for this account.
func (o LookupAzureResultOutput) Regions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAzureResult) []string { return v.Regions }).(pulumi.StringArrayOutput)
}

// Azure Subscription ID.
func (o LookupAzureResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

// A set of tag keys and optional values that were set on this resource.
// example:[ { "key" : "vmware", "value": "provider" } ]
func (o LookupAzureResultOutput) Tags() GetAzureTagArrayOutput {
	return o.ApplyT(func(v LookupAzureResult) []GetAzureTag { return v.Tags }).(GetAzureTagArrayOutput)
}

// Azure Tenant ID.
func (o LookupAzureResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// Date when the entity was last updated. The date is ISO 8601 and UTC.
func (o LookupAzureResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAzureResultOutput{})
}