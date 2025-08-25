package events

/*type Filtering struct {
	Include []map[string]any `pulumi:"include,optional" json:"include,omitempty"`
	Exclude []map[string]any `pulumi:"exclude,optional" json:"exclude,omitempty"`
}

type DataExporter struct {
	common.GenericResource[
		DataExporter,
		struct {
			Location      *common.LocationInputs `pulumi:"_loc,optional" json:"_loc,omitempty" computed:"true"`
			Name          string                 `pulumi:"name" json:"name" computed:"true"`
			Type          string                 `pulumi:"type" json:"type" computed:"true"`
			Enabled       *bool                  `pulumi:"enabled,optional" json:"enabled,omitempty" computed:"true"`
			Description   *string                `pulumi:"desc,optional" json:"desc,omitempty" computed:"true"`
			Tags          []string               `pulumi:"tags,optional" json:"tags,omitempty" computed:"true"`
			Metadata      map[string]string      `pulumi:"metadata,optional" json:"metadata,omitempty" computed:"true"`
			BufferSize    *int                   `pulumi:"bufferSize,optional" json:"bufferSize,omitempty" computed:"true"`
			JSONWorkers   *int                   `pulumi:"jsonWorkers,optional" json:"jsonWorkers,omitempty" computed:"true"`
			SendWorkers   *int                   `pulumi:"sendWorkers,optional" json:"sendWorkers,omitempty" computed:"true"`
			GroupSize     *int                   `pulumi:"groupSize,optional" json:"groupSize,omitempty" computed:"true"`
			GroupDuration *int                   `pulumi:"groupDuration,optional" json:"groupDuration,omitempty" computed:"true"`
			Projection    map[string]any         `pulumi:"projection,optional" json:"projection,omitempty" computed:"true"`
			Filtering     *Filtering             `pulumi:"filtering,optional" json:"filtering,omitempty" computed:"true"`
			Config        map[string]any         `pulumi:"config,optional" json:"config,omitempty" computed:"true"`
		},
		struct {
			Location      *common.LocationOutput `pulumi:"_loc,optional" json:"_loc,omitempty" computed:"true"`
			Name          *string                `pulumi:"name,optional" json:"name,omitempty" computed:"true"`
			Type          *string                `pulumi:"type,optional" json:"type,omitempty" computed:"true"`
			Enabled       *bool                  `pulumi:"enabled,optional" json:"enabled,omitempty" computed:"true"`
			Description   *string                `pulumi:"desc,optional" json:"desc,omitempty" computed:"true"`
			Tags          []string               `pulumi:"tags,optional" json:"tags,omitempty" computed:"true"`
			Metadata      map[string]string      `pulumi:"metadata,optional" json:"metadata,omitempty" computed:"true"`
			BufferSize    *int                   `pulumi:"bufferSize,optional" json:"bufferSize,omitempty" computed:"true"`
			JSONWorkers   *int                   `pulumi:"jsonWorkers,optional" json:"jsonWorkers,omitempty" computed:"true"`
			SendWorkers   *int                   `pulumi:"sendWorkers,optional" json:"sendWorkers,omitempty" computed:"true"`
			GroupSize     *int                   `pulumi:"groupSize,optional" json:"groupSize,omitempty" computed:"true"`
			GroupDuration *int                   `pulumi:"groupDuration,optional" json:"groupDuration,omitempty" computed:"true"`
			Projection    map[string]any         `pulumi:"projection,optional" json:"projection,omitempty" computed:"true"`
			Filtering     *Filtering             `pulumi:"filtering,optional" json:"filtering,omitempty" computed:"true"`
			Config        map[string]any         `pulumi:"config,optional" json:"config,omitempty" computed:"true"`
		},
	]
}

func (DataExporter) PathTemplate() string {
	return "/apis/events.otoroshi.io/v1/data-exporters"
}*/
