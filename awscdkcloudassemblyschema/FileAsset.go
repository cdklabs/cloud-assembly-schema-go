package awscdkcloudassemblyschema


// A file asset.
type FileAsset struct {
	// Destinations for this file asset.
	Destinations *map[string]*FileDestination `field:"required" json:"destinations" yaml:"destinations"`
	// Source description for file assets.
	Source *FileSource `field:"required" json:"source" yaml:"source"`
	// A display name for this asset.
	// Default: - The identifier will be used as the display name.
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

