package lastseen

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/metadata"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

func batchUpdateEndDeviceLastSeen(ctx context.Context, endDevices map[string]*ttnpb.EndDevice, cls metadata.ClusterPeerAccess) error {
	conn, err := cls.GetPeerConn(ctx, ttnpb.ClusterRole_ENTITY_REGISTRY, nil)
	if err != nil {
		log.FromContext(ctx).WithError(err).Warn("Failed to get Identity Server peer")
		return err
	}
	cl := ttnpb.NewEndDeviceRegistryClient(conn)
	deviceLastSeenList := make([]*ttnpb.BatchUpdateEndDeviceLastSeenRequest_EndDeviceLastSeenUpdate, 0, len(endDevices))
	for _, dev := range endDevices {
		deviceLastSeenList = append(deviceLastSeenList, &ttnpb.BatchUpdateEndDeviceLastSeenRequest_EndDeviceLastSeenUpdate{
			Ids:        dev.Ids,
			LastSeenAt: dev.LastSeenAt,
		})
	}
	_, err = cl.BatchUpdateLastSeen(ctx, &ttnpb.BatchUpdateEndDeviceLastSeenRequest{
		Updates: deviceLastSeenList,
	}, cls.WithClusterAuth())
	return err
}
