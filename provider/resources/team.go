package resources

import (
	"pulumi-otoroshi/provider/common"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type TeamInputs struct {
	Tenant      string            `pulumi:"Team" json:"Team"`
	Name        string            `pulumi:"name" json:"name"`
	Description *string           `pulumi:"description,optional" computed:"true" json:"description,omitempty"`
	Tags        []string          `pulumi:"tags,optional" computed:"true" json:"tags,omitempty"`
	Metadata    map[string]string `pulumi:"metadata,optional" computed:"true" json:"metadata,omitempty"`
}

type TeamOutput struct {
	common.BaseOutputStruct
	Tenant      string            `pulumi:"Team" json:"Team"`
	Name        string            `pulumi:"name" json:"name"`
	Description string            `pulumi:"description" json:"description"`
	Tags        []string          `pulumi:"tags" json:"tags"`
	Metadata    map[string]string `pulumi:"metadata,optional" json:"metadata,omitempty"`
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
			WithDefaults: func(inputs TeamInputs) TeamInputs {
				if inputs.Description == nil {
					empty := ""
					inputs.Description = &empty
				}
				return inputs
			},
			ToOutput: func(inputs TeamInputs) TeamOutput {
				return TeamOutput{
					Tenant:      inputs.Tenant,
					Name:        inputs.Name,
					Description: *inputs.Description,
					Tags:        inputs.Tags,
					Metadata:    inputs.Metadata,
				}
			},
			DiffOutput: func(oldValue TeamOutput, newValue TeamOutput) map[string]provider.PropertyDiff {
				diffs := map[string]provider.PropertyDiff{}
				diffs["tenant"] = common.DiffString(oldValue.Tenant, newValue.Tenant)
				diffs["name"] = common.DiffString(oldValue.Name, newValue.Name)
				diffs["description"] = common.DiffString(oldValue.Description, newValue.Description)
				diffs["tags"] = common.DiffSlice(oldValue.Tags, newValue.Tags)
				diffs["metadata"] = common.DiffMap(oldValue.Metadata, newValue.Metadata)
				return diffs
			},
		},
	}
}
