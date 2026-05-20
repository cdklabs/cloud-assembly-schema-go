package awscdkcloudassemblyschema


// The top-level structure of the policy validation report file.
type PolicyValidationReportJson struct {
	// Reports from all validation plugins that ran during synthesis.
	PluginReports *[]*PluginReportJson `field:"required" json:"pluginReports" yaml:"pluginReports"`
	// Protocol version.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Report title, if present.
	Title *string `field:"optional" json:"title" yaml:"title"`
}

