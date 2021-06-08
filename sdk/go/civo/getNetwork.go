// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package civo

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetwork(ctx *pulumi.Context, args *LookupNetworkArgs, opts ...pulumi.InvokeOption) (*LookupNetworkResult, error) {
	var rv LookupNetworkResult
	err := ctx.Invoke("civo:index/getNetwork:getNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetwork.
type LookupNetworkArgs struct {
	// The unique identifier of an existing Network.
	Id *string `pulumi:"id"`
	// The label of an existing Network.
	Label *string `pulumi:"label"`
	// The region of an existing Network.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getNetwork.
type LookupNetworkResult struct {
	// If is the default network.
	Default bool `pulumi:"default"`
	// A unique ID that can be used to identify and reference a Network.
	Id *string `pulumi:"id"`
	// The label used in the configuration.
	Label *string `pulumi:"label"`
	// The name of the network.
	Name   string  `pulumi:"name"`
	Region *string `pulumi:"region"`
}
