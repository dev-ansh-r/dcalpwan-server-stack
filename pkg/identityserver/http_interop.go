package identityserver

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/identityserver/store"
	"go.thethings.network/lorawan-stack/v3/pkg/interop"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
)

type interopServer struct {
	*IdentityServer
	interop.Authorizer
}

func (srv *interopServer) HomeNSRequest(ctx context.Context, in *interop.HomeNSReq) (*interop.TTIHomeNSAns, error) {
	ctx = log.NewContextWithField(ctx, "namespace", "identityserver/interop")
	if err := srv.RequireAuthorized(ctx); err != nil {
		return nil, err
	}

	ids := &ttnpb.EndDeviceIdentifiers{
		JoinEui: (*types.EUI64)(&in.ReceiverID).Bytes(),
		DevEui:  (*types.EUI64)(&in.DevEUI).Bytes(),
	}

	var dev *ttnpb.EndDevice
	err := srv.store.Transact(ctx, func(ctx context.Context, st store.Store) (err error) {
		dev, err = st.GetEndDevice(ctx, ids, []string{"network_server_address"})
		return err
	})
	if err != nil {
		if errors.IsNotFound(err) {
			return nil, interop.ErrUnknownDevEUI.WithCause(err)
		}
		return nil, err
	}

	var (
		conf   = srv.configFromContext(ctx)
		hNetID = conf.Network.NetID
		hNSID  = conf.Network.NSID
	)

	header, err := in.AnswerHeader()
	if err != nil {
		return nil, interop.ErrMalformedMessage.WithCause(err)
	}
	ans := &interop.TTIHomeNSAns{
		HomeNSAns: interop.HomeNSAns{
			JsNsMessageHeader: interop.JsNsMessageHeader{
				MessageHeader: header,
				SenderID:      in.ReceiverID,
				ReceiverID:    in.SenderID,
				ReceiverNSID:  in.SenderNSID,
			},
			Result: interop.Result{
				ResultCode: interop.ResultSuccess,
			},
			HNetID: interop.NetID(hNetID),
		},
		TTIVSExtension: interop.TTIVSExtension{
			HNSAddress: dev.NetworkServerAddress,
			HTenantID:  conf.Network.TenantID,
		},
	}
	if in.ProtocolVersion.RequiresNSID() {
		ans.HNSID = (*interop.EUI64)(hNSID)
	}
	return ans, nil
}
