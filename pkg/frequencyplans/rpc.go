package frequencyplans

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// RPCServer is the RPC server that serves frequency plan information.
type RPCServer struct {
	store *Store
}

// NewRPCServer returns a new RPC server that serves frequency plan information.
func NewRPCServer(store *Store) *RPCServer { return &RPCServer{store: store} }

// ListFrequencyPlans lists frequency plans for the requested base frequency.
func (s *RPCServer) ListFrequencyPlans(ctx context.Context, req *ttnpb.ListFrequencyPlansRequest) (*ttnpb.ListFrequencyPlansResponse, error) {
	descriptions, err := s.store.descriptions()
	if err != nil {
		return nil, err
	}
	res := &ttnpb.ListFrequencyPlansResponse{}
	for _, desc := range descriptions {
		if req.BaseFrequency != 0 && uint16(req.BaseFrequency) != desc.BaseFrequency {
			continue
		}
		res.FrequencyPlans = append(res.FrequencyPlans, &ttnpb.FrequencyPlanDescription{
			Id:            desc.ID,
			BaseId:        desc.BaseID,
			Name:          desc.Name,
			BaseFrequency: uint32(desc.BaseFrequency),
		})
	}
	return res, nil
}
