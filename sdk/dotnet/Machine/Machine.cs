// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Vra.Machine
{
    /// <summary>
    /// Creates a VMware vRealize Automation machine resource.
    /// 
    /// ## Example Usage
    /// ### S
    /// 
    /// The following example shows how to create a machine resource.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Vra = Pulumiverse.Vra;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @this = new Vra.Machine.Machine("this", new()
    ///     {
    ///         Description = "terrafrom test machine",
    ///         ProjectId = data.Vra_project.This.Id,
    ///         Image = "ubuntu2",
    ///         Flavor = "medium",
    ///         BootConfig = new Vra.Machine.Inputs.MachineBootConfigArgs
    ///         {
    ///             Content = @"#cloud-config
    ///   users:
    ///   - default
    ///   - name: myuser
    ///     sudo: ['ALL=(ALL) NOPASSWD:ALL']
    ///     groups: [wheel, sudo, admin]
    ///     shell: '/bin/bash'
    ///     ssh-authorized-keys: |
    ///       ssh-rsa your-ssh-rsa:
    ///     - sudo sed -e 's/.*PasswordAuthentication yes.*/PasswordAuthentication no/' -i /etc/ssh/sshd_config
    ///     - sudo service sshd restart
    /// ",
    ///         },
    ///         Nics = new[]
    ///         {
    ///             new Vra.Machine.Inputs.MachineNicArgs
    ///             {
    ///                 NetworkId = data.Vra_network.This.Id,
    ///             },
    ///         },
    ///         Constraints = new[]
    ///         {
    ///             new Vra.Machine.Inputs.MachineConstraintArgs
    ///             {
    ///                 Mandatory = true,
    ///                 Expression = "AWS",
    ///             },
    ///         },
    ///         Tags = new[]
    ///         {
    ///             new Vra.Machine.Inputs.MachineTagArgs
    ///             {
    ///                 Key = "foo",
    ///                 Value = "bar",
    ///             },
    ///         },
    ///         Disks = new[]
    ///         {
    ///             new Vra.Machine.Inputs.MachineDiskArgs
    ///             {
    ///                 BlockDeviceId = vra_block_device.Disk1.Id,
    ///             },
    ///             new Vra.Machine.Inputs.MachineDiskArgs
    ///             {
    ///                 BlockDeviceId = vra_block_device.Disk2.Id,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// A machine resource supports the following resource:
    /// </summary>
    [VraResourceType("vra:machine/machine:Machine")]
    public partial class Machine : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Primary address allocated or in use by this machine. The actual type of the address depends on the adapter type. Typically it is either the public or the external IP address.
        /// </summary>
        [Output("address")]
        public Output<string> Address { get; private set; } = null!;

        /// <summary>
        /// Machine boot config that will be passed to the instance. Used to perform common automated configuration tasks and even run scripts after instance starts.
        /// </summary>
        [Output("bootConfig")]
        public Output<Outputs.MachineBootConfig?> BootConfig { get; private set; } = null!;

        /// <summary>
        /// Constraints used to drive placement policies for the virtual machine produced from the specification. Constraint expressions are matched against tags on existing placement targets.  
        /// Example:[{"mandatory" : "true", "expression": "environment":"prod"}, {"mandatory" : "false", "expression": "pci"}]
        /// </summary>
        [Output("constraints")]
        public Output<ImmutableArray<Outputs.MachineConstraint>> Constraints { get; private set; } = null!;

        /// <summary>
        /// Date when the entity was created. Date and time format is ISO 8601 and UTC.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Additional properties that may be used to extend the base type.
        /// </summary>
        [Output("customProperties")]
        public Output<ImmutableDictionary<string, object>> CustomProperties { get; private set; } = null!;

        /// <summary>
        /// Describes machine within the scope of your organization and is not propagated to the cloud.
        /// </summary>
        [Output("deploymentId")]
        public Output<string> DeploymentId { get; private set; } = null!;

        /// <summary>
        /// Human-friendly description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Specification for attaching/detaching disks to a machine.
        /// </summary>
        [Output("disks")]
        public Output<ImmutableArray<Outputs.MachineDisk>> Disks { get; private set; } = null!;

        /// <summary>
        /// List of all disks attached to a machine including boot disk, and additional block devices attached using the disks attribute.
        /// </summary>
        [Output("disksLists")]
        public Output<ImmutableArray<Outputs.MachineDisksList>> DisksLists { get; private set; } = null!;

        /// <summary>
        /// External entity ID on the provider side.
        /// </summary>
        [Output("externalId")]
        public Output<string> ExternalId { get; private set; } = null!;

        /// <summary>
        /// External regionId of the resource.
        /// </summary>
        [Output("externalRegionId")]
        public Output<string> ExternalRegionId { get; private set; } = null!;

        /// <summary>
        /// External zoneId of the resource.
        /// </summary>
        [Output("externalZoneId")]
        public Output<string> ExternalZoneId { get; private set; } = null!;

        /// <summary>
        /// Flavor of machine instance.
        /// </summary>
        [Output("flavor")]
        public Output<string> Flavor { get; private set; } = null!;

        /// <summary>
        /// Type of image used for this machine.
        /// </summary>
        [Output("image")]
        public Output<string?> Image { get; private set; } = null!;

        /// <summary>
        /// Constraints that are used to drive placement policies for the image disk. Constraint expressions are matched against tags on existing placement targets. example:[{"mandatory" : "true", "expression": "environment:prod"}, {"mandatory" : "false", "expression": "pci"}]. It is nested argument with the following properties.
        /// </summary>
        [Output("imageDiskConstraints")]
        public Output<ImmutableArray<Outputs.MachineImageDiskConstraint>> ImageDiskConstraints { get; private set; } = null!;

        /// <summary>
        /// Direct image reference used for this machine (name, path, location, uri, etc.). Valid if no image property is provided
        /// </summary>
        [Output("imageRef")]
        public Output<string?> ImageRef { get; private set; } = null!;

        /// <summary>
        /// HATEOAS of the entity
        /// </summary>
        [Output("links")]
        public Output<ImmutableArray<Outputs.MachineLink>> Links { get; private set; } = null!;

        /// <summary>
        /// Human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Set of network interface controller specifications for this machine. If not specified, then a default network connection will be created.
        /// </summary>
        [Output("nics")]
        public Output<ImmutableArray<Outputs.MachineNic>> Nics { get; private set; } = null!;

        /// <summary>
        /// ID of the organization this entity belongs to.
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// Email of entity owner.
        /// </summary>
        [Output("owner")]
        public Output<string> Owner { get; private set; } = null!;

        /// <summary>
        /// Power state of machine.
        /// </summary>
        [Output("powerState")]
        public Output<string> PowerState { get; private set; } = null!;

        /// <summary>
        /// ID of project that resource belongs to.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Set of tag keys and optional values that should be set on any resource that is produced from this specification. example:[ { "key" : "ownedBy", "value": "Rainpole" } ]. It is nested argument with the following properties.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.MachineTag>> Tags { get; private set; } = null!;

        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a Machine resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Machine(string name, MachineArgs args, CustomResourceOptions? options = null)
            : base("vra:machine/machine:Machine", name, args ?? new MachineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Machine(string name, Input<string> id, MachineState? state = null, CustomResourceOptions? options = null)
            : base("vra:machine/machine:Machine", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Machine resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Machine Get(string name, Input<string> id, MachineState? state = null, CustomResourceOptions? options = null)
        {
            return new Machine(name, id, state, options);
        }
    }

    public sealed class MachineArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Machine boot config that will be passed to the instance. Used to perform common automated configuration tasks and even run scripts after instance starts.
        /// </summary>
        [Input("bootConfig")]
        public Input<Inputs.MachineBootConfigArgs>? BootConfig { get; set; }

        [Input("constraints")]
        private InputList<Inputs.MachineConstraintArgs>? _constraints;

        /// <summary>
        /// Constraints used to drive placement policies for the virtual machine produced from the specification. Constraint expressions are matched against tags on existing placement targets.  
        /// Example:[{"mandatory" : "true", "expression": "environment":"prod"}, {"mandatory" : "false", "expression": "pci"}]
        /// </summary>
        public InputList<Inputs.MachineConstraintArgs> Constraints
        {
            get => _constraints ?? (_constraints = new InputList<Inputs.MachineConstraintArgs>());
            set => _constraints = value;
        }

        [Input("customProperties")]
        private InputMap<object>? _customProperties;

        /// <summary>
        /// Additional properties that may be used to extend the base type.
        /// </summary>
        public InputMap<object> CustomProperties
        {
            get => _customProperties ?? (_customProperties = new InputMap<object>());
            set => _customProperties = value;
        }

        /// <summary>
        /// Describes machine within the scope of your organization and is not propagated to the cloud.
        /// </summary>
        [Input("deploymentId")]
        public Input<string>? DeploymentId { get; set; }

        /// <summary>
        /// Human-friendly description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("disks")]
        private InputList<Inputs.MachineDiskArgs>? _disks;

        /// <summary>
        /// Specification for attaching/detaching disks to a machine.
        /// </summary>
        public InputList<Inputs.MachineDiskArgs> Disks
        {
            get => _disks ?? (_disks = new InputList<Inputs.MachineDiskArgs>());
            set => _disks = value;
        }

        /// <summary>
        /// Flavor of machine instance.
        /// </summary>
        [Input("flavor", required: true)]
        public Input<string> Flavor { get; set; } = null!;

        /// <summary>
        /// Type of image used for this machine.
        /// </summary>
        [Input("image")]
        public Input<string>? Image { get; set; }

        [Input("imageDiskConstraints")]
        private InputList<Inputs.MachineImageDiskConstraintArgs>? _imageDiskConstraints;

        /// <summary>
        /// Constraints that are used to drive placement policies for the image disk. Constraint expressions are matched against tags on existing placement targets. example:[{"mandatory" : "true", "expression": "environment:prod"}, {"mandatory" : "false", "expression": "pci"}]. It is nested argument with the following properties.
        /// </summary>
        public InputList<Inputs.MachineImageDiskConstraintArgs> ImageDiskConstraints
        {
            get => _imageDiskConstraints ?? (_imageDiskConstraints = new InputList<Inputs.MachineImageDiskConstraintArgs>());
            set => _imageDiskConstraints = value;
        }

        /// <summary>
        /// Direct image reference used for this machine (name, path, location, uri, etc.). Valid if no image property is provided
        /// </summary>
        [Input("imageRef")]
        public Input<string>? ImageRef { get; set; }

        /// <summary>
        /// Human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nics")]
        private InputList<Inputs.MachineNicArgs>? _nics;

        /// <summary>
        /// Set of network interface controller specifications for this machine. If not specified, then a default network connection will be created.
        /// </summary>
        public InputList<Inputs.MachineNicArgs> Nics
        {
            get => _nics ?? (_nics = new InputList<Inputs.MachineNicArgs>());
            set => _nics = value;
        }

        /// <summary>
        /// ID of project that resource belongs to.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.MachineTagArgs>? _tags;

        /// <summary>
        /// Set of tag keys and optional values that should be set on any resource that is produced from this specification. example:[ { "key" : "ownedBy", "value": "Rainpole" } ]. It is nested argument with the following properties.
        /// </summary>
        public InputList<Inputs.MachineTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.MachineTagArgs>());
            set => _tags = value;
        }

        public MachineArgs()
        {
        }
        public static new MachineArgs Empty => new MachineArgs();
    }

    public sealed class MachineState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Primary address allocated or in use by this machine. The actual type of the address depends on the adapter type. Typically it is either the public or the external IP address.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// Machine boot config that will be passed to the instance. Used to perform common automated configuration tasks and even run scripts after instance starts.
        /// </summary>
        [Input("bootConfig")]
        public Input<Inputs.MachineBootConfigGetArgs>? BootConfig { get; set; }

        [Input("constraints")]
        private InputList<Inputs.MachineConstraintGetArgs>? _constraints;

        /// <summary>
        /// Constraints used to drive placement policies for the virtual machine produced from the specification. Constraint expressions are matched against tags on existing placement targets.  
        /// Example:[{"mandatory" : "true", "expression": "environment":"prod"}, {"mandatory" : "false", "expression": "pci"}]
        /// </summary>
        public InputList<Inputs.MachineConstraintGetArgs> Constraints
        {
            get => _constraints ?? (_constraints = new InputList<Inputs.MachineConstraintGetArgs>());
            set => _constraints = value;
        }

        /// <summary>
        /// Date when the entity was created. Date and time format is ISO 8601 and UTC.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("customProperties")]
        private InputMap<object>? _customProperties;

        /// <summary>
        /// Additional properties that may be used to extend the base type.
        /// </summary>
        public InputMap<object> CustomProperties
        {
            get => _customProperties ?? (_customProperties = new InputMap<object>());
            set => _customProperties = value;
        }

        /// <summary>
        /// Describes machine within the scope of your organization and is not propagated to the cloud.
        /// </summary>
        [Input("deploymentId")]
        public Input<string>? DeploymentId { get; set; }

        /// <summary>
        /// Human-friendly description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("disks")]
        private InputList<Inputs.MachineDiskGetArgs>? _disks;

        /// <summary>
        /// Specification for attaching/detaching disks to a machine.
        /// </summary>
        public InputList<Inputs.MachineDiskGetArgs> Disks
        {
            get => _disks ?? (_disks = new InputList<Inputs.MachineDiskGetArgs>());
            set => _disks = value;
        }

        [Input("disksLists")]
        private InputList<Inputs.MachineDisksListGetArgs>? _disksLists;

        /// <summary>
        /// List of all disks attached to a machine including boot disk, and additional block devices attached using the disks attribute.
        /// </summary>
        public InputList<Inputs.MachineDisksListGetArgs> DisksLists
        {
            get => _disksLists ?? (_disksLists = new InputList<Inputs.MachineDisksListGetArgs>());
            set => _disksLists = value;
        }

        /// <summary>
        /// External entity ID on the provider side.
        /// </summary>
        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        /// <summary>
        /// External regionId of the resource.
        /// </summary>
        [Input("externalRegionId")]
        public Input<string>? ExternalRegionId { get; set; }

        /// <summary>
        /// External zoneId of the resource.
        /// </summary>
        [Input("externalZoneId")]
        public Input<string>? ExternalZoneId { get; set; }

        /// <summary>
        /// Flavor of machine instance.
        /// </summary>
        [Input("flavor")]
        public Input<string>? Flavor { get; set; }

        /// <summary>
        /// Type of image used for this machine.
        /// </summary>
        [Input("image")]
        public Input<string>? Image { get; set; }

        [Input("imageDiskConstraints")]
        private InputList<Inputs.MachineImageDiskConstraintGetArgs>? _imageDiskConstraints;

        /// <summary>
        /// Constraints that are used to drive placement policies for the image disk. Constraint expressions are matched against tags on existing placement targets. example:[{"mandatory" : "true", "expression": "environment:prod"}, {"mandatory" : "false", "expression": "pci"}]. It is nested argument with the following properties.
        /// </summary>
        public InputList<Inputs.MachineImageDiskConstraintGetArgs> ImageDiskConstraints
        {
            get => _imageDiskConstraints ?? (_imageDiskConstraints = new InputList<Inputs.MachineImageDiskConstraintGetArgs>());
            set => _imageDiskConstraints = value;
        }

        /// <summary>
        /// Direct image reference used for this machine (name, path, location, uri, etc.). Valid if no image property is provided
        /// </summary>
        [Input("imageRef")]
        public Input<string>? ImageRef { get; set; }

        [Input("links")]
        private InputList<Inputs.MachineLinkGetArgs>? _links;

        /// <summary>
        /// HATEOAS of the entity
        /// </summary>
        public InputList<Inputs.MachineLinkGetArgs> Links
        {
            get => _links ?? (_links = new InputList<Inputs.MachineLinkGetArgs>());
            set => _links = value;
        }

        /// <summary>
        /// Human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nics")]
        private InputList<Inputs.MachineNicGetArgs>? _nics;

        /// <summary>
        /// Set of network interface controller specifications for this machine. If not specified, then a default network connection will be created.
        /// </summary>
        public InputList<Inputs.MachineNicGetArgs> Nics
        {
            get => _nics ?? (_nics = new InputList<Inputs.MachineNicGetArgs>());
            set => _nics = value;
        }

        /// <summary>
        /// ID of the organization this entity belongs to.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// Email of entity owner.
        /// </summary>
        [Input("owner")]
        public Input<string>? Owner { get; set; }

        /// <summary>
        /// Power state of machine.
        /// </summary>
        [Input("powerState")]
        public Input<string>? PowerState { get; set; }

        /// <summary>
        /// ID of project that resource belongs to.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("tags")]
        private InputList<Inputs.MachineTagGetArgs>? _tags;

        /// <summary>
        /// Set of tag keys and optional values that should be set on any resource that is produced from this specification. example:[ { "key" : "ownedBy", "value": "Rainpole" } ]. It is nested argument with the following properties.
        /// </summary>
        public InputList<Inputs.MachineTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.MachineTagGetArgs>());
            set => _tags = value;
        }

        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public MachineState()
        {
        }
        public static new MachineState Empty => new MachineState();
    }
}