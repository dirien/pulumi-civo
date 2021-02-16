// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class LoadBalancer extends pulumi.CustomResource {
    /**
     * Get an existing LoadBalancer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoadBalancerState, opts?: pulumi.CustomResourceOptions): LoadBalancer {
        return new LoadBalancer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'civo:index/loadBalancer:LoadBalancer';

    /**
     * Returns true if the given object is an instance of LoadBalancer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoadBalancer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoadBalancer.__pulumiType;
    }

    /**
     * a list of backend instances, each containing an instance_id, protocol (http or https) and port
     */
    public readonly backends!: pulumi.Output<outputs.LoadBalancerBackend[]>;
    /**
     * how long to wait in seconds before determining a backend has failed, defaults to 30
     */
    public readonly failTimeout!: pulumi.Output<number>;
    /**
     * what URL should be used on the backends to determine if it's OK (2xx/3xx status), defaults to /
     */
    public readonly healthCheckPath!: pulumi.Output<string | undefined>;
    /**
     * the hostname to receive traffic for, e.g. www.example.com (optional: sets hostname to loadbalancer-uuid.civo.com if
     * blank)
     */
    public readonly hostname!: pulumi.Output<string>;
    /**
     * should self-signed/invalid certificates be ignored from the backend servers, defaults to true
     */
    public readonly ignoreInvalidBackendTls!: pulumi.Output<boolean | undefined>;
    /**
     * how many concurrent connections can each backend handle, defaults to 10
     */
    public readonly maxConns!: pulumi.Output<number>;
    /**
     * the size in megabytes of the maximum request content that will be accepted, defaults to 20
     */
    public readonly maxRequestSize!: pulumi.Output<number>;
    /**
     * one of: least_conn (sends new requests to the least busy server), random (sends new requests to a random backend),
     * round_robin (sends new requests to the next backend in order), ip_hash (sends requests from a given IP address to the
     * same backend), default is random
     */
    public readonly policy!: pulumi.Output<string>;
    /**
     * you can listen on any port, the default is 80 to match the default protocol of http,if not you must specify it here
     * (commonly 80 for HTTP or 443 for HTTPS)
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * either http or https. If you specify https then you must also provide the next two fields, the default is http
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * if your protocol is https then you should send the TLS certificate in Base64-encoded PEM format
     */
    public readonly tlsCertificate!: pulumi.Output<string | undefined>;
    /**
     * if your protocol is https then you should send the TLS private key in Base64-encoded PEM format
     */
    public readonly tlsKey!: pulumi.Output<string | undefined>;

    /**
     * Create a LoadBalancer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LoadBalancerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LoadBalancerArgs | LoadBalancerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LoadBalancerState | undefined;
            inputs["backends"] = state ? state.backends : undefined;
            inputs["failTimeout"] = state ? state.failTimeout : undefined;
            inputs["healthCheckPath"] = state ? state.healthCheckPath : undefined;
            inputs["hostname"] = state ? state.hostname : undefined;
            inputs["ignoreInvalidBackendTls"] = state ? state.ignoreInvalidBackendTls : undefined;
            inputs["maxConns"] = state ? state.maxConns : undefined;
            inputs["maxRequestSize"] = state ? state.maxRequestSize : undefined;
            inputs["policy"] = state ? state.policy : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["tlsCertificate"] = state ? state.tlsCertificate : undefined;
            inputs["tlsKey"] = state ? state.tlsKey : undefined;
        } else {
            const args = argsOrState as LoadBalancerArgs | undefined;
            if ((!args || args.backends === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backends'");
            }
            if ((!args || args.failTimeout === undefined) && !opts.urn) {
                throw new Error("Missing required property 'failTimeout'");
            }
            if ((!args || args.hostname === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostname'");
            }
            if ((!args || args.maxConns === undefined) && !opts.urn) {
                throw new Error("Missing required property 'maxConns'");
            }
            if ((!args || args.maxRequestSize === undefined) && !opts.urn) {
                throw new Error("Missing required property 'maxRequestSize'");
            }
            if ((!args || args.policy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policy'");
            }
            if ((!args || args.port === undefined) && !opts.urn) {
                throw new Error("Missing required property 'port'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            inputs["backends"] = args ? args.backends : undefined;
            inputs["failTimeout"] = args ? args.failTimeout : undefined;
            inputs["healthCheckPath"] = args ? args.healthCheckPath : undefined;
            inputs["hostname"] = args ? args.hostname : undefined;
            inputs["ignoreInvalidBackendTls"] = args ? args.ignoreInvalidBackendTls : undefined;
            inputs["maxConns"] = args ? args.maxConns : undefined;
            inputs["maxRequestSize"] = args ? args.maxRequestSize : undefined;
            inputs["policy"] = args ? args.policy : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["tlsCertificate"] = args ? args.tlsCertificate : undefined;
            inputs["tlsKey"] = args ? args.tlsKey : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(LoadBalancer.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LoadBalancer resources.
 */
export interface LoadBalancerState {
    /**
     * a list of backend instances, each containing an instance_id, protocol (http or https) and port
     */
    readonly backends?: pulumi.Input<pulumi.Input<inputs.LoadBalancerBackend>[]>;
    /**
     * how long to wait in seconds before determining a backend has failed, defaults to 30
     */
    readonly failTimeout?: pulumi.Input<number>;
    /**
     * what URL should be used on the backends to determine if it's OK (2xx/3xx status), defaults to /
     */
    readonly healthCheckPath?: pulumi.Input<string>;
    /**
     * the hostname to receive traffic for, e.g. www.example.com (optional: sets hostname to loadbalancer-uuid.civo.com if
     * blank)
     */
    readonly hostname?: pulumi.Input<string>;
    /**
     * should self-signed/invalid certificates be ignored from the backend servers, defaults to true
     */
    readonly ignoreInvalidBackendTls?: pulumi.Input<boolean>;
    /**
     * how many concurrent connections can each backend handle, defaults to 10
     */
    readonly maxConns?: pulumi.Input<number>;
    /**
     * the size in megabytes of the maximum request content that will be accepted, defaults to 20
     */
    readonly maxRequestSize?: pulumi.Input<number>;
    /**
     * one of: least_conn (sends new requests to the least busy server), random (sends new requests to a random backend),
     * round_robin (sends new requests to the next backend in order), ip_hash (sends requests from a given IP address to the
     * same backend), default is random
     */
    readonly policy?: pulumi.Input<string>;
    /**
     * you can listen on any port, the default is 80 to match the default protocol of http,if not you must specify it here
     * (commonly 80 for HTTP or 443 for HTTPS)
     */
    readonly port?: pulumi.Input<number>;
    /**
     * either http or https. If you specify https then you must also provide the next two fields, the default is http
     */
    readonly protocol?: pulumi.Input<string>;
    /**
     * if your protocol is https then you should send the TLS certificate in Base64-encoded PEM format
     */
    readonly tlsCertificate?: pulumi.Input<string>;
    /**
     * if your protocol is https then you should send the TLS private key in Base64-encoded PEM format
     */
    readonly tlsKey?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LoadBalancer resource.
 */
export interface LoadBalancerArgs {
    /**
     * a list of backend instances, each containing an instance_id, protocol (http or https) and port
     */
    readonly backends: pulumi.Input<pulumi.Input<inputs.LoadBalancerBackend>[]>;
    /**
     * how long to wait in seconds before determining a backend has failed, defaults to 30
     */
    readonly failTimeout: pulumi.Input<number>;
    /**
     * what URL should be used on the backends to determine if it's OK (2xx/3xx status), defaults to /
     */
    readonly healthCheckPath?: pulumi.Input<string>;
    /**
     * the hostname to receive traffic for, e.g. www.example.com (optional: sets hostname to loadbalancer-uuid.civo.com if
     * blank)
     */
    readonly hostname: pulumi.Input<string>;
    /**
     * should self-signed/invalid certificates be ignored from the backend servers, defaults to true
     */
    readonly ignoreInvalidBackendTls?: pulumi.Input<boolean>;
    /**
     * how many concurrent connections can each backend handle, defaults to 10
     */
    readonly maxConns: pulumi.Input<number>;
    /**
     * the size in megabytes of the maximum request content that will be accepted, defaults to 20
     */
    readonly maxRequestSize: pulumi.Input<number>;
    /**
     * one of: least_conn (sends new requests to the least busy server), random (sends new requests to a random backend),
     * round_robin (sends new requests to the next backend in order), ip_hash (sends requests from a given IP address to the
     * same backend), default is random
     */
    readonly policy: pulumi.Input<string>;
    /**
     * you can listen on any port, the default is 80 to match the default protocol of http,if not you must specify it here
     * (commonly 80 for HTTP or 443 for HTTPS)
     */
    readonly port: pulumi.Input<number>;
    /**
     * either http or https. If you specify https then you must also provide the next two fields, the default is http
     */
    readonly protocol: pulumi.Input<string>;
    /**
     * if your protocol is https then you should send the TLS certificate in Base64-encoded PEM format
     */
    readonly tlsCertificate?: pulumi.Input<string>;
    /**
     * if your protocol is https then you should send the TLS private key in Base64-encoded PEM format
     */
    readonly tlsKey?: pulumi.Input<string>;
}
