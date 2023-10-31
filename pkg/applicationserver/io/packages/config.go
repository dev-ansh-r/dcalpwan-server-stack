package packages

import "time"

// Config contains configuration options for application packages.
type Config struct {
	Workers int           `name:"workers" description:"Number of workers per application package"`
	Timeout time.Duration `name:"timeout" description:"Message processing timeout"`
}
