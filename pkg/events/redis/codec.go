package redis

import (
	"encoding/base64"
	"strings"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/unique"
	"google.golang.org/protobuf/proto"
)

var binaryEncoding = base64.RawStdEncoding

func encodeBinary(b []byte) string {
	return binaryEncoding.EncodeToString(b)
}

func decodeBinary(s string) ([]byte, error) {
	var found bool
	if s, found = strings.CutSuffix(s, "="); found {
		s = strings.TrimSuffix(s, "=")
	}
	return binaryEncoding.DecodeString(s)
}

func encodeEventData(evt events.Event) (string, error) {
	pb, err := events.Proto(evt)
	if err != nil {
		return "", err
	}
	bpb, err := proto.Marshal(pb)
	if err != nil {
		return "", err
	}
	return protoEncodingPrefix + encodeBinary(bpb), nil
}

var errUnknownEncoding = errors.DefineInvalidArgument("unknown_encoding", "unknown encoding")

func decodeEventData(enc string, evt *ttnpb.Event) error {
	if !strings.HasPrefix(enc, protoEncodingPrefix) {
		return errUnknownEncoding.New()
	}
	bpb, err := decodeBinary(strings.TrimPrefix(enc, protoEncodingPrefix))
	if err != nil {
		return err
	}
	// NOTE: We override the contents of `evt` with the contents `enc` explicitly.
	// proto.UnmarshalMerge has append semantics for slices, and as such identifiers
	// and event visibility will be duplicated, as `evt` most likely contains them
	// from the sparse event representation.
	return proto.Unmarshal(bpb, evt)
}

const (
	eventUIDKey    = "event_uid"
	eventNameKey   = "event_name"
	entityIDKey    = "entity_id"
	entityTypeKey  = "entity_type"
	eventSparseKey = "event_sparse"
)

func encodeEventMeta(evt events.Event, id *ttnpb.EntityIdentifiers) ([]string, error) {
	meta := []string{
		eventUIDKey, evt.UniqueID(),
		eventNameKey, evt.Name(),
		entityIDKey, unique.ID(evt.Context(), id),
		entityTypeKey, id.EntityType(),
	}
	ePB, err := events.Proto(evt)
	if err != nil {
		return nil, err
	}
	sparsePB := &ttnpb.Event{
		// Name is stored in eventNameKey
		Time: ePB.Time,
		// Identifiers is stored in entityTypeKey+entityIDKey
		Context:    ePB.Context,
		Visibility: ePB.Visibility,
		// UniqueID is stored in eventUIDKey
	}
	rb, err := proto.Marshal(sparsePB)
	if err != nil {
		return nil, err
	}
	return append(meta, eventSparseKey, encodeBinary(rb)), nil
}

func decodeEventMeta(values map[string]any) (*ttnpb.Event, error) {
	var (
		sparseEvent string
		pb          ttnpb.Event
		ok          bool
	)
	if sparseEvent, ok = values[eventSparseKey].(string); ok {
		b64, err := decodeBinary(sparseEvent)
		if err != nil {
			return nil, err
		}
		if err = proto.Unmarshal(b64, &pb); err != nil {
			return nil, err
		}
	}
	if pb.UniqueId, ok = values[eventUIDKey].(string); !ok {
		return nil, errUnknownEncoding.New()
	}
	if pb.Name, ok = values[eventNameKey].(string); !ok {
		return nil, errUnknownEncoding.New()
	}
	if entityType, ok := values[entityTypeKey].(string); ok {
		if uid, ok := values[entityIDKey].(string); ok {
			switch entityType {
			case "application":
				id, err := unique.ToApplicationID(uid)
				if err != nil {
					return nil, err
				}
				pb.Identifiers = append(pb.Identifiers, id.GetEntityIdentifiers())
			case "client":
				id, err := unique.ToClientID(uid)
				if err != nil {
					return nil, err
				}
				pb.Identifiers = append(pb.Identifiers, id.GetEntityIdentifiers())
			case "end device":
				id, err := unique.ToDeviceID(uid)
				if err != nil {
					return nil, err
				}
				pb.Identifiers = append(pb.Identifiers, id.GetEntityIdentifiers())
			case "gateway":
				id, err := unique.ToGatewayID(uid)
				if err != nil {
					return nil, err
				}
				pb.Identifiers = append(pb.Identifiers, id.GetEntityIdentifiers())
			case "organization":
				id, err := unique.ToOrganizationID(uid)
				if err != nil {
					return nil, err
				}
				pb.Identifiers = append(pb.Identifiers, id.GetEntityIdentifiers())
			case "user":
				id, err := unique.ToUserID(uid)
				if err != nil {
					return nil, err
				}
				pb.Identifiers = append(pb.Identifiers, id.GetEntityIdentifiers())
			}
		}
	}
	return &pb, nil
}
