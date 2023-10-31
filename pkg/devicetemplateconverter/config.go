// Package devicetemplateconverter provides device template services.
package devicetemplateconverter

// Config represents the DeviceTemplateConverter configuration.
type Config struct {
	Enabled []string `name:"enabled" description:"Enabled converters"`
}
