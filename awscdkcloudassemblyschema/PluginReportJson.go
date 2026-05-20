package awscdkcloudassemblyschema


// A report from a single validation plugin.
type PluginReportJson struct {
	// Whether the plugin's validation passed or failed.
	Conclusion *string `field:"required" json:"conclusion" yaml:"conclusion"`
	// The name of the plugin that produced this report.
	PluginName *string `field:"required" json:"pluginName" yaml:"pluginName"`
	// Violations found by this plugin.
	Violations *[]*PolicyViolationJson `field:"required" json:"violations" yaml:"violations"`
	// Additional plugin-specific metadata.
	// Default: - no metadata.
	//
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// Version of the plugin that produced this report.
	// Default: - no version.
	//
	PluginVersion *string `field:"optional" json:"pluginVersion" yaml:"pluginVersion"`
}

