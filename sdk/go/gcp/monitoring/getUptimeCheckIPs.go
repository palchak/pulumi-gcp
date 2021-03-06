// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func GetUptimeCheckIPs(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetUptimeCheckIPsResult, error) {
	var rv GetUptimeCheckIPsResult
	err := ctx.Invoke("gcp:monitoring/getUptimeCheckIPs:getUptimeCheckIPs", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getUptimeCheckIPs.
type GetUptimeCheckIPsResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id             string                           `pulumi:"id"`
	UptimeCheckIps []GetUptimeCheckIPsUptimeCheckIp `pulumi:"uptimeCheckIps"`
}
