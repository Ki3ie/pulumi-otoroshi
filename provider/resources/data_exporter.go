package resources

import (
	"pulumi-otoroshi/provider/common"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type DataExporterInputs struct {
	common.BaseInputStruct
	common.LocationInputsStruct
}

type DataExporterOutput struct {
	common.BaseOutputStruct
	common.LocationOutputStruct
}

type DataExporter struct {
	common.BaseResource[DataExporterInputs, DataExporterOutput]
}

func (DataExporter) Annotate(a infer.Annotator) {
	a.SetToken("events", "DataExporter")
}

func NewDataExporter() DataExporter {
	return DataExporter{
		common.BaseResource[DataExporterInputs, DataExporterOutput]{
			Path: "/apis/events.otoroshi.io/v1/data-exporters",
			CreateOutput: func() DataExporterOutput {
				return DataExporterOutput{}
			},
			WithDefaults: func(inputs DataExporterInputs) DataExporterInputs {
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
			ExtraToOutput: func(inputs DataExporterInputs, output *DataExporterOutput) {
				output.BaseOutputStruct = inputs.BaseInputStruct.ToOutput()
				output.LocationOutputStruct = inputs.LocationInputsStruct.ToOutput()
			},
			ExtraDiff: func(oldValue DataExporterOutput, newValue DataExporterOutput, diffs map[string]provider.PropertyDiff) {
				diffs["location"] = common.DiffLocation(oldValue.Location, newValue.Location)
			},
		},
	}
}
