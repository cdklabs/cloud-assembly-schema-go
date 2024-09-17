package awscdkcloudassemblyschema


// Where to publish docker images.
type DockerImageDestination struct {
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
	// Tag of the image to publish.
	ImageTag *string `field:"required" json:"imageTag" yaml:"imageTag"`
	// Name of the ECR repository to publish to.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
}

