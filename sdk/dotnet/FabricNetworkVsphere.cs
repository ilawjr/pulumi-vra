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
    /// Updates a VMware vRealize Automation fabric_network_vsphere resource.
    /// 
    /// ## Example Usage
    /// ### S
    /// 
    /// You cannot create a vSphere fabric network resource, however you can import using the command specified in the import section below.
    /// Once a resource is imported, you can update it as shown below:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Vra = schmidtw.Vra;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var simple = new Vra.FabricNetworkVsphere("simple", new()
    ///     {
    ///         Cidr = @var.Cidr,
    ///         DefaultGateway = @var.Gateway,
    ///         Domain = @var.Domain,
    ///         Tags = new[]
    ///         {
    ///             new Vra.Inputs.FabricNetworkVsphereTagArgs
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
    /// ## Import
    /// 
    /// To import the vSphere fabric network resource, use the ID as in the following example
    /// 
    /// ```sh
    ///  $ pulumi import vra:index/fabricNetworkVsphere:FabricNetworkVsphere new_fabric_network_vsphere 05956583-6488-4e7d-84c9-92a7b7219a15`
    /// ```
    /// </summary>
    [VraResourceType("vra:index/fabricNetworkVsphere:FabricNetworkVsphere")]
    public partial class FabricNetworkVsphere : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Network CIDR to be used.
        /// </summary>
        [Output("cidr")]
        public Output<string> Cidr { get; private set; } = null!;

        /// <summary>
        /// Set of ids of the cloud accounts this entity belongs to.
        /// </summary>
        [Output("cloudAccountIds")]
        public Output<ImmutableArray<string>> CloudAccountIds { get; private set; } = null!;

        /// <summary>
        /// Date when the entity was created. The date is in ISO 6801 and UTC.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        [Output("customProperties")]
        public Output<ImmutableDictionary<string, object>> CustomProperties { get; private set; } = null!;

        /// <summary>
        /// IPv4 default gateway to be used.
        /// </summary>
        [Output("defaultGateway")]
        public Output<string> DefaultGateway { get; private set; } = null!;

        /// <summary>
        /// IPv6 default gateway to be used.
        /// </summary>
        [Output("defaultIpv6Gateway")]
        public Output<string> DefaultIpv6Gateway { get; private set; } = null!;

        /// <summary>
        /// List of dns search domains for the vSphere network.
        /// </summary>
        [Output("dnsSearchDomains")]
        public Output<ImmutableArray<string>> DnsSearchDomains { get; private set; } = null!;

        /// <summary>
        /// A human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Output("dnsServerAddresses")]
        public Output<ImmutableArray<string>> DnsServerAddresses { get; private set; } = null!;

        /// <summary>
        /// Domain for the vSphere network.
        /// </summary>
        [Output("domain")]
        public Output<string?> Domain { get; private set; } = null!;

        /// <summary>
        /// External entity Id on the provider side.
        /// </summary>
        [Output("externalId")]
        public Output<string> ExternalId { get; private set; } = null!;

        /// <summary>
        /// The id of the region for which this network is defined.
        /// </summary>
        [Output("externalRegionId")]
        public Output<string> ExternalRegionId { get; private set; } = null!;

        /// <summary>
        /// Network IPv6 CIDR to be used.
        /// </summary>
        [Output("ipv6Cidr")]
        public Output<string?> Ipv6Cidr { get; private set; } = null!;

        /// <summary>
        /// Indicates whether this is the default subnet for the zone.
        /// </summary>
        [Output("isDefault")]
        public Output<bool?> IsDefault { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the sub-network supports public IP assignment.
        /// </summary>
        [Output("isPublic")]
        public Output<bool?> IsPublic { get; private set; } = null!;

        /// <summary>
        /// HATEOAS of the entity
        /// </summary>
        [Output("links")]
        public Output<ImmutableArray<Outputs.FabricNetworkVsphereLink>> Links { get; private set; } = null!;

        /// <summary>
        /// A human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;

        /// <summary>
        /// ID of organization that entity belongs to.
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// Set of tag keys and values to apply to the resource.
        /// Example:[ { "key" : "vmware", "value": "provider" } ]
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.FabricNetworkVsphereTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// Date when the entity was last updated. The date is ISO 8601 and UTC.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a FabricNetworkVsphere resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FabricNetworkVsphere(string name, FabricNetworkVsphereArgs? args = null, CustomResourceOptions? options = null)
            : base("vra:index/fabricNetworkVsphere:FabricNetworkVsphere", name, args ?? new FabricNetworkVsphereArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FabricNetworkVsphere(string name, Input<string> id, FabricNetworkVsphereState? state = null, CustomResourceOptions? options = null)
            : base("vra:index/fabricNetworkVsphere:FabricNetworkVsphere", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FabricNetworkVsphere resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FabricNetworkVsphere Get(string name, Input<string> id, FabricNetworkVsphereState? state = null, CustomResourceOptions? options = null)
        {
            return new FabricNetworkVsphere(name, id, state, options);
        }
    }

    public sealed class FabricNetworkVsphereArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Network CIDR to be used.
        /// </summary>
        [Input("cidr")]
        public Input<string>? Cidr { get; set; }

        /// <summary>
        /// IPv4 default gateway to be used.
        /// </summary>
        [Input("defaultGateway")]
        public Input<string>? DefaultGateway { get; set; }

        /// <summary>
        /// IPv6 default gateway to be used.
        /// </summary>
        [Input("defaultIpv6Gateway")]
        public Input<string>? DefaultIpv6Gateway { get; set; }

        [Input("dnsSearchDomains")]
        private InputList<string>? _dnsSearchDomains;

        /// <summary>
        /// List of dns search domains for the vSphere network.
        /// </summary>
        public InputList<string> DnsSearchDomains
        {
            get => _dnsSearchDomains ?? (_dnsSearchDomains = new InputList<string>());
            set => _dnsSearchDomains = value;
        }

        [Input("dnsServerAddresses")]
        private InputList<string>? _dnsServerAddresses;

        /// <summary>
        /// A human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        public InputList<string> DnsServerAddresses
        {
            get => _dnsServerAddresses ?? (_dnsServerAddresses = new InputList<string>());
            set => _dnsServerAddresses = value;
        }

        /// <summary>
        /// Domain for the vSphere network.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Network IPv6 CIDR to be used.
        /// </summary>
        [Input("ipv6Cidr")]
        public Input<string>? Ipv6Cidr { get; set; }

        /// <summary>
        /// Indicates whether this is the default subnet for the zone.
        /// </summary>
        [Input("isDefault")]
        public Input<bool>? IsDefault { get; set; }

        /// <summary>
        /// Indicates whether the sub-network supports public IP assignment.
        /// </summary>
        [Input("isPublic")]
        public Input<bool>? IsPublic { get; set; }

        [Input("tags")]
        private InputList<Inputs.FabricNetworkVsphereTagArgs>? _tags;

        /// <summary>
        /// Set of tag keys and values to apply to the resource.
        /// Example:[ { "key" : "vmware", "value": "provider" } ]
        /// </summary>
        public InputList<Inputs.FabricNetworkVsphereTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.FabricNetworkVsphereTagArgs>());
            set => _tags = value;
        }

        public FabricNetworkVsphereArgs()
        {
        }
        public static new FabricNetworkVsphereArgs Empty => new FabricNetworkVsphereArgs();
    }

    public sealed class FabricNetworkVsphereState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Network CIDR to be used.
        /// </summary>
        [Input("cidr")]
        public Input<string>? Cidr { get; set; }

        [Input("cloudAccountIds")]
        private InputList<string>? _cloudAccountIds;

        /// <summary>
        /// Set of ids of the cloud accounts this entity belongs to.
        /// </summary>
        public InputList<string> CloudAccountIds
        {
            get => _cloudAccountIds ?? (_cloudAccountIds = new InputList<string>());
            set => _cloudAccountIds = value;
        }

        /// <summary>
        /// Date when the entity was created. The date is in ISO 6801 and UTC.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("customProperties")]
        private InputMap<object>? _customProperties;
        public InputMap<object> CustomProperties
        {
            get => _customProperties ?? (_customProperties = new InputMap<object>());
            set => _customProperties = value;
        }

        /// <summary>
        /// IPv4 default gateway to be used.
        /// </summary>
        [Input("defaultGateway")]
        public Input<string>? DefaultGateway { get; set; }

        /// <summary>
        /// IPv6 default gateway to be used.
        /// </summary>
        [Input("defaultIpv6Gateway")]
        public Input<string>? DefaultIpv6Gateway { get; set; }

        [Input("dnsSearchDomains")]
        private InputList<string>? _dnsSearchDomains;

        /// <summary>
        /// List of dns search domains for the vSphere network.
        /// </summary>
        public InputList<string> DnsSearchDomains
        {
            get => _dnsSearchDomains ?? (_dnsSearchDomains = new InputList<string>());
            set => _dnsSearchDomains = value;
        }

        [Input("dnsServerAddresses")]
        private InputList<string>? _dnsServerAddresses;

        /// <summary>
        /// A human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        public InputList<string> DnsServerAddresses
        {
            get => _dnsServerAddresses ?? (_dnsServerAddresses = new InputList<string>());
            set => _dnsServerAddresses = value;
        }

        /// <summary>
        /// Domain for the vSphere network.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// External entity Id on the provider side.
        /// </summary>
        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        /// <summary>
        /// The id of the region for which this network is defined.
        /// </summary>
        [Input("externalRegionId")]
        public Input<string>? ExternalRegionId { get; set; }

        /// <summary>
        /// Network IPv6 CIDR to be used.
        /// </summary>
        [Input("ipv6Cidr")]
        public Input<string>? Ipv6Cidr { get; set; }

        /// <summary>
        /// Indicates whether this is the default subnet for the zone.
        /// </summary>
        [Input("isDefault")]
        public Input<bool>? IsDefault { get; set; }

        /// <summary>
        /// Indicates whether the sub-network supports public IP assignment.
        /// </summary>
        [Input("isPublic")]
        public Input<bool>? IsPublic { get; set; }

        [Input("links")]
        private InputList<Inputs.FabricNetworkVsphereLinkGetArgs>? _links;

        /// <summary>
        /// HATEOAS of the entity
        /// </summary>
        public InputList<Inputs.FabricNetworkVsphereLinkGetArgs> Links
        {
            get => _links ?? (_links = new InputList<Inputs.FabricNetworkVsphereLinkGetArgs>());
            set => _links = value;
        }

        /// <summary>
        /// A human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// ID of organization that entity belongs to.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        [Input("tags")]
        private InputList<Inputs.FabricNetworkVsphereTagGetArgs>? _tags;

        /// <summary>
        /// Set of tag keys and values to apply to the resource.
        /// Example:[ { "key" : "vmware", "value": "provider" } ]
        /// </summary>
        public InputList<Inputs.FabricNetworkVsphereTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.FabricNetworkVsphereTagGetArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// Date when the entity was last updated. The date is ISO 8601 and UTC.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public FabricNetworkVsphereState()
        {
        }
        public static new FabricNetworkVsphereState Empty => new FabricNetworkVsphereState();
    }
}