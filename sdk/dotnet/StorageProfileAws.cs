// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace schmidtw.Vra
{
    /// <summary>
    /// ## Example Usage
    /// ### S
    /// This is an example of how to create a storage profile aws resource.
    /// 
    /// **Vra storage profile aws:**
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Vra = schmidtw.Vra;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // AWS storage profile using vra_storage_profile_aws resource and EBS disk type.
    ///     var thisStorageProfileAws = new Vra.StorageProfileAws("thisStorageProfileAws", new()
    ///     {
    ///         Description = "AWS Storage Profile with instance store device type.",
    ///         RegionId = data.Vra_region.This.Id,
    ///         DefaultItem = false,
    ///         SupportsEncryption = false,
    ///         DeviceType = "ebs",
    ///         VolumeType = "io1",
    ///         Iops = "1000",
    ///         Tags = new[]
    ///         {
    ///             new Vra.Inputs.StorageProfileAwsTagArgs
    ///             {
    ///                 Key = "foo",
    ///                 Value = "bar",
    ///             },
    ///         },
    ///     });
    /// 
    ///     // AWS storage profile using vra_storage_profile_aws resource and instance store disk type.
    ///     var thisIndex_storageProfileAwsStorageProfileAws = new Vra.StorageProfileAws("thisIndex/storageProfileAwsStorageProfileAws", new()
    ///     {
    ///         Description = "AWS Storage Profile with instance store device type.",
    ///         RegionId = data.Vra_region.This.Id,
    ///         DefaultItem = false,
    ///         DeviceType = "instance-store",
    ///         Tags = new[]
    ///         {
    ///             new Vra.Inputs.StorageProfileAwsTagArgs
    ///             {
    ///                 Key = "foo",
    ///                 Value = "bar",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// A storage profile aws resource supports the following arguments:
    /// </summary>
    [VraResourceType("vra:index/storageProfileAws:StorageProfileAws")]
    public partial class StorageProfileAws : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Date when the entity was created. The date is in ISO 6801 and UTC.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Indicates if this storage profile is a default profile.
        /// </summary>
        [Output("defaultItem")]
        public Output<bool> DefaultItem { get; private set; } = null!;

        /// <summary>
        /// A human-friendly description.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Indicates the type of storage device.
        /// </summary>
        [Output("deviceType")]
        public Output<string> DeviceType { get; private set; } = null!;

        /// <summary>
        /// The id of the region as seen in the cloud provider for which this profile is defined.
        /// </summary>
        [Output("externalRegionId")]
        public Output<string> ExternalRegionId { get; private set; } = null!;

        /// <summary>
        /// Indicates maximum I/O operations per second in range(1-20,000).
        /// </summary>
        [Output("iops")]
        public Output<string> Iops { get; private set; } = null!;

        /// <summary>
        /// HATEOAS of the entity
        /// </summary>
        [Output("links")]
        public Output<ImmutableArray<Outputs.StorageProfileAwsLink>> Links { get; private set; } = null!;

        /// <summary>
        /// A human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The id of the organization this entity belongs to.
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// Email of the user that owns the entity.
        /// </summary>
        [Output("owner")]
        public Output<string> Owner { get; private set; } = null!;

        /// <summary>
        /// A link to the region that is associated with the storage profile.
        /// </summary>
        [Output("regionId")]
        public Output<string> RegionId { get; private set; } = null!;

        /// <summary>
        /// Indicates whether this storage profile supports encryption or not.
        /// </summary>
        [Output("supportsEncryption")]
        public Output<bool> SupportsEncryption { get; private set; } = null!;

        /// <summary>
        /// A set of tag keys and optional values that were set on this Network Profile.
        /// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.StorageProfileAwsTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// Date when the entity was last updated. The date is ISO 8601 and UTC.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// Indicates the type of volume associated with type of storage.
        /// </summary>
        [Output("volumeType")]
        public Output<string> VolumeType { get; private set; } = null!;


        /// <summary>
        /// Create a StorageProfileAws resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StorageProfileAws(string name, StorageProfileAwsArgs args, CustomResourceOptions? options = null)
            : base("vra:index/storageProfileAws:StorageProfileAws", name, args ?? new StorageProfileAwsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StorageProfileAws(string name, Input<string> id, StorageProfileAwsState? state = null, CustomResourceOptions? options = null)
            : base("vra:index/storageProfileAws:StorageProfileAws", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/schmidtw/pulumi-vra/releases/download/v${VERSION}",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing StorageProfileAws resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StorageProfileAws Get(string name, Input<string> id, StorageProfileAwsState? state = null, CustomResourceOptions? options = null)
        {
            return new StorageProfileAws(name, id, state, options);
        }
    }

    public sealed class StorageProfileAwsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates if this storage profile is a default profile.
        /// </summary>
        [Input("defaultItem", required: true)]
        public Input<bool> DefaultItem { get; set; } = null!;

        /// <summary>
        /// A human-friendly description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicates the type of storage device.
        /// </summary>
        [Input("deviceType")]
        public Input<string>? DeviceType { get; set; }

        /// <summary>
        /// Indicates maximum I/O operations per second in range(1-20,000).
        /// </summary>
        [Input("iops")]
        public Input<string>? Iops { get; set; }

        /// <summary>
        /// A human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A link to the region that is associated with the storage profile.
        /// </summary>
        [Input("regionId", required: true)]
        public Input<string> RegionId { get; set; } = null!;

        /// <summary>
        /// Indicates whether this storage profile supports encryption or not.
        /// </summary>
        [Input("supportsEncryption")]
        public Input<bool>? SupportsEncryption { get; set; }

        [Input("tags")]
        private InputList<Inputs.StorageProfileAwsTagArgs>? _tags;

        /// <summary>
        /// A set of tag keys and optional values that were set on this Network Profile.
        /// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        /// </summary>
        public InputList<Inputs.StorageProfileAwsTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.StorageProfileAwsTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// Indicates the type of volume associated with type of storage.
        /// </summary>
        [Input("volumeType")]
        public Input<string>? VolumeType { get; set; }

        public StorageProfileAwsArgs()
        {
        }
        public static new StorageProfileAwsArgs Empty => new StorageProfileAwsArgs();
    }

    public sealed class StorageProfileAwsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Date when the entity was created. The date is in ISO 6801 and UTC.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Indicates if this storage profile is a default profile.
        /// </summary>
        [Input("defaultItem")]
        public Input<bool>? DefaultItem { get; set; }

        /// <summary>
        /// A human-friendly description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicates the type of storage device.
        /// </summary>
        [Input("deviceType")]
        public Input<string>? DeviceType { get; set; }

        /// <summary>
        /// The id of the region as seen in the cloud provider for which this profile is defined.
        /// </summary>
        [Input("externalRegionId")]
        public Input<string>? ExternalRegionId { get; set; }

        /// <summary>
        /// Indicates maximum I/O operations per second in range(1-20,000).
        /// </summary>
        [Input("iops")]
        public Input<string>? Iops { get; set; }

        [Input("links")]
        private InputList<Inputs.StorageProfileAwsLinkGetArgs>? _links;

        /// <summary>
        /// HATEOAS of the entity
        /// </summary>
        public InputList<Inputs.StorageProfileAwsLinkGetArgs> Links
        {
            get => _links ?? (_links = new InputList<Inputs.StorageProfileAwsLinkGetArgs>());
            set => _links = value;
        }

        /// <summary>
        /// A human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The id of the organization this entity belongs to.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// Email of the user that owns the entity.
        /// </summary>
        [Input("owner")]
        public Input<string>? Owner { get; set; }

        /// <summary>
        /// A link to the region that is associated with the storage profile.
        /// </summary>
        [Input("regionId")]
        public Input<string>? RegionId { get; set; }

        /// <summary>
        /// Indicates whether this storage profile supports encryption or not.
        /// </summary>
        [Input("supportsEncryption")]
        public Input<bool>? SupportsEncryption { get; set; }

        [Input("tags")]
        private InputList<Inputs.StorageProfileAwsTagGetArgs>? _tags;

        /// <summary>
        /// A set of tag keys and optional values that were set on this Network Profile.
        /// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        /// </summary>
        public InputList<Inputs.StorageProfileAwsTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.StorageProfileAwsTagGetArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// Date when the entity was last updated. The date is ISO 8601 and UTC.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// Indicates the type of volume associated with type of storage.
        /// </summary>
        [Input("volumeType")]
        public Input<string>? VolumeType { get; set; }

        public StorageProfileAwsState()
        {
        }
        public static new StorageProfileAwsState Empty => new StorageProfileAwsState();
    }
}