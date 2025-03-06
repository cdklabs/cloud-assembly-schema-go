package awscdkcloudassemblyschema


// A file asset.
type DockerImageAsset struct {
	// Destinations for this container asset.
	Destinations *map[string]*DockerImageDestination `field:"required" json:"destinations" yaml:"destinations"`
	// Source description for container assets.
	Source *DockerImageSource `field:"required" json:"source" yaml:"source"`
	// A display name for this asset.
	// Default: - The identifier will be used as the display name.
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

