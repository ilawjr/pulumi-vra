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
    /// The provider type for the vra package. By default, resources use package-wide configuration
    /// settings, however an explicit `Provider` instance may be created and passed during resource
    /// construction to achieve fine-grained programmatic control over provider settings. See the
    /// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// </summary>
    [VraResourceType("pulumi:providers:vra")]
    public partial class Provider : global::Pulumi.ProviderResource
    {
        /// <summary>
        /// The access token for API operations.
        /// </summary>
        [Output("accessToken")]
        public Output<string> AccessToken { get; private set; } = null!;

        /// <summary>
        /// Specify timeout for how often to reauthorize the access token
        /// </summary>
        [Output("reauthorizeTimeout")]
        public Output<string?> ReauthorizeTimeout { get; private set; } = null!;

        /// <summary>
        /// The refresh token for API operations.
        /// </summary>
        [Output("refreshToken")]
        public Output<string?> RefreshToken { get; private set; } = null!;

        /// <summary>
        /// The base url for API operations.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs args, CustomResourceOptions? options = null)
            : base("vra", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/schmidtw/pulumi-vra/releases/download/v${VERSION}",
                AdditionalSecretOutputs =
                {
                    "accessToken",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class ProviderArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessToken", required: true)]
        private Input<string>? _accessToken;

        /// <summary>
        /// The access token for API operations.
        /// </summary>
        public Input<string>? AccessToken
        {
            get => _accessToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Specify whether to validate TLS certificates.
        /// </summary>
        [Input("insecure", json: true)]
        public Input<bool>? Insecure { get; set; }

        /// <summary>
        /// Specify timeout for how often to reauthorize the access token
        /// </summary>
        [Input("reauthorizeTimeout")]
        public Input<string>? ReauthorizeTimeout { get; set; }

        /// <summary>
        /// The refresh token for API operations.
        /// </summary>
        [Input("refreshToken")]
        public Input<string>? RefreshToken { get; set; }

        /// <summary>
        /// The base url for API operations.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public ProviderArgs()
        {
        }
        public static new ProviderArgs Empty => new ProviderArgs();
    }
}