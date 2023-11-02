package mac

import (
	"go.thethings.network/lorawan-stack/v3/pkg/band"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

func legacyDeviceUseADR(dev *ttnpb.EndDevice, defaults *ttnpb.MACSettings, phy *band.Band) bool {
	switch {
	case !phy.SupportsDynamicADR:
		return false
	case dev.GetMulticast():
		return false
	case dev.GetMacSettings().GetUseAdr() != nil:
		return dev.MacSettings.UseAdr.Value
	case defaults.UseAdr != nil:
		return defaults.UseAdr.Value
	default:
		return true
	}
}

func legacyDeviceShouldAdaptDataRate(dev *ttnpb.EndDevice, defaults *ttnpb.MACSettings, phy *band.Band) (adaptDataRate bool, resetDesiredParameters bool, staticSettings *ttnpb.ADRSettings_StaticMode) {
	useADR := legacyDeviceUseADR(dev, defaults, phy)
	return useADR, useADR, nil
}

func legacyDeviceADRMargin(dev *ttnpb.EndDevice, defaults *ttnpb.MACSettings) float32 {
	switch {
	case dev.GetMacSettings().GetAdrMargin() != nil:
		return dev.MacSettings.AdrMargin.Value
	case defaults.GetAdrMargin() != nil:
		return defaults.AdrMargin.Value
	default:
		return DefaultADRMargin
	}
}
