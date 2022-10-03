// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This data source provides information about a deployment in vRA.
 *
 * ## Example Usage
 * ### S
 *
 * This is an example of how to get a vRA deployment by its name.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vra from "@pulumi/vra";
 *
 * const this = vra.getDeployment({
 *     name: _var.deployment_name,
 * });
 * ```
 *
 * This is an example of how to get a vRA cloud template by its id.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vra from "@pulumi/vra";
 *
 * const this = vra.getDeployment({
 *     id: _var.deployment_id,
 * });
 * ```
 */
export function getDeployment(args?: GetDeploymentArgs, opts?: pulumi.InvokeOptions): Promise<GetDeploymentResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vra:index/getDeployment:getDeployment", {
        "blueprintContent": args.blueprintContent,
        "blueprintId": args.blueprintId,
        "blueprintVersion": args.blueprintVersion,
        "catalogItemId": args.catalogItemId,
        "catalogItemVersion": args.catalogItemVersion,
        "createdAt": args.createdAt,
        "createdBy": args.createdBy,
        "description": args.description,
        "expandLastRequest": args.expandLastRequest,
        "expandProject": args.expandProject,
        "expandResources": args.expandResources,
        "id": args.id,
        "inputs": args.inputs,
        "lastUpdatedAt": args.lastUpdatedAt,
        "lastUpdatedBy": args.lastUpdatedBy,
        "leaseExpireAt": args.leaseExpireAt,
        "name": args.name,
        "owner": args.owner,
        "projectId": args.projectId,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getDeployment.
 */
export interface GetDeploymentArgs {
    /**
     * vRA Cloud template content. Conflicts with `blueprintId` and `catalogItemId`.
     */
    blueprintContent?: string;
    /**
     * Identifier of the requested blueprint in the form ‘UUID:version’.
     */
    blueprintId?: string;
    /**
     * The version of the vRA cloud template to request the deployment. Used only when `blueprintId` is provided.
     */
    blueprintVersion?: string;
    /**
     * Identifier of the requested catalog item in the form ‘UUID:version’.
     */
    catalogItemId?: string;
    /**
     * The version of the vRA catalog item to request the deployment. Used only when `catalogItemId` is provided.
     */
    catalogItemVersion?: string;
    /**
     * Creation time (e.g. date format ‘2019-07-13T23:16:49.310Z’).
     */
    createdAt?: string;
    /**
     * The user the entity was created by.
     */
    createdBy?: string;
    /**
     * A description of the resource.
     */
    description?: string;
    /**
     * Flag to indicate whether to expand last request on the deployment.
     */
    expandLastRequest?: boolean;
    /**
     * Flag to indicate whether to expand project information.
     */
    expandProject?: boolean;
    /**
     * Flag to indicate whether to expand resources in the deployment.
     */
    expandResources?: boolean;
    /**
     * The id of the deployment. One of `id` or `name` must be provided.
     */
    id?: string;
    /**
     * List of request inputs.
     */
    inputs?: {[key: string]: string};
    /**
     * Time at which the deployment was last updated.
     */
    lastUpdatedAt?: string;
    /**
     * The user that last updated the deployment.
     */
    lastUpdatedBy?: string;
    /**
     * Time at which the deployment lease expires.
     */
    leaseExpireAt?: string;
    /**
     * Name of the deployment. One of `id` or `name` must be provided.
     */
    name?: string;
    /**
     * The user this deployment belongs to.
     */
    owner?: string;
    /**
     * The id of the project this deployment belongs to.
     */
    projectId?: string;
    /**
     * Deployment status. Supported values are: `CREATE_SUCCESSFUL`, `CREATE_INPROGRESS`, `CREATE_FAILED`, `UPDATE_SUCCESSFUL`, `UPDATE_INPROGRESS`, `UPDATE_FAILED`, `DELETE_SUCCESSFUL`, `DELETE_INPROGRESS`, `DELETE_FAILED`, `ACTION_SUCCESSFUL`, `ACTION_INPROGRESS`, `ACTION_FAILED`.
     */
    status?: string;
}

/**
 * A collection of values returned by getDeployment.
 */
export interface GetDeploymentResult {
    /**
     * vRA Cloud template content. Conflicts with `blueprintId` and `catalogItemId`.
     */
    readonly blueprintContent: string;
    /**
     * Identifier of the requested blueprint in the form ‘UUID:version’.
     */
    readonly blueprintId: string;
    /**
     * The version of the vRA cloud template to request the deployment. Used only when `blueprintId` is provided.
     */
    readonly blueprintVersion: string;
    /**
     * Identifier of the requested catalog item in the form ‘UUID:version’.
     */
    readonly catalogItemId: string;
    /**
     * The version of the vRA catalog item to request the deployment. Used only when `catalogItemId` is provided.
     */
    readonly catalogItemVersion: string;
    /**
     * Creation time (e.g. date format ‘2019-07-13T23:16:49.310Z’).
     */
    readonly createdAt: string;
    /**
     * The user the entity was created by.
     */
    readonly createdBy: string;
    /**
     * A description of the resource.
     */
    readonly description: string;
    readonly expandLastRequest?: boolean;
    readonly expandProject?: boolean;
    readonly expandResources?: boolean;
    /**
     * Expense incurred for this resource.
     */
    readonly expenses: outputs.GetDeploymentExpense[];
    /**
     * Unique identifier of the resource.
     */
    readonly id: string;
    /**
     * List of request inputs.
     */
    readonly inputs: {[key: string]: string};
    /**
     * Represents deployment requests.
     */
    readonly lastRequests: outputs.GetDeploymentLastRequest[];
    /**
     * Time at which the deployment was last updated.
     */
    readonly lastUpdatedAt: string;
    /**
     * The user that last updated the deployment.
     */
    readonly lastUpdatedBy: string;
    /**
     * Time at which the deployment lease expires.
     */
    readonly leaseExpireAt: string;
    /**
     * Name of the resource.
     */
    readonly name: string;
    /**
     * The ID of the organization this deployment belongs to.
     */
    readonly orgId: string;
    /**
     * The user this deployment belongs to.
     */
    readonly owner?: string;
    /**
     * The id of the project this deployment belongs to.
     */
    readonly projectId: string;
    /**
     * The project this entity belongs to.
     */
    readonly projects: outputs.GetDeploymentProject[];
    /**
     * Expanded resources for the deployment. Content of this property will not be maintained backward compatible.
     */
    readonly resources: outputs.GetDeploymentResource[];
    /**
     * Deployment status. Supported values are: `CREATE_SUCCESSFUL`, `CREATE_INPROGRESS`, `CREATE_FAILED`, `UPDATE_SUCCESSFUL`, `UPDATE_INPROGRESS`, `UPDATE_FAILED`, `DELETE_SUCCESSFUL`, `DELETE_INPROGRESS`, `DELETE_FAILED`, `ACTION_SUCCESSFUL`, `ACTION_INPROGRESS`, `ACTION_FAILED`.
     */
    readonly status: string;
}

export function getDeploymentOutput(args?: GetDeploymentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDeploymentResult> {
    return pulumi.output(args).apply(a => getDeployment(a, opts))
}

/**
 * A collection of arguments for invoking getDeployment.
 */
export interface GetDeploymentOutputArgs {
    /**
     * vRA Cloud template content. Conflicts with `blueprintId` and `catalogItemId`.
     */
    blueprintContent?: pulumi.Input<string>;
    /**
     * Identifier of the requested blueprint in the form ‘UUID:version’.
     */
    blueprintId?: pulumi.Input<string>;
    /**
     * The version of the vRA cloud template to request the deployment. Used only when `blueprintId` is provided.
     */
    blueprintVersion?: pulumi.Input<string>;
    /**
     * Identifier of the requested catalog item in the form ‘UUID:version’.
     */
    catalogItemId?: pulumi.Input<string>;
    /**
     * The version of the vRA catalog item to request the deployment. Used only when `catalogItemId` is provided.
     */
    catalogItemVersion?: pulumi.Input<string>;
    /**
     * Creation time (e.g. date format ‘2019-07-13T23:16:49.310Z’).
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The user the entity was created by.
     */
    createdBy?: pulumi.Input<string>;
    /**
     * A description of the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Flag to indicate whether to expand last request on the deployment.
     */
    expandLastRequest?: pulumi.Input<boolean>;
    /**
     * Flag to indicate whether to expand project information.
     */
    expandProject?: pulumi.Input<boolean>;
    /**
     * Flag to indicate whether to expand resources in the deployment.
     */
    expandResources?: pulumi.Input<boolean>;
    /**
     * The id of the deployment. One of `id` or `name` must be provided.
     */
    id?: pulumi.Input<string>;
    /**
     * List of request inputs.
     */
    inputs?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Time at which the deployment was last updated.
     */
    lastUpdatedAt?: pulumi.Input<string>;
    /**
     * The user that last updated the deployment.
     */
    lastUpdatedBy?: pulumi.Input<string>;
    /**
     * Time at which the deployment lease expires.
     */
    leaseExpireAt?: pulumi.Input<string>;
    /**
     * Name of the deployment. One of `id` or `name` must be provided.
     */
    name?: pulumi.Input<string>;
    /**
     * The user this deployment belongs to.
     */
    owner?: pulumi.Input<string>;
    /**
     * The id of the project this deployment belongs to.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Deployment status. Supported values are: `CREATE_SUCCESSFUL`, `CREATE_INPROGRESS`, `CREATE_FAILED`, `UPDATE_SUCCESSFUL`, `UPDATE_INPROGRESS`, `UPDATE_FAILED`, `DELETE_SUCCESSFUL`, `DELETE_INPROGRESS`, `DELETE_FAILED`, `ACTION_SUCCESSFUL`, `ACTION_INPROGRESS`, `ACTION_FAILED`.
     */
    status?: pulumi.Input<string>;
}