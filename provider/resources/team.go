package resources

import (
	"pulumi-otoroshi/provider/common"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type TeamInputs struct {
	common.BaseInputStruct
	Tenant string `pulumi:"Team" json:"Team"`
}

type TeamOutput struct {
	common.BaseOutputStruct
	Tenant string `pulumi:"Team" json:"Team"`
}

type Team struct {
	common.BaseResource[TeamInputs, TeamOutput]
}

func (Team) Annotate(a infer.Annotator) {
	a.SetToken("organize", "Team")
}

func NewTeam() Team {
	return Team{
		common.BaseResource[TeamInputs, TeamOutput]{
			Path: "/apis/organize.otoroshi.io/v1/teams",
			CreateOutput: func() TeamOutput {
				return TeamOutput{}
			},
			WithDefaults: func(inputs TeamInputs) TeamInputs {
				if inputs.Description == nil {
					empty := ""
					inputs.Description = &empty
				}
				return inputs
			},
			ExtraToOutput: func(inputs TeamInputs, output *TeamOutput) {
				output.BaseOutputStruct = inputs.BaseInputStruct.ToOutput()
				output.Tenant = inputs.Tenant
			},
			ExtraDiff: func(oldValue TeamOutput, newValue TeamOutput, diffs map[string]provider.PropertyDiff) {
				diffs["tenant"] = common.DiffString(oldValue.Tenant, newValue.Tenant)
			},
		},
	}
}
