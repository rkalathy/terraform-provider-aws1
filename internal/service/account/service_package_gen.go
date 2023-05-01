// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package account

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceAlternateContact,
			TypeName: "aws_account_alternate_contact",
		},
		{
			Factory:  ResourcePrimaryContact,
			TypeName: "aws_account_primary_contact",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Account
}

var ServicePackage = &servicePackage{}
