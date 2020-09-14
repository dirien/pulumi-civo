// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package civo

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Civo dns domain record resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-civo/sdk/go/civo"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := civo.NewDnsDomainRecord(ctx, "www", &civo.DnsDomainRecordArgs{
// 			DomainId: pulumi.Any(civo_dns_domain_name.Main.Id),
// 			Type:     pulumi.String("A"),
// 			Value:    pulumi.Any(civo_instance.Foo.Public_ip),
// 			Ttl:      pulumi.Int(600),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			civo_dns_domain_name.Main,
// 			civo_instance.Foo,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DnsDomainRecord struct {
	pulumi.CustomResourceState

	// The id account of the domain
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The date when it was created in UTC format
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The id of the domain
	DomainId pulumi.StringOutput `pulumi:"domainId"`
	// The portion before the domain name (e.g. www) or an @ for the apex/root domain (you cannot use an A record with an amex/root domain)
	Name pulumi.StringOutput `pulumi:"name"`
	// Useful for MX records only, the priority mail should be attempted it (defaults to 10)
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// How long caching DNS servers should cache this record for, in seconds (the minimum is 600 and the default if unspecified is 600)
	Ttl pulumi.IntOutput `pulumi:"ttl"`
	// The choice of record type from A, CNAME, MX, SRV or TXT
	Type pulumi.StringOutput `pulumi:"type"`
	// The date when it was updated in UTC format
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// The IP address (A or MX), hostname (CNAME or MX) or text value (TXT) to serve for this record
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewDnsDomainRecord registers a new resource with the given unique name, arguments, and options.
func NewDnsDomainRecord(ctx *pulumi.Context,
	name string, args *DnsDomainRecordArgs, opts ...pulumi.ResourceOption) (*DnsDomainRecord, error) {
	if args == nil || args.DomainId == nil {
		return nil, errors.New("missing required argument 'DomainId'")
	}
	if args == nil || args.Ttl == nil {
		return nil, errors.New("missing required argument 'Ttl'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil || args.Value == nil {
		return nil, errors.New("missing required argument 'Value'")
	}
	if args == nil {
		args = &DnsDomainRecordArgs{}
	}
	var resource DnsDomainRecord
	err := ctx.RegisterResource("civo:index/dnsDomainRecord:DnsDomainRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsDomainRecord gets an existing DnsDomainRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsDomainRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsDomainRecordState, opts ...pulumi.ResourceOption) (*DnsDomainRecord, error) {
	var resource DnsDomainRecord
	err := ctx.ReadResource("civo:index/dnsDomainRecord:DnsDomainRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsDomainRecord resources.
type dnsDomainRecordState struct {
	// The id account of the domain
	AccountId *string `pulumi:"accountId"`
	// The date when it was created in UTC format
	CreatedAt *string `pulumi:"createdAt"`
	// The id of the domain
	DomainId *string `pulumi:"domainId"`
	// The portion before the domain name (e.g. www) or an @ for the apex/root domain (you cannot use an A record with an amex/root domain)
	Name *string `pulumi:"name"`
	// Useful for MX records only, the priority mail should be attempted it (defaults to 10)
	Priority *int `pulumi:"priority"`
	// How long caching DNS servers should cache this record for, in seconds (the minimum is 600 and the default if unspecified is 600)
	Ttl *int `pulumi:"ttl"`
	// The choice of record type from A, CNAME, MX, SRV or TXT
	Type *string `pulumi:"type"`
	// The date when it was updated in UTC format
	UpdatedAt *string `pulumi:"updatedAt"`
	// The IP address (A or MX), hostname (CNAME or MX) or text value (TXT) to serve for this record
	Value *string `pulumi:"value"`
}

type DnsDomainRecordState struct {
	// The id account of the domain
	AccountId pulumi.StringPtrInput
	// The date when it was created in UTC format
	CreatedAt pulumi.StringPtrInput
	// The id of the domain
	DomainId pulumi.StringPtrInput
	// The portion before the domain name (e.g. www) or an @ for the apex/root domain (you cannot use an A record with an amex/root domain)
	Name pulumi.StringPtrInput
	// Useful for MX records only, the priority mail should be attempted it (defaults to 10)
	Priority pulumi.IntPtrInput
	// How long caching DNS servers should cache this record for, in seconds (the minimum is 600 and the default if unspecified is 600)
	Ttl pulumi.IntPtrInput
	// The choice of record type from A, CNAME, MX, SRV or TXT
	Type pulumi.StringPtrInput
	// The date when it was updated in UTC format
	UpdatedAt pulumi.StringPtrInput
	// The IP address (A or MX), hostname (CNAME or MX) or text value (TXT) to serve for this record
	Value pulumi.StringPtrInput
}

func (DnsDomainRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsDomainRecordState)(nil)).Elem()
}

type dnsDomainRecordArgs struct {
	// The id of the domain
	DomainId string `pulumi:"domainId"`
	// The portion before the domain name (e.g. www) or an @ for the apex/root domain (you cannot use an A record with an amex/root domain)
	Name *string `pulumi:"name"`
	// Useful for MX records only, the priority mail should be attempted it (defaults to 10)
	Priority *int `pulumi:"priority"`
	// How long caching DNS servers should cache this record for, in seconds (the minimum is 600 and the default if unspecified is 600)
	Ttl int `pulumi:"ttl"`
	// The choice of record type from A, CNAME, MX, SRV or TXT
	Type string `pulumi:"type"`
	// The IP address (A or MX), hostname (CNAME or MX) or text value (TXT) to serve for this record
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a DnsDomainRecord resource.
type DnsDomainRecordArgs struct {
	// The id of the domain
	DomainId pulumi.StringInput
	// The portion before the domain name (e.g. www) or an @ for the apex/root domain (you cannot use an A record with an amex/root domain)
	Name pulumi.StringPtrInput
	// Useful for MX records only, the priority mail should be attempted it (defaults to 10)
	Priority pulumi.IntPtrInput
	// How long caching DNS servers should cache this record for, in seconds (the minimum is 600 and the default if unspecified is 600)
	Ttl pulumi.IntInput
	// The choice of record type from A, CNAME, MX, SRV or TXT
	Type pulumi.StringInput
	// The IP address (A or MX), hostname (CNAME or MX) or text value (TXT) to serve for this record
	Value pulumi.StringInput
}

func (DnsDomainRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsDomainRecordArgs)(nil)).Elem()
}
