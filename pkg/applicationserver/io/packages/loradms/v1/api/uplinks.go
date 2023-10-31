package api

import (
	"context"
	"net/http"

	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io/packages/loradms/v1/api/objects"
)

// Uplinks is an API client for the Uplink API.
type Uplinks struct {
	cl *Client
}

const uplinkEntity = "uplink"

// Send sends the given uplink to the Device Management service.
func (u *Uplinks) Send(ctx context.Context, uplinks objects.DeviceUplinks) (objects.DeviceUplinkResponses, error) {
	resp, err := u.cl.Do(ctx, http.MethodPost, uplinkEntity, "", sendOperation, uplinks)
	if err != nil {
		return nil, err
	}
	response := make(objects.DeviceUplinkResponses)
	err = parse(&response, resp)
	if err != nil {
		return nil, err
	}
	return response, nil
}
