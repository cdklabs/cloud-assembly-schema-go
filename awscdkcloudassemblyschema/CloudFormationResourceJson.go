package awscdkcloudassemblyschema


// A node in the construct creation stack trace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cloud_assembly_schema "github.com/cdklabs/cloud-assembly-schema-go/awscdkcloudassemblyschema"
//
//   cloudFormationResourceJson := &CloudFormationResourceJson{
//   	LogicalId: jsii.String("logicalId"),
//   	TemplatePath: jsii.String("templatePath"),
//
//   	// the properties below are optional
//   	PropertyPaths: []*string{
//   		jsii.String("propertyPaths"),
//   	},
//   }
//
type CloudFormationResourceJson struct {
	// The logical ID of the resource in the CloudFormation template.
	LogicalId *string `field:"required" json:"logicalId" yaml:"logicalId"`
	// The path to the CloudFormation template containing this resource.
	//
	// This path is relative to the Cloud Assembly root directory.
	TemplatePath *string `field:"required" json:"templatePath" yaml:"templatePath"`
	// Properties within the construct where the violation was detected.
	//
	// Either a single component, in which case it regards a top-level property
	// name, or a JSON path (starting with `$.`) to indicate a deeper property.
	// Default: - no locations.
	//
	PropertyPaths *[]*string `field:"optional" json:"propertyPaths" yaml:"propertyPaths"`
}

