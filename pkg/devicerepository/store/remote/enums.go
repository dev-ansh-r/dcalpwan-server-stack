package remote

import (
	"go.thethings.network/lorawan-stack/v3/pkg/band"
)

// regionToBandID maps LoRaWAN schema regions to TTS Band IDs.
var regionToBandID = map[string]string{
	"EU863-870":      band.EU_863_870,
	"US902-928":      band.US_902_928,
	"CN779-787":      band.CN_779_787,
	"EU433":          band.EU_433,
	"AU915-928":      band.AU_915_928,
	"CN470-510":      band.CN_470_510,
	"CN470-510-20-A": band.CN_470_510_20_A,
	"CN470-510-20-B": band.CN_470_510_20_B,
	"CN470-510-26-A": band.CN_470_510_26_A,
	"CN470-510-26-B": band.CN_470_510_26_B,
	"AS923":          band.AS_923,
	"AS923-2":        band.AS_923_2,
	"AS923-3":        band.AS_923_3,
	"AS923-4":        band.AS_923_4,
	"KR920-923":      band.KR_920_923,
	"IN865-867":      band.IN_865_867,
	"RU864-870":      band.RU_864_870,
}

// bandIDToRegion is the inverse mapping of regionToBandID.
var bandIDToRegion map[string]string

func init() {
	bandIDToRegion = make(map[string]string, len(regionToBandID))
	for region, bandID := range regionToBandID {
		bandIDToRegion[bandID] = region
	}
}
