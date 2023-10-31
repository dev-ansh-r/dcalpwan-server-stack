package webui

import "encoding/json"

// APIConfig for upstream APIs.
type APIConfig struct {
	Enabled bool   `json:"enabled" name:"enabled" description:"Enable this API"`
	BaseURL string `json:"base_url" name:"base-url" description:"Base URL to the HTTP API"`
}

// MarshalJSON implements json.Marshaler.
func (c APIConfig) MarshalJSON() ([]byte, error) {
	out := struct {
		Enabled bool   `json:"enabled"`
		BaseURL string `json:"base_url,omitempty"`
	}{
		Enabled: c.Enabled,
		BaseURL: c.BaseURL,
	}
	if !out.Enabled {
		out.BaseURL = ""
	}
	return json.Marshal(out)
}
