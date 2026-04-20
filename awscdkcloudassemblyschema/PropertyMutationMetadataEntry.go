package awscdkcloudassemblyschema


// Metadata type of a PropertyMutation.
type PropertyMutationMetadataEntry struct {
	// Name of the property.
	PropertyName *string `field:"required" json:"propertyName" yaml:"propertyName"`
	// Stack trace of the mutation.
	StackTrace *[]*string `field:"required" json:"stackTrace" yaml:"stackTrace"`
}

