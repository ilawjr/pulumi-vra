# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['StorageProfileArgs', 'StorageProfile']

@pulumi.input_type
class StorageProfileArgs:
    def __init__(__self__, *,
                 default_item: pulumi.Input[bool],
                 region_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 disk_properties: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 disk_target_properties: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 supports_encryption: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['StorageProfileTagArgs']]]] = None):
        """
        The set of arguments for constructing a StorageProfile resource.
        :param pulumi.Input[bool] default_item: Indicates if this storage profile is a default profile.
        :param pulumi.Input[str] region_id: The id of the region for which this profile is defined as in vRealize Automation(vRA).
        :param pulumi.Input[str] description: A human-friendly description.
        :param pulumi.Input[Mapping[str, Any]] disk_properties: Map of storage properties that are to be applied on disk while provisioning.
        :param pulumi.Input[Mapping[str, Any]] disk_target_properties: Map of storage placements to know where the disk is provisioned.
        :param pulumi.Input[str] name: A human-friendly name used as an identifier in APIs that support this option.
        :param pulumi.Input[bool] supports_encryption: Indicates whether this storage profile supports encryption or not.
        :param pulumi.Input[Sequence[pulumi.Input['StorageProfileTagArgs']]] tags: A set of tag keys and optional values that were set on this Network Profile.
               example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        """
        pulumi.set(__self__, "default_item", default_item)
        pulumi.set(__self__, "region_id", region_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disk_properties is not None:
            pulumi.set(__self__, "disk_properties", disk_properties)
        if disk_target_properties is not None:
            pulumi.set(__self__, "disk_target_properties", disk_target_properties)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if supports_encryption is not None:
            pulumi.set(__self__, "supports_encryption", supports_encryption)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="defaultItem")
    def default_item(self) -> pulumi.Input[bool]:
        """
        Indicates if this storage profile is a default profile.
        """
        return pulumi.get(self, "default_item")

    @default_item.setter
    def default_item(self, value: pulumi.Input[bool]):
        pulumi.set(self, "default_item", value)

    @property
    @pulumi.getter(name="regionId")
    def region_id(self) -> pulumi.Input[str]:
        """
        The id of the region for which this profile is defined as in vRealize Automation(vRA).
        """
        return pulumi.get(self, "region_id")

    @region_id.setter
    def region_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "region_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A human-friendly description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="diskProperties")
    def disk_properties(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Map of storage properties that are to be applied on disk while provisioning.
        """
        return pulumi.get(self, "disk_properties")

    @disk_properties.setter
    def disk_properties(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "disk_properties", value)

    @property
    @pulumi.getter(name="diskTargetProperties")
    def disk_target_properties(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Map of storage placements to know where the disk is provisioned.
        """
        return pulumi.get(self, "disk_target_properties")

    @disk_target_properties.setter
    def disk_target_properties(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "disk_target_properties", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A human-friendly name used as an identifier in APIs that support this option.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="supportsEncryption")
    def supports_encryption(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether this storage profile supports encryption or not.
        """
        return pulumi.get(self, "supports_encryption")

    @supports_encryption.setter
    def supports_encryption(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "supports_encryption", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['StorageProfileTagArgs']]]]:
        """
        A set of tag keys and optional values that were set on this Network Profile.
        example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['StorageProfileTagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _StorageProfileState:
    def __init__(__self__, *,
                 cloud_account_id: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 default_item: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disk_properties: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 disk_target_properties: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 external_region_id: Optional[pulumi.Input[str]] = None,
                 links: Optional[pulumi.Input[Sequence[pulumi.Input['StorageProfileLinkArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 owner: Optional[pulumi.Input[str]] = None,
                 region_id: Optional[pulumi.Input[str]] = None,
                 supports_encryption: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['StorageProfileTagArgs']]]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering StorageProfile resources.
        :param pulumi.Input[str] cloud_account_id: Id of the cloud account this storage profile belongs to.
        :param pulumi.Input[str] created_at: Date when the entity was created. The date is in ISO 6801 and UTC.
        :param pulumi.Input[bool] default_item: Indicates if this storage profile is a default profile.
        :param pulumi.Input[str] description: A human-friendly description.
        :param pulumi.Input[Mapping[str, Any]] disk_properties: Map of storage properties that are to be applied on disk while provisioning.
        :param pulumi.Input[Mapping[str, Any]] disk_target_properties: Map of storage placements to know where the disk is provisioned.
        :param pulumi.Input[str] external_region_id: The id of the region as seen in the cloud provider for which this profile is defined.
        :param pulumi.Input[Sequence[pulumi.Input['StorageProfileLinkArgs']]] links: HATEOAS of the entity
        :param pulumi.Input[str] name: A human-friendly name used as an identifier in APIs that support this option.
        :param pulumi.Input[str] org_id: The id of the organization this entity belongs to.
        :param pulumi.Input[str] owner: Email of the user that owns the entity.
        :param pulumi.Input[str] region_id: The id of the region for which this profile is defined as in vRealize Automation(vRA).
        :param pulumi.Input[bool] supports_encryption: Indicates whether this storage profile supports encryption or not.
        :param pulumi.Input[Sequence[pulumi.Input['StorageProfileTagArgs']]] tags: A set of tag keys and optional values that were set on this Network Profile.
               example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        :param pulumi.Input[str] updated_at: Date when the entity was last updated. The date is ISO 8601 and UTC.
        """
        if cloud_account_id is not None:
            pulumi.set(__self__, "cloud_account_id", cloud_account_id)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if default_item is not None:
            pulumi.set(__self__, "default_item", default_item)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disk_properties is not None:
            pulumi.set(__self__, "disk_properties", disk_properties)
        if disk_target_properties is not None:
            pulumi.set(__self__, "disk_target_properties", disk_target_properties)
        if external_region_id is not None:
            pulumi.set(__self__, "external_region_id", external_region_id)
        if links is not None:
            pulumi.set(__self__, "links", links)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if owner is not None:
            pulumi.set(__self__, "owner", owner)
        if region_id is not None:
            pulumi.set(__self__, "region_id", region_id)
        if supports_encryption is not None:
            pulumi.set(__self__, "supports_encryption", supports_encryption)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="cloudAccountId")
    def cloud_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        Id of the cloud account this storage profile belongs to.
        """
        return pulumi.get(self, "cloud_account_id")

    @cloud_account_id.setter
    def cloud_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cloud_account_id", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        Date when the entity was created. The date is in ISO 6801 and UTC.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="defaultItem")
    def default_item(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates if this storage profile is a default profile.
        """
        return pulumi.get(self, "default_item")

    @default_item.setter
    def default_item(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "default_item", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A human-friendly description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="diskProperties")
    def disk_properties(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Map of storage properties that are to be applied on disk while provisioning.
        """
        return pulumi.get(self, "disk_properties")

    @disk_properties.setter
    def disk_properties(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "disk_properties", value)

    @property
    @pulumi.getter(name="diskTargetProperties")
    def disk_target_properties(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Map of storage placements to know where the disk is provisioned.
        """
        return pulumi.get(self, "disk_target_properties")

    @disk_target_properties.setter
    def disk_target_properties(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "disk_target_properties", value)

    @property
    @pulumi.getter(name="externalRegionId")
    def external_region_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the region as seen in the cloud provider for which this profile is defined.
        """
        return pulumi.get(self, "external_region_id")

    @external_region_id.setter
    def external_region_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "external_region_id", value)

    @property
    @pulumi.getter
    def links(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['StorageProfileLinkArgs']]]]:
        """
        HATEOAS of the entity
        """
        return pulumi.get(self, "links")

    @links.setter
    def links(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['StorageProfileLinkArgs']]]]):
        pulumi.set(self, "links", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A human-friendly name used as an identifier in APIs that support this option.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the organization this entity belongs to.
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter
    def owner(self) -> Optional[pulumi.Input[str]]:
        """
        Email of the user that owns the entity.
        """
        return pulumi.get(self, "owner")

    @owner.setter
    def owner(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner", value)

    @property
    @pulumi.getter(name="regionId")
    def region_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the region for which this profile is defined as in vRealize Automation(vRA).
        """
        return pulumi.get(self, "region_id")

    @region_id.setter
    def region_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region_id", value)

    @property
    @pulumi.getter(name="supportsEncryption")
    def supports_encryption(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether this storage profile supports encryption or not.
        """
        return pulumi.get(self, "supports_encryption")

    @supports_encryption.setter
    def supports_encryption(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "supports_encryption", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['StorageProfileTagArgs']]]]:
        """
        A set of tag keys and optional values that were set on this Network Profile.
        example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['StorageProfileTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        Date when the entity was last updated. The date is ISO 8601 and UTC.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class StorageProfile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_item: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disk_properties: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 disk_target_properties: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region_id: Optional[pulumi.Input[str]] = None,
                 supports_encryption: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StorageProfileTagArgs']]]]] = None,
                 __props__=None):
        """
        ## Example Usage
        ### S
        This is an example of how to create a storage profile resource.

        **Vra storage profile:**

        ```python
        import pulumi
        import pulumiverse_vra as vra

        # vSphere storage profile using generic vra_storage_profile resource.
        this_storage_profile = vra.storageprofile.StorageProfile("thisStorageProfile",
            description="vSphere Storage Profile with standard independent non-persistent disk.",
            region_id=data["vra_region"]["this"]["id"],
            default_item=False,
            disk_properties={
                "independent": "true",
                "persistent": "false",
                "limitIops": "2000",
                "provisioningType": "eagerZeroedThick",
                "sharesLevel": "custom",
                "shares": "1500",
            },
            disk_target_properties={
                "datastoreId": data["vra_fabric_datastore_vsphere"]["this"]["id"],
                "storagePolicyId": data["vra_fabric_storage_policy_vsphere"]["this"]["id"],
            },
            tags=[vra.storageprofile.StorageProfileTagArgs(
                key="foo",
                value="bar",
            )])
        # AWS storage profile using generic vra_storage_profile resource.
        this_storageprofile_storage_profile_storage_profile = vra.storageprofile.StorageProfile("thisStorageprofile/storageProfileStorageProfile",
            description="AWS Storage Profile with instance store device type.",
            region_id=data["vra_region"]["this"]["id"],
            default_item=False,
            disk_properties={
                "deviceType": "instance-store",
            },
            tags=[vra.storageprofile.StorageProfileTagArgs(
                key="foo",
                value="bar",
            )])
        # Azure storage profile using generic vra_storage_profile resource.
        this_vra_storageprofile_storage_profile_storage_profile = vra.storageprofile.StorageProfile("thisVraStorageprofile/storageProfileStorageProfile",
            description="Azure Storage Profile with managed disks.",
            region_id=data["vra_region"]["this"]["id"],
            default_item=False,
            supports_encryption=False,
            disk_properties={
                "azureDataDiskCaching": "None",
                "azureManagedDiskType": "Standard_LRS",
                "azureOsDiskCaching": "None",
            },
            tags=[vra.storageprofile.StorageProfileTagArgs(
                key="foo",
                value="bar",
            )])
        ```

        A storage profile resource supports the following arguments:

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] default_item: Indicates if this storage profile is a default profile.
        :param pulumi.Input[str] description: A human-friendly description.
        :param pulumi.Input[Mapping[str, Any]] disk_properties: Map of storage properties that are to be applied on disk while provisioning.
        :param pulumi.Input[Mapping[str, Any]] disk_target_properties: Map of storage placements to know where the disk is provisioned.
        :param pulumi.Input[str] name: A human-friendly name used as an identifier in APIs that support this option.
        :param pulumi.Input[str] region_id: The id of the region for which this profile is defined as in vRealize Automation(vRA).
        :param pulumi.Input[bool] supports_encryption: Indicates whether this storage profile supports encryption or not.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StorageProfileTagArgs']]]] tags: A set of tag keys and optional values that were set on this Network Profile.
               example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StorageProfileArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage
        ### S
        This is an example of how to create a storage profile resource.

        **Vra storage profile:**

        ```python
        import pulumi
        import pulumiverse_vra as vra

        # vSphere storage profile using generic vra_storage_profile resource.
        this_storage_profile = vra.storageprofile.StorageProfile("thisStorageProfile",
            description="vSphere Storage Profile with standard independent non-persistent disk.",
            region_id=data["vra_region"]["this"]["id"],
            default_item=False,
            disk_properties={
                "independent": "true",
                "persistent": "false",
                "limitIops": "2000",
                "provisioningType": "eagerZeroedThick",
                "sharesLevel": "custom",
                "shares": "1500",
            },
            disk_target_properties={
                "datastoreId": data["vra_fabric_datastore_vsphere"]["this"]["id"],
                "storagePolicyId": data["vra_fabric_storage_policy_vsphere"]["this"]["id"],
            },
            tags=[vra.storageprofile.StorageProfileTagArgs(
                key="foo",
                value="bar",
            )])
        # AWS storage profile using generic vra_storage_profile resource.
        this_storageprofile_storage_profile_storage_profile = vra.storageprofile.StorageProfile("thisStorageprofile/storageProfileStorageProfile",
            description="AWS Storage Profile with instance store device type.",
            region_id=data["vra_region"]["this"]["id"],
            default_item=False,
            disk_properties={
                "deviceType": "instance-store",
            },
            tags=[vra.storageprofile.StorageProfileTagArgs(
                key="foo",
                value="bar",
            )])
        # Azure storage profile using generic vra_storage_profile resource.
        this_vra_storageprofile_storage_profile_storage_profile = vra.storageprofile.StorageProfile("thisVraStorageprofile/storageProfileStorageProfile",
            description="Azure Storage Profile with managed disks.",
            region_id=data["vra_region"]["this"]["id"],
            default_item=False,
            supports_encryption=False,
            disk_properties={
                "azureDataDiskCaching": "None",
                "azureManagedDiskType": "Standard_LRS",
                "azureOsDiskCaching": "None",
            },
            tags=[vra.storageprofile.StorageProfileTagArgs(
                key="foo",
                value="bar",
            )])
        ```

        A storage profile resource supports the following arguments:

        :param str resource_name: The name of the resource.
        :param StorageProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StorageProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_item: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disk_properties: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 disk_target_properties: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region_id: Optional[pulumi.Input[str]] = None,
                 supports_encryption: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StorageProfileTagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StorageProfileArgs.__new__(StorageProfileArgs)

            if default_item is None and not opts.urn:
                raise TypeError("Missing required property 'default_item'")
            __props__.__dict__["default_item"] = default_item
            __props__.__dict__["description"] = description
            __props__.__dict__["disk_properties"] = disk_properties
            __props__.__dict__["disk_target_properties"] = disk_target_properties
            __props__.__dict__["name"] = name
            if region_id is None and not opts.urn:
                raise TypeError("Missing required property 'region_id'")
            __props__.__dict__["region_id"] = region_id
            __props__.__dict__["supports_encryption"] = supports_encryption
            __props__.__dict__["tags"] = tags
            __props__.__dict__["cloud_account_id"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["external_region_id"] = None
            __props__.__dict__["links"] = None
            __props__.__dict__["org_id"] = None
            __props__.__dict__["owner"] = None
            __props__.__dict__["updated_at"] = None
        super(StorageProfile, __self__).__init__(
            'vra:storageprofile/storageProfile:StorageProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cloud_account_id: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            default_item: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            disk_properties: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            disk_target_properties: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            external_region_id: Optional[pulumi.Input[str]] = None,
            links: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StorageProfileLinkArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            owner: Optional[pulumi.Input[str]] = None,
            region_id: Optional[pulumi.Input[str]] = None,
            supports_encryption: Optional[pulumi.Input[bool]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StorageProfileTagArgs']]]]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'StorageProfile':
        """
        Get an existing StorageProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cloud_account_id: Id of the cloud account this storage profile belongs to.
        :param pulumi.Input[str] created_at: Date when the entity was created. The date is in ISO 6801 and UTC.
        :param pulumi.Input[bool] default_item: Indicates if this storage profile is a default profile.
        :param pulumi.Input[str] description: A human-friendly description.
        :param pulumi.Input[Mapping[str, Any]] disk_properties: Map of storage properties that are to be applied on disk while provisioning.
        :param pulumi.Input[Mapping[str, Any]] disk_target_properties: Map of storage placements to know where the disk is provisioned.
        :param pulumi.Input[str] external_region_id: The id of the region as seen in the cloud provider for which this profile is defined.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StorageProfileLinkArgs']]]] links: HATEOAS of the entity
        :param pulumi.Input[str] name: A human-friendly name used as an identifier in APIs that support this option.
        :param pulumi.Input[str] org_id: The id of the organization this entity belongs to.
        :param pulumi.Input[str] owner: Email of the user that owns the entity.
        :param pulumi.Input[str] region_id: The id of the region for which this profile is defined as in vRealize Automation(vRA).
        :param pulumi.Input[bool] supports_encryption: Indicates whether this storage profile supports encryption or not.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StorageProfileTagArgs']]]] tags: A set of tag keys and optional values that were set on this Network Profile.
               example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        :param pulumi.Input[str] updated_at: Date when the entity was last updated. The date is ISO 8601 and UTC.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StorageProfileState.__new__(_StorageProfileState)

        __props__.__dict__["cloud_account_id"] = cloud_account_id
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["default_item"] = default_item
        __props__.__dict__["description"] = description
        __props__.__dict__["disk_properties"] = disk_properties
        __props__.__dict__["disk_target_properties"] = disk_target_properties
        __props__.__dict__["external_region_id"] = external_region_id
        __props__.__dict__["links"] = links
        __props__.__dict__["name"] = name
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["owner"] = owner
        __props__.__dict__["region_id"] = region_id
        __props__.__dict__["supports_encryption"] = supports_encryption
        __props__.__dict__["tags"] = tags
        __props__.__dict__["updated_at"] = updated_at
        return StorageProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cloudAccountId")
    def cloud_account_id(self) -> pulumi.Output[str]:
        """
        Id of the cloud account this storage profile belongs to.
        """
        return pulumi.get(self, "cloud_account_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Date when the entity was created. The date is in ISO 6801 and UTC.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="defaultItem")
    def default_item(self) -> pulumi.Output[bool]:
        """
        Indicates if this storage profile is a default profile.
        """
        return pulumi.get(self, "default_item")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A human-friendly description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="diskProperties")
    def disk_properties(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Map of storage properties that are to be applied on disk while provisioning.
        """
        return pulumi.get(self, "disk_properties")

    @property
    @pulumi.getter(name="diskTargetProperties")
    def disk_target_properties(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Map of storage placements to know where the disk is provisioned.
        """
        return pulumi.get(self, "disk_target_properties")

    @property
    @pulumi.getter(name="externalRegionId")
    def external_region_id(self) -> pulumi.Output[str]:
        """
        The id of the region as seen in the cloud provider for which this profile is defined.
        """
        return pulumi.get(self, "external_region_id")

    @property
    @pulumi.getter
    def links(self) -> pulumi.Output[Sequence['outputs.StorageProfileLink']]:
        """
        HATEOAS of the entity
        """
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A human-friendly name used as an identifier in APIs that support this option.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[str]:
        """
        The id of the organization this entity belongs to.
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter
    def owner(self) -> pulumi.Output[str]:
        """
        Email of the user that owns the entity.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter(name="regionId")
    def region_id(self) -> pulumi.Output[str]:
        """
        The id of the region for which this profile is defined as in vRealize Automation(vRA).
        """
        return pulumi.get(self, "region_id")

    @property
    @pulumi.getter(name="supportsEncryption")
    def supports_encryption(self) -> pulumi.Output[bool]:
        """
        Indicates whether this storage profile supports encryption or not.
        """
        return pulumi.get(self, "supports_encryption")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Sequence['outputs.StorageProfileTag']]:
        """
        A set of tag keys and optional values that were set on this Network Profile.
        example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        Date when the entity was last updated. The date is ISO 8601 and UTC.
        """
        return pulumi.get(self, "updated_at")
