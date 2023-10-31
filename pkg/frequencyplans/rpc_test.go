package frequencyplans_test

import (
	"context"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/fetch"
	"go.thethings.network/lorawan-stack/v3/pkg/frequencyplans"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestRPCServer(t *testing.T) {
	a := assertions.New(t)

	store := frequencyplans.NewStore(fetch.NewMemFetcher(map[string][]byte{
		"frequency-plans.yml": []byte(`- id: A
  description: Frequency Plan A
  base-frequency: 868
  file: A.yml
- id: B
  base-id: A
  description: Frequency Plan B
  file: B.yml
- id: C
  description: Frequency Plan C
  base-frequency: 915
  file: C.yml`),
	}))

	server := frequencyplans.NewRPCServer(store)

	all, err := server.ListFrequencyPlans(context.Background(), &ttnpb.ListFrequencyPlansRequest{})
	a.So(err, should.BeNil)
	a.So(all.FrequencyPlans, should.HaveLength, 3)

	base915, err := server.ListFrequencyPlans(context.Background(), &ttnpb.ListFrequencyPlansRequest{
		BaseFrequency: 868,
	})
	a.So(err, should.BeNil)
	a.So(base915.FrequencyPlans, should.HaveLength, 2)
}
