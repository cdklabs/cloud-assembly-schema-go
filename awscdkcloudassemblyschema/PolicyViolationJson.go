package awscdkcloudassemblyschema


// A single policy violation found by a validation plugin.
type PolicyViolationJson struct {
	// A description of the violation.
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name of the rule that was violated.
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// The severity of the violation.
	Severity *string `field:"required" json:"severity" yaml:"severity"`
	// Constructs that violated the rule.
	ViolatingConstructs *[]*ViolatingConstructJson `field:"required" json:"violatingConstructs" yaml:"violatingConstructs"`
	// If the plugin wants to report using a non-standard severity, put it here.
	CustomSeverity *string `field:"optional" json:"customSeverity" yaml:"customSeverity"`
	// Additional rule-specific metadata.
	// Default: - no metadata.
	//
	RuleMetadata *map[string]*string `field:"optional" json:"ruleMetadata" yaml:"ruleMetadata"`
	// How to fix the violation.
	// Default: - no fix provided.
	//
	SuggestedFix *string `field:"optional" json:"suggestedFix" yaml:"suggestedFix"`
}

