package resources

import (
	"pulumi-otoroshi/provider/common"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type TenantInputs struct {
	Name        string            `pulumi:"name" json:"name"`
	Description *string           `pulumi:"description,optional" computed:"true" json:"description,omitempty"`
	Tags        []string          `pulumi:"tags,optional" computed:"true" json:"tags,omitempty"`
	Metadata    map[string]string `pulumi:"metadata,optional" computed:"true" json:"metadata,omitempty"`
}

type TenantOutput struct {
	common.BaseOutputStruct
	Name        string            `pulumi:"name" json:"name"`
	Description string            `pulumi:"description" json:"description"`
	Tags        []string          `pulumi:"tags" json:"tags"`
	Metadata    map[string]string `pulumi:"metadata,optional" json:"metadata,omitempty"`
}

type Tenant struct {
	common.BaseResource[TenantInputs, TenantOutput]
}

func (Tenant) Annotate(a infer.Annotator) {
	a.SetToken("organize", "Tenant")
}

func NewTenant() Tenant {
	return Tenant{
		common.BaseResource[TenantInputs, TenantOutput]{
			Path: "/apis/organize.otoroshi.io/v1/tenants",
			WithDefaults: func(inputs TenantInputs) TenantInputs {
				if inputs.Description == nil {
					empty := ""
					inputs.Description = &empty
				}
				return inputs
			},
			ToOutput: func(inputs TenantInputs) TenantOutput {
				return TenantOutput{
					Name:        inputs.Name,
					Description: *inputs.Description,
					Tags:        inputs.Tags,
					Metadata:    inputs.Metadata,
				}
			},
			DiffOutput: func(oldValue TenantOutput, newValue TenantOutput) map[string]provider.PropertyDiff {
				diffs := map[string]provider.PropertyDiff{}
				diffs["name"] = common.DiffString(oldValue.Name, newValue.Name)
				diffs["description"] = common.DiffString(oldValue.Description, newValue.Description)
				diffs["tags"] = common.DiffSlice(oldValue.Tags, newValue.Tags)
				diffs["metadata"] = common.DiffMap(oldValue.Metadata, newValue.Metadata)
				return diffs
			},
		},
	}
}
