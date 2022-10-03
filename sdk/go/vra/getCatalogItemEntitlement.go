// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vra

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides information about a catalog item entitlement in vRA.
//
// ## Example Usage
// ### S
//
// This is an example of how to get a vRA catalog item entitlement by its id:
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
//			_, err := vra.LookupCatalogItemEntitlement(ctx, &GetCatalogItemEntitlementArgs{
//				Id:        pulumi.StringRef(_var.Catalog_item_entitlement_id),
//				ProjectId: _var.Project_id,
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
// This is an example of how to get a vRA catalog item entitlement by its catalog item id:
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
//			_, err := vra.LookupCatalogItemEntitlement(ctx, &GetCatalogItemEntitlementArgs{
//				CatalogItemId: pulumi.StringRef(_var.Catalog_item_id),
//				ProjectId:     _var.Project_id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupCatalogItemEntitlement(ctx *pulumi.Context, args *LookupCatalogItemEntitlementArgs, opts ...pulumi.InvokeOption) (*LookupCatalogItemEntitlementResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupCatalogItemEntitlementResult
	err := ctx.Invoke("vra:index/getCatalogItemEntitlement:getCatalogItemEntitlement", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCatalogItemEntitlement.
type LookupCatalogItemEntitlementArgs struct {
	// The id of the catalog item to find the entitlement. One of `catalogItemId` or `id` must be provided.
	CatalogItemId *string `pulumi:"catalogItemId"`
	// The id of entitlement. One of `catalogItemId` or `id` must be provided.
	Id *string `pulumi:"id"`
	// The id of the project that this entitlement belongs to.
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getCatalogItemEntitlement.
type LookupCatalogItemEntitlementResult struct {
	CatalogItemId *string `pulumi:"catalogItemId"`
	// Represents a catalog item that is linked to a project via an entitlement.
	Definitions []GetCatalogItemEntitlementDefinition `pulumi:"definitions"`
	// Id of the catalog item.
	Id        *string `pulumi:"id"`
	ProjectId string  `pulumi:"projectId"`
}

func LookupCatalogItemEntitlementOutput(ctx *pulumi.Context, args LookupCatalogItemEntitlementOutputArgs, opts ...pulumi.InvokeOption) LookupCatalogItemEntitlementResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCatalogItemEntitlementResult, error) {
			args := v.(LookupCatalogItemEntitlementArgs)
			r, err := LookupCatalogItemEntitlement(ctx, &args, opts...)
			var s LookupCatalogItemEntitlementResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCatalogItemEntitlementResultOutput)
}

// A collection of arguments for invoking getCatalogItemEntitlement.
type LookupCatalogItemEntitlementOutputArgs struct {
	// The id of the catalog item to find the entitlement. One of `catalogItemId` or `id` must be provided.
	CatalogItemId pulumi.StringPtrInput `pulumi:"catalogItemId"`
	// The id of entitlement. One of `catalogItemId` or `id` must be provided.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The id of the project that this entitlement belongs to.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (LookupCatalogItemEntitlementOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCatalogItemEntitlementArgs)(nil)).Elem()
}

// A collection of values returned by getCatalogItemEntitlement.
type LookupCatalogItemEntitlementResultOutput struct{ *pulumi.OutputState }

func (LookupCatalogItemEntitlementResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCatalogItemEntitlementResult)(nil)).Elem()
}

func (o LookupCatalogItemEntitlementResultOutput) ToLookupCatalogItemEntitlementResultOutput() LookupCatalogItemEntitlementResultOutput {
	return o
}

func (o LookupCatalogItemEntitlementResultOutput) ToLookupCatalogItemEntitlementResultOutputWithContext(ctx context.Context) LookupCatalogItemEntitlementResultOutput {
	return o
}

func (o LookupCatalogItemEntitlementResultOutput) CatalogItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCatalogItemEntitlementResult) *string { return v.CatalogItemId }).(pulumi.StringPtrOutput)
}

// Represents a catalog item that is linked to a project via an entitlement.
func (o LookupCatalogItemEntitlementResultOutput) Definitions() GetCatalogItemEntitlementDefinitionArrayOutput {
	return o.ApplyT(func(v LookupCatalogItemEntitlementResult) []GetCatalogItemEntitlementDefinition { return v.Definitions }).(GetCatalogItemEntitlementDefinitionArrayOutput)
}

// Id of the catalog item.
func (o LookupCatalogItemEntitlementResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCatalogItemEntitlementResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupCatalogItemEntitlementResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCatalogItemEntitlementResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCatalogItemEntitlementResultOutput{})
}