// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fabric

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### S
//
// This is an example of how to lookup fabric computes.
//
// **Fabric compute data source by Id:**
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vra/sdk/go/vra/fabric"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumiverse/pulumi-vra/sdk/go/vra/fabric"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fabric.LookupCompute(ctx, &fabric.LookupComputeArgs{
// 			Id: pulumi.StringRef(_var.Fabric_compute_id),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// **Fabric compute data source by filter query:**
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-vra/sdk/go/vra/fabric"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumiverse/pulumi-vra/sdk/go/vra/fabric"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fabric.LookupCompute(ctx, &fabric.LookupComputeArgs{
// 			Filter: pulumi.StringRef(fmt.Sprintf("name eq '%v'", _var.Fabric_compute_name)),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// A fabric compute data source supports the following arguments:
func LookupCompute(ctx *pulumi.Context, args *LookupComputeArgs, opts ...pulumi.InvokeOption) (*LookupComputeResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupComputeResult
	err := ctx.Invoke("vra:fabric/getCompute:getCompute", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCompute.
type LookupComputeArgs struct {
	// Search criteria to narrow down the fabric compute resource instance. Only one of 'id' or 'filter' must be specified.
	Filter *string `pulumi:"filter"`
	// The id of the fabric compute resource instance. Only one of 'id' or 'filter' must be specified.
	Id *string `pulumi:"id"`
	// A set of tag keys and optional values that were set on this resource:
	Tags []GetComputeTag `pulumi:"tags"`
}

// A collection of values returned by getCompute.
type LookupComputeResult struct {
	// Date when the entity was created. The date is in ISO 8601 and UTC.
	CreatedAt string `pulumi:"createdAt"`
	// A list of key value pair of custom properties for the fabric compute resource.
	CustomProperties map[string]interface{} `pulumi:"customProperties"`
	// A human-friendly description.
	Description string `pulumi:"description"`
	// The id of the external entity on the provider side.
	ExternalId string `pulumi:"externalId"`
	// The external region id of the fabric compute.
	ExternalRegionId string `pulumi:"externalRegionId"`
	// The external zone id of the fabric compute.
	ExternalZoneId string  `pulumi:"externalZoneId"`
	Filter         *string `pulumi:"filter"`
	Id             string  `pulumi:"id"`
	// Lifecycle status of the compute instance.
	LifecycleState string `pulumi:"lifecycleState"`
	// HATEOAS of the entity.
	Links []GetComputeLink `pulumi:"links"`
	// A human-friendly name used as an identifier for the fabric compute resource instance.
	Name string `pulumi:"name"`
	// The id of the organization this entity belongs to.
	OrgId string `pulumi:"orgId"`
	// Email of the user that owns the entity.
	Owner string `pulumi:"owner"`
	// Power state of fabric compute instance.
	PowerState string `pulumi:"powerState"`
	// A set of tag keys and optional values that were set on this resource:
	Tags []GetComputeTag `pulumi:"tags"`
	// Type of the fabric compute instance.
	Type string `pulumi:"type"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt string `pulumi:"updatedAt"`
}

func LookupComputeOutput(ctx *pulumi.Context, args LookupComputeOutputArgs, opts ...pulumi.InvokeOption) LookupComputeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupComputeResult, error) {
			args := v.(LookupComputeArgs)
			r, err := LookupCompute(ctx, &args, opts...)
			var s LookupComputeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupComputeResultOutput)
}

// A collection of arguments for invoking getCompute.
type LookupComputeOutputArgs struct {
	// Search criteria to narrow down the fabric compute resource instance. Only one of 'id' or 'filter' must be specified.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// The id of the fabric compute resource instance. Only one of 'id' or 'filter' must be specified.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// A set of tag keys and optional values that were set on this resource:
	Tags GetComputeTagArrayInput `pulumi:"tags"`
}

func (LookupComputeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeArgs)(nil)).Elem()
}

// A collection of values returned by getCompute.
type LookupComputeResultOutput struct{ *pulumi.OutputState }

func (LookupComputeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeResult)(nil)).Elem()
}

func (o LookupComputeResultOutput) ToLookupComputeResultOutput() LookupComputeResultOutput {
	return o
}

func (o LookupComputeResultOutput) ToLookupComputeResultOutputWithContext(ctx context.Context) LookupComputeResultOutput {
	return o
}

// Date when the entity was created. The date is in ISO 8601 and UTC.
func (o LookupComputeResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// A list of key value pair of custom properties for the fabric compute resource.
func (o LookupComputeResultOutput) CustomProperties() pulumi.MapOutput {
	return o.ApplyT(func(v LookupComputeResult) map[string]interface{} { return v.CustomProperties }).(pulumi.MapOutput)
}

// A human-friendly description.
func (o LookupComputeResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeResult) string { return v.Description }).(pulumi.StringOutput)
}

// The id of the external entity on the provider side.
func (o LookupComputeResultOutput) ExternalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeResult) string { return v.ExternalId }).(pulumi.StringOutput)
}

// The external region id of the fabric compute.
func (o LookupComputeResultOutput) ExternalRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeResult) string { return v.ExternalRegionId }).(pulumi.StringOutput)
}

// The external zone id of the fabric compute.
func (o LookupComputeResultOutput) ExternalZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeResult) string { return v.ExternalZoneId }).(pulumi.StringOutput)
}

func (o LookupComputeResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComputeResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

func (o LookupComputeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeResult) string { return v.Id }).(pulumi.StringOutput)
}

// Lifecycle status of the compute instance.
func (o LookupComputeResultOutput) LifecycleState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeResult) string { return v.LifecycleState }).(pulumi.StringOutput)
}

// HATEOAS of the entity.
func (o LookupComputeResultOutput) Links() GetComputeLinkArrayOutput {
	return o.ApplyT(func(v LookupComputeResult) []GetComputeLink { return v.Links }).(GetComputeLinkArrayOutput)
}

// A human-friendly name used as an identifier for the fabric compute resource instance.
func (o LookupComputeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeResult) string { return v.Name }).(pulumi.StringOutput)
}

// The id of the organization this entity belongs to.
func (o LookupComputeResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeResult) string { return v.OrgId }).(pulumi.StringOutput)
}

// Email of the user that owns the entity.
func (o LookupComputeResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeResult) string { return v.Owner }).(pulumi.StringOutput)
}

// Power state of fabric compute instance.
func (o LookupComputeResultOutput) PowerState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeResult) string { return v.PowerState }).(pulumi.StringOutput)
}

// A set of tag keys and optional values that were set on this resource:
func (o LookupComputeResultOutput) Tags() GetComputeTagArrayOutput {
	return o.ApplyT(func(v LookupComputeResult) []GetComputeTag { return v.Tags }).(GetComputeTagArrayOutput)
}

// Type of the fabric compute instance.
func (o LookupComputeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeResult) string { return v.Type }).(pulumi.StringOutput)
}

// Date when the entity was last updated. The date is ISO 8601 and UTC.
func (o LookupComputeResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComputeResultOutput{})
}
