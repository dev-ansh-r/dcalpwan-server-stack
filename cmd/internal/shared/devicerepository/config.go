package devicerepository

import (
	"go.thethings.network/lorawan-stack/v3/pkg/devicerepository"
	"go.thethings.network/lorawan-stack/v3/pkg/devicerepository/store/bleve"
)

// DefaultDeviceRepositoryConfig is the default configuration for the Device Repository.
var DefaultDeviceRepositoryConfig = devicerepository.Config{
	Source:    "directory",
	Directory: "data/lorawan-devices",

	Store: devicerepository.StoreConfig{
		Bleve: bleve.Config{
			SearchPaths: []string{"lorawan-devices-index", "/srv/ttn-lorawan/lorawan-devices-index"},
		},
	},

	AssetsBaseURL: "https://raw.githubusercontent.com/TheThingsNetwork/lorawan-devices/master",
}
