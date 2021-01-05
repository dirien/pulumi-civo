# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['Template']


class Template(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cloud_config: Optional[pulumi.Input[str]] = None,
                 code: Optional[pulumi.Input[str]] = None,
                 default_username: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 short_description: Optional[pulumi.Input[str]] = None,
                 volume_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Civo Template resource.
        This can be used to create, modify, and delete Templates.

        ## Import

        Template can be imported using the template `code`, e.g.

        ```sh
         $ pulumi import civo:index/template:Template my-custom-template my-template-code
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cloud_config: Commonly referred to as 'user-data', this is a customisation script that is run after
               the instance is first booted. We recommend using cloud-config as it's a great distribution-agnostic
               way of configuring cloud servers. If you put `$INITIAL_USER` in your script, this will automatically
               be replaced by the initial user chosen when creating the instance, `$INITIAL_PASSWORD` will be
               replaced with the random password generated by the system, `$HOSTNAME` is the fully qualified
               domain name of the instance and `$SSH_KEY` will be the content of the SSH public key.
               (this is technically optional, but you won't really be able to use instances without it -
               see our learn guide on templates for more information).
        :param pulumi.Input[str] code: This is a unqiue, alphanumerical, short, human readable code for the template.
        :param pulumi.Input[str] default_username: The default username to suggest that the user creates
        :param pulumi.Input[str] description: A multi-line description of the template, in Markdown format
        :param pulumi.Input[str] image_id: This is the Image ID of any default template or the ID of another template
               either owned by you or global (optional; but must be specified if no volume_id is specified).
        :param pulumi.Input[str] name: This is a short human readable name for the template
        :param pulumi.Input[str] short_description: A one line description of the template
        :param pulumi.Input[str] volume_id: This is the ID of a bootable volume, either owned by you or global
               (optional; but must be specified if no image_id is specified)
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['cloud_config'] = cloud_config
            if code is None:
                raise TypeError("Missing required property 'code'")
            __props__['code'] = code
            __props__['default_username'] = default_username
            __props__['description'] = description
            __props__['image_id'] = image_id
            __props__['name'] = name
            __props__['short_description'] = short_description
            __props__['volume_id'] = volume_id
        super(Template, __self__).__init__(
            'civo:index/template:Template',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cloud_config: Optional[pulumi.Input[str]] = None,
            code: Optional[pulumi.Input[str]] = None,
            default_username: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            image_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            short_description: Optional[pulumi.Input[str]] = None,
            volume_id: Optional[pulumi.Input[str]] = None) -> 'Template':
        """
        Get an existing Template resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cloud_config: Commonly referred to as 'user-data', this is a customisation script that is run after
               the instance is first booted. We recommend using cloud-config as it's a great distribution-agnostic
               way of configuring cloud servers. If you put `$INITIAL_USER` in your script, this will automatically
               be replaced by the initial user chosen when creating the instance, `$INITIAL_PASSWORD` will be
               replaced with the random password generated by the system, `$HOSTNAME` is the fully qualified
               domain name of the instance and `$SSH_KEY` will be the content of the SSH public key.
               (this is technically optional, but you won't really be able to use instances without it -
               see our learn guide on templates for more information).
        :param pulumi.Input[str] code: This is a unqiue, alphanumerical, short, human readable code for the template.
        :param pulumi.Input[str] default_username: The default username to suggest that the user creates
        :param pulumi.Input[str] description: A multi-line description of the template, in Markdown format
        :param pulumi.Input[str] image_id: This is the Image ID of any default template or the ID of another template
               either owned by you or global (optional; but must be specified if no volume_id is specified).
        :param pulumi.Input[str] name: This is a short human readable name for the template
        :param pulumi.Input[str] short_description: A one line description of the template
        :param pulumi.Input[str] volume_id: This is the ID of a bootable volume, either owned by you or global
               (optional; but must be specified if no image_id is specified)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cloud_config"] = cloud_config
        __props__["code"] = code
        __props__["default_username"] = default_username
        __props__["description"] = description
        __props__["image_id"] = image_id
        __props__["name"] = name
        __props__["short_description"] = short_description
        __props__["volume_id"] = volume_id
        return Template(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cloudConfig")
    def cloud_config(self) -> pulumi.Output[Optional[str]]:
        """
        Commonly referred to as 'user-data', this is a customisation script that is run after
        the instance is first booted. We recommend using cloud-config as it's a great distribution-agnostic
        way of configuring cloud servers. If you put `$INITIAL_USER` in your script, this will automatically
        be replaced by the initial user chosen when creating the instance, `$INITIAL_PASSWORD` will be
        replaced with the random password generated by the system, `$HOSTNAME` is the fully qualified
        domain name of the instance and `$SSH_KEY` will be the content of the SSH public key.
        (this is technically optional, but you won't really be able to use instances without it -
        see our learn guide on templates for more information).
        """
        return pulumi.get(self, "cloud_config")

    @property
    @pulumi.getter
    def code(self) -> pulumi.Output[str]:
        """
        This is a unqiue, alphanumerical, short, human readable code for the template.
        """
        return pulumi.get(self, "code")

    @property
    @pulumi.getter(name="defaultUsername")
    def default_username(self) -> pulumi.Output[Optional[str]]:
        """
        The default username to suggest that the user creates
        """
        return pulumi.get(self, "default_username")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A multi-line description of the template, in Markdown format
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> pulumi.Output[Optional[str]]:
        """
        This is the Image ID of any default template or the ID of another template
        either owned by you or global (optional; but must be specified if no volume_id is specified).
        """
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        This is a short human readable name for the template
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="shortDescription")
    def short_description(self) -> pulumi.Output[Optional[str]]:
        """
        A one line description of the template
        """
        return pulumi.get(self, "short_description")

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> pulumi.Output[Optional[str]]:
        """
        This is the ID of a bootable volume, either owned by you or global
        (optional; but must be specified if no image_id is specified)
        """
        return pulumi.get(self, "volume_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

