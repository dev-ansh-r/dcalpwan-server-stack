package mockis

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type mockEntityAccess struct {
	ttnpb.UnimplementedEntityAccessServer

	authInfoResponse *ttnpb.AuthInfoResponse
	err              error
}

func newEntityAccess() *mockEntityAccess {
	return &mockEntityAccess{}
}

func (ea *mockEntityAccess) AuthInfo(context.Context, *emptypb.Empty) (*ttnpb.AuthInfoResponse, error) {
	if ea.err != nil {
		return &ttnpb.AuthInfoResponse{}, ea.err
	}
	return ea.authInfoResponse, nil
}

func (ea *mockEntityAccess) SetAuthInfoResponse(authInfo *ttnpb.AuthInfoResponse) {
	ea.authInfoResponse = authInfo
}

func (ea *mockEntityAccess) SetErr(err error) {
	ea.err = err
}
