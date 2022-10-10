// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securitygroup

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### S
// This is an example of how to lookup security groups.
//
// **Security groups by filter query:**
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-vra/sdk/go/vra/securitygroup"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumiverse/pulumi-vra/sdk/go/vra/securitygroup"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := securitygroup.GetSecurityGroup(ctx, &securitygroup.GetSecurityGroupArgs{
// 			Filter: fmt.Sprintf("name eq '%v'", _var.Name),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// A Security group supports the following arguments:
func GetSecurityGroup(ctx *pulumi.Context, args *GetSecurityGroupArgs, opts ...pulumi.InvokeOption) (*GetSecurityGroupResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSecurityGroupResult
	err := ctx.Invoke("vra:securitygroup/getSecurityGroup:getSecurityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecurityGroup.
type GetSecurityGroupArgs struct {
	// Search criteria to narrow down the Security groups.
	Filter string `pulumi:"filter"`
	// List of security rules.
	Rules []GetSecurityGroupRule `pulumi:"rules"`
}

// A collection of values returned by getSecurityGroup.
type GetSecurityGroupResult struct {
	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt string `pulumi:"createdAt"`
	// A human-friendly description of the security groups.
	Description string `pulumi:"description"`
	// External entity Id on the provider side.
	ExternalId string `pulumi:"externalId"`
	// The id of the region for which this entity is defined.
	ExternalRegionId string `pulumi:"externalRegionId"`
	Filter           string `pulumi:"filter"`
	// ID of the security group.
	Id string `pulumi:"id"`
	// HATEOAS of the entity
	Links []GetSecurityGroupLink `pulumi:"links"`
	// Name of the security group.
	Name string `pulumi:"name"`
	// ID of organization that entity belongs to.
	OrganizationId string `pulumi:"organizationId"`
	Owner          string `pulumi:"owner"`
	// List of security rules.
	Rules []GetSecurityGroupRule `pulumi:"rules"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt string `pulumi:"updatedAt"`
}

func GetSecurityGroupOutput(ctx *pulumi.Context, args GetSecurityGroupOutputArgs, opts ...pulumi.InvokeOption) GetSecurityGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSecurityGroupResult, error) {
			args := v.(GetSecurityGroupArgs)
			r, err := GetSecurityGroup(ctx, &args, opts...)
			var s GetSecurityGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSecurityGroupResultOutput)
}

// A collection of arguments for invoking getSecurityGroup.
type GetSecurityGroupOutputArgs struct {
	// Search criteria to narrow down the Security groups.
	Filter pulumi.StringInput `pulumi:"filter"`
	// List of security rules.
	Rules GetSecurityGroupRuleArrayInput `pulumi:"rules"`
}

func (GetSecurityGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecurityGroupArgs)(nil)).Elem()
}

// A collection of values returned by getSecurityGroup.
type GetSecurityGroupResultOutput struct{ *pulumi.OutputState }

func (GetSecurityGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecurityGroupResult)(nil)).Elem()
}

func (o GetSecurityGroupResultOutput) ToGetSecurityGroupResultOutput() GetSecurityGroupResultOutput {
	return o
}

func (o GetSecurityGroupResultOutput) ToGetSecurityGroupResultOutputWithContext(ctx context.Context) GetSecurityGroupResultOutput {
	return o
}

// Date when the entity was created. The date is in ISO 6801 and UTC.
func (o GetSecurityGroupResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityGroupResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// A human-friendly description of the security groups.
func (o GetSecurityGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

// External entity Id on the provider side.
func (o GetSecurityGroupResultOutput) ExternalId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityGroupResult) string { return v.ExternalId }).(pulumi.StringOutput)
}

// The id of the region for which this entity is defined.
func (o GetSecurityGroupResultOutput) ExternalRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityGroupResult) string { return v.ExternalRegionId }).(pulumi.StringOutput)
}

func (o GetSecurityGroupResultOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityGroupResult) string { return v.Filter }).(pulumi.StringOutput)
}

// ID of the security group.
func (o GetSecurityGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// HATEOAS of the entity
func (o GetSecurityGroupResultOutput) Links() GetSecurityGroupLinkArrayOutput {
	return o.ApplyT(func(v GetSecurityGroupResult) []GetSecurityGroupLink { return v.Links }).(GetSecurityGroupLinkArrayOutput)
}

// Name of the security group.
func (o GetSecurityGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// ID of organization that entity belongs to.
func (o GetSecurityGroupResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityGroupResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o GetSecurityGroupResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityGroupResult) string { return v.Owner }).(pulumi.StringOutput)
}

// List of security rules.
func (o GetSecurityGroupResultOutput) Rules() GetSecurityGroupRuleArrayOutput {
	return o.ApplyT(func(v GetSecurityGroupResult) []GetSecurityGroupRule { return v.Rules }).(GetSecurityGroupRuleArrayOutput)
}

// Date when the entity was last updated. The date is ISO 8601 and UTC.
func (o GetSecurityGroupResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityGroupResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSecurityGroupResultOutput{})
}
