package awscdkcloudassemblyschema


// A manifest for a single artifact within the cloud assembly.
type ArtifactManifest struct {
	// The type of artifact.
	Type ArtifactType `field:"required" json:"type" yaml:"type"`
	// A file with additional metadata entries.
	//
	// The schema of this file is exactly the same as the type of the `metadata` field.
	// In other words, that file contains an object mapping construct paths to arrays
	// of metadata entries.
	// Default: - no additional metadata.
	//
	AdditionalMetadataFile *string `field:"optional" json:"additionalMetadataFile" yaml:"additionalMetadataFile"`
	// IDs of artifacts that must be deployed before this artifact.
	// Default: - no dependencies.
	//
	Dependencies *[]*string `field:"optional" json:"dependencies" yaml:"dependencies"`
	// A string that can be shown to a user to uniquely identify this artifact inside a cloud assembly tree.
	//
	// Is used by the CLI to present a list of stacks to the user in a way that
	// makes sense to them. Even though the property name "display name" doesn't
	// imply it, this field is used to select stacks as well, so all stacks should
	// have a unique display name.
	// Default: - no display name.
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The environment into which this artifact is deployed.
	// Default: - no envrionment.
	//
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// Associated metadata.
	//
	// Metadata can be stored directly in the assembly manifest, as well as in a
	// separate file (see `additionalMetadataFile`). It should prefer to be stored
	// in the additional file, as that will reduce the size of the assembly
	// manifest in cases of a lot of metdata (which CDK does emit by default).
	// Default: - no metadata.
	//
	Metadata *map[string]*[]*MetadataEntry `field:"optional" json:"metadata" yaml:"metadata"`
	// The set of properties for this artifact (depends on type).
	// Default: - no properties.
	//
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
}

