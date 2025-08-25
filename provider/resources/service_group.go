package resources

import (
	"pulumi-otoroshi/provider/common"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type ServiceGroupInputs struct {
	Name        string                 `pulumi:"name" json:"name"`
	Location    *common.LocationInputs `pulumi:"location,optional"  computed:"true" json:"_loc,omitempty"`
	Description *string                `pulumi:"description,optional" computed:"true" json:"description,omitempty"`
	Tags        []string               `pulumi:"tags,optional" computed:"true" json:"tags,omitempty"`
	Metadata    map[string]string      `pulumi:"metadata,optional" computed:"true" json:"metadata,omitempty"`
}

type ServiceGroupOutput struct {
	common.BaseOutputStruct
	Name        string                `pulumi:"name" json:"name"`
	Location    common.LocationOutput `pulumi:"location" json:"_loc"`
	Description string                `pulumi:"description" json:"description"`
	Tags        []string              `pulumi:"tags" json:"tags"`
	Metadata    map[string]string     `pulumi:"metadata,optional" json:"metadata,omitempty"`
}

type ServiceGroup struct {
	common.BaseResource[ServiceGroupInputs, ServiceGroupOutput]
}

func (ServiceGroup) Annotate(a infer.Annotator) {
	a.SetToken("organize", "ServiceGroup")
}

func NewServiceGroup() ServiceGroup {
	return ServiceGroup{
		common.BaseResource[ServiceGroupInputs, ServiceGroupOutput]{
			Path: "/apis/organize.otoroshi.io/v1/service-groups",
			WithDefaults: func(inputs ServiceGroupInputs) ServiceGroupInputs {
				// TODO : Location
				if inputs.Description == nil {
					empty := ""
					inputs.Description = &empty
				}
				return inputs
			},
			ToOutput: func(inputs ServiceGroupInputs) ServiceGroupOutput {
				return ServiceGroupOutput{
					Name:        inputs.Name,
					Location:    inputs.Location.ToOutput(),
					Description: *inputs.Description,
					Tags:        inputs.Tags,
					Metadata:    inputs.Metadata,
				}
			},
			DiffOutput: func(oldValue ServiceGroupOutput, newValue ServiceGroupOutput) map[string]provider.PropertyDiff {
				diffs := map[string]provider.PropertyDiff{}
				diffs["name"] = common.DiffString(oldValue.Name, newValue.Name)
				diffs["location"] = common.DiffLocation(oldValue.Location, newValue.Location)
				diffs["description"] = common.DiffString(oldValue.Description, newValue.Description)
				diffs["tags"] = common.DiffSlice(oldValue.Tags, newValue.Tags)
				diffs["metadata"] = common.DiffMap(oldValue.Metadata, newValue.Metadata)
				return diffs
			},
		},
	}
}
