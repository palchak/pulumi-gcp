// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provide access to a Backend Service's attribute. For more information
// see [the official documentation](https://cloud.google.com/compute/docs/load-balancing/http/backend-service)
// and the [API](https://cloud.google.com/compute/docs/reference/latest/backendServices).
func LookupBackendService(ctx *pulumi.Context, args *LookupBackendServiceArgs, opts ...pulumi.InvokeOption) (*LookupBackendServiceResult, error) {
	var rv LookupBackendServiceResult
	err := ctx.Invoke("gcp:compute/getBackendService:getBackendService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBackendService.
type LookupBackendServiceArgs struct {
	// The name of the Backend Service.
	Name string `pulumi:"name"`
	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getBackendService.
type LookupBackendServiceResult struct {
	AffinityCookieTtlSec int `pulumi:"affinityCookieTtlSec"`
	// The set of backends that serve this Backend Service.
	Backends        []GetBackendServiceBackend        `pulumi:"backends"`
	CdnPolicies     []GetBackendServiceCdnPolicy      `pulumi:"cdnPolicies"`
	CircuitBreakers []GetBackendServiceCircuitBreaker `pulumi:"circuitBreakers"`
	// Time for which instance will be drained (not accept new connections, but still work to finish started ones).
	ConnectionDrainingTimeoutSec int                               `pulumi:"connectionDrainingTimeoutSec"`
	ConsistentHash               []GetBackendServiceConsistentHash `pulumi:"consistentHash"`
	CreationTimestamp            string                            `pulumi:"creationTimestamp"`
	CustomRequestHeaders         []string                          `pulumi:"customRequestHeaders"`
	// Textual description for the Backend Service.
	Description string `pulumi:"description"`
	// Whether or not Cloud CDN is enabled on the Backend Service.
	EnableCdn bool `pulumi:"enableCdn"`
	// The fingerprint of the Backend Service.
	Fingerprint string `pulumi:"fingerprint"`
	// The set of HTTP/HTTPS health checks used by the Backend Service.
	HealthChecks []string               `pulumi:"healthChecks"`
	Iaps         []GetBackendServiceIap `pulumi:"iaps"`
	// id is the provider-assigned unique ID for this managed resource.
	Id                  string                              `pulumi:"id"`
	LoadBalancingScheme string                              `pulumi:"loadBalancingScheme"`
	LocalityLbPolicy    string                              `pulumi:"localityLbPolicy"`
	LogConfigs          []GetBackendServiceLogConfig        `pulumi:"logConfigs"`
	Name                string                              `pulumi:"name"`
	OutlierDetections   []GetBackendServiceOutlierDetection `pulumi:"outlierDetections"`
	// The name of a service that has been added to an instance group in this backend.
	PortName string  `pulumi:"portName"`
	Project  *string `pulumi:"project"`
	// The protocol for incoming requests.
	Protocol       string `pulumi:"protocol"`
	SecurityPolicy string `pulumi:"securityPolicy"`
	// The URI of the Backend Service.
	SelfLink string `pulumi:"selfLink"`
	// The Backend Service session stickiness configuration.
	SessionAffinity string `pulumi:"sessionAffinity"`
	// The number of seconds to wait for a backend to respond to a request before considering the request failed.
	TimeoutSec int `pulumi:"timeoutSec"`
}
