// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Vra.Project.Outputs
{

    [OutputType]
    public sealed class ProjectConstraints
    {
        public readonly ImmutableArray<Outputs.ProjectConstraintsExtensibility> Extensibilities;
        public readonly ImmutableArray<Outputs.ProjectConstraintsNetwork> Networks;
        public readonly ImmutableArray<Outputs.ProjectConstraintsStorage> Storages;

        [OutputConstructor]
        private ProjectConstraints(
            ImmutableArray<Outputs.ProjectConstraintsExtensibility> extensibilities,

            ImmutableArray<Outputs.ProjectConstraintsNetwork> networks,

            ImmutableArray<Outputs.ProjectConstraintsStorage> storages)
        {
            Extensibilities = extensibilities;
            Networks = networks;
            Storages = storages;
        }
    }
}