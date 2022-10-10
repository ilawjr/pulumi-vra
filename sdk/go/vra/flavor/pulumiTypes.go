// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package flavor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProfileFlavorMapping struct {
	CpuCount     *int    `pulumi:"cpuCount"`
	InstanceType *string `pulumi:"instanceType"`
	Memory       *int    `pulumi:"memory"`
	// A human-friendly name used as an identifier in APIs that support this option.
	Name string `pulumi:"name"`
}

// ProfileFlavorMappingInput is an input type that accepts ProfileFlavorMappingArgs and ProfileFlavorMappingOutput values.
// You can construct a concrete instance of `ProfileFlavorMappingInput` via:
//
//          ProfileFlavorMappingArgs{...}
type ProfileFlavorMappingInput interface {
	pulumi.Input

	ToProfileFlavorMappingOutput() ProfileFlavorMappingOutput
	ToProfileFlavorMappingOutputWithContext(context.Context) ProfileFlavorMappingOutput
}

type ProfileFlavorMappingArgs struct {
	CpuCount     pulumi.IntPtrInput    `pulumi:"cpuCount"`
	InstanceType pulumi.StringPtrInput `pulumi:"instanceType"`
	Memory       pulumi.IntPtrInput    `pulumi:"memory"`
	// A human-friendly name used as an identifier in APIs that support this option.
	Name pulumi.StringInput `pulumi:"name"`
}

func (ProfileFlavorMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileFlavorMapping)(nil)).Elem()
}

func (i ProfileFlavorMappingArgs) ToProfileFlavorMappingOutput() ProfileFlavorMappingOutput {
	return i.ToProfileFlavorMappingOutputWithContext(context.Background())
}

func (i ProfileFlavorMappingArgs) ToProfileFlavorMappingOutputWithContext(ctx context.Context) ProfileFlavorMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileFlavorMappingOutput)
}

// ProfileFlavorMappingArrayInput is an input type that accepts ProfileFlavorMappingArray and ProfileFlavorMappingArrayOutput values.
// You can construct a concrete instance of `ProfileFlavorMappingArrayInput` via:
//
//          ProfileFlavorMappingArray{ ProfileFlavorMappingArgs{...} }
type ProfileFlavorMappingArrayInput interface {
	pulumi.Input

	ToProfileFlavorMappingArrayOutput() ProfileFlavorMappingArrayOutput
	ToProfileFlavorMappingArrayOutputWithContext(context.Context) ProfileFlavorMappingArrayOutput
}

type ProfileFlavorMappingArray []ProfileFlavorMappingInput

func (ProfileFlavorMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileFlavorMapping)(nil)).Elem()
}

func (i ProfileFlavorMappingArray) ToProfileFlavorMappingArrayOutput() ProfileFlavorMappingArrayOutput {
	return i.ToProfileFlavorMappingArrayOutputWithContext(context.Background())
}

func (i ProfileFlavorMappingArray) ToProfileFlavorMappingArrayOutputWithContext(ctx context.Context) ProfileFlavorMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileFlavorMappingArrayOutput)
}

type ProfileFlavorMappingOutput struct{ *pulumi.OutputState }

func (ProfileFlavorMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileFlavorMapping)(nil)).Elem()
}

func (o ProfileFlavorMappingOutput) ToProfileFlavorMappingOutput() ProfileFlavorMappingOutput {
	return o
}

func (o ProfileFlavorMappingOutput) ToProfileFlavorMappingOutputWithContext(ctx context.Context) ProfileFlavorMappingOutput {
	return o
}

func (o ProfileFlavorMappingOutput) CpuCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ProfileFlavorMapping) *int { return v.CpuCount }).(pulumi.IntPtrOutput)
}

func (o ProfileFlavorMappingOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileFlavorMapping) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o ProfileFlavorMappingOutput) Memory() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ProfileFlavorMapping) *int { return v.Memory }).(pulumi.IntPtrOutput)
}

// A human-friendly name used as an identifier in APIs that support this option.
func (o ProfileFlavorMappingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ProfileFlavorMapping) string { return v.Name }).(pulumi.StringOutput)
}

type ProfileFlavorMappingArrayOutput struct{ *pulumi.OutputState }

func (ProfileFlavorMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileFlavorMapping)(nil)).Elem()
}

func (o ProfileFlavorMappingArrayOutput) ToProfileFlavorMappingArrayOutput() ProfileFlavorMappingArrayOutput {
	return o
}

func (o ProfileFlavorMappingArrayOutput) ToProfileFlavorMappingArrayOutputWithContext(ctx context.Context) ProfileFlavorMappingArrayOutput {
	return o
}

func (o ProfileFlavorMappingArrayOutput) Index(i pulumi.IntInput) ProfileFlavorMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProfileFlavorMapping {
		return vs[0].([]ProfileFlavorMapping)[vs[1].(int)]
	}).(ProfileFlavorMappingOutput)
}

type ProfileLink struct {
	Href  *string  `pulumi:"href"`
	Hrefs []string `pulumi:"hrefs"`
	Rel   string   `pulumi:"rel"`
}

// ProfileLinkInput is an input type that accepts ProfileLinkArgs and ProfileLinkOutput values.
// You can construct a concrete instance of `ProfileLinkInput` via:
//
//          ProfileLinkArgs{...}
type ProfileLinkInput interface {
	pulumi.Input

	ToProfileLinkOutput() ProfileLinkOutput
	ToProfileLinkOutputWithContext(context.Context) ProfileLinkOutput
}

type ProfileLinkArgs struct {
	Href  pulumi.StringPtrInput   `pulumi:"href"`
	Hrefs pulumi.StringArrayInput `pulumi:"hrefs"`
	Rel   pulumi.StringInput      `pulumi:"rel"`
}

func (ProfileLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileLink)(nil)).Elem()
}

func (i ProfileLinkArgs) ToProfileLinkOutput() ProfileLinkOutput {
	return i.ToProfileLinkOutputWithContext(context.Background())
}

func (i ProfileLinkArgs) ToProfileLinkOutputWithContext(ctx context.Context) ProfileLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileLinkOutput)
}

// ProfileLinkArrayInput is an input type that accepts ProfileLinkArray and ProfileLinkArrayOutput values.
// You can construct a concrete instance of `ProfileLinkArrayInput` via:
//
//          ProfileLinkArray{ ProfileLinkArgs{...} }
type ProfileLinkArrayInput interface {
	pulumi.Input

	ToProfileLinkArrayOutput() ProfileLinkArrayOutput
	ToProfileLinkArrayOutputWithContext(context.Context) ProfileLinkArrayOutput
}

type ProfileLinkArray []ProfileLinkInput

func (ProfileLinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileLink)(nil)).Elem()
}

func (i ProfileLinkArray) ToProfileLinkArrayOutput() ProfileLinkArrayOutput {
	return i.ToProfileLinkArrayOutputWithContext(context.Background())
}

func (i ProfileLinkArray) ToProfileLinkArrayOutputWithContext(ctx context.Context) ProfileLinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileLinkArrayOutput)
}

type ProfileLinkOutput struct{ *pulumi.OutputState }

func (ProfileLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileLink)(nil)).Elem()
}

func (o ProfileLinkOutput) ToProfileLinkOutput() ProfileLinkOutput {
	return o
}

func (o ProfileLinkOutput) ToProfileLinkOutputWithContext(ctx context.Context) ProfileLinkOutput {
	return o
}

func (o ProfileLinkOutput) Href() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProfileLink) *string { return v.Href }).(pulumi.StringPtrOutput)
}

func (o ProfileLinkOutput) Hrefs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ProfileLink) []string { return v.Hrefs }).(pulumi.StringArrayOutput)
}

func (o ProfileLinkOutput) Rel() pulumi.StringOutput {
	return o.ApplyT(func(v ProfileLink) string { return v.Rel }).(pulumi.StringOutput)
}

type ProfileLinkArrayOutput struct{ *pulumi.OutputState }

func (ProfileLinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileLink)(nil)).Elem()
}

func (o ProfileLinkArrayOutput) ToProfileLinkArrayOutput() ProfileLinkArrayOutput {
	return o
}

func (o ProfileLinkArrayOutput) ToProfileLinkArrayOutputWithContext(ctx context.Context) ProfileLinkArrayOutput {
	return o
}

func (o ProfileLinkArrayOutput) Index(i pulumi.IntInput) ProfileLinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProfileLink {
		return vs[0].([]ProfileLink)[vs[1].(int)]
	}).(ProfileLinkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileFlavorMappingInput)(nil)).Elem(), ProfileFlavorMappingArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileFlavorMappingArrayInput)(nil)).Elem(), ProfileFlavorMappingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileLinkInput)(nil)).Elem(), ProfileLinkArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileLinkArrayInput)(nil)).Elem(), ProfileLinkArray{})
	pulumi.RegisterOutputType(ProfileFlavorMappingOutput{})
	pulumi.RegisterOutputType(ProfileFlavorMappingArrayOutput{})
	pulumi.RegisterOutputType(ProfileLinkOutput{})
	pulumi.RegisterOutputType(ProfileLinkArrayOutput{})
}
