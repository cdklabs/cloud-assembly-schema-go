package awscdkcloudassemblyschema


// A file asset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cloud_assembly_schema "github.com/cdklabs/cloud-assembly-schema-go/awscdkcloudassemblyschema"
//
//   var assumeRoleAdditionalOptions interface{}
//
//   fileAsset := &FileAsset{
//   	Destinations: map[string]FileDestination{
//   		"destinationsKey": &FileDestination{
//   			"bucketName": jsii.String("bucketName"),
//   			"objectKey": jsii.String("objectKey"),
//
//   			// the properties below are optional
//   			"assumeRoleAdditionalOptions": map[string]interface{}{
//   				"assumeRoleAdditionalOptionsKey": assumeRoleAdditionalOptions,
//   			},
//   			"assumeRoleArn": jsii.String("assumeRoleArn"),
//   			"assumeRoleExternalId": jsii.String("assumeRoleExternalId"),
//   			"region": jsii.String("region"),
//   		},
//   	},
//   	Source: &FileSource{
//   		Executable: []*string{
//   			jsii.String("executable"),
//   		},
//   		Packaging: cloud_assembly_schema.FileAssetPackaging_FILE,
//   		Path: jsii.String("path"),
//   	},
//
//   	// the properties below are optional
//   	DisplayName: jsii.String("displayName"),
//   }
//
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

