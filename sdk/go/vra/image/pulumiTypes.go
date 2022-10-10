// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package image

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProfileImageMapping struct {
	CloudConfig *string                         `pulumi:"cloudConfig"`
	Constraints []ProfileImageMappingConstraint `pulumi:"constraints"`
	// A human-friendly description.
	Description *string `pulumi:"description"`
	ExternalId  *string `pulumi:"externalId"`
	// The external regionId of the resource.
	ExternalRegionId *string `pulumi:"externalRegionId"`
	ImageId          *string `pulumi:"imageId"`
	ImageName        *string `pulumi:"imageName"`
	// A human-friendly name used as an identifier in APIs that support this option.
	Name         string  `pulumi:"name"`
	Organization *string `pulumi:"organization"`
	OsFamily     *string `pulumi:"osFamily"`
	// Email of the user that owns the entity.
	Owner   *string `pulumi:"owner"`
	Private *bool   `pulumi:"private"`
}

// ProfileImageMappingInput is an input type that accepts ProfileImageMappingArgs and ProfileImageMappingOutput values.
// You can construct a concrete instance of `ProfileImageMappingInput` via:
//
//          ProfileImageMappingArgs{...}
type ProfileImageMappingInput interface {
	pulumi.Input

	ToProfileImageMappingOutput() ProfileImageMappingOutput
	ToProfileImageMappingOutputWithContext(context.Context) ProfileImageMappingOutput
}

type ProfileImageMappingArgs struct {
	CloudConfig pulumi.StringPtrInput                   `pulumi:"cloudConfig"`
	Constraints ProfileImageMappingConstraintArrayInput `pulumi:"constraints"`
	// A human-friendly description.
	Description pulumi.StringPtrInput `pulumi:"description"`
	ExternalId  pulumi.StringPtrInput `pulumi:"externalId"`
	// The external regionId of the resource.
	ExternalRegionId pulumi.StringPtrInput `pulumi:"externalRegionId"`
	ImageId          pulumi.StringPtrInput `pulumi:"imageId"`
	ImageName        pulumi.StringPtrInput `pulumi:"imageName"`
	// A human-friendly name used as an identifier in APIs that support this option.
	Name         pulumi.StringInput    `pulumi:"name"`
	Organization pulumi.StringPtrInput `pulumi:"organization"`
	OsFamily     pulumi.StringPtrInput `pulumi:"osFamily"`
	// Email of the user that owns the entity.
	Owner   pulumi.StringPtrInput `pulumi:"owner"`
	Private pulumi.BoolPtrInput   `pulumi:"private"`
}

func (ProfileImageMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileImageMapping)(nil)).Elem()
}

func (i ProfileImageMappingArgs) ToProfileImageMappingOutput() ProfileImageMappingOutput {
	return i.ToProfileImageMappingOutputWithContext(context.Background())
}

func (i ProfileImageMappingArgs) ToProfileImageMappingOutputWithContext(ctx context.Context) ProfileImageMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileImageMappingOutput)
}

// ProfileImageMappingArrayInput is an input type that accepts ProfileImageMappingArray and ProfileImageMappingArrayOutput values.
// You can construct a concrete instance of `ProfileImageMappingArrayInput` via:
//
//          ProfileImageMappingArray{ ProfileImageMappingArgs{...} }
type ProfileImageMappingArrayInput interface {
	pulumi.Input

	ToProfileImageMappingArrayOutput() ProfileImageMappingArrayOutput
	ToProfileImageMappingArrayOutputWithContext(context.Context) ProfileImageMappingArrayOutput
}

type ProfileImageMappingArray []ProfileImageMappingInput

func (ProfileImageMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileImageMapping)(nil)).Elem()
}

func (i ProfileImageMappingArray) ToProfileImageMappingArrayOutput() ProfileImageMappingArrayOutput {
	return i.ToProfileImageMappingArrayOutputWithContext(context.Background())
}

func (i ProfileImageMappingArray) ToProfileImageMappingArrayOutputWithContext(ctx context.Context) ProfileImageMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileImageMappingArrayOutput)
}

type ProfileImageMappingOutput struct{ *pulumi.OutputState }

func (ProfileImageMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileImageMapping)(nil)).Elem()
}

func (o ProfileImageMappingOutput) ToProfileImageMappingOutput() ProfileImageMappingOutput {
	return o
}

func (o ProfileImageMappingOutput) ToProfileImageMappingOutputWithContext(ctx context.Context) ProfileImageMappingOutput {
	return o
}

func (o ProfileImageMappingOutput) CloudConfig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileImageMapping) *string { return v.CloudConfig }).(pulumi.StringPtrOutput)
}

func (o ProfileImageMappingOutput) Constraints() ProfileImageMappingConstraintArrayOutput {
	return o.ApplyT(func(v ProfileImageMapping) []ProfileImageMappingConstraint { return v.Constraints }).(ProfileImageMappingConstraintArrayOutput)
}

// A human-friendly description.
func (o ProfileImageMappingOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileImageMapping) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ProfileImageMappingOutput) ExternalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileImageMapping) *string { return v.ExternalId }).(pulumi.StringPtrOutput)
}

// The external regionId of the resource.
func (o ProfileImageMappingOutput) ExternalRegionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileImageMapping) *string { return v.ExternalRegionId }).(pulumi.StringPtrOutput)
}

func (o ProfileImageMappingOutput) ImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileImageMapping) *string { return v.ImageId }).(pulumi.StringPtrOutput)
}

func (o ProfileImageMappingOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileImageMapping) *string { return v.ImageName }).(pulumi.StringPtrOutput)
}

// A human-friendly name used as an identifier in APIs that support this option.
func (o ProfileImageMappingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ProfileImageMapping) string { return v.Name }).(pulumi.StringOutput)
}

func (o ProfileImageMappingOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileImageMapping) *string { return v.Organization }).(pulumi.StringPtrOutput)
}

func (o ProfileImageMappingOutput) OsFamily() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileImageMapping) *string { return v.OsFamily }).(pulumi.StringPtrOutput)
}

// Email of the user that owns the entity.
func (o ProfileImageMappingOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileImageMapping) *string { return v.Owner }).(pulumi.StringPtrOutput)
}

func (o ProfileImageMappingOutput) Private() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ProfileImageMapping) *bool { return v.Private }).(pulumi.BoolPtrOutput)
}

type ProfileImageMappingArrayOutput struct{ *pulumi.OutputState }

func (ProfileImageMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileImageMapping)(nil)).Elem()
}

func (o ProfileImageMappingArrayOutput) ToProfileImageMappingArrayOutput() ProfileImageMappingArrayOutput {
	return o
}

func (o ProfileImageMappingArrayOutput) ToProfileImageMappingArrayOutputWithContext(ctx context.Context) ProfileImageMappingArrayOutput {
	return o
}

func (o ProfileImageMappingArrayOutput) Index(i pulumi.IntInput) ProfileImageMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProfileImageMapping {
		return vs[0].([]ProfileImageMapping)[vs[1].(int)]
	}).(ProfileImageMappingOutput)
}

type ProfileImageMappingConstraint struct {
	Expression string `pulumi:"expression"`
	Mandatory  bool   `pulumi:"mandatory"`
}

// ProfileImageMappingConstraintInput is an input type that accepts ProfileImageMappingConstraintArgs and ProfileImageMappingConstraintOutput values.
// You can construct a concrete instance of `ProfileImageMappingConstraintInput` via:
//
//          ProfileImageMappingConstraintArgs{...}
type ProfileImageMappingConstraintInput interface {
	pulumi.Input

	ToProfileImageMappingConstraintOutput() ProfileImageMappingConstraintOutput
	ToProfileImageMappingConstraintOutputWithContext(context.Context) ProfileImageMappingConstraintOutput
}

type ProfileImageMappingConstraintArgs struct {
	Expression pulumi.StringInput `pulumi:"expression"`
	Mandatory  pulumi.BoolInput   `pulumi:"mandatory"`
}

func (ProfileImageMappingConstraintArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileImageMappingConstraint)(nil)).Elem()
}

func (i ProfileImageMappingConstraintArgs) ToProfileImageMappingConstraintOutput() ProfileImageMappingConstraintOutput {
	return i.ToProfileImageMappingConstraintOutputWithContext(context.Background())
}

func (i ProfileImageMappingConstraintArgs) ToProfileImageMappingConstraintOutputWithContext(ctx context.Context) ProfileImageMappingConstraintOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileImageMappingConstraintOutput)
}

// ProfileImageMappingConstraintArrayInput is an input type that accepts ProfileImageMappingConstraintArray and ProfileImageMappingConstraintArrayOutput values.
// You can construct a concrete instance of `ProfileImageMappingConstraintArrayInput` via:
//
//          ProfileImageMappingConstraintArray{ ProfileImageMappingConstraintArgs{...} }
type ProfileImageMappingConstraintArrayInput interface {
	pulumi.Input

	ToProfileImageMappingConstraintArrayOutput() ProfileImageMappingConstraintArrayOutput
	ToProfileImageMappingConstraintArrayOutputWithContext(context.Context) ProfileImageMappingConstraintArrayOutput
}

type ProfileImageMappingConstraintArray []ProfileImageMappingConstraintInput

func (ProfileImageMappingConstraintArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileImageMappingConstraint)(nil)).Elem()
}

func (i ProfileImageMappingConstraintArray) ToProfileImageMappingConstraintArrayOutput() ProfileImageMappingConstraintArrayOutput {
	return i.ToProfileImageMappingConstraintArrayOutputWithContext(context.Background())
}

func (i ProfileImageMappingConstraintArray) ToProfileImageMappingConstraintArrayOutputWithContext(ctx context.Context) ProfileImageMappingConstraintArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileImageMappingConstraintArrayOutput)
}

type ProfileImageMappingConstraintOutput struct{ *pulumi.OutputState }

func (ProfileImageMappingConstraintOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileImageMappingConstraint)(nil)).Elem()
}

func (o ProfileImageMappingConstraintOutput) ToProfileImageMappingConstraintOutput() ProfileImageMappingConstraintOutput {
	return o
}

func (o ProfileImageMappingConstraintOutput) ToProfileImageMappingConstraintOutputWithContext(ctx context.Context) ProfileImageMappingConstraintOutput {
	return o
}

func (o ProfileImageMappingConstraintOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v ProfileImageMappingConstraint) string { return v.Expression }).(pulumi.StringOutput)
}

func (o ProfileImageMappingConstraintOutput) Mandatory() pulumi.BoolOutput {
	return o.ApplyT(func(v ProfileImageMappingConstraint) bool { return v.Mandatory }).(pulumi.BoolOutput)
}

type ProfileImageMappingConstraintArrayOutput struct{ *pulumi.OutputState }

func (ProfileImageMappingConstraintArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileImageMappingConstraint)(nil)).Elem()
}

func (o ProfileImageMappingConstraintArrayOutput) ToProfileImageMappingConstraintArrayOutput() ProfileImageMappingConstraintArrayOutput {
	return o
}

func (o ProfileImageMappingConstraintArrayOutput) ToProfileImageMappingConstraintArrayOutputWithContext(ctx context.Context) ProfileImageMappingConstraintArrayOutput {
	return o
}

func (o ProfileImageMappingConstraintArrayOutput) Index(i pulumi.IntInput) ProfileImageMappingConstraintOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProfileImageMappingConstraint {
		return vs[0].([]ProfileImageMappingConstraint)[vs[1].(int)]
	}).(ProfileImageMappingConstraintOutput)
}

type GetProfileImageMapping struct {
	CloudConfig *string                            `pulumi:"cloudConfig"`
	Constraints []GetProfileImageMappingConstraint `pulumi:"constraints"`
	// A human-friendly description.
	Description string `pulumi:"description"`
	ExternalId  string `pulumi:"externalId"`
	// The external regionId of the resource.
	ExternalRegionId string  `pulumi:"externalRegionId"`
	ImageId          *string `pulumi:"imageId"`
	ImageName        *string `pulumi:"imageName"`
	// A human-friendly name used as an identifier in APIs that support this option.
	Name         string `pulumi:"name"`
	Organization string `pulumi:"organization"`
	OsFamily     string `pulumi:"osFamily"`
	// Email of the user that owns the entity.
	Owner   string `pulumi:"owner"`
	Private bool   `pulumi:"private"`
}

// GetProfileImageMappingInput is an input type that accepts GetProfileImageMappingArgs and GetProfileImageMappingOutput values.
// You can construct a concrete instance of `GetProfileImageMappingInput` via:
//
//          GetProfileImageMappingArgs{...}
type GetProfileImageMappingInput interface {
	pulumi.Input

	ToGetProfileImageMappingOutput() GetProfileImageMappingOutput
	ToGetProfileImageMappingOutputWithContext(context.Context) GetProfileImageMappingOutput
}

type GetProfileImageMappingArgs struct {
	CloudConfig pulumi.StringPtrInput                      `pulumi:"cloudConfig"`
	Constraints GetProfileImageMappingConstraintArrayInput `pulumi:"constraints"`
	// A human-friendly description.
	Description pulumi.StringInput `pulumi:"description"`
	ExternalId  pulumi.StringInput `pulumi:"externalId"`
	// The external regionId of the resource.
	ExternalRegionId pulumi.StringInput    `pulumi:"externalRegionId"`
	ImageId          pulumi.StringPtrInput `pulumi:"imageId"`
	ImageName        pulumi.StringPtrInput `pulumi:"imageName"`
	// A human-friendly name used as an identifier in APIs that support this option.
	Name         pulumi.StringInput `pulumi:"name"`
	Organization pulumi.StringInput `pulumi:"organization"`
	OsFamily     pulumi.StringInput `pulumi:"osFamily"`
	// Email of the user that owns the entity.
	Owner   pulumi.StringInput `pulumi:"owner"`
	Private pulumi.BoolInput   `pulumi:"private"`
}

func (GetProfileImageMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProfileImageMapping)(nil)).Elem()
}

func (i GetProfileImageMappingArgs) ToGetProfileImageMappingOutput() GetProfileImageMappingOutput {
	return i.ToGetProfileImageMappingOutputWithContext(context.Background())
}

func (i GetProfileImageMappingArgs) ToGetProfileImageMappingOutputWithContext(ctx context.Context) GetProfileImageMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProfileImageMappingOutput)
}

// GetProfileImageMappingArrayInput is an input type that accepts GetProfileImageMappingArray and GetProfileImageMappingArrayOutput values.
// You can construct a concrete instance of `GetProfileImageMappingArrayInput` via:
//
//          GetProfileImageMappingArray{ GetProfileImageMappingArgs{...} }
type GetProfileImageMappingArrayInput interface {
	pulumi.Input

	ToGetProfileImageMappingArrayOutput() GetProfileImageMappingArrayOutput
	ToGetProfileImageMappingArrayOutputWithContext(context.Context) GetProfileImageMappingArrayOutput
}

type GetProfileImageMappingArray []GetProfileImageMappingInput

func (GetProfileImageMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProfileImageMapping)(nil)).Elem()
}

func (i GetProfileImageMappingArray) ToGetProfileImageMappingArrayOutput() GetProfileImageMappingArrayOutput {
	return i.ToGetProfileImageMappingArrayOutputWithContext(context.Background())
}

func (i GetProfileImageMappingArray) ToGetProfileImageMappingArrayOutputWithContext(ctx context.Context) GetProfileImageMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProfileImageMappingArrayOutput)
}

type GetProfileImageMappingOutput struct{ *pulumi.OutputState }

func (GetProfileImageMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProfileImageMapping)(nil)).Elem()
}

func (o GetProfileImageMappingOutput) ToGetProfileImageMappingOutput() GetProfileImageMappingOutput {
	return o
}

func (o GetProfileImageMappingOutput) ToGetProfileImageMappingOutputWithContext(ctx context.Context) GetProfileImageMappingOutput {
	return o
}

func (o GetProfileImageMappingOutput) CloudConfig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProfileImageMapping) *string { return v.CloudConfig }).(pulumi.StringPtrOutput)
}

func (o GetProfileImageMappingOutput) Constraints() GetProfileImageMappingConstraintArrayOutput {
	return o.ApplyT(func(v GetProfileImageMapping) []GetProfileImageMappingConstraint { return v.Constraints }).(GetProfileImageMappingConstraintArrayOutput)
}

// A human-friendly description.
func (o GetProfileImageMappingOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetProfileImageMapping) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetProfileImageMappingOutput) ExternalId() pulumi.StringOutput {
	return o.ApplyT(func(v GetProfileImageMapping) string { return v.ExternalId }).(pulumi.StringOutput)
}

// The external regionId of the resource.
func (o GetProfileImageMappingOutput) ExternalRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetProfileImageMapping) string { return v.ExternalRegionId }).(pulumi.StringOutput)
}

func (o GetProfileImageMappingOutput) ImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProfileImageMapping) *string { return v.ImageId }).(pulumi.StringPtrOutput)
}

func (o GetProfileImageMappingOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProfileImageMapping) *string { return v.ImageName }).(pulumi.StringPtrOutput)
}

// A human-friendly name used as an identifier in APIs that support this option.
func (o GetProfileImageMappingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetProfileImageMapping) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetProfileImageMappingOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v GetProfileImageMapping) string { return v.Organization }).(pulumi.StringOutput)
}

func (o GetProfileImageMappingOutput) OsFamily() pulumi.StringOutput {
	return o.ApplyT(func(v GetProfileImageMapping) string { return v.OsFamily }).(pulumi.StringOutput)
}

// Email of the user that owns the entity.
func (o GetProfileImageMappingOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v GetProfileImageMapping) string { return v.Owner }).(pulumi.StringOutput)
}

func (o GetProfileImageMappingOutput) Private() pulumi.BoolOutput {
	return o.ApplyT(func(v GetProfileImageMapping) bool { return v.Private }).(pulumi.BoolOutput)
}

type GetProfileImageMappingArrayOutput struct{ *pulumi.OutputState }

func (GetProfileImageMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProfileImageMapping)(nil)).Elem()
}

func (o GetProfileImageMappingArrayOutput) ToGetProfileImageMappingArrayOutput() GetProfileImageMappingArrayOutput {
	return o
}

func (o GetProfileImageMappingArrayOutput) ToGetProfileImageMappingArrayOutputWithContext(ctx context.Context) GetProfileImageMappingArrayOutput {
	return o
}

func (o GetProfileImageMappingArrayOutput) Index(i pulumi.IntInput) GetProfileImageMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetProfileImageMapping {
		return vs[0].([]GetProfileImageMapping)[vs[1].(int)]
	}).(GetProfileImageMappingOutput)
}

type GetProfileImageMappingConstraint struct {
	Expression string `pulumi:"expression"`
	Mandatory  bool   `pulumi:"mandatory"`
}

// GetProfileImageMappingConstraintInput is an input type that accepts GetProfileImageMappingConstraintArgs and GetProfileImageMappingConstraintOutput values.
// You can construct a concrete instance of `GetProfileImageMappingConstraintInput` via:
//
//          GetProfileImageMappingConstraintArgs{...}
type GetProfileImageMappingConstraintInput interface {
	pulumi.Input

	ToGetProfileImageMappingConstraintOutput() GetProfileImageMappingConstraintOutput
	ToGetProfileImageMappingConstraintOutputWithContext(context.Context) GetProfileImageMappingConstraintOutput
}

type GetProfileImageMappingConstraintArgs struct {
	Expression pulumi.StringInput `pulumi:"expression"`
	Mandatory  pulumi.BoolInput   `pulumi:"mandatory"`
}

func (GetProfileImageMappingConstraintArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProfileImageMappingConstraint)(nil)).Elem()
}

func (i GetProfileImageMappingConstraintArgs) ToGetProfileImageMappingConstraintOutput() GetProfileImageMappingConstraintOutput {
	return i.ToGetProfileImageMappingConstraintOutputWithContext(context.Background())
}

func (i GetProfileImageMappingConstraintArgs) ToGetProfileImageMappingConstraintOutputWithContext(ctx context.Context) GetProfileImageMappingConstraintOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProfileImageMappingConstraintOutput)
}

// GetProfileImageMappingConstraintArrayInput is an input type that accepts GetProfileImageMappingConstraintArray and GetProfileImageMappingConstraintArrayOutput values.
// You can construct a concrete instance of `GetProfileImageMappingConstraintArrayInput` via:
//
//          GetProfileImageMappingConstraintArray{ GetProfileImageMappingConstraintArgs{...} }
type GetProfileImageMappingConstraintArrayInput interface {
	pulumi.Input

	ToGetProfileImageMappingConstraintArrayOutput() GetProfileImageMappingConstraintArrayOutput
	ToGetProfileImageMappingConstraintArrayOutputWithContext(context.Context) GetProfileImageMappingConstraintArrayOutput
}

type GetProfileImageMappingConstraintArray []GetProfileImageMappingConstraintInput

func (GetProfileImageMappingConstraintArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProfileImageMappingConstraint)(nil)).Elem()
}

func (i GetProfileImageMappingConstraintArray) ToGetProfileImageMappingConstraintArrayOutput() GetProfileImageMappingConstraintArrayOutput {
	return i.ToGetProfileImageMappingConstraintArrayOutputWithContext(context.Background())
}

func (i GetProfileImageMappingConstraintArray) ToGetProfileImageMappingConstraintArrayOutputWithContext(ctx context.Context) GetProfileImageMappingConstraintArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProfileImageMappingConstraintArrayOutput)
}

type GetProfileImageMappingConstraintOutput struct{ *pulumi.OutputState }

func (GetProfileImageMappingConstraintOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProfileImageMappingConstraint)(nil)).Elem()
}

func (o GetProfileImageMappingConstraintOutput) ToGetProfileImageMappingConstraintOutput() GetProfileImageMappingConstraintOutput {
	return o
}

func (o GetProfileImageMappingConstraintOutput) ToGetProfileImageMappingConstraintOutputWithContext(ctx context.Context) GetProfileImageMappingConstraintOutput {
	return o
}

func (o GetProfileImageMappingConstraintOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v GetProfileImageMappingConstraint) string { return v.Expression }).(pulumi.StringOutput)
}

func (o GetProfileImageMappingConstraintOutput) Mandatory() pulumi.BoolOutput {
	return o.ApplyT(func(v GetProfileImageMappingConstraint) bool { return v.Mandatory }).(pulumi.BoolOutput)
}

type GetProfileImageMappingConstraintArrayOutput struct{ *pulumi.OutputState }

func (GetProfileImageMappingConstraintArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProfileImageMappingConstraint)(nil)).Elem()
}

func (o GetProfileImageMappingConstraintArrayOutput) ToGetProfileImageMappingConstraintArrayOutput() GetProfileImageMappingConstraintArrayOutput {
	return o
}

func (o GetProfileImageMappingConstraintArrayOutput) ToGetProfileImageMappingConstraintArrayOutputWithContext(ctx context.Context) GetProfileImageMappingConstraintArrayOutput {
	return o
}

func (o GetProfileImageMappingConstraintArrayOutput) Index(i pulumi.IntInput) GetProfileImageMappingConstraintOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetProfileImageMappingConstraint {
		return vs[0].([]GetProfileImageMappingConstraint)[vs[1].(int)]
	}).(GetProfileImageMappingConstraintOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileImageMappingInput)(nil)).Elem(), ProfileImageMappingArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileImageMappingArrayInput)(nil)).Elem(), ProfileImageMappingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileImageMappingConstraintInput)(nil)).Elem(), ProfileImageMappingConstraintArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileImageMappingConstraintArrayInput)(nil)).Elem(), ProfileImageMappingConstraintArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProfileImageMappingInput)(nil)).Elem(), GetProfileImageMappingArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProfileImageMappingArrayInput)(nil)).Elem(), GetProfileImageMappingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProfileImageMappingConstraintInput)(nil)).Elem(), GetProfileImageMappingConstraintArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProfileImageMappingConstraintArrayInput)(nil)).Elem(), GetProfileImageMappingConstraintArray{})
	pulumi.RegisterOutputType(ProfileImageMappingOutput{})
	pulumi.RegisterOutputType(ProfileImageMappingArrayOutput{})
	pulumi.RegisterOutputType(ProfileImageMappingConstraintOutput{})
	pulumi.RegisterOutputType(ProfileImageMappingConstraintArrayOutput{})
	pulumi.RegisterOutputType(GetProfileImageMappingOutput{})
	pulumi.RegisterOutputType(GetProfileImageMappingArrayOutput{})
	pulumi.RegisterOutputType(GetProfileImageMappingConstraintOutput{})
	pulumi.RegisterOutputType(GetProfileImageMappingConstraintArrayOutput{})
}
