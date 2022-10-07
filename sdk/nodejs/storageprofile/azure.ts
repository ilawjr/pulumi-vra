// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 * ### S
 * This is an example of how to create a storage profile azure resource.
 *
 * **Vra storage profile azure:**
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vra from "@pulumiverse/vra";
 *
 * // Azure storage profile using vra_storage_profile_azure resource with managed disk.
 * const thisAzure = new vra.storageprofile.Azure("thisAzure", {
 *     description: "Azure Storage Profile with managed disks.",
 *     regionId: data.vra_region["this"].id,
 *     defaultItem: false,
 *     supportsEncryption: false,
 *     dataDiskCaching: "None",
 *     diskType: "Standard_LRS",
 *     osDiskCaching: "None",
 *     tags: [{
 *         key: "foo",
 *         value: "bar",
 *     }],
 * });
 * // Azure storage profile using vra_storage_profile_azure resource with unmanaged disk.
 * const thisStorageprofile_azureAzure = new vra.storageprofile.Azure("thisStorageprofile/azureAzure", {
 *     description: "Azure Storage Profile with unmanaged disks.",
 *     regionId: data.vra_region["this"].id,
 *     defaultItem: false,
 *     supportsEncryption: false,
 *     dataDiskCaching: "None",
 *     osDiskCaching: "None",
 *     tags: [{
 *         key: "foo",
 *         value: "bar",
 *     }],
 * });
 * ```
 *
 * A storage profile azure resource supports the following arguments:
 */
export class Azure extends pulumi.CustomResource {
    /**
     * Get an existing Azure resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AzureState, opts?: pulumi.CustomResourceOptions): Azure {
        return new Azure(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vra:storageprofile/azure:Azure';

    /**
     * Returns true if the given object is an instance of Azure.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Azure {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Azure.__pulumiType;
    }

    /**
     * Date when the entity was created. The date is in ISO 6801 and UTC.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Indicates the caching mechanism for additional disk.
     */
    public readonly dataDiskCaching!: pulumi.Output<string>;
    /**
     * Indicates if this storage profile is a default profile.
     */
    public readonly defaultItem!: pulumi.Output<boolean>;
    /**
     * A human-friendly description.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.
     */
    public readonly diskType!: pulumi.Output<string>;
    /**
     * The id of the region as seen in the cloud provider for which this profile is defined.
     */
    public /*out*/ readonly externalRegionId!: pulumi.Output<string>;
    /**
     * HATEOAS of the entity
     */
    public /*out*/ readonly links!: pulumi.Output<outputs.storageprofile.AzureLink[]>;
    /**
     * A human-friendly name used as an identifier in APIs that support this option.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The id of the organization this entity belongs to.
     */
    public /*out*/ readonly organizationId!: pulumi.Output<string>;
    /**
     * Indicates the caching mechanism for OS disk. Default policy for OS disks is Read/Write.
     */
    public readonly osDiskCaching!: pulumi.Output<string>;
    /**
     * Email of the user that owns the entity.
     */
    public /*out*/ readonly owner!: pulumi.Output<string>;
    /**
     * A link to the region that is associated with the storage profile.
     */
    public readonly regionId!: pulumi.Output<string>;
    /**
     * Id of a storage account where in the disk is placed.
     */
    public readonly storageAccountId!: pulumi.Output<string>;
    /**
     * Indicates whether this storage policy should support encryption or not.
     */
    public readonly supportsEncryption!: pulumi.Output<boolean>;
    /**
     * A set of tag keys and optional values that were set on this Network Profile.
     * example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
     */
    public readonly tags!: pulumi.Output<outputs.storageprofile.AzureTag[]>;
    /**
     * Date when the entity was last updated. The date is ISO 8601 and UTC.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a Azure resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AzureArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AzureArgs | AzureState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AzureState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["dataDiskCaching"] = state ? state.dataDiskCaching : undefined;
            resourceInputs["defaultItem"] = state ? state.defaultItem : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["diskType"] = state ? state.diskType : undefined;
            resourceInputs["externalRegionId"] = state ? state.externalRegionId : undefined;
            resourceInputs["links"] = state ? state.links : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["osDiskCaching"] = state ? state.osDiskCaching : undefined;
            resourceInputs["owner"] = state ? state.owner : undefined;
            resourceInputs["regionId"] = state ? state.regionId : undefined;
            resourceInputs["storageAccountId"] = state ? state.storageAccountId : undefined;
            resourceInputs["supportsEncryption"] = state ? state.supportsEncryption : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as AzureArgs | undefined;
            if ((!args || args.defaultItem === undefined) && !opts.urn) {
                throw new Error("Missing required property 'defaultItem'");
            }
            if ((!args || args.regionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'regionId'");
            }
            resourceInputs["dataDiskCaching"] = args ? args.dataDiskCaching : undefined;
            resourceInputs["defaultItem"] = args ? args.defaultItem : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["diskType"] = args ? args.diskType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["osDiskCaching"] = args ? args.osDiskCaching : undefined;
            resourceInputs["regionId"] = args ? args.regionId : undefined;
            resourceInputs["storageAccountId"] = args ? args.storageAccountId : undefined;
            resourceInputs["supportsEncryption"] = args ? args.supportsEncryption : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["externalRegionId"] = undefined /*out*/;
            resourceInputs["links"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["owner"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Azure.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Azure resources.
 */
export interface AzureState {
    /**
     * Date when the entity was created. The date is in ISO 6801 and UTC.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Indicates the caching mechanism for additional disk.
     */
    dataDiskCaching?: pulumi.Input<string>;
    /**
     * Indicates if this storage profile is a default profile.
     */
    defaultItem?: pulumi.Input<boolean>;
    /**
     * A human-friendly description.
     */
    description?: pulumi.Input<string>;
    /**
     * Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.
     */
    diskType?: pulumi.Input<string>;
    /**
     * The id of the region as seen in the cloud provider for which this profile is defined.
     */
    externalRegionId?: pulumi.Input<string>;
    /**
     * HATEOAS of the entity
     */
    links?: pulumi.Input<pulumi.Input<inputs.storageprofile.AzureLink>[]>;
    /**
     * A human-friendly name used as an identifier in APIs that support this option.
     */
    name?: pulumi.Input<string>;
    /**
     * The id of the organization this entity belongs to.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * Indicates the caching mechanism for OS disk. Default policy for OS disks is Read/Write.
     */
    osDiskCaching?: pulumi.Input<string>;
    /**
     * Email of the user that owns the entity.
     */
    owner?: pulumi.Input<string>;
    /**
     * A link to the region that is associated with the storage profile.
     */
    regionId?: pulumi.Input<string>;
    /**
     * Id of a storage account where in the disk is placed.
     */
    storageAccountId?: pulumi.Input<string>;
    /**
     * Indicates whether this storage policy should support encryption or not.
     */
    supportsEncryption?: pulumi.Input<boolean>;
    /**
     * A set of tag keys and optional values that were set on this Network Profile.
     * example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
     */
    tags?: pulumi.Input<pulumi.Input<inputs.storageprofile.AzureTag>[]>;
    /**
     * Date when the entity was last updated. The date is ISO 8601 and UTC.
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Azure resource.
 */
export interface AzureArgs {
    /**
     * Indicates the caching mechanism for additional disk.
     */
    dataDiskCaching?: pulumi.Input<string>;
    /**
     * Indicates if this storage profile is a default profile.
     */
    defaultItem: pulumi.Input<boolean>;
    /**
     * A human-friendly description.
     */
    description?: pulumi.Input<string>;
    /**
     * Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.
     */
    diskType?: pulumi.Input<string>;
    /**
     * A human-friendly name used as an identifier in APIs that support this option.
     */
    name?: pulumi.Input<string>;
    /**
     * Indicates the caching mechanism for OS disk. Default policy for OS disks is Read/Write.
     */
    osDiskCaching?: pulumi.Input<string>;
    /**
     * A link to the region that is associated with the storage profile.
     */
    regionId: pulumi.Input<string>;
    /**
     * Id of a storage account where in the disk is placed.
     */
    storageAccountId?: pulumi.Input<string>;
    /**
     * Indicates whether this storage policy should support encryption or not.
     */
    supportsEncryption?: pulumi.Input<boolean>;
    /**
     * A set of tag keys and optional values that were set on this Network Profile.
     * example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
     */
    tags?: pulumi.Input<pulumi.Input<inputs.storageprofile.AzureTag>[]>;
}