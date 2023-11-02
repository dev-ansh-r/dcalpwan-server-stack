package experimental

// Config is the configuration for experimental features.
type Config struct {
	Features []string `name:"features" description:"Experimental features to activate"`
}
