package proxy

import "pulumi-otoroshi/provider/common"

type RouteInputs struct {
	common.BaseInputStruct
	common.LocationInputsStruct
	Enabled         *bool           `pulumi:"enabled,optional" computed:"true" json:"enabled,omitempty"`
	DebugFlow       *bool           `pulumi:"debug_flow,optional" computed:"true" json:"debug_flow,omitempty"`
	ExportReporting *bool           `pulumi:"export_reporting,optional" computed:"true" json:"export_reporting,omitempty"`
	Capture         *bool           `pulumi:"capture,optional" computed:"true" json:"capture,omitempty"`
	Groups          []string        `pulumi:"groups,optional" computed:"true" json:"groups,omitempty"`
	BoundListeners  []string        `pulumi:"bound_listeners,optional" computed:"true" json:"bound_listeners,omitempty"`
	Frontend        *FrontendInputs `pulumi:"frontend,optional" computed:"true" json:"frontend,omitempty"`
	Backend         *BackendInputs  `pulumi:"backend,optional" computed:"true" json:"backend,omitempty"`
	BackendRef      *string         `pulumi:"backend_ref,optional" computed:"true" json:"backend_ref,omitempty"`
	// TODO Plugins
}

type FrontendInputs struct {
	Domains   []string          `pulumi:"domains,optional" computed:"true" json:"domains,omitempty"`
	StripPath *bool             `pulumi:"strip_path,optional" computed:"true" json:"strip_path,omitempty"`
	Exact     *bool             `pulumi:"exact,optional" computed:"true" json:"exact,omitempty"`
	Headers   map[string]string `pulumi:"headers,optional" computed:"true" json:"headers,omitempty"`
	Cookies   map[string]string `pulumi:"cookies,optional" computed:"true" json:"cookies,omitempty"`
	Query     map[string]string `pulumi:"query,optional" computed:"true" json:"query,omitempty"`
	Methods   []string          `pulumi:"methods,optional" computed:"true" json:"methods,omitempty"`
}

type BackendInputs struct {
	// TODO Targets
	Root    *string `pulumi:"root,optional" computed:"true" json:"root,omitempty"`
	Rewrite *bool   `pulumi:"rewrite,optional" computed:"true" json:"rewrite,omitempty"`
	// TODO LoadBalancing
	// TODO Client
	// TODO HealthCheck
}
