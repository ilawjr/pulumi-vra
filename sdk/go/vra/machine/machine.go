// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machine

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a VMware vRealize Automation machine resource.
//
// ## Example Usage
// ### S
//
// The following example shows how to create a machine resource.
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-vra/sdk/go/vra/machine"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumiverse/pulumi-vra/sdk/go/vra/machine"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := machine.NewMachine(ctx, "this", &machine.MachineArgs{
// 			Description: pulumi.String("terrafrom test machine"),
// 			ProjectId:   pulumi.Any(data.Vra_project.This.Id),
// 			Image:       pulumi.String("ubuntu2"),
// 			Flavor:      pulumi.String("medium"),
// 			BootConfig: &machine.MachineBootConfigArgs{
// 				Content: pulumi.String(fmt.Sprintf(`#cloud-config
//   users:
//   - default
//   - name: myuser
//     sudo: ['ALL=(ALL) NOPASSWD:ALL']
//     groups: [wheel, sudo, admin]
//     shell: '/bin/bash'
//     ssh-authorized-keys: |
//       ssh-rsa your-ssh-rsa:
//     - sudo sed -e 's/.*PasswordAuthentication yes.*/PasswordAuthentication no/' -i /etc/ssh/sshd_config
//     - sudo service sshd restart
// `)),
// 			},
// 			Nics: machine.MachineNicArray{
// 				&machine.MachineNicArgs{
// 					NetworkId: pulumi.Any(data.Vra_network.This.Id),
// 				},
// 			},
// 			Constraints: machine.MachineConstraintArray{
// 				&machine.MachineConstraintArgs{
// 					Mandatory:  pulumi.Bool(true),
// 					Expression: pulumi.String("AWS"),
// 				},
// 			},
// 			Tags: machine.MachineTagArray{
// 				&machine.MachineTagArgs{
// 					Key:   pulumi.String("foo"),
// 					Value: pulumi.String("bar"),
// 				},
// 			},
// 			Disks: machine.MachineDiskArray{
// 				&machine.MachineDiskArgs{
// 					BlockDeviceId: pulumi.Any(vra_block_device.Disk1.Id),
// 				},
// 				&machine.MachineDiskArgs{
// 					BlockDeviceId: pulumi.Any(vra_block_device.Disk2.Id),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// A machine resource supports the following resource:
type Machine struct {
	pulumi.CustomResourceState

	// Primary address allocated or in use by this machine. The actual type of the address depends on the adapter type. Typically it is either the public or the external IP address.
	Address pulumi.StringOutput `pulumi:"address"`
	// Machine boot config that will be passed to the instance. Used to perform common automated configuration tasks and even run scripts after instance starts.
	BootConfig MachineBootConfigPtrOutput `pulumi:"bootConfig"`
	// Constraints used to drive placement policies for the virtual machine produced from the specification. Constraint expressions are matched against tags on existing placement targets.\
	// Example:[{"mandatory" : "true", "expression": "environment":"prod"}, {"mandatory" : "false", "expression": "pci"}]
	Constraints MachineConstraintArrayOutput `pulumi:"constraints"`
	// Date when the entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Additional properties that may be used to extend the base type.
	CustomProperties pulumi.MapOutput `pulumi:"customProperties"`
	// Describes machine within the scope of your organization and is not propagated to the cloud.
	DeploymentId pulumi.StringOutput `pulumi:"deploymentId"`
	// Human-friendly description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specification for attaching/detaching disks to a machine.
	Disks MachineDiskArrayOutput `pulumi:"disks"`
	// List of all disks attached to a machine including boot disk, and additional block devices attached using the disks attribute.
	DisksLists MachineDisksListArrayOutput `pulumi:"disksLists"`
	// External entity ID on the provider side.
	ExternalId pulumi.StringOutput `pulumi:"externalId"`
	// External regionId of the resource.
	ExternalRegionId pulumi.StringOutput `pulumi:"externalRegionId"`
	// External zoneId of the resource.
	ExternalZoneId pulumi.StringOutput `pulumi:"externalZoneId"`
	// Flavor of machine instance.
	Flavor pulumi.StringOutput `pulumi:"flavor"`
	// Type of image used for this machine.
	Image pulumi.StringPtrOutput `pulumi:"image"`
	// Constraints that are used to drive placement policies for the image disk. Constraint expressions are matched against tags on existing placement targets. example:[{"mandatory" : "true", "expression": "environment:prod"}, {"mandatory" : "false", "expression": "pci"}]. It is nested argument with the following properties.
	ImageDiskConstraints MachineImageDiskConstraintArrayOutput `pulumi:"imageDiskConstraints"`
	// Direct image reference used for this machine (name, path, location, uri, etc.). Valid if no image property is provided
	ImageRef pulumi.StringPtrOutput `pulumi:"imageRef"`
	// HATEOAS of the entity
	Links MachineLinkArrayOutput `pulumi:"links"`
	// Human-friendly name used as an identifier in APIs that support this option.
	Name pulumi.StringOutput `pulumi:"name"`
	// Set of network interface controller specifications for this machine. If not specified, then a default network connection will be created.
	Nics MachineNicArrayOutput `pulumi:"nics"`
	// ID of the organization this entity belongs to.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// Email of entity owner.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// Power state of machine.
	PowerState pulumi.StringOutput `pulumi:"powerState"`
	// ID of project that resource belongs to.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Set of tag keys and optional values that should be set on any resource that is produced from this specification. example:[ { "key" : "ownedBy", "value": "Rainpole" } ]. It is nested argument with the following properties.
	Tags      MachineTagArrayOutput `pulumi:"tags"`
	UpdatedAt pulumi.StringOutput   `pulumi:"updatedAt"`
}

// NewMachine registers a new resource with the given unique name, arguments, and options.
func NewMachine(ctx *pulumi.Context,
	name string, args *MachineArgs, opts ...pulumi.ResourceOption) (*Machine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Flavor == nil {
		return nil, errors.New("invalid value for required argument 'Flavor'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Machine
	err := ctx.RegisterResource("vra:machine/machine:Machine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMachine gets an existing Machine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineState, opts ...pulumi.ResourceOption) (*Machine, error) {
	var resource Machine
	err := ctx.ReadResource("vra:machine/machine:Machine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Machine resources.
type machineState struct {
	// Primary address allocated or in use by this machine. The actual type of the address depends on the adapter type. Typically it is either the public or the external IP address.
	Address *string `pulumi:"address"`
	// Machine boot config that will be passed to the instance. Used to perform common automated configuration tasks and even run scripts after instance starts.
	BootConfig *MachineBootConfig `pulumi:"bootConfig"`
	// Constraints used to drive placement policies for the virtual machine produced from the specification. Constraint expressions are matched against tags on existing placement targets.\
	// Example:[{"mandatory" : "true", "expression": "environment":"prod"}, {"mandatory" : "false", "expression": "pci"}]
	Constraints []MachineConstraint `pulumi:"constraints"`
	// Date when the entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt *string `pulumi:"createdAt"`
	// Additional properties that may be used to extend the base type.
	CustomProperties map[string]interface{} `pulumi:"customProperties"`
	// Describes machine within the scope of your organization and is not propagated to the cloud.
	DeploymentId *string `pulumi:"deploymentId"`
	// Human-friendly description.
	Description *string `pulumi:"description"`
	// Specification for attaching/detaching disks to a machine.
	Disks []MachineDisk `pulumi:"disks"`
	// List of all disks attached to a machine including boot disk, and additional block devices attached using the disks attribute.
	DisksLists []MachineDisksList `pulumi:"disksLists"`
	// External entity ID on the provider side.
	ExternalId *string `pulumi:"externalId"`
	// External regionId of the resource.
	ExternalRegionId *string `pulumi:"externalRegionId"`
	// External zoneId of the resource.
	ExternalZoneId *string `pulumi:"externalZoneId"`
	// Flavor of machine instance.
	Flavor *string `pulumi:"flavor"`
	// Type of image used for this machine.
	Image *string `pulumi:"image"`
	// Constraints that are used to drive placement policies for the image disk. Constraint expressions are matched against tags on existing placement targets. example:[{"mandatory" : "true", "expression": "environment:prod"}, {"mandatory" : "false", "expression": "pci"}]. It is nested argument with the following properties.
	ImageDiskConstraints []MachineImageDiskConstraint `pulumi:"imageDiskConstraints"`
	// Direct image reference used for this machine (name, path, location, uri, etc.). Valid if no image property is provided
	ImageRef *string `pulumi:"imageRef"`
	// HATEOAS of the entity
	Links []MachineLink `pulumi:"links"`
	// Human-friendly name used as an identifier in APIs that support this option.
	Name *string `pulumi:"name"`
	// Set of network interface controller specifications for this machine. If not specified, then a default network connection will be created.
	Nics []MachineNic `pulumi:"nics"`
	// ID of the organization this entity belongs to.
	OrganizationId *string `pulumi:"organizationId"`
	// Email of entity owner.
	Owner *string `pulumi:"owner"`
	// Power state of machine.
	PowerState *string `pulumi:"powerState"`
	// ID of project that resource belongs to.
	ProjectId *string `pulumi:"projectId"`
	// Set of tag keys and optional values that should be set on any resource that is produced from this specification. example:[ { "key" : "ownedBy", "value": "Rainpole" } ]. It is nested argument with the following properties.
	Tags      []MachineTag `pulumi:"tags"`
	UpdatedAt *string      `pulumi:"updatedAt"`
}

type MachineState struct {
	// Primary address allocated or in use by this machine. The actual type of the address depends on the adapter type. Typically it is either the public or the external IP address.
	Address pulumi.StringPtrInput
	// Machine boot config that will be passed to the instance. Used to perform common automated configuration tasks and even run scripts after instance starts.
	BootConfig MachineBootConfigPtrInput
	// Constraints used to drive placement policies for the virtual machine produced from the specification. Constraint expressions are matched against tags on existing placement targets.\
	// Example:[{"mandatory" : "true", "expression": "environment":"prod"}, {"mandatory" : "false", "expression": "pci"}]
	Constraints MachineConstraintArrayInput
	// Date when the entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt pulumi.StringPtrInput
	// Additional properties that may be used to extend the base type.
	CustomProperties pulumi.MapInput
	// Describes machine within the scope of your organization and is not propagated to the cloud.
	DeploymentId pulumi.StringPtrInput
	// Human-friendly description.
	Description pulumi.StringPtrInput
	// Specification for attaching/detaching disks to a machine.
	Disks MachineDiskArrayInput
	// List of all disks attached to a machine including boot disk, and additional block devices attached using the disks attribute.
	DisksLists MachineDisksListArrayInput
	// External entity ID on the provider side.
	ExternalId pulumi.StringPtrInput
	// External regionId of the resource.
	ExternalRegionId pulumi.StringPtrInput
	// External zoneId of the resource.
	ExternalZoneId pulumi.StringPtrInput
	// Flavor of machine instance.
	Flavor pulumi.StringPtrInput
	// Type of image used for this machine.
	Image pulumi.StringPtrInput
	// Constraints that are used to drive placement policies for the image disk. Constraint expressions are matched against tags on existing placement targets. example:[{"mandatory" : "true", "expression": "environment:prod"}, {"mandatory" : "false", "expression": "pci"}]. It is nested argument with the following properties.
	ImageDiskConstraints MachineImageDiskConstraintArrayInput
	// Direct image reference used for this machine (name, path, location, uri, etc.). Valid if no image property is provided
	ImageRef pulumi.StringPtrInput
	// HATEOAS of the entity
	Links MachineLinkArrayInput
	// Human-friendly name used as an identifier in APIs that support this option.
	Name pulumi.StringPtrInput
	// Set of network interface controller specifications for this machine. If not specified, then a default network connection will be created.
	Nics MachineNicArrayInput
	// ID of the organization this entity belongs to.
	OrganizationId pulumi.StringPtrInput
	// Email of entity owner.
	Owner pulumi.StringPtrInput
	// Power state of machine.
	PowerState pulumi.StringPtrInput
	// ID of project that resource belongs to.
	ProjectId pulumi.StringPtrInput
	// Set of tag keys and optional values that should be set on any resource that is produced from this specification. example:[ { "key" : "ownedBy", "value": "Rainpole" } ]. It is nested argument with the following properties.
	Tags      MachineTagArrayInput
	UpdatedAt pulumi.StringPtrInput
}

func (MachineState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineState)(nil)).Elem()
}

type machineArgs struct {
	// Machine boot config that will be passed to the instance. Used to perform common automated configuration tasks and even run scripts after instance starts.
	BootConfig *MachineBootConfig `pulumi:"bootConfig"`
	// Constraints used to drive placement policies for the virtual machine produced from the specification. Constraint expressions are matched against tags on existing placement targets.\
	// Example:[{"mandatory" : "true", "expression": "environment":"prod"}, {"mandatory" : "false", "expression": "pci"}]
	Constraints []MachineConstraint `pulumi:"constraints"`
	// Additional properties that may be used to extend the base type.
	CustomProperties map[string]interface{} `pulumi:"customProperties"`
	// Describes machine within the scope of your organization and is not propagated to the cloud.
	DeploymentId *string `pulumi:"deploymentId"`
	// Human-friendly description.
	Description *string `pulumi:"description"`
	// Specification for attaching/detaching disks to a machine.
	Disks []MachineDisk `pulumi:"disks"`
	// Flavor of machine instance.
	Flavor string `pulumi:"flavor"`
	// Type of image used for this machine.
	Image *string `pulumi:"image"`
	// Constraints that are used to drive placement policies for the image disk. Constraint expressions are matched against tags on existing placement targets. example:[{"mandatory" : "true", "expression": "environment:prod"}, {"mandatory" : "false", "expression": "pci"}]. It is nested argument with the following properties.
	ImageDiskConstraints []MachineImageDiskConstraint `pulumi:"imageDiskConstraints"`
	// Direct image reference used for this machine (name, path, location, uri, etc.). Valid if no image property is provided
	ImageRef *string `pulumi:"imageRef"`
	// Human-friendly name used as an identifier in APIs that support this option.
	Name *string `pulumi:"name"`
	// Set of network interface controller specifications for this machine. If not specified, then a default network connection will be created.
	Nics []MachineNic `pulumi:"nics"`
	// ID of project that resource belongs to.
	ProjectId string `pulumi:"projectId"`
	// Set of tag keys and optional values that should be set on any resource that is produced from this specification. example:[ { "key" : "ownedBy", "value": "Rainpole" } ]. It is nested argument with the following properties.
	Tags []MachineTag `pulumi:"tags"`
}

// The set of arguments for constructing a Machine resource.
type MachineArgs struct {
	// Machine boot config that will be passed to the instance. Used to perform common automated configuration tasks and even run scripts after instance starts.
	BootConfig MachineBootConfigPtrInput
	// Constraints used to drive placement policies for the virtual machine produced from the specification. Constraint expressions are matched against tags on existing placement targets.\
	// Example:[{"mandatory" : "true", "expression": "environment":"prod"}, {"mandatory" : "false", "expression": "pci"}]
	Constraints MachineConstraintArrayInput
	// Additional properties that may be used to extend the base type.
	CustomProperties pulumi.MapInput
	// Describes machine within the scope of your organization and is not propagated to the cloud.
	DeploymentId pulumi.StringPtrInput
	// Human-friendly description.
	Description pulumi.StringPtrInput
	// Specification for attaching/detaching disks to a machine.
	Disks MachineDiskArrayInput
	// Flavor of machine instance.
	Flavor pulumi.StringInput
	// Type of image used for this machine.
	Image pulumi.StringPtrInput
	// Constraints that are used to drive placement policies for the image disk. Constraint expressions are matched against tags on existing placement targets. example:[{"mandatory" : "true", "expression": "environment:prod"}, {"mandatory" : "false", "expression": "pci"}]. It is nested argument with the following properties.
	ImageDiskConstraints MachineImageDiskConstraintArrayInput
	// Direct image reference used for this machine (name, path, location, uri, etc.). Valid if no image property is provided
	ImageRef pulumi.StringPtrInput
	// Human-friendly name used as an identifier in APIs that support this option.
	Name pulumi.StringPtrInput
	// Set of network interface controller specifications for this machine. If not specified, then a default network connection will be created.
	Nics MachineNicArrayInput
	// ID of project that resource belongs to.
	ProjectId pulumi.StringInput
	// Set of tag keys and optional values that should be set on any resource that is produced from this specification. example:[ { "key" : "ownedBy", "value": "Rainpole" } ]. It is nested argument with the following properties.
	Tags MachineTagArrayInput
}

func (MachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineArgs)(nil)).Elem()
}

type MachineInput interface {
	pulumi.Input

	ToMachineOutput() MachineOutput
	ToMachineOutputWithContext(ctx context.Context) MachineOutput
}

func (*Machine) ElementType() reflect.Type {
	return reflect.TypeOf((**Machine)(nil)).Elem()
}

func (i *Machine) ToMachineOutput() MachineOutput {
	return i.ToMachineOutputWithContext(context.Background())
}

func (i *Machine) ToMachineOutputWithContext(ctx context.Context) MachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineOutput)
}

// MachineArrayInput is an input type that accepts MachineArray and MachineArrayOutput values.
// You can construct a concrete instance of `MachineArrayInput` via:
//
//          MachineArray{ MachineArgs{...} }
type MachineArrayInput interface {
	pulumi.Input

	ToMachineArrayOutput() MachineArrayOutput
	ToMachineArrayOutputWithContext(context.Context) MachineArrayOutput
}

type MachineArray []MachineInput

func (MachineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Machine)(nil)).Elem()
}

func (i MachineArray) ToMachineArrayOutput() MachineArrayOutput {
	return i.ToMachineArrayOutputWithContext(context.Background())
}

func (i MachineArray) ToMachineArrayOutputWithContext(ctx context.Context) MachineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineArrayOutput)
}

// MachineMapInput is an input type that accepts MachineMap and MachineMapOutput values.
// You can construct a concrete instance of `MachineMapInput` via:
//
//          MachineMap{ "key": MachineArgs{...} }
type MachineMapInput interface {
	pulumi.Input

	ToMachineMapOutput() MachineMapOutput
	ToMachineMapOutputWithContext(context.Context) MachineMapOutput
}

type MachineMap map[string]MachineInput

func (MachineMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Machine)(nil)).Elem()
}

func (i MachineMap) ToMachineMapOutput() MachineMapOutput {
	return i.ToMachineMapOutputWithContext(context.Background())
}

func (i MachineMap) ToMachineMapOutputWithContext(ctx context.Context) MachineMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineMapOutput)
}

type MachineOutput struct{ *pulumi.OutputState }

func (MachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Machine)(nil)).Elem()
}

func (o MachineOutput) ToMachineOutput() MachineOutput {
	return o
}

func (o MachineOutput) ToMachineOutputWithContext(ctx context.Context) MachineOutput {
	return o
}

// Primary address allocated or in use by this machine. The actual type of the address depends on the adapter type. Typically it is either the public or the external IP address.
func (o MachineOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// Machine boot config that will be passed to the instance. Used to perform common automated configuration tasks and even run scripts after instance starts.
func (o MachineOutput) BootConfig() MachineBootConfigPtrOutput {
	return o.ApplyT(func(v *Machine) MachineBootConfigPtrOutput { return v.BootConfig }).(MachineBootConfigPtrOutput)
}

// Constraints used to drive placement policies for the virtual machine produced from the specification. Constraint expressions are matched against tags on existing placement targets.\
// Example:[{"mandatory" : "true", "expression": "environment":"prod"}, {"mandatory" : "false", "expression": "pci"}]
func (o MachineOutput) Constraints() MachineConstraintArrayOutput {
	return o.ApplyT(func(v *Machine) MachineConstraintArrayOutput { return v.Constraints }).(MachineConstraintArrayOutput)
}

// Date when the entity was created. Date and time format is ISO 8601 and UTC.
func (o MachineOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Additional properties that may be used to extend the base type.
func (o MachineOutput) CustomProperties() pulumi.MapOutput {
	return o.ApplyT(func(v *Machine) pulumi.MapOutput { return v.CustomProperties }).(pulumi.MapOutput)
}

// Describes machine within the scope of your organization and is not propagated to the cloud.
func (o MachineOutput) DeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.DeploymentId }).(pulumi.StringOutput)
}

// Human-friendly description.
func (o MachineOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specification for attaching/detaching disks to a machine.
func (o MachineOutput) Disks() MachineDiskArrayOutput {
	return o.ApplyT(func(v *Machine) MachineDiskArrayOutput { return v.Disks }).(MachineDiskArrayOutput)
}

// List of all disks attached to a machine including boot disk, and additional block devices attached using the disks attribute.
func (o MachineOutput) DisksLists() MachineDisksListArrayOutput {
	return o.ApplyT(func(v *Machine) MachineDisksListArrayOutput { return v.DisksLists }).(MachineDisksListArrayOutput)
}

// External entity ID on the provider side.
func (o MachineOutput) ExternalId() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.ExternalId }).(pulumi.StringOutput)
}

// External regionId of the resource.
func (o MachineOutput) ExternalRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.ExternalRegionId }).(pulumi.StringOutput)
}

// External zoneId of the resource.
func (o MachineOutput) ExternalZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.ExternalZoneId }).(pulumi.StringOutput)
}

// Flavor of machine instance.
func (o MachineOutput) Flavor() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.Flavor }).(pulumi.StringOutput)
}

// Type of image used for this machine.
func (o MachineOutput) Image() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringPtrOutput { return v.Image }).(pulumi.StringPtrOutput)
}

// Constraints that are used to drive placement policies for the image disk. Constraint expressions are matched against tags on existing placement targets. example:[{"mandatory" : "true", "expression": "environment:prod"}, {"mandatory" : "false", "expression": "pci"}]. It is nested argument with the following properties.
func (o MachineOutput) ImageDiskConstraints() MachineImageDiskConstraintArrayOutput {
	return o.ApplyT(func(v *Machine) MachineImageDiskConstraintArrayOutput { return v.ImageDiskConstraints }).(MachineImageDiskConstraintArrayOutput)
}

// Direct image reference used for this machine (name, path, location, uri, etc.). Valid if no image property is provided
func (o MachineOutput) ImageRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringPtrOutput { return v.ImageRef }).(pulumi.StringPtrOutput)
}

// HATEOAS of the entity
func (o MachineOutput) Links() MachineLinkArrayOutput {
	return o.ApplyT(func(v *Machine) MachineLinkArrayOutput { return v.Links }).(MachineLinkArrayOutput)
}

// Human-friendly name used as an identifier in APIs that support this option.
func (o MachineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Set of network interface controller specifications for this machine. If not specified, then a default network connection will be created.
func (o MachineOutput) Nics() MachineNicArrayOutput {
	return o.ApplyT(func(v *Machine) MachineNicArrayOutput { return v.Nics }).(MachineNicArrayOutput)
}

// ID of the organization this entity belongs to.
func (o MachineOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// Email of entity owner.
func (o MachineOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// Power state of machine.
func (o MachineOutput) PowerState() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.PowerState }).(pulumi.StringOutput)
}

// ID of project that resource belongs to.
func (o MachineOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Set of tag keys and optional values that should be set on any resource that is produced from this specification. example:[ { "key" : "ownedBy", "value": "Rainpole" } ]. It is nested argument with the following properties.
func (o MachineOutput) Tags() MachineTagArrayOutput {
	return o.ApplyT(func(v *Machine) MachineTagArrayOutput { return v.Tags }).(MachineTagArrayOutput)
}

func (o MachineOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type MachineArrayOutput struct{ *pulumi.OutputState }

func (MachineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Machine)(nil)).Elem()
}

func (o MachineArrayOutput) ToMachineArrayOutput() MachineArrayOutput {
	return o
}

func (o MachineArrayOutput) ToMachineArrayOutputWithContext(ctx context.Context) MachineArrayOutput {
	return o
}

func (o MachineArrayOutput) Index(i pulumi.IntInput) MachineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Machine {
		return vs[0].([]*Machine)[vs[1].(int)]
	}).(MachineOutput)
}

type MachineMapOutput struct{ *pulumi.OutputState }

func (MachineMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Machine)(nil)).Elem()
}

func (o MachineMapOutput) ToMachineMapOutput() MachineMapOutput {
	return o
}

func (o MachineMapOutput) ToMachineMapOutputWithContext(ctx context.Context) MachineMapOutput {
	return o
}

func (o MachineMapOutput) MapIndex(k pulumi.StringInput) MachineOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Machine {
		return vs[0].(map[string]*Machine)[vs[1].(string)]
	}).(MachineOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MachineInput)(nil)).Elem(), &Machine{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineArrayInput)(nil)).Elem(), MachineArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineMapInput)(nil)).Elem(), MachineMap{})
	pulumi.RegisterOutputType(MachineOutput{})
	pulumi.RegisterOutputType(MachineArrayOutput{})
	pulumi.RegisterOutputType(MachineMapOutput{})
}
