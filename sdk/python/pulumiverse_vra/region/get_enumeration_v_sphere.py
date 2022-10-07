# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetEnumerationVSphereResult',
    'AwaitableGetEnumerationVSphereResult',
    'get_enumeration_v_sphere',
    'get_enumeration_v_sphere_output',
]

@pulumi.output_type
class GetEnumerationVSphereResult:
    """
    A collection of values returned by getEnumerationVSphere.
    """
    def __init__(__self__, accept_self_signed_cert=None, dcid=None, hostname=None, id=None, password=None, regions=None, username=None):
        if accept_self_signed_cert and not isinstance(accept_self_signed_cert, bool):
            raise TypeError("Expected argument 'accept_self_signed_cert' to be a bool")
        pulumi.set(__self__, "accept_self_signed_cert", accept_self_signed_cert)
        if dcid and not isinstance(dcid, str):
            raise TypeError("Expected argument 'dcid' to be a str")
        pulumi.set(__self__, "dcid", dcid)
        if hostname and not isinstance(hostname, str):
            raise TypeError("Expected argument 'hostname' to be a str")
        pulumi.set(__self__, "hostname", hostname)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if password and not isinstance(password, str):
            raise TypeError("Expected argument 'password' to be a str")
        pulumi.set(__self__, "password", password)
        if regions and not isinstance(regions, list):
            raise TypeError("Expected argument 'regions' to be a list")
        pulumi.set(__self__, "regions", regions)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="acceptSelfSignedCert")
    def accept_self_signed_cert(self) -> Optional[bool]:
        return pulumi.get(self, "accept_self_signed_cert")

    @property
    @pulumi.getter
    def dcid(self) -> Optional[str]:
        return pulumi.get(self, "dcid")

    @property
    @pulumi.getter
    def hostname(self) -> str:
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def password(self) -> str:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def regions(self) -> Sequence[str]:
        """
        A set of datacenter managed object reference identifiers to enable provisioning on. Example: Datacenter:datacenter-2
        """
        return pulumi.get(self, "regions")

    @property
    @pulumi.getter
    def username(self) -> str:
        return pulumi.get(self, "username")


class AwaitableGetEnumerationVSphereResult(GetEnumerationVSphereResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEnumerationVSphereResult(
            accept_self_signed_cert=self.accept_self_signed_cert,
            dcid=self.dcid,
            hostname=self.hostname,
            id=self.id,
            password=self.password,
            regions=self.regions,
            username=self.username)


def get_enumeration_v_sphere(accept_self_signed_cert: Optional[bool] = None,
                             dcid: Optional[str] = None,
                             hostname: Optional[str] = None,
                             password: Optional[str] = None,
                             username: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEnumerationVSphereResult:
    """
    ## Example Usage


    :param bool accept_self_signed_cert: Accept self signed certificate when connecting to vSphere. Example: false
    :param str dcid: ID of a data collector vm deployed in the on premise infrastructure. Example: d5316b00-f3b8-4895-9e9a-c4b98649c2ca
    :param str hostname: Host name for the cloud account endpoint. Example: dc1-lnd.mycompany.com
    :param str password: Password for the user used to authenticate with the cloud Account
    :param str username: Username to authenticate with the cloud account
    """
    __args__ = dict()
    __args__['acceptSelfSignedCert'] = accept_self_signed_cert
    __args__['dcid'] = dcid
    __args__['hostname'] = hostname
    __args__['password'] = password
    __args__['username'] = username
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vra:region/getEnumerationVSphere:getEnumerationVSphere', __args__, opts=opts, typ=GetEnumerationVSphereResult).value

    return AwaitableGetEnumerationVSphereResult(
        accept_self_signed_cert=__ret__.accept_self_signed_cert,
        dcid=__ret__.dcid,
        hostname=__ret__.hostname,
        id=__ret__.id,
        password=__ret__.password,
        regions=__ret__.regions,
        username=__ret__.username)


@_utilities.lift_output_func(get_enumeration_v_sphere)
def get_enumeration_v_sphere_output(accept_self_signed_cert: Optional[pulumi.Input[Optional[bool]]] = None,
                                    dcid: Optional[pulumi.Input[Optional[str]]] = None,
                                    hostname: Optional[pulumi.Input[str]] = None,
                                    password: Optional[pulumi.Input[str]] = None,
                                    username: Optional[pulumi.Input[str]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEnumerationVSphereResult]:
    """
    ## Example Usage


    :param bool accept_self_signed_cert: Accept self signed certificate when connecting to vSphere. Example: false
    :param str dcid: ID of a data collector vm deployed in the on premise infrastructure. Example: d5316b00-f3b8-4895-9e9a-c4b98649c2ca
    :param str hostname: Host name for the cloud account endpoint. Example: dc1-lnd.mycompany.com
    :param str password: Password for the user used to authenticate with the cloud Account
    :param str username: Username to authenticate with the cloud account
    """
    ...