package resources

import (
	"pulumi-otoroshi/provider/common"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type OrganizationInputs struct {
	common.BaseInputStruct
}

type OrganizationOutput struct {
	common.BaseOutputStruct
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
			CreateOutput: func() OrganizationOutput {
				return OrganizationOutput{}
			},
			WithDefaults: func(inputs OrganizationInputs) OrganizationInputs {
				if inputs.Description == nil {
					inputs.Description = common.StringEmpty
				}
				return inputs
			},
			ExtraToOutput: func(inputs OrganizationInputs, output *OrganizationOutput) {
				output.BaseOutputStruct = inputs.BaseInputStruct.ToOutput()
			},
		},
	}
}
