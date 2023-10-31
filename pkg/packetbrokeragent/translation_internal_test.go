package packetbrokeragent

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"testing"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/unique"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func WrapUplinkTokens(gateway, forwarder []byte, agent *ttnpb.PacketBrokerAgentUplinkToken) ([]byte, error) {
	return wrapUplinkTokens(gateway, forwarder, agent)
}

func TestWrapGatewayUplinkToken(t *testing.T) {
	a, ctx := test.New(t)
	key := bytes.Repeat([]byte{0x42}, 16)
	forwarderData := []byte("000013:tnt:eu1")
	blockCipher, err := aes.NewCipher(key)
	if !a.So(err, should.BeNil) {
		t.FailNow()
	}
	aead, err := cipher.NewGCM(blockCipher)
	if !a.So(err, should.BeNil) {
		t.FailNow()
	}

	wrappedToken, err := wrapGatewayUplinkToken(
		ctx,
		&ttnpb.GatewayIdentifiers{GatewayId: "test-gateway"},
		[]byte{0x1, 0x2, 0x3},
		forwarderData,
		aead,
	)
	if !a.So(err, should.BeNil) {
		t.FailNow()
	}
	t.Logf("Wrapped token: %q", base64.RawStdEncoding.EncodeToString(wrappedToken))

	uid, gtwToken, err := unwrapGatewayUplinkToken(wrappedToken, forwarderData, aead, key)
	if !a.So(err, should.BeNil) {
		t.FailNow()
	}
	a.So(uid, should.Resemble, unique.ID(ctx, &ttnpb.GatewayIdentifiers{GatewayId: "test-gateway"}))
	a.So(gtwToken, should.Resemble, []byte{0x1, 0x2, 0x3})
}

func TestFrequencyPlan(t *testing.T) {
	a, _ := test.New(t)

	fp, err := toPBFrequencyPlan(test.FrequencyPlan("EU_863_870"))
	a.So(err, should.BeNil)
	a.So(fp.LoraSingleSfChannels, should.BeEmpty)
	a.So(fp.FskChannel, should.BeNil)
	a.So(fp.LoraMultiSfChannels, should.HaveLength, 8)
}
