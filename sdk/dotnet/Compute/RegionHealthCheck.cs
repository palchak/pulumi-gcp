// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_region_health_check.html.markdown.
    /// </summary>
    public partial class RegionHealthCheck : Pulumi.CustomResource
    {
        /// <summary>
        /// How often (in seconds) to send a health check. The default value is 5 seconds.
        /// </summary>
        [Output("checkIntervalSec")]
        public Output<int?> CheckIntervalSec { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value
        /// is 2.
        /// </summary>
        [Output("healthyThreshold")]
        public Output<int?> HealthyThreshold { get; private set; } = null!;

        /// <summary>
        /// A nested object resource
        /// </summary>
        [Output("http2HealthCheck")]
        public Output<Outputs.RegionHealthCheckHttp2HealthCheck?> Http2HealthCheck { get; private set; } = null!;

        /// <summary>
        /// A nested object resource
        /// </summary>
        [Output("httpHealthCheck")]
        public Output<Outputs.RegionHealthCheckHttpHealthCheck?> HttpHealthCheck { get; private set; } = null!;

        /// <summary>
        /// A nested object resource
        /// </summary>
        [Output("httpsHealthCheck")]
        public Output<Outputs.RegionHealthCheckHttpsHealthCheck?> HttpsHealthCheck { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters
        /// long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular
        /// expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all
        /// following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be
        /// a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The Region in which the created health check should reside. If it is not provided, the provider region is
        /// used.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// A nested object resource
        /// </summary>
        [Output("sslHealthCheck")]
        public Output<Outputs.RegionHealthCheckSslHealthCheck?> SslHealthCheck { get; private set; } = null!;

        /// <summary>
        /// A nested object resource
        /// </summary>
        [Output("tcpHealthCheck")]
        public Output<Outputs.RegionHealthCheckTcpHealthCheck?> TcpHealthCheck { get; private set; } = null!;

        /// <summary>
        /// How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for
        /// timeoutSec to have greater value than checkIntervalSec.
        /// </summary>
        [Output("timeoutSec")]
        public Output<int?> TimeoutSec { get; private set; } = null!;

        /// <summary>
        /// The type of the health check. One of HTTP, HTTP2, HTTPS, TCP, or SSL.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value
        /// is 2.
        /// </summary>
        [Output("unhealthyThreshold")]
        public Output<int?> UnhealthyThreshold { get; private set; } = null!;


        /// <summary>
        /// Create a RegionHealthCheck resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RegionHealthCheck(string name, RegionHealthCheckArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:compute/regionHealthCheck:RegionHealthCheck", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private RegionHealthCheck(string name, Input<string> id, RegionHealthCheckState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/regionHealthCheck:RegionHealthCheck", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RegionHealthCheck resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RegionHealthCheck Get(string name, Input<string> id, RegionHealthCheckState? state = null, CustomResourceOptions? options = null)
        {
            return new RegionHealthCheck(name, id, state, options);
        }
    }

    public sealed class RegionHealthCheckArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// How often (in seconds) to send a health check. The default value is 5 seconds.
        /// </summary>
        [Input("checkIntervalSec")]
        public Input<int>? CheckIntervalSec { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value
        /// is 2.
        /// </summary>
        [Input("healthyThreshold")]
        public Input<int>? HealthyThreshold { get; set; }

        /// <summary>
        /// A nested object resource
        /// </summary>
        [Input("http2HealthCheck")]
        public Input<Inputs.RegionHealthCheckHttp2HealthCheckArgs>? Http2HealthCheck { get; set; }

        /// <summary>
        /// A nested object resource
        /// </summary>
        [Input("httpHealthCheck")]
        public Input<Inputs.RegionHealthCheckHttpHealthCheckArgs>? HttpHealthCheck { get; set; }

        /// <summary>
        /// A nested object resource
        /// </summary>
        [Input("httpsHealthCheck")]
        public Input<Inputs.RegionHealthCheckHttpsHealthCheckArgs>? HttpsHealthCheck { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters
        /// long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular
        /// expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all
        /// following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be
        /// a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The Region in which the created health check should reside. If it is not provided, the provider region is
        /// used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// A nested object resource
        /// </summary>
        [Input("sslHealthCheck")]
        public Input<Inputs.RegionHealthCheckSslHealthCheckArgs>? SslHealthCheck { get; set; }

        /// <summary>
        /// A nested object resource
        /// </summary>
        [Input("tcpHealthCheck")]
        public Input<Inputs.RegionHealthCheckTcpHealthCheckArgs>? TcpHealthCheck { get; set; }

        /// <summary>
        /// How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for
        /// timeoutSec to have greater value than checkIntervalSec.
        /// </summary>
        [Input("timeoutSec")]
        public Input<int>? TimeoutSec { get; set; }

        /// <summary>
        /// A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value
        /// is 2.
        /// </summary>
        [Input("unhealthyThreshold")]
        public Input<int>? UnhealthyThreshold { get; set; }

        public RegionHealthCheckArgs()
        {
        }
    }

    public sealed class RegionHealthCheckState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// How often (in seconds) to send a health check. The default value is 5 seconds.
        /// </summary>
        [Input("checkIntervalSec")]
        public Input<int>? CheckIntervalSec { get; set; }

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value
        /// is 2.
        /// </summary>
        [Input("healthyThreshold")]
        public Input<int>? HealthyThreshold { get; set; }

        /// <summary>
        /// A nested object resource
        /// </summary>
        [Input("http2HealthCheck")]
        public Input<Inputs.RegionHealthCheckHttp2HealthCheckGetArgs>? Http2HealthCheck { get; set; }

        /// <summary>
        /// A nested object resource
        /// </summary>
        [Input("httpHealthCheck")]
        public Input<Inputs.RegionHealthCheckHttpHealthCheckGetArgs>? HttpHealthCheck { get; set; }

        /// <summary>
        /// A nested object resource
        /// </summary>
        [Input("httpsHealthCheck")]
        public Input<Inputs.RegionHealthCheckHttpsHealthCheckGetArgs>? HttpsHealthCheck { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters
        /// long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular
        /// expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all
        /// following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be
        /// a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The Region in which the created health check should reside. If it is not provided, the provider region is
        /// used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// A nested object resource
        /// </summary>
        [Input("sslHealthCheck")]
        public Input<Inputs.RegionHealthCheckSslHealthCheckGetArgs>? SslHealthCheck { get; set; }

        /// <summary>
        /// A nested object resource
        /// </summary>
        [Input("tcpHealthCheck")]
        public Input<Inputs.RegionHealthCheckTcpHealthCheckGetArgs>? TcpHealthCheck { get; set; }

        /// <summary>
        /// How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for
        /// timeoutSec to have greater value than checkIntervalSec.
        /// </summary>
        [Input("timeoutSec")]
        public Input<int>? TimeoutSec { get; set; }

        /// <summary>
        /// The type of the health check. One of HTTP, HTTP2, HTTPS, TCP, or SSL.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value
        /// is 2.
        /// </summary>
        [Input("unhealthyThreshold")]
        public Input<int>? UnhealthyThreshold { get; set; }

        public RegionHealthCheckState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class RegionHealthCheckHttp2HealthCheckArgs : Pulumi.ResourceArgs
    {
        [Input("host")]
        public Input<string>? Host { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("portName")]
        public Input<string>? PortName { get; set; }

        [Input("portSpecification")]
        public Input<string>? PortSpecification { get; set; }

        [Input("proxyHeader")]
        public Input<string>? ProxyHeader { get; set; }

        [Input("requestPath")]
        public Input<string>? RequestPath { get; set; }

        [Input("response")]
        public Input<string>? Response { get; set; }

        public RegionHealthCheckHttp2HealthCheckArgs()
        {
        }
    }

    public sealed class RegionHealthCheckHttp2HealthCheckGetArgs : Pulumi.ResourceArgs
    {
        [Input("host")]
        public Input<string>? Host { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("portName")]
        public Input<string>? PortName { get; set; }

        [Input("portSpecification")]
        public Input<string>? PortSpecification { get; set; }

        [Input("proxyHeader")]
        public Input<string>? ProxyHeader { get; set; }

        [Input("requestPath")]
        public Input<string>? RequestPath { get; set; }

        [Input("response")]
        public Input<string>? Response { get; set; }

        public RegionHealthCheckHttp2HealthCheckGetArgs()
        {
        }
    }

    public sealed class RegionHealthCheckHttpHealthCheckArgs : Pulumi.ResourceArgs
    {
        [Input("host")]
        public Input<string>? Host { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("portName")]
        public Input<string>? PortName { get; set; }

        [Input("portSpecification")]
        public Input<string>? PortSpecification { get; set; }

        [Input("proxyHeader")]
        public Input<string>? ProxyHeader { get; set; }

        [Input("requestPath")]
        public Input<string>? RequestPath { get; set; }

        [Input("response")]
        public Input<string>? Response { get; set; }

        public RegionHealthCheckHttpHealthCheckArgs()
        {
        }
    }

    public sealed class RegionHealthCheckHttpHealthCheckGetArgs : Pulumi.ResourceArgs
    {
        [Input("host")]
        public Input<string>? Host { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("portName")]
        public Input<string>? PortName { get; set; }

        [Input("portSpecification")]
        public Input<string>? PortSpecification { get; set; }

        [Input("proxyHeader")]
        public Input<string>? ProxyHeader { get; set; }

        [Input("requestPath")]
        public Input<string>? RequestPath { get; set; }

        [Input("response")]
        public Input<string>? Response { get; set; }

        public RegionHealthCheckHttpHealthCheckGetArgs()
        {
        }
    }

    public sealed class RegionHealthCheckHttpsHealthCheckArgs : Pulumi.ResourceArgs
    {
        [Input("host")]
        public Input<string>? Host { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("portName")]
        public Input<string>? PortName { get; set; }

        [Input("portSpecification")]
        public Input<string>? PortSpecification { get; set; }

        [Input("proxyHeader")]
        public Input<string>? ProxyHeader { get; set; }

        [Input("requestPath")]
        public Input<string>? RequestPath { get; set; }

        [Input("response")]
        public Input<string>? Response { get; set; }

        public RegionHealthCheckHttpsHealthCheckArgs()
        {
        }
    }

    public sealed class RegionHealthCheckHttpsHealthCheckGetArgs : Pulumi.ResourceArgs
    {
        [Input("host")]
        public Input<string>? Host { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("portName")]
        public Input<string>? PortName { get; set; }

        [Input("portSpecification")]
        public Input<string>? PortSpecification { get; set; }

        [Input("proxyHeader")]
        public Input<string>? ProxyHeader { get; set; }

        [Input("requestPath")]
        public Input<string>? RequestPath { get; set; }

        [Input("response")]
        public Input<string>? Response { get; set; }

        public RegionHealthCheckHttpsHealthCheckGetArgs()
        {
        }
    }

    public sealed class RegionHealthCheckSslHealthCheckArgs : Pulumi.ResourceArgs
    {
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("portName")]
        public Input<string>? PortName { get; set; }

        [Input("portSpecification")]
        public Input<string>? PortSpecification { get; set; }

        [Input("proxyHeader")]
        public Input<string>? ProxyHeader { get; set; }

        [Input("request")]
        public Input<string>? Request { get; set; }

        [Input("response")]
        public Input<string>? Response { get; set; }

        public RegionHealthCheckSslHealthCheckArgs()
        {
        }
    }

    public sealed class RegionHealthCheckSslHealthCheckGetArgs : Pulumi.ResourceArgs
    {
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("portName")]
        public Input<string>? PortName { get; set; }

        [Input("portSpecification")]
        public Input<string>? PortSpecification { get; set; }

        [Input("proxyHeader")]
        public Input<string>? ProxyHeader { get; set; }

        [Input("request")]
        public Input<string>? Request { get; set; }

        [Input("response")]
        public Input<string>? Response { get; set; }

        public RegionHealthCheckSslHealthCheckGetArgs()
        {
        }
    }

    public sealed class RegionHealthCheckTcpHealthCheckArgs : Pulumi.ResourceArgs
    {
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("portName")]
        public Input<string>? PortName { get; set; }

        [Input("portSpecification")]
        public Input<string>? PortSpecification { get; set; }

        [Input("proxyHeader")]
        public Input<string>? ProxyHeader { get; set; }

        [Input("request")]
        public Input<string>? Request { get; set; }

        [Input("response")]
        public Input<string>? Response { get; set; }

        public RegionHealthCheckTcpHealthCheckArgs()
        {
        }
    }

    public sealed class RegionHealthCheckTcpHealthCheckGetArgs : Pulumi.ResourceArgs
    {
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("portName")]
        public Input<string>? PortName { get; set; }

        [Input("portSpecification")]
        public Input<string>? PortSpecification { get; set; }

        [Input("proxyHeader")]
        public Input<string>? ProxyHeader { get; set; }

        [Input("request")]
        public Input<string>? Request { get; set; }

        [Input("response")]
        public Input<string>? Response { get; set; }

        public RegionHealthCheckTcpHealthCheckGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class RegionHealthCheckHttp2HealthCheck
    {
        public readonly string? Host;
        public readonly int? Port;
        public readonly string? PortName;
        public readonly string? PortSpecification;
        public readonly string? ProxyHeader;
        public readonly string? RequestPath;
        public readonly string? Response;

        [OutputConstructor]
        private RegionHealthCheckHttp2HealthCheck(
            string? host,
            int? port,
            string? portName,
            string? portSpecification,
            string? proxyHeader,
            string? requestPath,
            string? response)
        {
            Host = host;
            Port = port;
            PortName = portName;
            PortSpecification = portSpecification;
            ProxyHeader = proxyHeader;
            RequestPath = requestPath;
            Response = response;
        }
    }

    [OutputType]
    public sealed class RegionHealthCheckHttpHealthCheck
    {
        public readonly string? Host;
        public readonly int? Port;
        public readonly string? PortName;
        public readonly string? PortSpecification;
        public readonly string? ProxyHeader;
        public readonly string? RequestPath;
        public readonly string? Response;

        [OutputConstructor]
        private RegionHealthCheckHttpHealthCheck(
            string? host,
            int? port,
            string? portName,
            string? portSpecification,
            string? proxyHeader,
            string? requestPath,
            string? response)
        {
            Host = host;
            Port = port;
            PortName = portName;
            PortSpecification = portSpecification;
            ProxyHeader = proxyHeader;
            RequestPath = requestPath;
            Response = response;
        }
    }

    [OutputType]
    public sealed class RegionHealthCheckHttpsHealthCheck
    {
        public readonly string? Host;
        public readonly int? Port;
        public readonly string? PortName;
        public readonly string? PortSpecification;
        public readonly string? ProxyHeader;
        public readonly string? RequestPath;
        public readonly string? Response;

        [OutputConstructor]
        private RegionHealthCheckHttpsHealthCheck(
            string? host,
            int? port,
            string? portName,
            string? portSpecification,
            string? proxyHeader,
            string? requestPath,
            string? response)
        {
            Host = host;
            Port = port;
            PortName = portName;
            PortSpecification = portSpecification;
            ProxyHeader = proxyHeader;
            RequestPath = requestPath;
            Response = response;
        }
    }

    [OutputType]
    public sealed class RegionHealthCheckSslHealthCheck
    {
        public readonly int? Port;
        public readonly string? PortName;
        public readonly string? PortSpecification;
        public readonly string? ProxyHeader;
        public readonly string? Request;
        public readonly string? Response;

        [OutputConstructor]
        private RegionHealthCheckSslHealthCheck(
            int? port,
            string? portName,
            string? portSpecification,
            string? proxyHeader,
            string? request,
            string? response)
        {
            Port = port;
            PortName = portName;
            PortSpecification = portSpecification;
            ProxyHeader = proxyHeader;
            Request = request;
            Response = response;
        }
    }

    [OutputType]
    public sealed class RegionHealthCheckTcpHealthCheck
    {
        public readonly int? Port;
        public readonly string? PortName;
        public readonly string? PortSpecification;
        public readonly string? ProxyHeader;
        public readonly string? Request;
        public readonly string? Response;

        [OutputConstructor]
        private RegionHealthCheckTcpHealthCheck(
            int? port,
            string? portName,
            string? portSpecification,
            string? proxyHeader,
            string? request,
            string? response)
        {
            Port = port;
            PortName = portName;
            PortSpecification = portSpecification;
            ProxyHeader = proxyHeader;
            Request = request;
            Response = response;
        }
    }
    }
}
