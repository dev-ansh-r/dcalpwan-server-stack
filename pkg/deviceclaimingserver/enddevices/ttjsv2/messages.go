package ttjsv2

// ClaimData contains information about the claim.
type ClaimData struct {
	HomeNetID string  `json:"homeNetID"`
	HomeNSID  *string `json:"homeNSID,omitempty"`
	Locked    bool    `json:"locked"`
}

// KEK contains the key encryption key.
type KEK struct {
	Label string `json:"label"`
	Key   string `json:"key"`
}

// ClaimRequest is the claim request message.
type ClaimRequest struct {
	OwnerToken           string  `json:"ownerToken"`
	HomeNetID            string  `json:"homeNetID"`
	HomeNSID             *string `json:"homeNSID,omitempty"`
	ASID                 string  `json:"asID"`
	RegenerateOwnerToken *bool   `json:"regenerateOwnerToken,omitempty"`
	Lock                 *bool   `json:"lock,omitempty"`
	KEK                  *KEK    `json:"kek,omitempty"`
}

// ErrorResponse is a message that may be returned by The Things Join Server in case of an error.
type ErrorResponse struct {
	Message string `json:"message"`
}

// EndDevicesErrors is a map of end device identifiers to error responses.
type EndDevicesErrors map[string]ErrorResponse

func boolValue(b bool) *bool {
	return &b
}

func stringValue(s string) *string {
	return &s
}
