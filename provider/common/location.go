package common

import (
	"reflect"

	provider "github.com/pulumi/pulumi-go-provider"
)

var LocationInputsDefault = &LocationInputs{
	Tenant: StringDefault,
	Teams:  []string{*StringDefault},
}

type LocationInputsStruct struct {
	Location *LocationInputs `pulumi:"location,optional" computed:"true" json:"_loc,omitempty"`
}

type LocationInputs struct {
	Tenant *string  `pulumi:"tenant,optional" json:"tenant,omitempty"`
	Teams  []string `pulumi:"teams,optional" json:"teams,omitempty"`
}

type LocationOutputStruct struct {
	Location LocationOutput `pulumi:"_loc" json:"_loc"`
}

func (i LocationInputsStruct) ToOutput() LocationOutputStruct {
	return LocationOutputStruct{
		Location: i.Location.toOutput(),
	}
}

type LocationOutput struct {
	Tenant string   `pulumi:"tenant" json:"tenant"`
	Teams  []string `pulumi:"teams" json:"teams"`
}

func (i LocationInputs) toOutput() LocationOutput {
	return LocationOutput{
		Tenant: *i.Tenant,
		Teams:  i.Teams,
	}
}

func (o LocationOutput) IsEmpty() bool {
	return o.Tenant == "" && isEmptySlice(o.Teams)
}

func DiffLocation(oldValue LocationOutput, newValue LocationOutput) provider.PropertyDiff {
	if reflect.DeepEqual(oldValue, newValue) {
		return provider.PropertyDiff{
			Kind:      provider.Stable,
			InputDiff: false,
		}
	}
	var kind provider.DiffKind
	switch {
	case oldValue.IsEmpty() && !newValue.IsEmpty():
		kind = provider.Add
	case !oldValue.IsEmpty() && newValue.IsEmpty():
		kind = provider.Delete
	default:
		kind = provider.Update
	}
	return provider.PropertyDiff{
		Kind:      kind,
		InputDiff: true,
	}
}
