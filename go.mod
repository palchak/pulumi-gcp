module github.com/pulumi/pulumi-gcp

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk v1.4.0
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.6.0
	github.com/pulumi/pulumi-terraform-bridge v1.5.3-0.20200116214742-68c8f56da38a
	github.com/stretchr/testify v1.4.0
	github.com/terraform-providers/terraform-provider-google-beta v0.0.0-00010101000000-000000000000
	github.com/xanzy/ssh-agent v0.2.1 // indirect
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
	github.com/terraform-providers/terraform-provider-google-beta => github.com/pulumi/terraform-provider-google-beta v0.0.0-20200108173009-0a5002629481
)
