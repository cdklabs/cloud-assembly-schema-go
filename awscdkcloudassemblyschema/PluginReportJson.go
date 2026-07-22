package awscdkcloudassemblyschema


// A report from a single validation plugin.
//
// Example:
//   import "github.com/cdklabs/cloud-assembly-schema-go/awscdkcloudassemblyschema"
//
//
//   report := &PluginReportJson{
//   	PluginName: jsii.String("my-plugin"),
//   	Conclusion: jsii.String("success"),
//   	Violations: []PolicyViolationJson{
//   	},
//   }
//
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
	// Any preamble text that the plugin wants to include in the report that is not a violation itself.
	// Default: - no preamble.
	//
	Preamble *string `field:"optional" json:"preamble" yaml:"preamble"`
	// Violations that were suppressed via acknowledgement.
	//
	// These violations matched an acknowledged rule ID and were excluded
	// from the active violations list. They are retained for audit
	// trail and reporting purposes.
	// Default: - no suppressed violations.
	//
	SuppressedViolations *[]*SuppressedViolationJson `field:"optional" json:"suppressedViolations" yaml:"suppressedViolations"`
}

