package awscdkcloudassemblyschema


// A single feature flag.
type FeatureFlag struct {
	// Explanation about the purpose of this flag that can be shown to the user.
	// Default: - No description.
	//
	Explanation *string `field:"optional" json:"explanation" yaml:"explanation"`
	// The library-recommended value for this flag, if any.
	//
	// It is possible that there is no recommended value.
	// Default: - No recommended value.
	//
	RecommendedValue interface{} `field:"optional" json:"recommendedValue" yaml:"recommendedValue"`
	// The value of the flag that produces the same behavior as when the flag is not configured at all.
	//
	// The structure of this field is a historical accident. The type of this field
	// should have been boolean, which should have contained the default value for
	// the flag appropriate for the *current* version of the CDK library. We are
	// not rectifying this accident because doing so
	//
	// Instead, the canonical way to access this value is by evaluating
	// `unconfiguredBehavesLike?.v2 ?? false`.
	// Default: false.
	//
	UnconfiguredBehavesLike *UnconfiguredBehavesLike `field:"optional" json:"unconfiguredBehavesLike" yaml:"unconfiguredBehavesLike"`
	// The value configured by the user.
	//
	// This is the value configured at the root of the tree. Users may also have
	// configured values at specific locations in the tree; we don't report on
	// those.
	// Default: - Not configured by the user.
	//
	UserValue interface{} `field:"optional" json:"userValue" yaml:"userValue"`
}

