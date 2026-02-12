package awscdkcloudassemblyschema

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cloud-assembly-schema-go/awscdkcloudassemblyschema/v52/jsii"
)

// Protocol utility class.
type Manifest interface {
}

// The jsii proxy struct for Manifest
type jsiiProxy_Manifest struct {
	_ byte // padding
}

func NewManifest_Override(m Manifest) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cloud-assembly-schema.Manifest",
		nil, // no parameters
		m,
	)
}

// Return the CLI version that supports this Cloud Assembly Schema version.
func Manifest_CliVersion() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"@aws-cdk/cloud-assembly-schema.Manifest",
		"cliVersion",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Deprecated.
// Deprecated: use `loadAssemblyManifest()`.
func Manifest_Load(filePath *string) *AssemblyManifest {
	_init_.Initialize()

	if err := validateManifest_LoadParameters(filePath); err != nil {
		panic(err)
	}
	var returns *AssemblyManifest

	_jsii_.StaticInvoke(
		"@aws-cdk/cloud-assembly-schema.Manifest",
		"load",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Load and validates the cloud assembly manifest from file.
func Manifest_LoadAssemblyManifest(filePath *string, options *LoadManifestOptions) *AssemblyManifest {
	_init_.Initialize()

	if err := validateManifest_LoadAssemblyManifestParameters(filePath, options); err != nil {
		panic(err)
	}
	var returns *AssemblyManifest

	_jsii_.StaticInvoke(
		"@aws-cdk/cloud-assembly-schema.Manifest",
		"loadAssemblyManifest",
		[]interface{}{filePath, options},
		&returns,
	)

	return returns
}

// Load and validates the asset manifest from file.
func Manifest_LoadAssetManifest(filePath *string) *AssetManifest {
	_init_.Initialize()

	if err := validateManifest_LoadAssetManifestParameters(filePath); err != nil {
		panic(err)
	}
	var returns *AssetManifest

	_jsii_.StaticInvoke(
		"@aws-cdk/cloud-assembly-schema.Manifest",
		"loadAssetManifest",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Load and validates the integ manifest from file.
func Manifest_LoadIntegManifest(filePath *string) *IntegManifest {
	_init_.Initialize()

	if err := validateManifest_LoadIntegManifestParameters(filePath); err != nil {
		panic(err)
	}
	var returns *IntegManifest

	_jsii_.StaticInvoke(
		"@aws-cdk/cloud-assembly-schema.Manifest",
		"loadIntegManifest",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Deprecated.
// Deprecated: use `saveAssemblyManifest()`.
func Manifest_Save(manifest *AssemblyManifest, filePath *string) {
	_init_.Initialize()

	if err := validateManifest_SaveParameters(manifest, filePath); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"@aws-cdk/cloud-assembly-schema.Manifest",
		"save",
		[]interface{}{manifest, filePath},
	)
}

// Validates and saves the cloud assembly manifest to file.
func Manifest_SaveAssemblyManifest(manifest *AssemblyManifest, filePath *string) {
	_init_.Initialize()

	if err := validateManifest_SaveAssemblyManifestParameters(manifest, filePath); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"@aws-cdk/cloud-assembly-schema.Manifest",
		"saveAssemblyManifest",
		[]interface{}{manifest, filePath},
	)
}

// Validates and saves the asset manifest to file.
func Manifest_SaveAssetManifest(manifest *AssetManifest, filePath *string) {
	_init_.Initialize()

	if err := validateManifest_SaveAssetManifestParameters(manifest, filePath); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"@aws-cdk/cloud-assembly-schema.Manifest",
		"saveAssetManifest",
		[]interface{}{manifest, filePath},
	)
}

// Validates and saves the integ manifest to file.
func Manifest_SaveIntegManifest(manifest *IntegManifest, filePath *string) {
	_init_.Initialize()

	if err := validateManifest_SaveIntegManifestParameters(manifest, filePath); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"@aws-cdk/cloud-assembly-schema.Manifest",
		"saveIntegManifest",
		[]interface{}{manifest, filePath},
	)
}

// Fetch the current schema version number.
func Manifest_Version() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"@aws-cdk/cloud-assembly-schema.Manifest",
		"version",
		nil, // no parameters
		&returns,
	)

	return returns
}

