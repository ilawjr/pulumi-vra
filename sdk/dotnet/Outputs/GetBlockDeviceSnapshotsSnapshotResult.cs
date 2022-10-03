// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace schmidtw.Vra.Outputs
{

    [OutputType]
    public sealed class GetBlockDeviceSnapshotsSnapshotResult
    {
        /// <summary>
        /// Date when the entity was created. The date is in ISO 6801 and UTC.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// A human-friendly description.
        /// </summary>
        public readonly string Description;
        public readonly string Id;
        /// <summary>
        /// Indicates whether this snapshot is the current snapshot on the block-device.
        /// </summary>
        public readonly bool IsCurrent;
        /// <summary>
        /// HATEOAS of the entity
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBlockDeviceSnapshotsSnapshotLinkResult> Links;
        /// <summary>
        /// A human-friendly name used as an identifier in APIs that support this option.  Only one of 'filter', 'id', 'name' or 'region_id' must be specified.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The id of the organization this entity belongs to.
        /// </summary>
        public readonly string OrgId;
        /// <summary>
        /// Email of the user that owns the entity.
        /// </summary>
        public readonly string Owner;
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetBlockDeviceSnapshotsSnapshotResult(
            string createdAt,

            string description,

            string id,

            bool isCurrent,

            ImmutableArray<Outputs.GetBlockDeviceSnapshotsSnapshotLinkResult> links,

            string name,

            string orgId,

            string owner,

            string updatedAt)
        {
            CreatedAt = createdAt;
            Description = description;
            Id = id;
            IsCurrent = isCurrent;
            Links = links;
            Name = name;
            OrgId = orgId;
            Owner = owner;
            UpdatedAt = updatedAt;
        }
    }
}