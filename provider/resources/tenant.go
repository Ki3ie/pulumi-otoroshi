package resources

import (
	"pulumi-otoroshi/provider/common"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type TenantInputs struct {
	common.BaseInputStruct
}

type TenantOutput struct {
	common.BaseOutputStruct
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
			CreateOutput: func() TenantOutput {
				return TenantOutput{}
			},
			WithDefaults: func(inputs TenantInputs) TenantInputs {
				if inputs.Description == nil {
					empty := ""
					inputs.Description = &empty
				}
				return inputs
			},
			ExtraToOutput: func(inputs TenantInputs, output *TenantOutput) {
				output.BaseOutputStruct = inputs.BaseInputStruct.ToOutput()
			},
		},
	}
}
