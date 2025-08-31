package proxy

import (
	"pulumi-otoroshi/provider/common"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Route struct {
	common.BaseResource[RouteInputs, RouteOutput]
}

func (Route) Annotate(a infer.Annotator) {
	a.SetToken("proxy", "Route")
}

func NewRoute() Route {
	return Route{
		common.BaseResource[RouteInputs, RouteOutput]{
			Path: "/apis/proxy.otoroshi.io/v1/routes",
			CreateOutput: func() RouteOutput {
				return RouteOutput{}
			},
			WithDefaults: func(inputs RouteInputs) RouteInputs {
				if inputs.Description == nil {
					inputs.Description = common.StringEmpty
				}
				if inputs.Location == nil {
					inputs.Location = common.LocationInputsDefault
				}
				if inputs.Enabled == nil {
					inputs.Enabled = common.BoolTrue
				}
				if inputs.DebugFlow == nil {
					inputs.DebugFlow = common.BoolFalse
				}
				if inputs.ExportReporting == nil {
					inputs.ExportReporting = common.BoolFalse
				}
				if inputs.Capture == nil {
					inputs.Capture = common.BoolFalse
				}
				if inputs.Groups == nil {
					inputs.Groups = []string{*common.StringDefault}
				}
				if inputs.Backend == nil {
					inputs.Backend = &BackendInputs{}
				}
				return inputs
			},
			ExtraToOutput: func(inputs RouteInputs, output *RouteOutput) {
				output.Location = inputs.Location.ToOutput()
				output.Enabled = *inputs.Enabled
				output.DebugFlow = *inputs.DebugFlow
				output.ExportReporting = *inputs.ExportReporting
				output.Capture = *inputs.Capture
				output.Groups = inputs.Groups
				output.BoundListeners = inputs.BoundListeners
				output.Frontend = FrontendOutput{} // TODO
				output.Backend = BackendOutput{}   // TODO
				output.BackendRef = inputs.BackendRef
				// TODO Plugins
			},
			ExtraDiff: func(oldValue RouteOutput, newValue RouteOutput, diffs map[string]provider.PropertyDiff) {
				diffs["location"] = common.DiffLocation(oldValue.Location, newValue.Location)

				diffs["enabled"] = common.DiffBool(oldValue.Enabled, newValue.Enabled)
				diffs["debug_flow"] = common.DiffBool(oldValue.DebugFlow, newValue.DebugFlow)
				diffs["export_reporting"] = common.DiffBool(oldValue.ExportReporting, newValue.ExportReporting)
				diffs["capture"] = common.DiffBool(oldValue.Capture, newValue.Capture)

				diffs["groups"] = common.DiffSlice(oldValue.Groups, newValue.Groups)
				diffs["bound_listeners"] = common.DiffSlice(oldValue.BoundListeners, newValue.BoundListeners)
				// TODO Frontend
				// TODO Backend
				diffs["backend_ref"] = common.DiffNilString(oldValue.BackendRef, newValue.BackendRef)
				// TODO Plugins
			},
		},
	}
}
