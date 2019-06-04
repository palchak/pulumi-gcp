// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package composer

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides access to available Cloud Composer versions in a region for a given project.
func LookupImageVersions(ctx *pulumi.Context, args *GetImageVersionsArgs) (*GetImageVersionsResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["project"] = args.Project
		inputs["region"] = args.Region
	}
	outputs, err := ctx.Invoke("gcp:composer/getImageVersions:getImageVersions", inputs)
	if err != nil {
		return nil, err
	}
	return &GetImageVersionsResult{
		ImageVersions: outputs["imageVersions"],
		Project: outputs["project"],
		Region: outputs["region"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getImageVersions.
type GetImageVersionsArgs struct {
	// The ID of the project to list versions in.
	// If it is not provided, the provider project is used.
	Project interface{}
	// The location to list versions in.
	// If it is not provider, the provider region is used.
	Region interface{}
}

// A collection of values returned by getImageVersions.
type GetImageVersionsResult struct {
	// A list of composer image versions available in the given project and location. Each `image_version` contains:
	ImageVersions interface{}
	Project interface{}
	Region interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
