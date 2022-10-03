// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vra

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Updates a VMware vRealize Automation fabricDatastoreVsphere resource.
//
// ## Example Usage
// ### S
//
// You cannot create a fabric datastore vSphere resource, however you can import it using the command specified in the import section below.
//
// Once a resource is imported, you can update it as shown below:
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
//			_, err := vra.NewFabricDatastoreVsphere(ctx, "this", &vra.FabricDatastoreVsphereArgs{
//				Tags: FabricDatastoreVsphereTagArray{
//					&FabricDatastoreVsphereTagArgs{
//						Key:   pulumi.String("foo"),
//						Value: pulumi.String("bar"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # To import the fabric datastore vSphere resource, use the ID as in the following example
//
// ```sh
//
//	$ pulumi import vra:index/fabricDatastoreVsphere:FabricDatastoreVsphere this 8e0c9a4c-3ab8-48e8-b9d5-0751c871e282
//
// ```
type FabricDatastoreVsphere struct {
	pulumi.CustomResourceState

	// Set of ids of the cloud accounts this entity belongs to.
	CloudAccountIds pulumi.StringArrayOutput `pulumi:"cloudAccountIds"`
	// Date when the entity was created. The date is in ISO 8601 and UTC.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// A human-friendly description.
	Description pulumi.StringOutput `pulumi:"description"`
	// External entity Id on the provider side.
	ExternalId pulumi.StringOutput `pulumi:"externalId"`
	// Id of datacenter in which the datastore is present.
	ExternalRegionId pulumi.StringOutput `pulumi:"externalRegionId"`
	// Indicates free size available in datastore.
	FreeSizeGb pulumi.StringOutput `pulumi:"freeSizeGb"`
	// HATEOAS of the entity
	Links FabricDatastoreVsphereLinkArrayOutput `pulumi:"links"`
	// A human-friendly name used as an identifier for the vSphere fabric datastore resource instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// The id of the organization this entity belongs to.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// Email of the user that owns the entity.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// A set of tag keys and optional values that were set on this resource:
	Tags FabricDatastoreVsphereTagArrayOutput `pulumi:"tags"`
	// Type of datastore.
	Type pulumi.StringOutput `pulumi:"type"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewFabricDatastoreVsphere registers a new resource with the given unique name, arguments, and options.
func NewFabricDatastoreVsphere(ctx *pulumi.Context,
	name string, args *FabricDatastoreVsphereArgs, opts ...pulumi.ResourceOption) (*FabricDatastoreVsphere, error) {
	if args == nil {
		args = &FabricDatastoreVsphereArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FabricDatastoreVsphere
	err := ctx.RegisterResource("vra:index/fabricDatastoreVsphere:FabricDatastoreVsphere", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFabricDatastoreVsphere gets an existing FabricDatastoreVsphere resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFabricDatastoreVsphere(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FabricDatastoreVsphereState, opts ...pulumi.ResourceOption) (*FabricDatastoreVsphere, error) {
	var resource FabricDatastoreVsphere
	err := ctx.ReadResource("vra:index/fabricDatastoreVsphere:FabricDatastoreVsphere", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FabricDatastoreVsphere resources.
type fabricDatastoreVsphereState struct {
	// Set of ids of the cloud accounts this entity belongs to.
	CloudAccountIds []string `pulumi:"cloudAccountIds"`
	// Date when the entity was created. The date is in ISO 8601 and UTC.
	CreatedAt *string `pulumi:"createdAt"`
	// A human-friendly description.
	Description *string `pulumi:"description"`
	// External entity Id on the provider side.
	ExternalId *string `pulumi:"externalId"`
	// Id of datacenter in which the datastore is present.
	ExternalRegionId *string `pulumi:"externalRegionId"`
	// Indicates free size available in datastore.
	FreeSizeGb *string `pulumi:"freeSizeGb"`
	// HATEOAS of the entity
	Links []FabricDatastoreVsphereLink `pulumi:"links"`
	// A human-friendly name used as an identifier for the vSphere fabric datastore resource instance.
	Name *string `pulumi:"name"`
	// The id of the organization this entity belongs to.
	OrgId *string `pulumi:"orgId"`
	// Email of the user that owns the entity.
	Owner *string `pulumi:"owner"`
	// A set of tag keys and optional values that were set on this resource:
	Tags []FabricDatastoreVsphereTag `pulumi:"tags"`
	// Type of datastore.
	Type *string `pulumi:"type"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type FabricDatastoreVsphereState struct {
	// Set of ids of the cloud accounts this entity belongs to.
	CloudAccountIds pulumi.StringArrayInput
	// Date when the entity was created. The date is in ISO 8601 and UTC.
	CreatedAt pulumi.StringPtrInput
	// A human-friendly description.
	Description pulumi.StringPtrInput
	// External entity Id on the provider side.
	ExternalId pulumi.StringPtrInput
	// Id of datacenter in which the datastore is present.
	ExternalRegionId pulumi.StringPtrInput
	// Indicates free size available in datastore.
	FreeSizeGb pulumi.StringPtrInput
	// HATEOAS of the entity
	Links FabricDatastoreVsphereLinkArrayInput
	// A human-friendly name used as an identifier for the vSphere fabric datastore resource instance.
	Name pulumi.StringPtrInput
	// The id of the organization this entity belongs to.
	OrgId pulumi.StringPtrInput
	// Email of the user that owns the entity.
	Owner pulumi.StringPtrInput
	// A set of tag keys and optional values that were set on this resource:
	Tags FabricDatastoreVsphereTagArrayInput
	// Type of datastore.
	Type pulumi.StringPtrInput
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt pulumi.StringPtrInput
}

func (FabricDatastoreVsphereState) ElementType() reflect.Type {
	return reflect.TypeOf((*fabricDatastoreVsphereState)(nil)).Elem()
}

type fabricDatastoreVsphereArgs struct {
	// A set of tag keys and optional values that were set on this resource:
	Tags []FabricDatastoreVsphereTag `pulumi:"tags"`
}

// The set of arguments for constructing a FabricDatastoreVsphere resource.
type FabricDatastoreVsphereArgs struct {
	// A set of tag keys and optional values that were set on this resource:
	Tags FabricDatastoreVsphereTagArrayInput
}

func (FabricDatastoreVsphereArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fabricDatastoreVsphereArgs)(nil)).Elem()
}

type FabricDatastoreVsphereInput interface {
	pulumi.Input

	ToFabricDatastoreVsphereOutput() FabricDatastoreVsphereOutput
	ToFabricDatastoreVsphereOutputWithContext(ctx context.Context) FabricDatastoreVsphereOutput
}

func (*FabricDatastoreVsphere) ElementType() reflect.Type {
	return reflect.TypeOf((**FabricDatastoreVsphere)(nil)).Elem()
}

func (i *FabricDatastoreVsphere) ToFabricDatastoreVsphereOutput() FabricDatastoreVsphereOutput {
	return i.ToFabricDatastoreVsphereOutputWithContext(context.Background())
}

func (i *FabricDatastoreVsphere) ToFabricDatastoreVsphereOutputWithContext(ctx context.Context) FabricDatastoreVsphereOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FabricDatastoreVsphereOutput)
}

// FabricDatastoreVsphereArrayInput is an input type that accepts FabricDatastoreVsphereArray and FabricDatastoreVsphereArrayOutput values.
// You can construct a concrete instance of `FabricDatastoreVsphereArrayInput` via:
//
//	FabricDatastoreVsphereArray{ FabricDatastoreVsphereArgs{...} }
type FabricDatastoreVsphereArrayInput interface {
	pulumi.Input

	ToFabricDatastoreVsphereArrayOutput() FabricDatastoreVsphereArrayOutput
	ToFabricDatastoreVsphereArrayOutputWithContext(context.Context) FabricDatastoreVsphereArrayOutput
}

type FabricDatastoreVsphereArray []FabricDatastoreVsphereInput

func (FabricDatastoreVsphereArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FabricDatastoreVsphere)(nil)).Elem()
}

func (i FabricDatastoreVsphereArray) ToFabricDatastoreVsphereArrayOutput() FabricDatastoreVsphereArrayOutput {
	return i.ToFabricDatastoreVsphereArrayOutputWithContext(context.Background())
}

func (i FabricDatastoreVsphereArray) ToFabricDatastoreVsphereArrayOutputWithContext(ctx context.Context) FabricDatastoreVsphereArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FabricDatastoreVsphereArrayOutput)
}

// FabricDatastoreVsphereMapInput is an input type that accepts FabricDatastoreVsphereMap and FabricDatastoreVsphereMapOutput values.
// You can construct a concrete instance of `FabricDatastoreVsphereMapInput` via:
//
//	FabricDatastoreVsphereMap{ "key": FabricDatastoreVsphereArgs{...} }
type FabricDatastoreVsphereMapInput interface {
	pulumi.Input

	ToFabricDatastoreVsphereMapOutput() FabricDatastoreVsphereMapOutput
	ToFabricDatastoreVsphereMapOutputWithContext(context.Context) FabricDatastoreVsphereMapOutput
}

type FabricDatastoreVsphereMap map[string]FabricDatastoreVsphereInput

func (FabricDatastoreVsphereMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FabricDatastoreVsphere)(nil)).Elem()
}

func (i FabricDatastoreVsphereMap) ToFabricDatastoreVsphereMapOutput() FabricDatastoreVsphereMapOutput {
	return i.ToFabricDatastoreVsphereMapOutputWithContext(context.Background())
}

func (i FabricDatastoreVsphereMap) ToFabricDatastoreVsphereMapOutputWithContext(ctx context.Context) FabricDatastoreVsphereMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FabricDatastoreVsphereMapOutput)
}

type FabricDatastoreVsphereOutput struct{ *pulumi.OutputState }

func (FabricDatastoreVsphereOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FabricDatastoreVsphere)(nil)).Elem()
}

func (o FabricDatastoreVsphereOutput) ToFabricDatastoreVsphereOutput() FabricDatastoreVsphereOutput {
	return o
}

func (o FabricDatastoreVsphereOutput) ToFabricDatastoreVsphereOutputWithContext(ctx context.Context) FabricDatastoreVsphereOutput {
	return o
}

// Set of ids of the cloud accounts this entity belongs to.
func (o FabricDatastoreVsphereOutput) CloudAccountIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FabricDatastoreVsphere) pulumi.StringArrayOutput { return v.CloudAccountIds }).(pulumi.StringArrayOutput)
}

// Date when the entity was created. The date is in ISO 8601 and UTC.
func (o FabricDatastoreVsphereOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *FabricDatastoreVsphere) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// A human-friendly description.
func (o FabricDatastoreVsphereOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *FabricDatastoreVsphere) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// External entity Id on the provider side.
func (o FabricDatastoreVsphereOutput) ExternalId() pulumi.StringOutput {
	return o.ApplyT(func(v *FabricDatastoreVsphere) pulumi.StringOutput { return v.ExternalId }).(pulumi.StringOutput)
}

// Id of datacenter in which the datastore is present.
func (o FabricDatastoreVsphereOutput) ExternalRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v *FabricDatastoreVsphere) pulumi.StringOutput { return v.ExternalRegionId }).(pulumi.StringOutput)
}

// Indicates free size available in datastore.
func (o FabricDatastoreVsphereOutput) FreeSizeGb() pulumi.StringOutput {
	return o.ApplyT(func(v *FabricDatastoreVsphere) pulumi.StringOutput { return v.FreeSizeGb }).(pulumi.StringOutput)
}

// HATEOAS of the entity
func (o FabricDatastoreVsphereOutput) Links() FabricDatastoreVsphereLinkArrayOutput {
	return o.ApplyT(func(v *FabricDatastoreVsphere) FabricDatastoreVsphereLinkArrayOutput { return v.Links }).(FabricDatastoreVsphereLinkArrayOutput)
}

// A human-friendly name used as an identifier for the vSphere fabric datastore resource instance.
func (o FabricDatastoreVsphereOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FabricDatastoreVsphere) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The id of the organization this entity belongs to.
func (o FabricDatastoreVsphereOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *FabricDatastoreVsphere) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// Email of the user that owns the entity.
func (o FabricDatastoreVsphereOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *FabricDatastoreVsphere) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// A set of tag keys and optional values that were set on this resource:
func (o FabricDatastoreVsphereOutput) Tags() FabricDatastoreVsphereTagArrayOutput {
	return o.ApplyT(func(v *FabricDatastoreVsphere) FabricDatastoreVsphereTagArrayOutput { return v.Tags }).(FabricDatastoreVsphereTagArrayOutput)
}

// Type of datastore.
func (o FabricDatastoreVsphereOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FabricDatastoreVsphere) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Date when the entity was last updated. The date is ISO 8601 and UTC.
func (o FabricDatastoreVsphereOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *FabricDatastoreVsphere) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type FabricDatastoreVsphereArrayOutput struct{ *pulumi.OutputState }

func (FabricDatastoreVsphereArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FabricDatastoreVsphere)(nil)).Elem()
}

func (o FabricDatastoreVsphereArrayOutput) ToFabricDatastoreVsphereArrayOutput() FabricDatastoreVsphereArrayOutput {
	return o
}

func (o FabricDatastoreVsphereArrayOutput) ToFabricDatastoreVsphereArrayOutputWithContext(ctx context.Context) FabricDatastoreVsphereArrayOutput {
	return o
}

func (o FabricDatastoreVsphereArrayOutput) Index(i pulumi.IntInput) FabricDatastoreVsphereOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FabricDatastoreVsphere {
		return vs[0].([]*FabricDatastoreVsphere)[vs[1].(int)]
	}).(FabricDatastoreVsphereOutput)
}

type FabricDatastoreVsphereMapOutput struct{ *pulumi.OutputState }

func (FabricDatastoreVsphereMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FabricDatastoreVsphere)(nil)).Elem()
}

func (o FabricDatastoreVsphereMapOutput) ToFabricDatastoreVsphereMapOutput() FabricDatastoreVsphereMapOutput {
	return o
}

func (o FabricDatastoreVsphereMapOutput) ToFabricDatastoreVsphereMapOutputWithContext(ctx context.Context) FabricDatastoreVsphereMapOutput {
	return o
}

func (o FabricDatastoreVsphereMapOutput) MapIndex(k pulumi.StringInput) FabricDatastoreVsphereOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FabricDatastoreVsphere {
		return vs[0].(map[string]*FabricDatastoreVsphere)[vs[1].(string)]
	}).(FabricDatastoreVsphereOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FabricDatastoreVsphereInput)(nil)).Elem(), &FabricDatastoreVsphere{})
	pulumi.RegisterInputType(reflect.TypeOf((*FabricDatastoreVsphereArrayInput)(nil)).Elem(), FabricDatastoreVsphereArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FabricDatastoreVsphereMapInput)(nil)).Elem(), FabricDatastoreVsphereMap{})
	pulumi.RegisterOutputType(FabricDatastoreVsphereOutput{})
	pulumi.RegisterOutputType(FabricDatastoreVsphereArrayOutput{})
	pulumi.RegisterOutputType(FabricDatastoreVsphereMapOutput{})
}