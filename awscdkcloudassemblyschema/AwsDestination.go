package awscdkcloudassemblyschema


// Destination for assets that need to be uploaded to AWS.
type AwsDestination struct {
	// Additional options to pass to STS when assuming the role.
	//
	// - `RoleArn` should not be used. Use the dedicated `assumeRoleArn` property instead.
	// - `ExternalId` should not be used. Use the dedicated `assumeRoleExternalId` instead.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/AWS/STS.html#assumeRole-property
	//
	// Default: - No additional options.
	//
	AssumeRoleAdditionalOptions *map[string]interface{} `field:"optional" json:"assumeRoleAdditionalOptions" yaml:"assumeRoleAdditionalOptions"`
	// The role that needs to be assumed while publishing this asset.
	// Default: - No role will be assumed.
	//
	AssumeRoleArn *string `field:"optional" json:"assumeRoleArn" yaml:"assumeRoleArn"`
	// The ExternalId that needs to be supplied while assuming this role.
	// Default: - No ExternalId will be supplied.
	//
	AssumeRoleExternalId *string `field:"optional" json:"assumeRoleExternalId" yaml:"assumeRoleExternalId"`
	// The region where this asset will need to be published.
	// Default: - Current region.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}

