package resources

import (
	"pulumi-otoroshi/provider/common"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type OrganizationInputs struct {
	Name        string            `pulumi:"name" json:"name"`
	Description *string           `pulumi:"description,optional" computed:"true" json:"description,omitempty"`
	Tags        []string          `pulumi:"tags,optional" computed:"true" json:"tags,omitempty"`
	Metadata    map[string]string `pulumi:"metadata,optional" computed:"true" json:"metadata,omitempty"`
}

type OrganizationOutput struct {
	common.BaseOutputStruct
	Name        string            `pulumi:"name" json:"name"`
	Description string            `pulumi:"description" json:"description"`
	Tags        []string          `pulumi:"tags" json:"tags"`
	Metadata    map[string]string `pulumi:"metadata,optional" json:"metadata,omitempty"`
}

type Organization struct {
	common.BaseResource[OrganizationInputs, OrganizationOutput]
}

func (Organization) Annotate(a infer.Annotator) {
	a.SetToken("organize", "Organization")
}

func NewOrganization() Organization {
	return Organization{
		common.BaseResource[OrganizationInputs, OrganizationOutput]{
			Path: "/apis/organize.otoroshi.io/v1/organizations",
			WithDefaults: func(inputs OrganizationInputs) OrganizationInputs {
				if inputs.Description == nil {
					empty := ""
					inputs.Description = &empty
				}
				return inputs
			},
			ToOutput: func(inputs OrganizationInputs) OrganizationOutput {
				return OrganizationOutput{
					Name:        inputs.Name,
					Description: *inputs.Description,
					Tags:        inputs.Tags,
					Metadata:    inputs.Metadata,
				}
			},
			DiffOutput: func(oldValue OrganizationOutput, newValue OrganizationOutput) map[string]provider.PropertyDiff {
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
