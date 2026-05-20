package awscdkcloudassemblyschema


// A construct that violated a policy rule.
type ViolatingConstructJson struct {
	// The construct path as defined in the application.
	// Default: - no construct path.
	//
	ConstructPath *string `field:"required" json:"constructPath" yaml:"constructPath"`
	// If this construct violation regards a CloudFormation resource, a reference to the resource details.
	CloudFormationResource *CloudFormationResourceJson `field:"optional" json:"cloudFormationResource" yaml:"cloudFormationResource"`
	// The fully qualified name of the construct class (includes the library name).
	// Default: - no construct info.
	//
	ConstructFqn *string `field:"optional" json:"constructFqn" yaml:"constructFqn"`
	// The version of the library that contains this construct.
	//
	// The library name is the first component of the construct FQN.
	// Default: - no version info.
	//
	LibraryVersion *string `field:"optional" json:"libraryVersion" yaml:"libraryVersion"`
	// Stack traces associated with this violation.
	//
	// This can be all the stack traces where a violating property got its value,
	// or just the construct creation stack trace.
	//
	// Every element of the array is a stack trace, where each stack trace is a `\n`-delimited string.
	// Default: - No stack traces.
	//
	StackTraces *[]*string `field:"optional" json:"stackTraces" yaml:"stackTraces"`
}

