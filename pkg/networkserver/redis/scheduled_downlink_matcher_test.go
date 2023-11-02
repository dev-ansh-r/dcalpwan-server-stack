package redis_test

import (
	"testing"

	"go.thethings.network/lorawan-stack/v3/pkg/band"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/networkserver/redis"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestScheduledDownlinkMatcher(t *testing.T) {
	a, ctx := test.New(t)

	cl, flush := test.NewRedis(ctx, "redis_test")
	defer flush()
	defer cl.Close()

	m := redis.ScheduledDownlinkMatcher{cl}

	ids := &ttnpb.EndDeviceIdentifiers{
		ApplicationIds: &ttnpb.ApplicationIdentifiers{
			ApplicationId: "app1",
		},
		DeviceId: "dev1",
	}

	stored := &ttnpb.DownlinkMessage{
		RawPayload:   []byte{1, 2, 3},
		EndDeviceIds: ids,
		Settings: &ttnpb.DownlinkMessage_Request{
			Request: &ttnpb.TxRequest{
				Class: ttnpb.Class_CLASS_A,
			},
		},
		CorrelationIds: []string{"corr1", "corr2", "ns:transmission:CORRELATIONID"},
	}

	ack := &ttnpb.TxAcknowledgment{
		Result: ttnpb.TxAcknowledgment_SUCCESS,
		DownlinkMessage: &ttnpb.DownlinkMessage{
			Settings: &ttnpb.DownlinkMessage_Scheduled{
				Scheduled: &ttnpb.TxSettings{
					DataRate: &ttnpb.DataRate{
						Modulation: &ttnpb.DataRate_Lora{
							Lora: &ttnpb.LoRaDataRate{
								SpreadingFactor: 7,
								Bandwidth:       125000,
								CodingRate:      band.Cr4_5,
							},
						},
					},
				},
			},
			CorrelationIds: []string{"corr1", "corr2", "ns:transmission:CORRELATIONID"},
		},
	}

	err := m.Add(test.Context(), stored)
	a.So(err, should.BeNil)

	t.Run("MissingCorrelationID", func(t *testing.T) {
		a, ctx := test.New(t)
		down, err := m.Match(ctx, &ttnpb.TxAcknowledgment{})
		a.So(errors.IsNotFound(err), should.BeTrue)
		a.So(down, should.BeNil)
	})

	t.Run("InvalidCorrelationID", func(t *testing.T) {
		a, ctx := test.New(t)
		down, err := m.Match(ctx, &ttnpb.TxAcknowledgment{
			DownlinkMessage: &ttnpb.DownlinkMessage{
				CorrelationIds: []string{"ns:transmission:OTHERCORRELATIONID"},
			},
		})
		a.So(errors.IsNotFound(err), should.BeTrue)
		a.So(down, should.BeNil)
	})

	t.Run("Match", func(t *testing.T) {
		a, ctx := test.New(t)
		down, err := m.Match(ctx, ack)
		a.So(err, should.BeNil)
		a.So(down, should.Resemble, stored)
	})

	t.Run("DoNotMatchTwice", func(t *testing.T) {
		a, ctx := test.New(t)
		down, err := m.Match(ctx, ack)
		a.So(errors.IsNotFound(err), should.BeTrue)
		a.So(down, should.BeNil)
	})
}
