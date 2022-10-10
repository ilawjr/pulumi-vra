// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## ---layout: "vra"
//
// page_title: "VMware vRealize Automation: network.Network"
// description: |-
//   Provides a data lookup for vra_network.
// ---
//
// # Data Source: network.Network
// ## Example Usage
// ### S
//
// This is an example of how to read a network data source.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vra/sdk/go/vra/network"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumiverse/pulumi-vra/sdk/go/vra/network"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := network.LookupNetwork(ctx, &network.LookupNetworkArgs{
// 			Name: pulumi.StringRef(_var.Network_name),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupNetwork(ctx *pulumi.Context, args *LookupNetworkArgs, opts ...pulumi.InvokeOption) (*LookupNetworkResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupNetworkResult
	err := ctx.Invoke("vra:network/getNetwork:getNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetwork.
type LookupNetworkArgs struct {
	// List of storage, network and extensibility constraints to be applied when provisioning through this project.
	Constraints []GetNetworkConstraint `pulumi:"constraints"`
	// The id of the image profile instance.
	Id *string `pulumi:"id"`
	// A human-friendly name used as an identifier in APIs that support this option.
	Name *string `pulumi:"name"`
	// A set of tag keys and optional values that were set on this resource.
	// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
	Tags []GetNetworkTag `pulumi:"tags"`
}

// A collection of values returned by getNetwork.
type LookupNetworkResult struct {
	// IPv4 address range of the network in CIDR format.
	Cidr string `pulumi:"cidr"`
	// List of storage, network and extensibility constraints to be applied when provisioning through this project.
	Constraints []GetNetworkConstraint `pulumi:"constraints"`
	// Additional properties that may be used to extend the base resource.
	CustomProperties map[string]interface{} `pulumi:"customProperties"`
	// Deployment id that is associated with this resource.
	DeploymentId string `pulumi:"deploymentId"`
	// A human-friendly description.
	Description string `pulumi:"description"`
	// External entity Id on the provider side.
	ExternalId string `pulumi:"externalId"`
	// The external zoneId of the resource.
	ExternalZoneId string `pulumi:"externalZoneId"`
	Id             string `pulumi:"id"`
	// HATEOAS of the entity
	Links []GetNetworkLink `pulumi:"links"`
	Name  string           `pulumi:"name"`
	// The id of the organization this entity belongs to.
	OrganizationId string `pulumi:"organizationId"`
	// Flag to indicate if the network needs to have outbound access or not. Default is true. This field will be ignored if there is proper input for networkType customProperty
	OutboundAccess bool `pulumi:"outboundAccess"`
	// Email of the user that owns the entity.
	Owner string `pulumi:"owner"`
	// The id of the project this resource belongs to.
	ProjectId string `pulumi:"projectId"`
	// Self link of this request.
	SelfLink string `pulumi:"selfLink"`
	// A set of tag keys and optional values that were set on this resource.
	// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
	Tags      []GetNetworkTag `pulumi:"tags"`
	UpdatedAt string          `pulumi:"updatedAt"`
}

func LookupNetworkOutput(ctx *pulumi.Context, args LookupNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkResult, error) {
			args := v.(LookupNetworkArgs)
			r, err := LookupNetwork(ctx, &args, opts...)
			var s LookupNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkResultOutput)
}

// A collection of arguments for invoking getNetwork.
type LookupNetworkOutputArgs struct {
	// List of storage, network and extensibility constraints to be applied when provisioning through this project.
	Constraints GetNetworkConstraintArrayInput `pulumi:"constraints"`
	// The id of the image profile instance.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// A human-friendly name used as an identifier in APIs that support this option.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// A set of tag keys and optional values that were set on this resource.
	// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
	Tags GetNetworkTagArrayInput `pulumi:"tags"`
}

func (LookupNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkArgs)(nil)).Elem()
}

// A collection of values returned by getNetwork.
type LookupNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkResult)(nil)).Elem()
}

func (o LookupNetworkResultOutput) ToLookupNetworkResultOutput() LookupNetworkResultOutput {
	return o
}

func (o LookupNetworkResultOutput) ToLookupNetworkResultOutputWithContext(ctx context.Context) LookupNetworkResultOutput {
	return o
}

// IPv4 address range of the network in CIDR format.
func (o LookupNetworkResultOutput) Cidr() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Cidr }).(pulumi.StringOutput)
}

// List of storage, network and extensibility constraints to be applied when provisioning through this project.
func (o LookupNetworkResultOutput) Constraints() GetNetworkConstraintArrayOutput {
	return o.ApplyT(func(v LookupNetworkResult) []GetNetworkConstraint { return v.Constraints }).(GetNetworkConstraintArrayOutput)
}

// Additional properties that may be used to extend the base resource.
func (o LookupNetworkResultOutput) CustomProperties() pulumi.MapOutput {
	return o.ApplyT(func(v LookupNetworkResult) map[string]interface{} { return v.CustomProperties }).(pulumi.MapOutput)
}

// Deployment id that is associated with this resource.
func (o LookupNetworkResultOutput) DeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.DeploymentId }).(pulumi.StringOutput)
}

// A human-friendly description.
func (o LookupNetworkResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Description }).(pulumi.StringOutput)
}

// External entity Id on the provider side.
func (o LookupNetworkResultOutput) ExternalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.ExternalId }).(pulumi.StringOutput)
}

// The external zoneId of the resource.
func (o LookupNetworkResultOutput) ExternalZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.ExternalZoneId }).(pulumi.StringOutput)
}

func (o LookupNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

// HATEOAS of the entity
func (o LookupNetworkResultOutput) Links() GetNetworkLinkArrayOutput {
	return o.ApplyT(func(v LookupNetworkResult) []GetNetworkLink { return v.Links }).(GetNetworkLinkArrayOutput)
}

func (o LookupNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

// The id of the organization this entity belongs to.
func (o LookupNetworkResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

// Flag to indicate if the network needs to have outbound access or not. Default is true. This field will be ignored if there is proper input for networkType customProperty
func (o LookupNetworkResultOutput) OutboundAccess() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNetworkResult) bool { return v.OutboundAccess }).(pulumi.BoolOutput)
}

// Email of the user that owns the entity.
func (o LookupNetworkResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Owner }).(pulumi.StringOutput)
}

// The id of the project this resource belongs to.
func (o LookupNetworkResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// Self link of this request.
func (o LookupNetworkResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// A set of tag keys and optional values that were set on this resource.
// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
func (o LookupNetworkResultOutput) Tags() GetNetworkTagArrayOutput {
	return o.ApplyT(func(v LookupNetworkResult) []GetNetworkTag { return v.Tags }).(GetNetworkTagArrayOutput)
}

func (o LookupNetworkResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkResultOutput{})
}
