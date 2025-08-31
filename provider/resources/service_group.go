package resources

import (
	"log"
	"pulumi-otoroshi/provider/common"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type ServiceGroupInputs struct {
	common.BaseInputStruct
	common.LocationInputsStruct
}

type ServiceGroupOutput struct {
	common.BaseOutputStruct
	common.LocationOutputStruct
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
			CreateOutput: func() ServiceGroupOutput {
				return ServiceGroupOutput{}
			},
			WithDefaults: func(inputs ServiceGroupInputs) ServiceGroupInputs {
				if inputs.Description == nil {
					empty := ""
					inputs.Description = &empty
				}
				if inputs.Location == nil {
					defaultValue := "default"
					inputs.Location = &common.LocationInputs{
						Tenant: &defaultValue,
						Teams:  []string{defaultValue},
					}
				}
				return inputs
			},
			ExtraToOutput: func(inputs ServiceGroupInputs, output *ServiceGroupOutput) {
				log.Printf("DEBUG ServiceGroupOutput output=%v\n", output)
				log.Printf("DEBUG ServiceGroupInputs inputs=%v\n", inputs)
				output.Location = inputs.Location.ToOutput()
			},
			ExtraDiff: func(oldValue ServiceGroupOutput, newValue ServiceGroupOutput, diffs map[string]provider.PropertyDiff) {
				diffs["location"] = common.DiffLocation(oldValue.Location, newValue.Location)
			},
		},
	}
}
