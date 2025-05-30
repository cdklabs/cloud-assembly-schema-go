package awscdkcloudassemblyschema


// Options for configuring the Docker cache backend.
type DockerCacheOption struct {
	// The type of cache to use.
	//
	// Refer to https://docs.docker.com/build/cache/backends/ for full list of backends.
	//
	// Example:
	//   "registry"
	//
	// Default: - unspecified.
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Any parameters to pass into the docker cache backend configuration.
	//
	// Refer to https://docs.docker.com/build/cache/backends/ for cache backend configuration.
	//
	// Example:
	//   var branch string
	//
	//
	//   params := map[string]interface{}{
	//   	"ref": fmt.Sprintf("12345678.dkr.ecr.us-west-2.amazonaws.com/cache:%v", branch),
	//   	"mode": jsii.String("max"),
	//   }
	//
	// Default: {} No options provided.
	//
	Params *map[string]*string `field:"optional" json:"params" yaml:"params"`
}

