// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package flavor

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### S
// This is an example of how to create a flavor profile resource.
//
// **Flavor profile:**
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vra/sdk/go/vra/flavor"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vra/sdk/go/vra/flavor"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := flavor.NewProfile(ctx, "my-flavor-profile", &flavor.ProfileArgs{
//				Description: pulumi.String("my flavor"),
//				FlavorMappings: flavor.ProfileFlavorMappingArray{
//					&flavor.ProfileFlavorMappingArgs{
//						InstanceType: pulumi.String("t2.small"),
//						Name:         pulumi.String("small"),
//					},
//					&flavor.ProfileFlavorMappingArgs{
//						InstanceType: pulumi.String("t2.medium"),
//						Name:         pulumi.String("medium"),
//					},
//				},
//				RegionId: pulumi.Any(data.Vra_region.Us - east - 1 - region.Id),
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
// An flavor profile resource supports the following arguments:
type Profile struct {
	pulumi.CustomResourceState

	// The ID of the cloud account this entity belongs to.
	CloudAccountId pulumi.StringOutput `pulumi:"cloudAccountId"`
	// Date when  entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// A human-friendly description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of the external region that is associated with the flavor profile.
	ExternalRegionId pulumi.StringOutput `pulumi:"externalRegionId"`
	// Map between global fabric flavor keys and fabric flavor descriptions.
	FlavorMappings ProfileFlavorMappingArrayOutput `pulumi:"flavorMappings"`
	// HATEOAS of entity.
	Links ProfileLinkArrayOutput `pulumi:"links"`
	// A human-friendly name used as an identifier in APIs that support this option.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of organization that entity belongs to.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// Email of entity owner.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// The id of the region for which this profile is defined as in vRealize Automation(vRA).
	RegionId pulumi.StringOutput `pulumi:"regionId"`
	// Date when the entity was last updated. Date and time format is ISO 8601 and UTC.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewProfile registers a new resource with the given unique name, arguments, and options.
func NewProfile(ctx *pulumi.Context,
	name string, args *ProfileArgs, opts ...pulumi.ResourceOption) (*Profile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RegionId == nil {
		return nil, errors.New("invalid value for required argument 'RegionId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Profile
	err := ctx.RegisterResource("vra:flavor/profile:Profile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProfile gets an existing Profile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfileState, opts ...pulumi.ResourceOption) (*Profile, error) {
	var resource Profile
	err := ctx.ReadResource("vra:flavor/profile:Profile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Profile resources.
type profileState struct {
	// The ID of the cloud account this entity belongs to.
	CloudAccountId *string `pulumi:"cloudAccountId"`
	// Date when  entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt *string `pulumi:"createdAt"`
	// A human-friendly description.
	Description *string `pulumi:"description"`
	// The ID of the external region that is associated with the flavor profile.
	ExternalRegionId *string `pulumi:"externalRegionId"`
	// Map between global fabric flavor keys and fabric flavor descriptions.
	FlavorMappings []ProfileFlavorMapping `pulumi:"flavorMappings"`
	// HATEOAS of entity.
	Links []ProfileLink `pulumi:"links"`
	// A human-friendly name used as an identifier in APIs that support this option.
	Name *string `pulumi:"name"`
	// ID of organization that entity belongs to.
	OrgId *string `pulumi:"orgId"`
	// Email of entity owner.
	Owner *string `pulumi:"owner"`
	// The id of the region for which this profile is defined as in vRealize Automation(vRA).
	RegionId *string `pulumi:"regionId"`
	// Date when the entity was last updated. Date and time format is ISO 8601 and UTC.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type ProfileState struct {
	// The ID of the cloud account this entity belongs to.
	CloudAccountId pulumi.StringPtrInput
	// Date when  entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt pulumi.StringPtrInput
	// A human-friendly description.
	Description pulumi.StringPtrInput
	// The ID of the external region that is associated with the flavor profile.
	ExternalRegionId pulumi.StringPtrInput
	// Map between global fabric flavor keys and fabric flavor descriptions.
	FlavorMappings ProfileFlavorMappingArrayInput
	// HATEOAS of entity.
	Links ProfileLinkArrayInput
	// A human-friendly name used as an identifier in APIs that support this option.
	Name pulumi.StringPtrInput
	// ID of organization that entity belongs to.
	OrgId pulumi.StringPtrInput
	// Email of entity owner.
	Owner pulumi.StringPtrInput
	// The id of the region for which this profile is defined as in vRealize Automation(vRA).
	RegionId pulumi.StringPtrInput
	// Date when the entity was last updated. Date and time format is ISO 8601 and UTC.
	UpdatedAt pulumi.StringPtrInput
}

func (ProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*profileState)(nil)).Elem()
}

type profileArgs struct {
	// A human-friendly description.
	Description *string `pulumi:"description"`
	// Map between global fabric flavor keys and fabric flavor descriptions.
	FlavorMappings []ProfileFlavorMapping `pulumi:"flavorMappings"`
	// A human-friendly name used as an identifier in APIs that support this option.
	Name *string `pulumi:"name"`
	// The id of the region for which this profile is defined as in vRealize Automation(vRA).
	RegionId string `pulumi:"regionId"`
}

// The set of arguments for constructing a Profile resource.
type ProfileArgs struct {
	// A human-friendly description.
	Description pulumi.StringPtrInput
	// Map between global fabric flavor keys and fabric flavor descriptions.
	FlavorMappings ProfileFlavorMappingArrayInput
	// A human-friendly name used as an identifier in APIs that support this option.
	Name pulumi.StringPtrInput
	// The id of the region for which this profile is defined as in vRealize Automation(vRA).
	RegionId pulumi.StringInput
}

func (ProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*profileArgs)(nil)).Elem()
}

type ProfileInput interface {
	pulumi.Input

	ToProfileOutput() ProfileOutput
	ToProfileOutputWithContext(ctx context.Context) ProfileOutput
}

func (*Profile) ElementType() reflect.Type {
	return reflect.TypeOf((**Profile)(nil)).Elem()
}

func (i *Profile) ToProfileOutput() ProfileOutput {
	return i.ToProfileOutputWithContext(context.Background())
}

func (i *Profile) ToProfileOutputWithContext(ctx context.Context) ProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileOutput)
}

// ProfileArrayInput is an input type that accepts ProfileArray and ProfileArrayOutput values.
// You can construct a concrete instance of `ProfileArrayInput` via:
//
//	ProfileArray{ ProfileArgs{...} }
type ProfileArrayInput interface {
	pulumi.Input

	ToProfileArrayOutput() ProfileArrayOutput
	ToProfileArrayOutputWithContext(context.Context) ProfileArrayOutput
}

type ProfileArray []ProfileInput

func (ProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Profile)(nil)).Elem()
}

func (i ProfileArray) ToProfileArrayOutput() ProfileArrayOutput {
	return i.ToProfileArrayOutputWithContext(context.Background())
}

func (i ProfileArray) ToProfileArrayOutputWithContext(ctx context.Context) ProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileArrayOutput)
}

// ProfileMapInput is an input type that accepts ProfileMap and ProfileMapOutput values.
// You can construct a concrete instance of `ProfileMapInput` via:
//
//	ProfileMap{ "key": ProfileArgs{...} }
type ProfileMapInput interface {
	pulumi.Input

	ToProfileMapOutput() ProfileMapOutput
	ToProfileMapOutputWithContext(context.Context) ProfileMapOutput
}

type ProfileMap map[string]ProfileInput

func (ProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Profile)(nil)).Elem()
}

func (i ProfileMap) ToProfileMapOutput() ProfileMapOutput {
	return i.ToProfileMapOutputWithContext(context.Background())
}

func (i ProfileMap) ToProfileMapOutputWithContext(ctx context.Context) ProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileMapOutput)
}

type ProfileOutput struct{ *pulumi.OutputState }

func (ProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Profile)(nil)).Elem()
}

func (o ProfileOutput) ToProfileOutput() ProfileOutput {
	return o
}

func (o ProfileOutput) ToProfileOutputWithContext(ctx context.Context) ProfileOutput {
	return o
}

// The ID of the cloud account this entity belongs to.
func (o ProfileOutput) CloudAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.CloudAccountId }).(pulumi.StringOutput)
}

// Date when  entity was created. Date and time format is ISO 8601 and UTC.
func (o ProfileOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// A human-friendly description.
func (o ProfileOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the external region that is associated with the flavor profile.
func (o ProfileOutput) ExternalRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.ExternalRegionId }).(pulumi.StringOutput)
}

// Map between global fabric flavor keys and fabric flavor descriptions.
func (o ProfileOutput) FlavorMappings() ProfileFlavorMappingArrayOutput {
	return o.ApplyT(func(v *Profile) ProfileFlavorMappingArrayOutput { return v.FlavorMappings }).(ProfileFlavorMappingArrayOutput)
}

// HATEOAS of entity.
func (o ProfileOutput) Links() ProfileLinkArrayOutput {
	return o.ApplyT(func(v *Profile) ProfileLinkArrayOutput { return v.Links }).(ProfileLinkArrayOutput)
}

// A human-friendly name used as an identifier in APIs that support this option.
func (o ProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of organization that entity belongs to.
func (o ProfileOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// Email of entity owner.
func (o ProfileOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// The id of the region for which this profile is defined as in vRealize Automation(vRA).
func (o ProfileOutput) RegionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.RegionId }).(pulumi.StringOutput)
}

// Date when the entity was last updated. Date and time format is ISO 8601 and UTC.
func (o ProfileOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type ProfileArrayOutput struct{ *pulumi.OutputState }

func (ProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Profile)(nil)).Elem()
}

func (o ProfileArrayOutput) ToProfileArrayOutput() ProfileArrayOutput {
	return o
}

func (o ProfileArrayOutput) ToProfileArrayOutputWithContext(ctx context.Context) ProfileArrayOutput {
	return o
}

func (o ProfileArrayOutput) Index(i pulumi.IntInput) ProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Profile {
		return vs[0].([]*Profile)[vs[1].(int)]
	}).(ProfileOutput)
}

type ProfileMapOutput struct{ *pulumi.OutputState }

func (ProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Profile)(nil)).Elem()
}

func (o ProfileMapOutput) ToProfileMapOutput() ProfileMapOutput {
	return o
}

func (o ProfileMapOutput) ToProfileMapOutputWithContext(ctx context.Context) ProfileMapOutput {
	return o
}

func (o ProfileMapOutput) MapIndex(k pulumi.StringInput) ProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Profile {
		return vs[0].(map[string]*Profile)[vs[1].(string)]
	}).(ProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileInput)(nil)).Elem(), &Profile{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileArrayInput)(nil)).Elem(), ProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileMapInput)(nil)).Elem(), ProfileMap{})
	pulumi.RegisterOutputType(ProfileOutput{})
	pulumi.RegisterOutputType(ProfileArrayOutput{})
	pulumi.RegisterOutputType(ProfileMapOutput{})
}