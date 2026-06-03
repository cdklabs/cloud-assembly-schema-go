package awscdkcloudassemblyschema


// Artifact properties for the Construct Tree Artifact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cloud_assembly_schema "github.com/cdklabs/cloud-assembly-schema-go/awscdkcloudassemblyschema"
//
//   treeArtifactProperties := &TreeArtifactProperties{
//   	File: jsii.String("file"),
//   }
//
type TreeArtifactProperties struct {
	// Filename of the tree artifact.
	File *string `field:"required" json:"file" yaml:"file"`
}

