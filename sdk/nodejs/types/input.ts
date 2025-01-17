// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

export namespace blockdevice {
    export interface BlockDeviceConstraint {
        expression: pulumi.Input<string>;
        mandatory: pulumi.Input<boolean>;
    }

    export interface BlockDeviceLink {
        href?: pulumi.Input<string>;
        hrefs?: pulumi.Input<pulumi.Input<string>[]>;
        rel: pulumi.Input<string>;
    }

    export interface BlockDeviceSnapshot {
        /**
         * Date when entity was created. Date and time format is ISO 8601 and UTC.
         */
        createdAt?: pulumi.Input<string>;
        /**
         * Describes machine within the scope of your organization and is not propagated to the cloud.
         */
        description?: pulumi.Input<string>;
        /**
         * ID of the block device snapshot.
         */
        id?: pulumi.Input<string>;
        /**
         * Indicates whether snapshot on block device is current.
         */
        isCurrent?: pulumi.Input<boolean>;
        /**
         * HATEOAS of the entity
         */
        links?: pulumi.Input<pulumi.Input<inputs.blockdevice.BlockDeviceSnapshotLink>[]>;
        /**
         * Human-friendly name used as an identifier in APIs that support this option.
         */
        name?: pulumi.Input<string>;
        /**
         * ID of organization that block device snapshot belongs to.
         */
        orgId?: pulumi.Input<string>;
        /**
         * Email of block device snapshot owner.
         */
        owner?: pulumi.Input<string>;
        /**
         * Date when entity was last updated. Date and time format is ISO 8601 and UTC.
         */
        updatedAt?: pulumi.Input<string>;
    }

    export interface BlockDeviceSnapshotLink {
        href?: pulumi.Input<string>;
        hrefs?: pulumi.Input<pulumi.Input<string>[]>;
        rel: pulumi.Input<string>;
    }

    export interface BlockDeviceTag {
        key: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

    export interface GetBlockDeviceTag {
        /**
         * Tag’s key.
         */
        key: string;
        /**
         * Tag’s value.
         */
        value: string;
    }

    export interface GetBlockDeviceTagArgs {
        /**
         * Tag’s key.
         */
        key: pulumi.Input<string>;
        /**
         * Tag’s value.
         */
        value: pulumi.Input<string>;
    }

    export interface SnapshotLink {
        href?: pulumi.Input<string>;
        hrefs?: pulumi.Input<pulumi.Input<string>[]>;
        rel: pulumi.Input<string>;
    }

}

export namespace blueprint {
    export interface BlueprintValidationMessage {
        message?: pulumi.Input<string>;
        metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        path?: pulumi.Input<string>;
        resourceName?: pulumi.Input<string>;
        type?: pulumi.Input<string>;
    }

}

export namespace catalog {
    export interface ItemEntitlementDefinition {
        /**
         * Description of the catalog item.
         */
        description?: pulumi.Input<string>;
        /**
         * Icon id of associated catalog item.
         */
        iconId?: pulumi.Input<string>;
        /**
         * Id of the catalog item.
         */
        id?: pulumi.Input<string>;
        /**
         * Name of the catalog item.
         */
        name?: pulumi.Input<string>;
        /**
         * Number of items in the associated catalog source.
         */
        numberOfItems?: pulumi.Input<number>;
        /**
         * Catalog source name.
         */
        sourceName?: pulumi.Input<string>;
        /**
         * Catalog source type.
         */
        sourceType?: pulumi.Input<string>;
        /**
         * Content definition type.
         */
        type?: pulumi.Input<string>;
    }

    export interface SourceEntitlementDefinition {
        description?: pulumi.Input<string>;
        iconId?: pulumi.Input<string>;
        id?: pulumi.Input<string>;
        name?: pulumi.Input<string>;
        numberOfItems?: pulumi.Input<number>;
        sourceName?: pulumi.Input<string>;
        sourceType?: pulumi.Input<string>;
        type?: pulumi.Input<string>;
    }

}

export namespace cloudaccount {
    export interface AwsLink {
        href?: pulumi.Input<string>;
        hrefs?: pulumi.Input<pulumi.Input<string>[]>;
        rel: pulumi.Input<string>;
    }

    export interface AwsTag {
        key: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

    export interface AzureLink {
        href?: pulumi.Input<string>;
        hrefs?: pulumi.Input<pulumi.Input<string>[]>;
        rel: pulumi.Input<string>;
    }

    export interface AzureTag {
        key: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

    export interface GcpLink {
        href?: pulumi.Input<string>;
        hrefs?: pulumi.Input<pulumi.Input<string>[]>;
        rel: pulumi.Input<string>;
    }

    export interface GcpTag {
        key: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

    export interface GetAwsTagArgs {
        /**
         * Tag’s key.
         */
        key: pulumi.Input<string>;
        /**
         * Tag’s value.
         */
        value: pulumi.Input<string>;
    }

    export interface GetAwsTag {
        /**
         * Tag’s key.
         */
        key: string;
        /**
         * Tag’s value.
         */
        value: string;
    }

    export interface GetAzureTag {
        /**
         * Tag’s key.
         */
        key: string;
        /**
         * Tag’s value.
         */
        value: string;
    }

    export interface GetAzureTagArgs {
        /**
         * Tag’s key.
         */
        key: pulumi.Input<string>;
        /**
         * Tag’s value.
         */
        value: pulumi.Input<string>;
    }

    export interface GetGcpTagArgs {
        /**
         * Tag’s key.
         */
        key: pulumi.Input<string>;
        /**
         * Tag’s value.
         */
        value: pulumi.Input<string>;
    }

    export interface GetGcpTag {
        /**
         * Tag’s key.
         */
        key: string;
        /**
         * Tag’s value.
         */
        value: string;
    }

    export interface GetNsxtTagArgs {
        /**
         * Tag’s key.
         */
        key: pulumi.Input<string>;
        /**
         * Tag’s value.
         */
        value: pulumi.Input<string>;
    }

    export interface GetNsxtTag {
        /**
         * Tag’s key.
         */
        key: string;
        /**
         * Tag’s value.
         */
        value: string;
    }

    export interface GetNsxvTagArgs {
        /**
         * Tag’s key.
         */
        key: pulumi.Input<string>;
        /**
         * Tag’s value.
         */
        value: pulumi.Input<string>;
    }

    export interface GetNsxvTag {
        /**
         * Tag’s key.
         */
        key: string;
        /**
         * Tag’s value.
         */
        value: string;
    }

    export interface GetVSphereTag {
        /**
         * Tag’s key.
         */
        key: string;
        /**
         * Tag’s value.
         */
        value: string;
    }

    export interface GetVSphereTagArgs {
        /**
         * Tag’s key.
         */
        key: pulumi.Input<string>;
        /**
         * Tag’s value.
         */
        value: pulumi.Input<string>;
    }

    export interface GetVmcTag {
        /**
         * Tag’s key.
         */
        key: string;
        /**
         * Tag’s value.
         */
        value: string;
    }

    export interface GetVmcTagArgs {
        /**
         * Tag’s key.
         */
        key: pulumi.Input<string>;
        /**
         * Tag’s value.
         */
        value: pulumi.Input<string>;
    }

    export interface NsxtLink {
        href?: pulumi.Input<string>;
        hrefs?: pulumi.Input<pulumi.Input<string>[]>;
        rel: pulumi.Input<string>;
    }

    export interface NsxtTag {
        key: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

    export interface NsxvLink {
        href?: pulumi.Input<string>;
        hrefs?: pulumi.Input<pulumi.Input<string>[]>;
        rel: pulumi.Input<string>;
    }

    export interface NsxvTag {
        key: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

    export interface VSphereLink {
        href?: pulumi.Input<string>;
        hrefs?: pulumi.Input<pulumi.Input<string>[]>;
        rel: pulumi.Input<string>;
    }

    export interface VSphereTag {
        key: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

    export interface VmcLink {
        href?: pulumi.Input<string>;
        hrefs?: pulumi.Input<pulumi.Input<string>[]>;
        rel: pulumi.Input<string>;
    }

    export interface VmcTag {
        key: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }
}

export namespace contentsource {
    export interface ContentSourceConfig {
        /**
         * Content source branch name.
         */
        branch?: pulumi.Input<string>;
        /**
         * Content source type. Supported values are `BLUEPRINT`, `IMAGE`, `ABX_SCRIPTS`, `TERRAFORM_CONFIGURATION`.
         */
        contentType?: pulumi.Input<string>;
        /**
         * Content source integration id as seen in vRA integrations.
         */
        integrationId: pulumi.Input<string>;
        /**
         * Path to refer to in the content source repository and branch.
         */
        path: pulumi.Input<string>;
        /**
         * Name of the project.
         */
        projectName: pulumi.Input<string>;
        /**
         * Content source repository.
         */
        repository?: pulumi.Input<string>;
    }
}

export namespace deployment {
    export interface DeploymentExpense {
        /**
         * Additional expense incurred for the resource.
         */
        additionalExpense?: pulumi.Input<number>;
        /**
         * Expense sync message code if any.
         */
        code?: pulumi.Input<string>;
        /**
         * Compute expense of the entity.
         */
        computeExpense?: pulumi.Input<number>;
        /**
         * Last expense sync time.
         */
        lastUpdateTime?: pulumi.Input<string>;
        /**
         * Expense sync message if any.
         */
        message?: pulumi.Input<string>;
        /**
         * Network expense of the entity.
         */
        networkExpense?: pulumi.Input<number>;
        /**
         * Storage expense of the entity.
         */
        storageExpense?: pulumi.Input<number>;
        /**
         * Total expense of the entity.
         */
        totalExpense?: pulumi.Input<number>;
        /**
         * Monetary unit.
         */
        unit?: pulumi.Input<string>;
    }

    export interface DeploymentLastRequest {
        /**
         * Identifier of the requested action.
         */
        actionId?: pulumi.Input<string>;
        /**
         * Time at which the request was approved.
         */
        approvedAt?: pulumi.Input<string>;
        /**
         * Identifier of the requested blueprint in the form ‘UUID:version’.
         */
        blueprintId?: pulumi.Input<string>;
        /**
         * Indicates whether request can be canceled or not.
         */
        cancelable?: pulumi.Input<boolean>;
        /**
         * The id of the vRA catalog item to request the deployment. Conflicts with `blueprintId` and `blueprintContent`.
         */
        catalogItemId?: pulumi.Input<string>;
        /**
         * Time at which the request completed.
         */
        completedAt?: pulumi.Input<string>;
        /**
         * The number of tasks completed while fulfilling this request.
         */
        completedTasks?: pulumi.Input<number>;
        /**
         * Creation time (e.g. date format ‘2019-07-13T23:16:49.310Z’).
         */
        createdAt?: pulumi.Input<string>;
        /**
         * Longer user-friendly details of the request.
         */
        details?: pulumi.Input<string>;
        /**
         * Indicates whether request is in dismissed state.
         */
        dismissed?: pulumi.Input<boolean>;
        /**
         * Unique identifier of the resource.
         */
        id?: pulumi.Input<string>;
        /**
         * Time at which the request was initialized.
         */
        initializedAt?: pulumi.Input<string>;
        /**
         * Inputs provided by the user. For inputs including those with default values, refer to `inputsIncludingDefaults`.
         */
        inputs?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        /**
         * A human-friendly name used as an identifier in APIs that support this option.
         */
        name?: pulumi.Input<string>;
        /**
         * Request outputs.
         */
        outputs?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        /**
         * The user that initiated the request.
         */
        requestedBy?: pulumi.Input<string>;
        resourceIds?: pulumi.Input<pulumi.Input<string>[]>;
        /**
         * Deployment status. Supported values are: `CREATE_SUCCESSFUL`, `CREATE_INPROGRESS`, `CREATE_FAILED`, `UPDATE_SUCCESSFUL`, `UPDATE_INPROGRESS`, `UPDATE_FAILED`, `DELETE_SUCCESSFUL`, `DELETE_INPROGRESS`, `DELETE_FAILED`, `ACTION_SUCCESSFUL`, `ACTION_INPROGRESS`, `ACTION_FAILED`.
         */
        status?: pulumi.Input<string>;
        totalTasks?: pulumi.Input<number>;
        /**
         * Last update time (e.g. date format ‘2019-07-13T23:16:49.310Z’).
         */
        updatedAt?: pulumi.Input<string>;
    }

    export interface DeploymentProject {
        /**
         * A human-friendly description.
         */
        description?: pulumi.Input<string>;
        /**
         * Unique identifier of the resource.
         */
        id?: pulumi.Input<string>;
        /**
         * A human-friendly name used as an identifier in APIs that support this option.
         */
        name?: pulumi.Input<string>;
        /**
         * Version of the entity, if applicable.
         */
        version?: pulumi.Input<string>;
    }

    export interface DeploymentResource {
        /**
         * Creation time (e.g. date format ‘2019-07-13T23:16:49.310Z’).
         */
        createdAt?: pulumi.Input<string>;
        /**
         * A list of other resources this resource depends on.
         */
        dependsOns?: pulumi.Input<pulumi.Input<string>[]>;
        /**
         * A human-friendly description.
         */
        description?: pulumi.Input<string>;
        /**
         * Expense incurred for the deployment.
         */
        expenses?: pulumi.Input<pulumi.Input<inputs.deployment.DeploymentResourceExpense>[]>;
        /**
         * Unique identifier of the resource.
         */
        id: pulumi.Input<string>;
        /**
         * A human-friendly name used as an identifier in APIs that support this option.
         */
        name: pulumi.Input<string>;
        /**
         * List of properties in the encoded JSON string format.
         */
        propertiesJson?: pulumi.Input<string>;
        /**
         * The current state of the resource. Supported values are `PARTIAL`, `TAINTED`, `OK.`
         */
        state?: pulumi.Input<string>;
        /**
         * The current sync status. Supported values are `SUCCESS`, `MISSING`, `STALE`.
         */
        syncStatus?: pulumi.Input<string>;
        /**
         * Type of the resource.
         */
        type?: pulumi.Input<string>;
    }

    export interface DeploymentResourceExpense {
        /**
         * Additional expense incurred for the resource.
         */
        additionalExpense?: pulumi.Input<number>;
        /**
         * Expense sync message code if any.
         */
        code?: pulumi.Input<string>;
        /**
         * Compute expense of the entity.
         */
        computeExpense?: pulumi.Input<number>;
        /**
         * Last expense sync time.
         */
        lastUpdateTime?: pulumi.Input<string>;
        /**
         * Expense sync message if any.
         */
        message?: pulumi.Input<string>;
        /**
         * Network expense of the entity.
         */
        networkExpense?: pulumi.Input<number>;
        /**
         * Storage expense of the entity.
         */
        storageExpense?: pulumi.Input<number>;
        /**
         * Total expense of the entity.
         */
        totalExpense?: pulumi.Input<number>;
        /**
         * Monetary unit.
         */
        unit?: pulumi.Input<string>;
    }

}

export namespace fabric {
    export interface ComputeLink {
        href?: pulumi.Input<string>;
        hrefs?: pulumi.Input<pulumi.Input<string>[]>;
        rel: pulumi.Input<string>;
    }

    export interface ComputeTag {
        /**
         * Tag’s key.
         */
        key: pulumi.Input<string>;
        /**
         * Tag’s value.
         */
        value: pulumi.Input<string>;
    }

    export interface DatastoreVSphereLink {
        href?: pulumi.Input<string>;
        hrefs?: pulumi.Input<pulumi.Input<string>[]>;
        rel: pulumi.Input<string>;
    }

    export interface DatastoreVSphereTag {
        /**
         * Tag’s key.
         */
        key: pulumi.Input<string>;
        /**
         * Tag’s value.
         */
        value: pulumi.Input<string>;
    }

    export interface GetComputeTagArgs {
        /**
         * Tag’s key.
         */
        key: pulumi.Input<string>;
        /**
         * Tag’s value.
         */
        value: pulumi.Input<string>;
    }

    export interface GetComputeTag {
        /**
         * Tag’s key.
         */
        key: string;
        /**
         * Tag’s value.
         */
        value: string;
    }

    export interface GetDatastoreVSphereTag {
        /**
         * Tag’s key.
         */
        key: string;
        /**
         * Tag’s value.
         */
        value: string;
    }

    export interface GetDatastoreVSphereTagArgs {
        /**
         * Tag’s key.
         */
        key: pulumi.Input<string>;
        /**
         * Tag’s value.
         */
        value: pulumi.Input<string>;
    }

    export interface GetNetworkTag {
        key: string;
        value: string;
    }

    export interface GetNetworkTagArgs {
        key: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

    export interface NetworkVSphereLink {
        href?: pulumi.Input<string>;
        hrefs?: pulumi.Input<pulumi.Input<string>[]>;
        rel: pulumi.Input<string>;
    }

    export interface NetworkVSphereTag {
        key: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }
}

export namespace flavor {
    export interface ProfileFlavorMapping {
        cpuCount?: pulumi.Input<number>;
        instanceType?: pulumi.Input<string>;
        memory?: pulumi.Input<number>;
        /**
         * A human-friendly name used as an identifier in APIs that support this option.
         */
        name: pulumi.Input<string>;
    }

    export interface ProfileLink {
        href?: pulumi.Input<string>;
        hrefs?: pulumi.Input<pulumi.Input<string>[]>;
        rel: pulumi.Input<string>;
    }
}

export namespace image {
    export interface GetProfileImageMapping {
        cloudConfig?: string;
        constraints?: inputs.image.GetProfileImageMappingConstraint[];
        /**
         * A human-friendly description.
         */
        description?: string;
        externalId?: string;
        /**
         * The external regionId of the resource.
         */
        externalRegionId?: string;
        imageId?: string;
        imageName?: string;
        /**
         * A human-friendly name used as an identifier in APIs that support this option.
         */
        name: string;
        organization?: string;
        osFamily?: string;
        /**
         * Email of the user that owns the entity.
         */
        owner?: string;
        private?: boolean;
    }

    export interface GetProfileImageMappingArgs {
        cloudConfig?: pulumi.Input<string>;
        constraints?: pulumi.Input<pulumi.Input<inputs.image.GetProfileImageMappingConstraintArgs>[]>;
        /**
         * A human-friendly description.
         */
        description?: pulumi.Input<string>;
        externalId?: pulumi.Input<string>;
        /**
         * The external regionId of the resource.
         */
        externalRegionId?: pulumi.Input<string>;
        imageId?: pulumi.Input<string>;
        imageName?: pulumi.Input<string>;
        /**
         * A human-friendly name used as an identifier in APIs that support this option.
         */
        name: pulumi.Input<string>;
        organization?: pulumi.Input<string>;
        osFamily?: pulumi.Input<string>;
        /**
         * Email of the user that owns the entity.
         */
        owner?: pulumi.Input<string>;
        private?: pulumi.Input<boolean>;
    }

    export interface GetProfileImageMappingConstraint {
        expression: string;
        mandatory: boolean;
    }

    export interface GetProfileImageMappingConstraintArgs {
        expression: pulumi.Input<string>;
        mandatory: pulumi.Input<boolean>;
    }

    export interface ProfileImageMapping {
        cloudConfig?: pulumi.Input<string>;
        constraints?: pulumi.Input<pulumi.Input<inputs.image.ProfileImageMappingConstraint>[]>;
        /**
         * A human-friendly description.
         */
        description?: pulumi.Input<string>;
        externalId?: pulumi.Input<string>;
        /**
         * The external regionId of the resource.
         */
        externalRegionId?: pulumi.Input<string>;
        imageId?: pulumi.Input<string>;
        imageName?: pulumi.Input<string>;
        /**
         * A human-friendly name used as an identifier in APIs that support this option.
         */
        name: pulumi.Input<string>;
        organization?: pulumi.Input<string>;
        osFamily?: pulumi.Input<string>;
        /**
         * Email of the user that owns the entity.
         */
        owner?: pulumi.Input<string>;
        private?: pulumi.Input<boolean>;
    }

    export interface ProfileImageMappingConstraint {
        expression: pulumi.Input<string>;
        mandatory: pulumi.Input<boolean>;
    }
}

export namespace loadbalancer {
    export interface LoadBalancerLink {
        href?: pulumi.Input<string>;
        hrefs?: pulumi.Input<pulumi.Input<string>[]>;
        rel: pulumi.Input<string>;
    }

    export interface LoadBalancerNic {
        addresses?: pulumi.Input<pulumi.Input<string>[]>;
        /**
         * Additional custom properties that may be used to extend the machine.
         */
        customProperties?: pulumi.Input<{[key: string]: any}>;
        /**
         * Describes machine within the scope of your organization and is not propagated to the cloud.
         */
        description?: pulumi.Input<string>;
        deviceIndex?: pulumi.Input<number>;
        /**
         * A human-friendly name used as an identifier in APIs that support this option.
         */
        name?: pulumi.Input<string>;
        networkId: pulumi.Input<string>;
        securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    }

    export interface LoadBalancerRoute {
        /**
         * Load balancer health check configuration.
         */
        healthCheckConfigurations?: pulumi.Input<pulumi.Input<inputs.loadbalancer.LoadBalancerRouteHealthCheckConfiguration>[]>;
        /**
         * Member port where the traffic is routed to.
         */
        memberPort: pulumi.Input<string>;
        /**
         * The protocol of the member traffic.
         */
        memberProtocol: pulumi.Input<string>;
        /**
         * Port which the load balancer is listening to.
         */
        port: pulumi.Input<string>;
        /**
         * The protocol of the incoming load balancer requests.
         */
        protocol: pulumi.Input<string>;
    }

    export interface LoadBalancerRouteHealthCheckConfiguration {
        /**
         * Number of consecutive successful checks before considering a particular back-end instance as healthy.
         */
        healthyThreshold?: pulumi.Input<number>;
        /**
         * Interval (in seconds) at which the health checks will be performed.
         */
        intervalSeconds?: pulumi.Input<number>;
        /**
         * Port which the load balancer is listening to.
         */
        port: pulumi.Input<string>;
        /**
         * The protocol of the incoming load balancer requests.
         */
        protocol: pulumi.Input<string>;
        /**
         * Timeout (in seconds) to wait for a response from the back-end instance.
         */
        timeoutSeconds?: pulumi.Input<number>;
        unhealthyThreshold?: pulumi.Input<number>;
        urlPath?: pulumi.Input<string>;
    }

    export interface LoadBalancerTag {
        /**
         * Tag’s key.
         */
        key: pulumi.Input<string>;
        /**
         * Tag’s value.
         */
        value: pulumi.Input<string>;
    }

    export interface LoadBalancerTarget {
        machineId: pulumi.Input<string>;
        networkInterfaceId?: pulumi.Input<string>;
    }
}

export namespace machine {
    export interface GetMachineTagArgs {
        key: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

    export interface GetMachineTag {
        key: string;
        value: string;
    }

    export interface MachineBootConfig {
        /**
         * Calid cloud config data in json-escaped yaml syntax.
         */
        content?: pulumi.Input<string>;
    }

    export interface MachineConstraint {
        /**
         * Constraint that is conveyed to the policy engine. An expression of the form "[!]tag-key[:[tag-value]]", used to indicate a constraint match on keys and values of tags.
         */
        expression: pulumi.Input<string>;
        /**
         * Indicates whether this constraint should be strictly enforced or not.
         */
        mandatory: pulumi.Input<boolean>;
    }

    export interface MachineDisk {
        /**
         * ID of the existing block device.
         */
        blockDeviceId: pulumi.Input<string>;
        /**
         * Human-friendly description.
         */
        description?: pulumi.Input<string>;
        /**
         * Human-friendly name used as an identifier in APIs that support this option.
         */
        name?: pulumi.Input<string>;
    }

    export interface MachineDisksList {
        /**
         * ID of the existing block device.
         */
        blockDeviceId: pulumi.Input<string>;
        /**
         * Human-friendly description.
         */
        description?: pulumi.Input<string>;
        /**
         * Human-friendly name used as an identifier in APIs that support this option.
         */
        name?: pulumi.Input<string>;
    }

    export interface MachineImageDiskConstraint {
        /**
         * Constraint that is conveyed to the policy engine. An expression of the form "[!]tag-key[:[tag-value]]", used to indicate a constraint match on keys and values of tags.
         */
        expression: pulumi.Input<string>;
        /**
         * Indicates whether this constraint should be strictly enforced or not.
         */
        mandatory: pulumi.Input<boolean>;
    }

    export interface MachineLink {
        href?: pulumi.Input<string>;
        hrefs?: pulumi.Input<pulumi.Input<string>[]>;
        rel: pulumi.Input<string>;
    }

    export interface MachineNic {
        /**
         * List of IP addresses allocated or in use by this network interface.
         * example:[ "10.1.2.190" ]
         */
        addresses?: pulumi.Input<pulumi.Input<string>[]>;
        /**
         * Additional properties that may be used to extend the base type.
         */
        customProperties?: pulumi.Input<{[key: string]: any}>;
        /**
         * Human-friendly description.
         */
        description?: pulumi.Input<string>;
        /**
         * The device index of this network interface.
         */
        deviceIndex?: pulumi.Input<number>;
        /**
         * Human-friendly name used as an identifier in APIs that support this option.
         */
        name?: pulumi.Input<string>;
        /**
         * ID of the network instance that this network interface plugs into.
         */
        networkId: pulumi.Input<string>;
        /**
         * List of security group ids which this network interface will be assigned to.
         */
        securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    }

    export interface MachineTag {
        /**
         * Tag’s key.
         */
        key: pulumi.Input<string>;
        /**
         * Tag’s value.
         */
        value: pulumi.Input<string>;
    }

}

export namespace network {
    export interface GetDomainTagArgs {
        key: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

    export interface GetDomainTag {
        key: string;
        value: string;
    }

    export interface GetNetworkConstraintArgs {
        expression: pulumi.Input<string>;
        mandatory: pulumi.Input<boolean>;
    }

    export interface GetNetworkConstraint {
        expression: string;
        mandatory: boolean;
    }

    export interface GetNetworkTagArgs {
        key: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

    export interface GetNetworkTag {
        key: string;
        value: string;
    }

    export interface GetProfileTag {
        key: string;
        value: string;
    }

    export interface GetProfileTagArgs {
        key: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

    export interface IpRangeLink {
        href?: pulumi.Input<string>;
        hrefs?: pulumi.Input<pulumi.Input<string>[]>;
        rel: pulumi.Input<string>;
    }

    export interface IpRangeTag {
        key: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

    export interface NetworkConstraint {
        expression: pulumi.Input<string>;
        mandatory: pulumi.Input<boolean>;
    }

    export interface NetworkLink {
        href?: pulumi.Input<string>;
        hrefs?: pulumi.Input<pulumi.Input<string>[]>;
        rel: pulumi.Input<string>;
    }

    export interface NetworkTag {
        key: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

    export interface ProfileLink {
        href?: pulumi.Input<string>;
        hrefs?: pulumi.Input<pulumi.Input<string>[]>;
        rel: pulumi.Input<string>;
    }

    export interface ProfileTag {
        key: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }
}

export namespace project {
    export interface GetProjectAdministratorRole {
        email: string;
        type?: string;
    }

    export interface GetProjectAdministratorRoleArgs {
        email: pulumi.Input<string>;
        type?: pulumi.Input<string>;
    }

    export interface GetProjectConstraints {
        extensibilities?: inputs.project.GetProjectConstraintsExtensibility[];
        networks?: inputs.project.GetProjectConstraintsNetwork[];
        storages?: inputs.project.GetProjectConstraintsStorage[];
    }

    export interface GetProjectConstraintsArgs {
        extensibilities?: pulumi.Input<pulumi.Input<inputs.project.GetProjectConstraintsExtensibilityArgs>[]>;
        networks?: pulumi.Input<pulumi.Input<inputs.project.GetProjectConstraintsNetworkArgs>[]>;
        storages?: pulumi.Input<pulumi.Input<inputs.project.GetProjectConstraintsStorageArgs>[]>;
    }

    export interface GetProjectConstraintsExtensibility {
        expression: string;
        mandatory: boolean;
    }

    export interface GetProjectConstraintsExtensibilityArgs {
        expression: pulumi.Input<string>;
        mandatory: pulumi.Input<boolean>;
    }

    export interface GetProjectConstraintsNetwork {
        expression: string;
        mandatory: boolean;
    }

    export interface GetProjectConstraintsNetworkArgs {
        expression: pulumi.Input<string>;
        mandatory: pulumi.Input<boolean>;
    }

    export interface GetProjectConstraintsStorageArgs {
        expression: pulumi.Input<string>;
        mandatory: pulumi.Input<boolean>;
    }

    export interface GetProjectConstraintsStorage {
        expression: string;
        mandatory: boolean;
    }

    export interface GetProjectMemberRole {
        email: string;
        type?: string;
    }

    export interface GetProjectMemberRoleArgs {
        email: pulumi.Input<string>;
        type?: pulumi.Input<string>;
    }

    export interface GetProjectViewerRole {
        email: string;
        type?: string;
    }

    export interface GetProjectViewerRoleArgs {
        email: pulumi.Input<string>;
        type?: pulumi.Input<string>;
    }

    export interface GetProjectZoneAssignment {
        cpuLimit?: number;
        maxInstances?: number;
        memoryLimitMb?: number;
        priority?: number;
        storageLimitGb?: number;
        zoneId?: string;
    }

    export interface GetProjectZoneAssignmentArgs {
        cpuLimit?: pulumi.Input<number>;
        maxInstances?: pulumi.Input<number>;
        memoryLimitMb?: pulumi.Input<number>;
        priority?: pulumi.Input<number>;
        storageLimitGb?: pulumi.Input<number>;
        zoneId?: pulumi.Input<string>;
    }

    export interface ProjectAdministratorRole {
        email: pulumi.Input<string>;
        type?: pulumi.Input<string>;
    }

    export interface ProjectConstraints {
        extensibilities?: pulumi.Input<pulumi.Input<inputs.project.ProjectConstraintsExtensibility>[]>;
        networks?: pulumi.Input<pulumi.Input<inputs.project.ProjectConstraintsNetwork>[]>;
        storages?: pulumi.Input<pulumi.Input<inputs.project.ProjectConstraintsStorage>[]>;
    }

    export interface ProjectConstraintsExtensibility {
        expression: pulumi.Input<string>;
        mandatory: pulumi.Input<boolean>;
    }

    export interface ProjectConstraintsNetwork {
        expression: pulumi.Input<string>;
        mandatory: pulumi.Input<boolean>;
    }

    export interface ProjectConstraintsStorage {
        expression: pulumi.Input<string>;
        mandatory: pulumi.Input<boolean>;
    }

    export interface ProjectMemberRole {
        email: pulumi.Input<string>;
        type?: pulumi.Input<string>;
    }

    export interface ProjectViewerRole {
        email: pulumi.Input<string>;
        type?: pulumi.Input<string>;
    }

    export interface ProjectZoneAssignment {
        cpuLimit?: pulumi.Input<number>;
        maxInstances?: pulumi.Input<number>;
        memoryLimitMb?: pulumi.Input<number>;
        priority?: pulumi.Input<number>;
        storageLimitGb?: pulumi.Input<number>;
        zoneId: pulumi.Input<string>;
    }
}

export namespace securitygroup {
    export interface GetSecurityGroupRule {
        access: string;
        direction: string;
        ipRangeCidr: number;
        /**
         * Name of the security group.
         */
        name?: string;
        ports: string;
        protocol: string;
        service?: string;
    }

    export interface GetSecurityGroupRuleArgs {
        access: pulumi.Input<string>;
        direction: pulumi.Input<string>;
        ipRangeCidr: pulumi.Input<number>;
        /**
         * Name of the security group.
         */
        name?: pulumi.Input<string>;
        ports: pulumi.Input<string>;
        protocol: pulumi.Input<string>;
        service?: pulumi.Input<string>;
    }
}

export namespace storageprofile {
    export interface AwsLink {
        href?: pulumi.Input<string>;
        hrefs?: pulumi.Input<pulumi.Input<string>[]>;
        rel: pulumi.Input<string>;
    }

    export interface AwsTag {
        key: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

    export interface AzureLink {
        href?: pulumi.Input<string>;
        hrefs?: pulumi.Input<pulumi.Input<string>[]>;
        rel: pulumi.Input<string>;
    }

    export interface AzureTag {
        key: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

    export interface GetAwsTagArgs {
        key: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

    export interface GetAwsTag {
        key: string;
        value: string;
    }

    export interface GetAzureTagArgs {
        key: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

    export interface GetAzureTag {
        key: string;
        value: string;
    }

    export interface GetStorageProfileTagArgs {
        key: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

    export interface GetStorageProfileTag {
        key: string;
        value: string;
    }

    export interface GetVSphereTag {
        key: string;
        value: string;
    }

    export interface GetVSphereTagArgs {
        key: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

    export interface StorageProfileLink {
        href?: pulumi.Input<string>;
        hrefs?: pulumi.Input<pulumi.Input<string>[]>;
        rel: pulumi.Input<string>;
    }

    export interface StorageProfileTag {
        key: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

    export interface VSphereLink {
        href?: pulumi.Input<string>;
        hrefs?: pulumi.Input<pulumi.Input<string>[]>;
        rel: pulumi.Input<string>;
    }

    export interface VSphereTag {
        key: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

}

export namespace zone {
    export interface GetZoneTag {
        /**
         * Tag’s key.
         */
        key: string;
        /**
         * Tag’s value.
         */
        value: string;
    }

    export interface GetZoneTagArgs {
        /**
         * Tag’s key.
         */
        key: pulumi.Input<string>;
        /**
         * Tag’s value.
         */
        value: pulumi.Input<string>;
    }

    export interface GetZoneTagsToMatch {
        /**
         * Tag’s key.
         */
        key: string;
        /**
         * Tag’s value.
         */
        value: string;
    }

    export interface GetZoneTagsToMatchArgs {
        /**
         * Tag’s key.
         */
        key: pulumi.Input<string>;
        /**
         * Tag’s value.
         */
        value: pulumi.Input<string>;
    }

    export interface ZoneLink {
        href?: pulumi.Input<string>;
        hrefs?: pulumi.Input<pulumi.Input<string>[]>;
        rel: pulumi.Input<string>;
    }

    export interface ZoneTag {
        /**
         * Tag’s key.
         */
        key: pulumi.Input<string>;
        /**
         * Tag’s value.
         */
        value: pulumi.Input<string>;
    }

    export interface ZoneTagsToMatch {
        /**
         * Tag’s key.
         */
        key: pulumi.Input<string>;
        /**
         * Tag’s value.
         */
        value: pulumi.Input<string>;
    }
}
