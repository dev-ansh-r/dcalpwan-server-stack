package config

// Directories
type Dir struct {
	DeviceRepo DeviceRepo
	Hugo       Hugo
}

type DeviceRepo struct {
	Templates string
	Vendor    string
}

type Hugo struct {
	Devices string
}
