package awscdkcloudassemblyschema


// Artifact properties for a feature flag report.
//
// A feature flag report is small enough that all the properties can be inlined
// here, and doesn't need an additional file.
type FeatureFlagReportProperties struct {
	// Information about every feature flag supported by this library.
	Flags *map[string]*FeatureFlag `field:"required" json:"flags" yaml:"flags"`
	// The library that this feature flag report applies to.
	Module *string `field:"required" json:"module" yaml:"module"`
}

