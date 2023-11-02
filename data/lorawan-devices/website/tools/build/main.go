package main

import (
	"log"

	"github.com/TheThingsNetwork/lorawan-devices/website/tools/build/pkg/config"
	"github.com/TheThingsNetwork/lorawan-devices/website/tools/build/pkg/devicerepository"
	"github.com/TheThingsNetwork/lorawan-devices/website/tools/build/pkg/hugo"
)

func main() {
	dir := config.Dir{
		DeviceRepo: config.DeviceRepo{
			Templates: "templates/",
			Vendor:    "../../../vendor/",
		},
		Hugo: config.Hugo{
			Devices: "../../content/devices/",
		},
	}

	dr, err := devicerepository.NewStore(dir)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = hugo.CreateContentSingleDevice(dir, dr)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = hugo.CreateContentDevices(dir)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
