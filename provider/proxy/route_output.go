package proxy

import "pulumi-otoroshi/provider/common"

type RouteOutput struct {
	common.BaseOutputStruct
	common.LocationOutputStruct
	Enabled         bool           `pulumi:"enabled" json:"enabled"`
	DebugFlow       bool           `pulumi:"debug_flow" json:"debug_flow"`
	ExportReporting bool           `pulumi:"export_reporting" json:"export_reporting"`
	Capture         bool           `pulumi:"capture" json:"capture"`
	Groups          []string       `pulumi:"groups" json:"groups"`
	BoundListeners  []string       `pulumi:"bound_listeners" json:"bound_listeners"`
	Frontend        FrontendOutput `pulumi:"frontend" json:"frontend"`
	Backend         BackendOutput  `pulumi:"backend" json:"backend"`
	BackendRef      *string        `pulumi:"backend_ref,optional" json:"backend_ref,omitempty"`
	Plugins         []PluginOutput `pulumi:"plugins" json:"plugins"` // TODO
}

type FrontendOutput struct {
	Domains   []string          `pulumi:"domains" json:"domains"`
	StripPath bool              `pulumi:"strip_path" json:"strip_path"`
	Exact     bool              `pulumi:"exact" json:"exact"`
	Headers   map[string]string `pulumi:"headers" json:"headers"`
	Cookies   map[string]string `pulumi:"cookies" json:"cookies"`
	Query     map[string]string `pulumi:"query" json:"query"`
	Methods   []string          `pulumi:"methods" json:"methods"`
}

type BackendOutput struct {
	Targets       []TargetOutput      `pulumi:"targets" json:"targets"`
	Root          string              `pulumi:"root" json:"root"`
	Rewrite       bool                `pulumi:"rewrite" json:"rewrite"`
	LoadBalancing LoadBalancingOutput `pulumi:"load_balancing" json:"load_balancing"`
	Client        ClientOutput        `pulumi:"client" json:"client"`
	HealthCheck   *HealthCheckOutput  `pulumi:"health_check,optional" json:"health_check,omitempty"`
}

type LoadBalancingOutput struct {
	Type  string `pulumi:"type" json:"type"`
	Ratio *int   `pulumi:"ratio,optional" json:"ratio,omitempty"`
}

type ClientOutput struct {
	Retries                 int                           `pulumi:"retries" json:"retries"`
	MaxErrors               int                           `pulumi:"max_errors" json:"max_errors"`
	RetryInitialDelay       int                           `pulumi:"retry_initial_delay" json:"retry_initial_delay"`
	BackoffFactor           int                           `pulumi:"backoff_factor" json:"backoff_factor"`
	CallTimeout             int                           `pulumi:"call_timeout" json:"call_timeout"`
	CallAndStreamTimeout    int                           `pulumi:"call_and_stream_timeout" json:"call_and_stream_timeout"`
	ConnectionTimeout       int                           `pulumi:"connection_timeout" json:"connection_timeout"`
	IdleTimeout             int                           `pulumi:"idle_timeout" json:"idle_timeout"`
	GlobalTimeout           int                           `pulumi:"global_timeout" json:"global_timeout"`
	SampleInterval          int                           `pulumi:"sample_interval" json:"sample_interval"`
	Proxy                   map[string]interface{}        `pulumi:"proxy" json:"proxy"`
	CustomTimeouts          []interface{}                 `pulumi:"custom_timeouts" json:"custom_timeouts"`
	CacheConnectionSettings CacheConnectionSettingsOutput `pulumi:"cache_connection_settings" json:"cache_connection_settings"`
}

type CacheConnectionSettingsOutput struct {
	Enabled   bool `pulumi:"enabled" json:"enabled"`
	QueueSize int  `pulumi:"queue_size" json:"queue_size"`
}

type HealthCheckOutput struct {
	Enabled           bool   `pulumi:"enabled" json:"enabled"`
	URL               string `pulumi:"url" json:"url"`
	Timeout           int    `pulumi:"timeout" json:"timeout"`
	HealthyStatuses   []int  `pulumi:"healthy_statuses" json:"healthyStatuses"`
	UnhealthyStatuses []int  `pulumi:"unhealthy_statuses" json:"unhealthyStatuses"`
}

type TargetOutput struct {
	Hostname  string          `pulumi:"hostname" json:"hostname"`
	Port      int             `pulumi:"port" json:"port"`
	TLS       bool            `pulumi:"tls" json:"tls"`
	Weight    int             `pulumi:"weight" json:"weight"`
	Backup    bool            `pulumi:"backup" json:"backup"`
	Predicate PredicateOutput `pulumi:"predicate" json:"predicate"`
	Protocol  string          `pulumi:"protocol" json:"protocol"`
	IPAddress *string         `pulumi:"ip_address" json:"ip_address"`
	TLSConfig TLSConfigOutput `pulumi:"tls_config" json:"tls_config"`
}

type PredicateOutput struct {
	Type string `pulumi:"type" json:"type"`
}

type TLSConfigOutput struct {
	Certs        []string `pulumi:"certs" json:"certs"`
	TrustedCerts []string `pulumi:"trusted_certs" json:"trusted_certs"`
	Enabled      bool     `pulumi:"enabled" json:"enabled"`
	Loose        bool     `pulumi:"loose" json:"loose"`
	TrustAll     bool     `pulumi:"trust_all" json:"trust_all"`
}

type PluginOutput struct {
	Enabled        bool                   `pulumi:"enabled" json:"enabled"`
	Debug          bool                   `pulumi:"debug" json:"debug"`
	Plugin         string                 `pulumi:"plugin" json:"plugin"`
	Include        []string               `pulumi:"include" json:"include"`
	Exclude        []string               `pulumi:"exclude" json:"exclude"`
	Config         map[string]interface{} `pulumi:"config" json:"config"`
	BoundListeners []string               `pulumi:"bound_listeners" json:"bound_listeners"`
	PluginIndex    map[string]int         `pulumi:"plugin_index" json:"plugin_index"`
}
