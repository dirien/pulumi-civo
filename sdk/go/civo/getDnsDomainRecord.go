// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package civo

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupDnsDomainRecord(ctx *pulumi.Context, args *LookupDnsDomainRecordArgs, opts ...pulumi.InvokeOption) (*LookupDnsDomainRecordResult, error) {
	var rv LookupDnsDomainRecordResult
	err := ctx.Invoke("civo:index/getDnsDomainRecord:getDnsDomainRecord", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDnsDomainRecord.
type LookupDnsDomainRecordArgs struct {
	// The domain id of the record.
	DomainId string `pulumi:"domainId"`
	// The name of the record.
	Name string `pulumi:"name"`
}

// A collection of values returned by getDnsDomainRecord.
type LookupDnsDomainRecordResult struct {
	// The id account of the domain.
	AccountId string `pulumi:"accountId"`
	// The date when it was created in UTC format
	CreatedAt string `pulumi:"createdAt"`
	// The id of the domain
	DomainId string `pulumi:"domainId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The portion before the domain name (e.g. www) or an @ for the apex/root domain (you cannot use an A record with an amex/root domain)
	Name string `pulumi:"name"`
	// The priority of the record.
	Priority int `pulumi:"priority"`
	// How long caching DNS servers should cache this record.
	Ttl int `pulumi:"ttl"`
	// The choice of record type from A, CNAME, MX, SRV or TXT
	Type string `pulumi:"type"`
	// The date when it was updated in UTC format
	UpdatedAt string `pulumi:"updatedAt"`
	// The IP address (A or MX), hostname (CNAME or MX) or text value (TXT) to serve for this record
	Value string `pulumi:"value"`
}
