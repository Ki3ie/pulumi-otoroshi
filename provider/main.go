package main

import (
	"context"
	"fmt"
	"os"
	"pulumi-otoroshi/provider/common"
	"pulumi-otoroshi/provider/resources"

	"github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string

// Name controls how this provider is referenced in package names and elsewhere.
const Name string = "otoroshi"

func Provider() provider.Provider {
	builder := infer.ProviderBuilder{}
	provider, err := builder.
		WithLanguageMap(map[string]any{
			"nodejs": map[string]any{
				"respectSchemaVersion": true,
			},
		}).
		WithDisplayName("otoroshi").
		WithNamespace("pulumi").
		WithConfig(infer.Config(&common.Config{})).
		WithResources(
			// infer.Resource(events.DataExporter{}),
			infer.Resource(resources.NewOrganization()),
			infer.Resource(resources.NewServiceGroup()),
			infer.Resource(resources.NewTeam()),
			infer.Resource(resources.NewTenant()),
		).Build()
	if err != nil {
		panic(fmt.Errorf("unable to build provider: %w", err))
	}
	return provider
}

// Serve the provider against Pulumi's Provider protocol.
func main() {
	err := Provider().Run(context.Background(), Name, Version)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
}
