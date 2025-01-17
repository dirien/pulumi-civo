// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a Civo volume which can be attached to a Instance in order to provide expanded storage.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as civo from "@pulumi/civo";
 *
 * const db = new civo.Volume("db", {
 *     bootable: false,
 *     sizeGb: 60,
 * });
 * ```
 *
 * ## Import
 *
 * Volumes can be imported using the `volume id`, e.g.
 *
 * ```sh
 *  $ pulumi import civo:index/volume:Volume db 506f78a4-e098-11e5-ad9f-000f53306ae1
 * ```
 */
export class Volume extends pulumi.CustomResource {
    /**
     * Get an existing Volume resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VolumeState, opts?: pulumi.CustomResourceOptions): Volume {
        return new Volume(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'civo:index/volume:Volume';

    /**
     * Returns true if the given object is an instance of Volume.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Volume {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Volume.__pulumiType;
    }

    /**
     * Mark the volume as bootable.
     */
    public readonly bootable!: pulumi.Output<boolean>;
    /**
     * The date of the creation of the volume.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The mount point of the volume.
     */
    public /*out*/ readonly mountPoint!: pulumi.Output<string>;
    /**
     * A name that you wish to use to refer to this volume .
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region for the volume
     */
    public readonly region!: pulumi.Output<string | undefined>;
    /**
     * A minimum of 1 and a maximum of your available disk space from your quota specifies the size of the volume in gigabytes .
     */
    public readonly sizeGb!: pulumi.Output<number>;

    /**
     * Create a Volume resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VolumeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VolumeArgs | VolumeState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VolumeState | undefined;
            inputs["bootable"] = state ? state.bootable : undefined;
            inputs["createdAt"] = state ? state.createdAt : undefined;
            inputs["mountPoint"] = state ? state.mountPoint : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["sizeGb"] = state ? state.sizeGb : undefined;
        } else {
            const args = argsOrState as VolumeArgs | undefined;
            if ((!args || args.bootable === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bootable'");
            }
            if ((!args || args.sizeGb === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sizeGb'");
            }
            inputs["bootable"] = args ? args.bootable : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["sizeGb"] = args ? args.sizeGb : undefined;
            inputs["createdAt"] = undefined /*out*/;
            inputs["mountPoint"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Volume.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Volume resources.
 */
export interface VolumeState {
    /**
     * Mark the volume as bootable.
     */
    readonly bootable?: pulumi.Input<boolean>;
    /**
     * The date of the creation of the volume.
     */
    readonly createdAt?: pulumi.Input<string>;
    /**
     * The mount point of the volume.
     */
    readonly mountPoint?: pulumi.Input<string>;
    /**
     * A name that you wish to use to refer to this volume .
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region for the volume
     */
    readonly region?: pulumi.Input<string>;
    /**
     * A minimum of 1 and a maximum of your available disk space from your quota specifies the size of the volume in gigabytes .
     */
    readonly sizeGb?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Volume resource.
 */
export interface VolumeArgs {
    /**
     * Mark the volume as bootable.
     */
    readonly bootable: pulumi.Input<boolean>;
    /**
     * A name that you wish to use to refer to this volume .
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region for the volume
     */
    readonly region?: pulumi.Input<string>;
    /**
     * A minimum of 1 and a maximum of your available disk space from your quota specifies the size of the volume in gigabytes .
     */
    readonly sizeGb: pulumi.Input<number>;
}
