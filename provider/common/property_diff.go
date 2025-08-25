package common

import (
	"maps"
	"slices"

	provider "github.com/pulumi/pulumi-go-provider"
)

func DiffString(oldValue string, newValue string) provider.PropertyDiff {
	if oldValue == newValue {
		return provider.PropertyDiff{
			Kind:      provider.Stable,
			InputDiff: false,
		}
	}
	var kind provider.DiffKind
	switch {
	case oldValue == "" && newValue != "":
		kind = provider.Add
	case oldValue != "" && newValue == "":
		kind = provider.Delete
	default:
		kind = provider.Update
	}
	return provider.PropertyDiff{
		Kind:      kind,
		InputDiff: true,
	}
}

func DiffSlice(oldValue []string, newValue []string) provider.PropertyDiff {
	if slices.Equal(oldValue, newValue) {
		return provider.PropertyDiff{
			Kind:      provider.Stable,
			InputDiff: false,
		}
	}
	var kind provider.DiffKind
	switch {
	case isEmptySlice(oldValue) && !isEmptySlice(newValue):
		kind = provider.Add
	case !isEmptySlice(oldValue) && isEmptySlice(newValue):
		kind = provider.Delete
	default:
		kind = provider.Update
	}
	return provider.PropertyDiff{
		Kind:      kind,
		InputDiff: true,
	}
}

func isEmptySlice(s []string) bool {
	return s == nil || len(s) == 0
}

func DiffMap(oldValue map[string]string, newValue map[string]string) provider.PropertyDiff {
	if _, ok := oldValue["created_at"]; ok {
		delete(oldValue, "created_at")
	}
	if _, ok := oldValue["updated_at"]; ok {
		delete(oldValue, "updated_at")
	}
	if maps.Equal(oldValue, newValue) {
		return provider.PropertyDiff{
			Kind:      provider.Stable,
			InputDiff: false,
		}
	}
	var kind provider.DiffKind
	switch {
	case isEmptyMap(oldValue) && !isEmptyMap(newValue):
		kind = provider.Add
	case !isEmptyMap(oldValue) && isEmptyMap(newValue):
		kind = provider.Delete
	default:
		kind = provider.Update
	}
	return provider.PropertyDiff{
		Kind:      kind,
		InputDiff: true,
	}
}

func isEmptyMap(m map[string]string) bool {
	return m == nil || len(m) == 0
}
