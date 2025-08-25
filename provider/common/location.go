package common

import (
	"reflect"

	provider "github.com/pulumi/pulumi-go-provider"
)

type LocationInputs struct {
	Tenant *string  `pulumi:"tenant,optional" json:"tenant,omitempty"`
	Teams  []string `pulumi:"teams,optional" json:"teams,omitempty"`
}

type LocationOutput struct {
	Tenant string   `pulumi:"tenant" json:"tenant"`
	Teams  []string `pulumi:"teams" json:"teams"`
}

func (i LocationInputs) ToOutput() LocationOutput {
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
