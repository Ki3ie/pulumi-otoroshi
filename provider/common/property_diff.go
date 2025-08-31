package common

import (
	"log"
	"maps"
	"slices"

	provider "github.com/pulumi/pulumi-go-provider"
)

func DiffNilString(oldValue *string, newValue *string) provider.PropertyDiff {
	if oldValue == newValue {
		return provider.PropertyDiff{
			Kind:      provider.Stable,
			InputDiff: false,
		}
	}
	var kind provider.DiffKind
	switch {
	case oldValue == nil && newValue != nil:
		kind = provider.Add
	case oldValue != nil && newValue == nil:
		kind = provider.Delete
	default:
		kind = provider.Update
	}
	return provider.PropertyDiff{
		Kind:      kind,
		InputDiff: true,
	}
}

func DiffString(oldValue string, newValue string) provider.PropertyDiff {
	log.Printf("DEBUG DiffString: oldValue=%q, newValue=%q", oldValue, newValue)
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

func DiffBool(oldValue bool, newValue bool) provider.PropertyDiff {
	if oldValue == newValue {
		return provider.PropertyDiff{
			Kind:      provider.Stable,
			InputDiff: false,
		}
	}
	var kind provider.DiffKind
	switch {
	case oldValue && !newValue:
		kind = provider.Add
	case !oldValue && newValue:
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
	ignoreKeys := []string{"created_at", "updated_at"}
	for _, key := range ignoreKeys {
		delete(oldValue, key)
		delete(newValue, key)
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
